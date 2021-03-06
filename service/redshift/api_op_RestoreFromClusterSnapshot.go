// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RestoreFromClusterSnapshotInput struct {
	_ struct{} `type:"structure"`

	// Reserved.
	AdditionalInfo *string `type:"string"`

	// If true, major version upgrades can be applied during the maintenance window
	// to the Amazon Redshift engine that is running on the cluster.
	//
	// Default: true
	AllowVersionUpgrade *bool `type:"boolean"`

	// The number of days that automated snapshots are retained. If the value is
	// 0, automated snapshots are disabled. Even if automated snapshots are disabled,
	// you can still create manual snapshots when you want with CreateClusterSnapshot.
	//
	// Default: The value selected for the cluster from which the snapshot was taken.
	//
	// Constraints: Must be a value from 0 to 35.
	AutomatedSnapshotRetentionPeriod *int64 `type:"integer"`

	// The Amazon EC2 Availability Zone in which to restore the cluster.
	//
	// Default: A random, system-chosen Availability Zone.
	//
	// Example: us-east-1a
	AvailabilityZone *string `type:"string"`

	// The identifier of the cluster that will be created from restoring the snapshot.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 alphanumeric characters or hyphens.
	//
	//    * Alphabetic characters must be lowercase.
	//
	//    * First character must be a letter.
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens.
	//
	//    * Must be unique for all clusters within an AWS account.
	//
	// ClusterIdentifier is a required field
	ClusterIdentifier *string `type:"string" required:"true"`

	// The name of the parameter group to be associated with this cluster.
	//
	// Default: The default Amazon Redshift cluster parameter group. For information
	// about the default parameter group, go to Working with Amazon Redshift Parameter
	// Groups (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html).
	//
	// Constraints:
	//
	//    * Must be 1 to 255 alphanumeric characters or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens.
	ClusterParameterGroupName *string `type:"string"`

	// A list of security groups to be associated with this cluster.
	//
	// Default: The default cluster security group for Amazon Redshift.
	//
	// Cluster security groups only apply to clusters outside of VPCs.
	ClusterSecurityGroups []string `locationNameList:"ClusterSecurityGroupName" type:"list"`

	// The name of the subnet group where you want to cluster restored.
	//
	// A snapshot of cluster in VPC can be restored only in VPC. Therefore, you
	// must provide subnet group name where you want the cluster restored.
	ClusterSubnetGroupName *string `type:"string"`

	// The elastic IP (EIP) address for the cluster.
	ElasticIp *string `type:"string"`

	// An option that specifies whether to create the cluster with enhanced VPC
	// routing enabled. To create a cluster that uses enhanced VPC routing, the
	// cluster must be in a VPC. For more information, see Enhanced VPC Routing
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/enhanced-vpc-routing.html)
	// in the Amazon Redshift Cluster Management Guide.
	//
	// If this option is true, enhanced VPC routing is enabled.
	//
	// Default: false
	EnhancedVpcRouting *bool `type:"boolean"`

	// Specifies the name of the HSM client certificate the Amazon Redshift cluster
	// uses to retrieve the data encryption keys stored in an HSM.
	HsmClientCertificateIdentifier *string `type:"string"`

	// Specifies the name of the HSM configuration that contains the information
	// the Amazon Redshift cluster can use to retrieve and store keys in an HSM.
	HsmConfigurationIdentifier *string `type:"string"`

	// A list of AWS Identity and Access Management (IAM) roles that can be used
	// by the cluster to access other AWS services. You must supply the IAM roles
	// in their Amazon Resource Name (ARN) format. You can supply up to 10 IAM roles
	// in a single request.
	//
	// A cluster can have up to 10 IAM roles associated at any time.
	IamRoles []string `locationNameList:"IamRoleArn" type:"list"`

	// The AWS Key Management Service (KMS) key ID of the encryption key that you
	// want to use to encrypt data in the cluster that you restore from a shared
	// snapshot.
	KmsKeyId *string `type:"string"`

	// The name of the maintenance track for the restored cluster. When you take
	// a snapshot, the snapshot inherits the MaintenanceTrack value from the cluster.
	// The snapshot might be on a different track than the cluster that was the
	// source for the snapshot. For example, suppose that you take a snapshot of
	// a cluster that is on the current track and then change the cluster to be
	// on the trailing track. In this case, the snapshot and the source cluster
	// are on different tracks.
	MaintenanceTrackName *string `type:"string"`

	// The default number of days to retain a manual snapshot. If the value is -1,
	// the snapshot is retained indefinitely. This setting doesn't change the retention
	// period of existing snapshots.
	//
	// The value must be either -1 or an integer between 1 and 3,653.
	ManualSnapshotRetentionPeriod *int64 `type:"integer"`

	// The node type that the restored cluster will be provisioned with.
	//
	// Default: The node type of the cluster from which the snapshot was taken.
	// You can modify this if you are using any DS node type. In that case, you
	// can choose to restore into another DS node type of the same size. For example,
	// you can restore ds1.8xlarge into ds2.8xlarge, or ds1.xlarge into ds2.xlarge.
	// If you have a DC instance type, you must restore into that same instance
	// type and size. In other words, you can only restore a dc1.large instance
	// type into another dc1.large instance type or dc2.large instance type. You
	// can't restore dc1.8xlarge to dc2.8xlarge. First restore to a dc1.8xlareg
	// cluster, then resize to a dc2.8large cluster. For more information about
	// node types, see About Clusters and Nodes (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#rs-about-clusters-and-nodes)
	// in the Amazon Redshift Cluster Management Guide.
	NodeType *string `type:"string"`

	// The number of nodes specified when provisioning the restored cluster.
	NumberOfNodes *int64 `type:"integer"`

	// The AWS customer account used to create or copy the snapshot. Required if
	// you are restoring a snapshot you do not own, optional if you own the snapshot.
	OwnerAccount *string `type:"string"`

	// The port number on which the cluster accepts connections.
	//
	// Default: The same port as the original cluster.
	//
	// Constraints: Must be between 1115 and 65535.
	Port *int64 `type:"integer"`

	// The weekly time range (in UTC) during which automated cluster maintenance
	// can occur.
	//
	// Format: ddd:hh24:mi-ddd:hh24:mi
	//
	// Default: The value selected for the cluster from which the snapshot was taken.
	// For more information about the time blocks for each region, see Maintenance
	// Windows (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#rs-maintenance-windows)
	// in Amazon Redshift Cluster Management Guide.
	//
	// Valid Days: Mon | Tue | Wed | Thu | Fri | Sat | Sun
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `type:"string"`

	// If true, the cluster can be accessed from a public network.
	PubliclyAccessible *bool `type:"boolean"`

	// The name of the cluster the source snapshot was created from. This parameter
	// is required if your IAM user has a policy containing a snapshot resource
	// element that specifies anything other than * for the cluster name.
	SnapshotClusterIdentifier *string `type:"string"`

	// The name of the snapshot from which to create the new cluster. This parameter
	// isn't case sensitive.
	//
	// Example: my-snapshot-id
	//
	// SnapshotIdentifier is a required field
	SnapshotIdentifier *string `type:"string" required:"true"`

	// A unique identifier for the snapshot schedule.
	SnapshotScheduleIdentifier *string `type:"string"`

	// A list of Virtual Private Cloud (VPC) security groups to be associated with
	// the cluster.
	//
	// Default: The default VPC security group is associated with the cluster.
	//
	// VPC security groups only apply to clusters in VPCs.
	VpcSecurityGroupIds []string `locationNameList:"VpcSecurityGroupId" type:"list"`
}

