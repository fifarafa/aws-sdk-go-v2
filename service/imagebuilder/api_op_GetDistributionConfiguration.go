// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetDistributionConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the distribution configuration that you
	// wish to retrieve.
	//
	// DistributionConfigurationArn is a required field
	DistributionConfigurationArn *string `location:"querystring" locationName:"distributionConfigurationArn" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDistributionConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDistributionConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDistributionConfigurationInput"}

	if s.DistributionConfigurationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DistributionConfigurationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDistributionConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DistributionConfigurationArn != nil {
		v := *s.DistributionConfigurationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "distributionConfigurationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetDistributionConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The distribution configuration object.
	DistributionConfiguration *DistributionConfiguration `locationName:"distributionConfiguration" type:"structure"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s GetDistributionConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDistributionConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DistributionConfiguration != nil {
		v := s.DistributionConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "distributionConfiguration", v, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetDistributionConfiguration = "GetDistributionConfiguration"

// GetDistributionConfigurationRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// Gets a distribution configuration.
//
//    // Example sending a request using GetDistributionConfigurationRequest.
//    req := client.GetDistributionConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/GetDistributionConfiguration
func (c *Client) GetDistributionConfigurationRequest(input *GetDistributionConfigurationInput) GetDistributionConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetDistributionConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/GetDistributionConfiguration",
	}

	if input == nil {
		input = &GetDistributionConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetDistributionConfigurationOutput{})
	return GetDistributionConfigurationRequest{Request: req, Input: input, Copy: c.GetDistributionConfigurationRequest}
}

// GetDistributionConfigurationRequest is the request type for the
// GetDistributionConfiguration API operation.
type GetDistributionConfigurationRequest struct {
	*aws.Request
	Input *GetDistributionConfigurationInput
	Copy  func(*GetDistributionConfigurationInput) GetDistributionConfigurationRequest
}

// Send marshals and sends the GetDistributionConfiguration API request.
func (r GetDistributionConfigurationRequest) Send(ctx context.Context) (*GetDistributionConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDistributionConfigurationResponse{
		GetDistributionConfigurationOutput: r.Request.Data.(*GetDistributionConfigurationOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDistributionConfigurationResponse is the response type for the
// GetDistributionConfiguration API operation.
type GetDistributionConfigurationResponse struct {
	*GetDistributionConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDistributionConfiguration request.
func (r *GetDistributionConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}