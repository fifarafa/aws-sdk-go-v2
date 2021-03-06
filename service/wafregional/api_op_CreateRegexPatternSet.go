// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/waf"
)

type CreateRegexPatternSetInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// A friendly name or description of the RegexPatternSet. You can't change Name
	// after you create a RegexPatternSet.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateRegexPatternSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRegexPatternSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRegexPatternSetInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateRegexPatternSetOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the CreateRegexPatternSet request.
	// You can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`

	// A RegexPatternSet that contains no objects.
	RegexPatternSet *waf.RegexPatternSet `type:"structure"`
}

// String returns the string representation
func (s CreateRegexPatternSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateRegexPatternSet = "CreateRegexPatternSet"

// CreateRegexPatternSetRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Creates a RegexPatternSet. You then use UpdateRegexPatternSet to specify
// the regular expression (regex) pattern that you want AWS WAF to search for,
// such as B[a@]dB[o0]t. You can then configure AWS WAF to reject those requests.
//
// To create and configure a RegexPatternSet, perform the following steps:
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a CreateRegexPatternSet request.
//
// Submit a CreateRegexPatternSet request.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateRegexPatternSet request.
//
// Submit an UpdateRegexPatternSet request to specify the string that you want
// AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using CreateRegexPatternSetRequest.
//    req := client.CreateRegexPatternSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/CreateRegexPatternSet
func (c *Client) CreateRegexPatternSetRequest(input *CreateRegexPatternSetInput) CreateRegexPatternSetRequest {
	op := &aws.Operation{
		Name:       opCreateRegexPatternSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateRegexPatternSetInput{}
	}

	req := c.newRequest(op, input, &CreateRegexPatternSetOutput{})
	return CreateRegexPatternSetRequest{Request: req, Input: input, Copy: c.CreateRegexPatternSetRequest}
}

// CreateRegexPatternSetRequest is the request type for the
// CreateRegexPatternSet API operation.
type CreateRegexPatternSetRequest struct {
	*aws.Request
	Input *CreateRegexPatternSetInput
	Copy  func(*CreateRegexPatternSetInput) CreateRegexPatternSetRequest
}

// Send marshals and sends the CreateRegexPatternSet API request.
func (r CreateRegexPatternSetRequest) Send(ctx context.Context) (*CreateRegexPatternSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRegexPatternSetResponse{
		CreateRegexPatternSetOutput: r.Request.Data.(*CreateRegexPatternSetOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRegexPatternSetResponse is the response type for the
// CreateRegexPatternSet API operation.
type CreateRegexPatternSetResponse struct {
	*CreateRegexPatternSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRegexPatternSet request.
func (r *CreateRegexPatternSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
