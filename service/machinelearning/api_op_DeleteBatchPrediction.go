// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteBatchPredictionInput struct {
	_ struct{} `type:"structure"`

	// A user-supplied ID that uniquely identifies the BatchPrediction.
	//
	// BatchPredictionId is a required field
	BatchPredictionId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBatchPredictionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBatchPredictionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBatchPredictionInput"}

	if s.BatchPredictionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BatchPredictionId"))
	}
	if s.BatchPredictionId != nil && len(*s.BatchPredictionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BatchPredictionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a DeleteBatchPrediction operation.
//
// You can use the GetBatchPrediction operation and check the value of the Status
// parameter to see whether a BatchPrediction is marked as DELETED.
type DeleteBatchPredictionOutput struct {
	_ struct{} `type:"structure"`

	// A user-supplied ID that uniquely identifies the BatchPrediction. This value
	// should be identical to the value of the BatchPredictionID in the request.
	BatchPredictionId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteBatchPredictionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteBatchPrediction = "DeleteBatchPrediction"

// DeleteBatchPredictionRequest returns a request value for making API operation for
// Amazon Machine Learning.
//
// Assigns the DELETED status to a BatchPrediction, rendering it unusable.
//
// After using the DeleteBatchPrediction operation, you can use the GetBatchPrediction
// operation to verify that the status of the BatchPrediction changed to DELETED.
//
// Caution: The result of the DeleteBatchPrediction operation is irreversible.
//
//    // Example sending a request using DeleteBatchPredictionRequest.
//    req := client.DeleteBatchPredictionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteBatchPredictionRequest(input *DeleteBatchPredictionInput) DeleteBatchPredictionRequest {
	op := &aws.Operation{
		Name:       opDeleteBatchPrediction,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteBatchPredictionInput{}
	}

	req := c.newRequest(op, input, &DeleteBatchPredictionOutput{})
	return DeleteBatchPredictionRequest{Request: req, Input: input, Copy: c.DeleteBatchPredictionRequest}
}

// DeleteBatchPredictionRequest is the request type for the
// DeleteBatchPrediction API operation.
type DeleteBatchPredictionRequest struct {
	*aws.Request
	Input *DeleteBatchPredictionInput
	Copy  func(*DeleteBatchPredictionInput) DeleteBatchPredictionRequest
}

// Send marshals and sends the DeleteBatchPrediction API request.
func (r DeleteBatchPredictionRequest) Send(ctx context.Context) (*DeleteBatchPredictionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBatchPredictionResponse{
		DeleteBatchPredictionOutput: r.Request.Data.(*DeleteBatchPredictionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBatchPredictionResponse is the response type for the
// DeleteBatchPrediction API operation.
type DeleteBatchPredictionResponse struct {
	*DeleteBatchPredictionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBatchPrediction request.
func (r *DeleteBatchPredictionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