// String returns the string representation
func (s RestoreFromClusterSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreFromClusterSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreFromClusterSnapshotInput"}

	if s.ClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterIdentifier"))
	}

	if s.SnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RestoreFromClusterSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// Describes a cluster.
	Cluster *Cluster `type:"structure"`
}

// String returns the string representation
func (s RestoreFromClusterSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opRestoreFromClusterSnapshot = "RestoreFromClusterSnapshot"

// RestoreFromClusterSnapshotRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Creates a new cluster from a snapshot. By default, Amazon Redshift creates
// the resulting cluster with the same configuration as the original cluster
// from which the snapshot was created, except that the new cluster is created
// with the default cluster security and parameter groups. After Amazon Redshift
// creates the cluster, you can use the ModifyCluster API to associate a different
// security group and different parameter group with the restored cluster. If
// you are using a DS node type, you can also choose to change to another DS
// node type of the same size during restore.
//
// If you restore a cluster into a VPC, you must provide a cluster subnet group
// where you want the cluster restored.
//
// For more information about working with snapshots, go to Amazon Redshift
// Snapshots (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html)
// in the Amazon Redshift Cluster Management Guide.
//
//    // Example sending a request using RestoreFromClusterSnapshotRequest.
//    req := client.RestoreFromClusterSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/RestoreFromClusterSnapshot
func (c *Client) RestoreFromClusterSnapshotRequest(input *RestoreFromClusterSnapshotInput) RestoreFromClusterSnapshotRequest {
	op := &aws.Operation{
		Name:       opRestoreFromClusterSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreFromClusterSnapshotInput{}
	}

	req := c.newRequest(op, input, &RestoreFromClusterSnapshotOutput{})
	return RestoreFromClusterSnapshotRequest{Request: req, Input: input, Copy: c.RestoreFromClusterSnapshotRequest}
}

// RestoreFromClusterSnapshotRequest is the request type for the
// RestoreFromClusterSnapshot API operation.
type RestoreFromClusterSnapshotRequest struct {
	*aws.Request
	Input *RestoreFromClusterSnapshotInput
	Copy  func(*RestoreFromClusterSnapshotInput) RestoreFromClusterSnapshotRequest
}

// Send marshals and sends the RestoreFromClusterSnapshot API request.
func (r RestoreFromClusterSnapshotRequest) Send(ctx context.Context) (*RestoreFromClusterSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreFromClusterSnapshotResponse{
		RestoreFromClusterSnapshotOutput: r.Request.Data.(*RestoreFromClusterSnapshotOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreFromClusterSnapshotResponse is the response type for the
// RestoreFromClusterSnapshot API operation.
type RestoreFromClusterSnapshotResponse struct {
	*RestoreFromClusterSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreFromClusterSnapshot request.
func (r *RestoreFromClusterSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
