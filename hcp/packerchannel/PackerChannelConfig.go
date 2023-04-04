package packerchannel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PackerChannelConfig struct {
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
	// The slug of the HCP Packer Registry image bucket where the channel should be created in.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/packer_channel#bucket_name PackerChannel#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The name of the channel being managed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/packer_channel#name PackerChannel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/packer_channel#id PackerChannel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// iteration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/packer_channel#iteration PackerChannel#iteration}
	Iteration *PackerChannelIteration `field:"optional" json:"iteration" yaml:"iteration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/packer_channel#timeouts PackerChannel#timeouts}
	Timeouts *PackerChannelTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

