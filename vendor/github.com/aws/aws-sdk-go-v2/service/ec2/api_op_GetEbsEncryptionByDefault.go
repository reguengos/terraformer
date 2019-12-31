// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetEbsEncryptionByDefaultInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`
}

// String returns the string representation
func (s GetEbsEncryptionByDefaultInput) String() string {
	return awsutil.Prettify(s)
}

type GetEbsEncryptionByDefaultOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether encryption by default is enabled.
	EbsEncryptionByDefault *bool `locationName:"ebsEncryptionByDefault" type:"boolean"`
}

// String returns the string representation
func (s GetEbsEncryptionByDefaultOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetEbsEncryptionByDefault = "GetEbsEncryptionByDefault"

// GetEbsEncryptionByDefaultRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes whether EBS encryption by default is enabled for your account in
// the current Region.
//
// For more information, see Amazon EBS Encryption (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using GetEbsEncryptionByDefaultRequest.
//    req := client.GetEbsEncryptionByDefaultRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetEbsEncryptionByDefault
func (c *Client) GetEbsEncryptionByDefaultRequest(input *GetEbsEncryptionByDefaultInput) GetEbsEncryptionByDefaultRequest {
	op := &aws.Operation{
		Name:       opGetEbsEncryptionByDefault,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetEbsEncryptionByDefaultInput{}
	}

	req := c.newRequest(op, input, &GetEbsEncryptionByDefaultOutput{})
	return GetEbsEncryptionByDefaultRequest{Request: req, Input: input, Copy: c.GetEbsEncryptionByDefaultRequest}
}

// GetEbsEncryptionByDefaultRequest is the request type for the
// GetEbsEncryptionByDefault API operation.
type GetEbsEncryptionByDefaultRequest struct {
	*aws.Request
	Input *GetEbsEncryptionByDefaultInput
	Copy  func(*GetEbsEncryptionByDefaultInput) GetEbsEncryptionByDefaultRequest
}

// Send marshals and sends the GetEbsEncryptionByDefault API request.
func (r GetEbsEncryptionByDefaultRequest) Send(ctx context.Context) (*GetEbsEncryptionByDefaultResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEbsEncryptionByDefaultResponse{
		GetEbsEncryptionByDefaultOutput: r.Request.Data.(*GetEbsEncryptionByDefaultOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEbsEncryptionByDefaultResponse is the response type for the
// GetEbsEncryptionByDefault API operation.
type GetEbsEncryptionByDefaultResponse struct {
	*GetEbsEncryptionByDefaultOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEbsEncryptionByDefault request.
func (r *GetEbsEncryptionByDefaultResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}