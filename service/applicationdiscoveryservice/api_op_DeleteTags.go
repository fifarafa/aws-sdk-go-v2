// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteTagsInput struct {
	_ struct{} `type:"structure"`

	// A list of configuration items with tags that you want to delete.
	//
	// ConfigurationIds is a required field
	ConfigurationIds []string `locationName:"configurationIds" type:"list" required:"true"`

	// Tags that you want to delete from one or more configuration items. Specify
	// the tags that you want to delete in a key-value format. For example:
	//
	// {"key": "serverType", "value": "webServer"}
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s DeleteTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTagsInput"}

	if s.ConfigurationIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationIds"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteTagsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTags = "DeleteTags"

// DeleteTagsRequest returns a request value for making API operation for
// AWS Application Discovery Service.
//
// Deletes the association between configuration items and one or more tags.
// This API accepts a list of multiple configuration items.
//
//    // Example sending a request using DeleteTagsRequest.
//    req := client.DeleteTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/DeleteTags
func (c *Client) DeleteTagsRequest(input *DeleteTagsInput) DeleteTagsRequest {
	op := &aws.Operation{
		Name:       opDeleteTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTagsInput{}
	}

	req := c.newRequest(op, input, &DeleteTagsOutput{})
	return DeleteTagsRequest{Request: req, Input: input, Copy: c.DeleteTagsRequest}
}

// DeleteTagsRequest is the request type for the
// DeleteTags API operation.
type DeleteTagsRequest struct {
	*aws.Request
	Input *DeleteTagsInput
	Copy  func(*DeleteTagsInput) DeleteTagsRequest
}

// Send marshals and sends the DeleteTags API request.
func (r DeleteTagsRequest) Send(ctx context.Context) (*DeleteTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTagsResponse{
		DeleteTagsOutput: r.Request.Data.(*DeleteTagsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTagsResponse is the response type for the
// DeleteTags API operation.
type DeleteTagsResponse struct {
	*DeleteTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTags request.
func (r *DeleteTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
