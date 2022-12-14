package datahcppackeriteration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataHcpPackerIterationConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/packer_iteration#bucket_name DataHcpPackerIteration#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The channel that points to the version of the image you want.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/packer_iteration#channel DataHcpPackerIteration#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/packer_iteration#id DataHcpPackerIteration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/packer_iteration#timeouts DataHcpPackerIteration#timeouts}
	Timeouts *DataHcpPackerIterationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

