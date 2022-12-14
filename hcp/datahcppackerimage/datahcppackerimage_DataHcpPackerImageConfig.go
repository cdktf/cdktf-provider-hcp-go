package datahcppackerimage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataHcpPackerImageConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The slug of the HCP Packer Registry image bucket to pull from.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/packer_image#bucket_name DataHcpPackerImage#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Name of the cloud provider this image is stored-in.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/packer_image#cloud_provider DataHcpPackerImage#cloud_provider}
	CloudProvider *string `field:"required" json:"cloudProvider" yaml:"cloudProvider"`
	// Region this image is stored in, if any.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/packer_image#region DataHcpPackerImage#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// The channel that points to the version of the image being retrieved.
	//
	// Either this or `iteration_id` must be specified. Note: will incur a billable request
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/packer_image#channel DataHcpPackerImage#channel}
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// Name of the builder that built this image. Ex: `amazon-ebs.example`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/packer_image#component_type DataHcpPackerImage#component_type}
	ComponentType *string `field:"optional" json:"componentType" yaml:"componentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/packer_image#id DataHcpPackerImage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The iteration from which to get the image. Either this or `channel` must be specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/packer_image#iteration_id DataHcpPackerImage#iteration_id}
	IterationId *string `field:"optional" json:"iterationId" yaml:"iterationId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/packer_image#timeouts DataHcpPackerImage#timeouts}
	Timeouts *DataHcpPackerImageTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

