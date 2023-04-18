package datahcppackerimageiteration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataHcpPackerImageIterationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.56.0/docs/data-sources/packer_image_iteration#bucket_name DataHcpPackerImageIteration#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The channel that points to the version of the image you want.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.56.0/docs/data-sources/packer_image_iteration#channel DataHcpPackerImageIteration#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.56.0/docs/data-sources/packer_image_iteration#id DataHcpPackerImageIteration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.56.0/docs/data-sources/packer_image_iteration#timeouts DataHcpPackerImageIteration#timeouts}
	Timeouts *DataHcpPackerImageIterationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

