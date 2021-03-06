// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetPushTemplateInput struct {
	_ struct{} `type:"structure"`

	// TemplateName is a required field
	TemplateName *string `location:"uri" locationName:"template-name" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPushTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPushTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPushTemplateInput"}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPushTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TemplateName != nil {
		v := *s.TemplateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "template-name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetPushTemplateOutput struct {
	_ struct{} `type:"structure" payload:"PushNotificationTemplateResponse"`

	// Provides information about the content and settings for a message template
	// that can be used in messages that are sent through a push notification channel.
	//
	// PushNotificationTemplateResponse is a required field
	PushNotificationTemplateResponse *PushNotificationTemplateResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetPushTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPushTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.PushNotificationTemplateResponse != nil {
		v := s.PushNotificationTemplateResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "PushNotificationTemplateResponse", v, metadata)
	}
	return nil
}

const opGetPushTemplate = "GetPushTemplate"

// GetPushTemplateRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves the content and settings for a message template that you can use
// in messages that are sent through a push notification channel.
//
//    // Example sending a request using GetPushTemplateRequest.
//    req := client.GetPushTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetPushTemplate
func (c *Client) GetPushTemplateRequest(input *GetPushTemplateInput) GetPushTemplateRequest {
	op := &aws.Operation{
		Name:       opGetPushTemplate,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/templates/{template-name}/push",
	}

	if input == nil {
		input = &GetPushTemplateInput{}
	}

	req := c.newRequest(op, input, &GetPushTemplateOutput{})
	return GetPushTemplateRequest{Request: req, Input: input, Copy: c.GetPushTemplateRequest}
}

// GetPushTemplateRequest is the request type for the
// GetPushTemplate API operation.
type GetPushTemplateRequest struct {
	*aws.Request
	Input *GetPushTemplateInput
	Copy  func(*GetPushTemplateInput) GetPushTemplateRequest
}

// Send marshals and sends the GetPushTemplate API request.
func (r GetPushTemplateRequest) Send(ctx context.Context) (*GetPushTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPushTemplateResponse{
		GetPushTemplateOutput: r.Request.Data.(*GetPushTemplateOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPushTemplateResponse is the response type for the
// GetPushTemplate API operation.
type GetPushTemplateResponse struct {
	*GetPushTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPushTemplate request.
func (r *GetPushTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
