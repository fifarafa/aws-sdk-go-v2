// +build codegen

package api

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func buildAPI() *API {
	a := &API{}

	stringShape := &Shape{
		API:       a,
		ShapeName: "string",
		Type:      "string",
	}
	stringShapeRef := &ShapeRef{
		API:       a,
		ShapeName: "string",
		Shape:     stringShape,
	}

	intShape := &Shape{
		API:       a,
		ShapeName: "int",
		Type:      "int",
	}
	intShapeRef := &ShapeRef{
		API:       a,
		ShapeName: "int",
		Shape:     intShape,
	}

	input := &Shape{
		API:       a,
		ShapeName: "FooInput",
		MemberRefs: map[string]*ShapeRef{
			"BarShape": stringShapeRef,
		},
		Type: "structure",
	}
	output := &Shape{
		API:       a,
		ShapeName: "FooOutput",
		MemberRefs: map[string]*ShapeRef{
			"BazShape": intShapeRef,
		},
		Type: "structure",
	}

	inputRef := ShapeRef{
		API:       a,
		ShapeName: "FooInput",
		Shape:     input,
	}
	outputRef := ShapeRef{
		API:       a,
		ShapeName: "FooOutput",
		Shape:     output,
	}

	operations := map[string]*Operation{
		"Foo": {
			API:          a,
			Name:         "Foo",
			ExportedName: "Foo",
			InputRef:     inputRef,
			OutputRef:    outputRef,
		},
	}

	a.Operations = operations
	a.Shapes = map[string]*Shape{
		"FooInput":  input,
		"FooOutput": output,
		"string":    stringShape,
		"int":       intShape,
	}
	a.Metadata = Metadata{
		ServiceAbbreviation: "FooService",
	}

	a.BaseImportPath = "github.com/aws/aws-sdk-go-v2/service/"

	a.Setup()
	return a
}

func TestExampleGeneration(t *testing.T) {
	example := `
{
  "version": "1.0",
  "examples": {
    "Foo": [
      {
        "input": {
          "BarShape": "Hello world"
        },
        "output": {
          "BazShape": 1
        },
        "comments": {
          "input": {
          },
          "output": {
          }
        },
        "description": "Foo bar baz qux",
        "title": "I pity the foo"
      }
    ]
  }
}
	`
	a := buildAPI()
	def := &ExamplesDefinition{}
	err := json.Unmarshal([]byte(example), def)
	if err != nil {
		t.Error(err)
	}
	def.API = a

	def.setup()
	expected := `
import (
	"fmt"
	"context"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/fooservice"
	
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// I pity the foo
//
// Foo bar baz qux
func ExampleClient_FooRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := fooservice.New(cfg)
	input := &fooservice.FooInput{
		BarShape: aws.String("Hello world"),
	}

	req := svc.FooRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
`
	if v := cmp.Diff(expected, a.ExamplesGoCode()); len(v) != 0 {
		t.Errorf(v)
	}
}

func TestBuildShape(t *testing.T) {
	a := buildAPI()
	cases := []struct {
		defs     map[string]interface{}
		expected string
	}{
		{
			defs: map[string]interface{}{
				"barShape": "Hello World",
			},
			expected: "BarShape: aws.String(\"Hello World\"),\n",
		},
		{
			defs: map[string]interface{}{
				"BarShape": "Hello World",
			},
			expected: "BarShape: aws.String(\"Hello World\"),\n",
		},
	}

	for _, c := range cases {
		ref := a.Operations["Foo"].InputRef
		shapeStr := defaultExamplesBuilder{}.BuildShape(&ref, c.defs, false, false)
		if c.expected != shapeStr {
			t.Errorf("Expected:\n%s\nReceived:\n%s", c.expected, shapeStr)
		}
	}
}
