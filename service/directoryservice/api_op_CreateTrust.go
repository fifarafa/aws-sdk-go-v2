// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// AWS Directory Service for Microsoft Active Directory allows you to configure
// trust relationships. For example, you can establish a trust between your
// AWS Managed Microsoft AD directory, and your existing on-premises Microsoft
// Active Directory. This would allow you to provide users and groups access
// to resources in either domain, with a single set of credentials.
//
// This action initiates the creation of the AWS side of a trust relationship
// between an AWS Managed Microsoft AD directory and an external domain.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/CreateTrustRequest
type CreateTrustInput struct {
	_ struct{} `type:"structure"`

	// The IP addresses of the remote DNS server associated with RemoteDomainName.
	ConditionalForwarderIpAddrs []string `type:"list"`

	// The Directory ID of the AWS Managed Microsoft AD directory for which to establish
	// the trust relationship.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`

	// The Fully Qualified Domain Name (FQDN) of the external domain for which to
	// create the trust relationship.
	//
	// RemoteDomainName is a required field
	RemoteDomainName *string `type:"string" required:"true"`

	// Optional parameter to enable selective authentication for the trust.
	SelectiveAuth SelectiveAuth `type:"string" enum:"true"`

	// The direction of the trust relationship.
	//
	// TrustDirection is a required field
	TrustDirection TrustDirection `type:"string" required:"true" enum:"true"`

	// The trust password. The must be the same password that was used when creating
	// the trust relationship on the external domain.
	//
	// TrustPassword is a required field
	TrustPassword *string `min:"1" type:"string" required:"true"`

	// The trust relationship type. Forest is the default.
	TrustType TrustType `type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateTrustInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTrustInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTrustInput"}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}

	if s.RemoteDomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RemoteDomainName"))
	}
	if len(s.TrustDirection) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("TrustDirection"))
	}

	if s.TrustPassword == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrustPassword"))
	}
	if s.TrustPassword != nil && len(*s.TrustPassword) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TrustPassword", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a CreateTrust request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/CreateTrustResult
type CreateTrustOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the trust relationship that was created.
	TrustId *string `type:"string"`
}

// String returns the string representation
func (s CreateTrustOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateTrust = "CreateTrust"

// CreateTrustRequest returns a request value for making API operation for
// AWS Directory Service.
//
// AWS Directory Service for Microsoft Active Directory allows you to configure
// trust relationships. For example, you can establish a trust between your
// AWS Managed Microsoft AD directory, and your existing on-premises Microsoft
// Active Directory. This would allow you to provide users and groups access
// to resources in either domain, with a single set of credentials.
//
// This action initiates the creation of the AWS side of a trust relationship
// between an AWS Managed Microsoft AD directory and an external domain. You
// can create either a forest trust or an external trust.
//
//    // Example sending a request using CreateTrustRequest.
//    req := client.CreateTrustRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/CreateTrust
func (c *Client) CreateTrustRequest(input *CreateTrustInput) CreateTrustRequest {
	op := &aws.Operation{
		Name:       opCreateTrust,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTrustInput{}
	}

	req := c.newRequest(op, input, &CreateTrustOutput{})
	return CreateTrustRequest{Request: req, Input: input, Copy: c.CreateTrustRequest}
}

// CreateTrustRequest is the request type for the
// CreateTrust API operation.
type CreateTrustRequest struct {
	*aws.Request
	Input *CreateTrustInput
	Copy  func(*CreateTrustInput) CreateTrustRequest
}

// Send marshals and sends the CreateTrust API request.
func (r CreateTrustRequest) Send(ctx context.Context) (*CreateTrustResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTrustResponse{
		CreateTrustOutput: r.Request.Data.(*CreateTrustOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTrustResponse is the response type for the
// CreateTrust API operation.
type CreateTrustResponse struct {
	*CreateTrustOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTrust request.
func (r *CreateTrustResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}