// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ForgetSmartHomeAppliancesInput struct {
	_ struct{} `type:"structure"`

	// The room that the appliances are associated with.
	//
	// RoomArn is a required field
	RoomArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ForgetSmartHomeAppliancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ForgetSmartHomeAppliancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ForgetSmartHomeAppliancesInput"}

	if s.RoomArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoomArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ForgetSmartHomeAppliancesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ForgetSmartHomeAppliancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opForgetSmartHomeAppliances = "ForgetSmartHomeAppliances"

// ForgetSmartHomeAppliancesRequest returns a request value for making API operation for
// Alexa For Business.
//
// Forgets smart home appliances associated to a room.
//
//    // Example sending a request using ForgetSmartHomeAppliancesRequest.
//    req := client.ForgetSmartHomeAppliancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/ForgetSmartHomeAppliances
func (c *Client) ForgetSmartHomeAppliancesRequest(input *ForgetSmartHomeAppliancesInput) ForgetSmartHomeAppliancesRequest {
	op := &aws.Operation{
		Name:       opForgetSmartHomeAppliances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ForgetSmartHomeAppliancesInput{}
	}

	req := c.newRequest(op, input, &ForgetSmartHomeAppliancesOutput{})
	return ForgetSmartHomeAppliancesRequest{Request: req, Input: input, Copy: c.ForgetSmartHomeAppliancesRequest}
}

// ForgetSmartHomeAppliancesRequest is the request type for the
// ForgetSmartHomeAppliances API operation.
type ForgetSmartHomeAppliancesRequest struct {
	*aws.Request
	Input *ForgetSmartHomeAppliancesInput
	Copy  func(*ForgetSmartHomeAppliancesInput) ForgetSmartHomeAppliancesRequest
}

// Send marshals and sends the ForgetSmartHomeAppliances API request.
func (r ForgetSmartHomeAppliancesRequest) Send(ctx context.Context) (*ForgetSmartHomeAppliancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ForgetSmartHomeAppliancesResponse{
		ForgetSmartHomeAppliancesOutput: r.Request.Data.(*ForgetSmartHomeAppliancesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ForgetSmartHomeAppliancesResponse is the response type for the
// ForgetSmartHomeAppliances API operation.
type ForgetSmartHomeAppliancesResponse struct {
	*ForgetSmartHomeAppliancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ForgetSmartHomeAppliances request.
func (r *ForgetSmartHomeAppliancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
