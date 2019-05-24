// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the request to delete a user pool.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/DeleteUserPoolRequest
type DeleteUserPoolInput struct {
	_ struct{} `type:"structure"`

	// The user pool ID for the user pool you want to delete.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteUserPoolInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteUserPoolInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteUserPoolInput"}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/DeleteUserPoolOutput
type DeleteUserPoolOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteUserPoolOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteUserPool = "DeleteUserPool"

// DeleteUserPoolRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Deletes the specified Amazon Cognito user pool.
//
//    // Example sending a request using DeleteUserPoolRequest.
//    req := client.DeleteUserPoolRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/DeleteUserPool
func (c *Client) DeleteUserPoolRequest(input *DeleteUserPoolInput) DeleteUserPoolRequest {
	op := &aws.Operation{
		Name:       opDeleteUserPool,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteUserPoolInput{}
	}

	req := c.newRequest(op, input, &DeleteUserPoolOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteUserPoolRequest{Request: req, Input: input, Copy: c.DeleteUserPoolRequest}
}

// DeleteUserPoolRequest is the request type for the
// DeleteUserPool API operation.
type DeleteUserPoolRequest struct {
	*aws.Request
	Input *DeleteUserPoolInput
	Copy  func(*DeleteUserPoolInput) DeleteUserPoolRequest
}

// Send marshals and sends the DeleteUserPool API request.
func (r DeleteUserPoolRequest) Send(ctx context.Context) (*DeleteUserPoolResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteUserPoolResponse{
		DeleteUserPoolOutput: r.Request.Data.(*DeleteUserPoolOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteUserPoolResponse is the response type for the
// DeleteUserPool API operation.
type DeleteUserPoolResponse struct {
	*DeleteUserPoolOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteUserPool request.
func (r *DeleteUserPoolResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}