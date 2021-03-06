// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteDiskSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The name of the disk snapshot you want to delete (e.g., my-disk-snapshot).
	//
	// DiskSnapshotName is a required field
	DiskSnapshotName *string `locationName:"diskSnapshotName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDiskSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDiskSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDiskSnapshotInput"}

	if s.DiskSnapshotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DiskSnapshotName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDiskSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// An object describing the API operations.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s DeleteDiskSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDiskSnapshot = "DeleteDiskSnapshot"

// DeleteDiskSnapshotRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Deletes the specified disk snapshot.
//
// When you make periodic snapshots of a disk, the snapshots are incremental,
// and only the blocks on the device that have changed since your last snapshot
// are saved in the new snapshot. When you delete a snapshot, only the data
// not needed for any other snapshot is removed. So regardless of which prior
// snapshots have been deleted, all active snapshots will have access to all
// the information needed to restore the disk.
//
// The delete disk snapshot operation supports tag-based access control via
// resource tags applied to the resource identified by disk snapshot name. For
// more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using DeleteDiskSnapshotRequest.
//    req := client.DeleteDiskSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteDiskSnapshot
func (c *Client) DeleteDiskSnapshotRequest(input *DeleteDiskSnapshotInput) DeleteDiskSnapshotRequest {
	op := &aws.Operation{
		Name:       opDeleteDiskSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDiskSnapshotInput{}
	}

	req := c.newRequest(op, input, &DeleteDiskSnapshotOutput{})
	return DeleteDiskSnapshotRequest{Request: req, Input: input, Copy: c.DeleteDiskSnapshotRequest}
}

// DeleteDiskSnapshotRequest is the request type for the
// DeleteDiskSnapshot API operation.
type DeleteDiskSnapshotRequest struct {
	*aws.Request
	Input *DeleteDiskSnapshotInput
	Copy  func(*DeleteDiskSnapshotInput) DeleteDiskSnapshotRequest
}

// Send marshals and sends the DeleteDiskSnapshot API request.
func (r DeleteDiskSnapshotRequest) Send(ctx context.Context) (*DeleteDiskSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDiskSnapshotResponse{
		DeleteDiskSnapshotOutput: r.Request.Data.(*DeleteDiskSnapshotOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDiskSnapshotResponse is the response type for the
// DeleteDiskSnapshot API operation.
type DeleteDiskSnapshotResponse struct {
	*DeleteDiskSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDiskSnapshot request.
func (r *DeleteDiskSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
