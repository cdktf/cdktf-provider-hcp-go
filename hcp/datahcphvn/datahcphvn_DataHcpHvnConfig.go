package datahcphvn

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataHcpHvnConfig struct {
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
	// The ID of the HashiCorp Virtual Network (HVN).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/hvn#hvn_id DataHcpHvn#hvn_id}
	HvnId *string `field:"required" json:"hvnId" yaml:"hvnId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/hvn#id DataHcpHvn#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/hvn#timeouts DataHcpHvn#timeouts}
	Timeouts *DataHcpHvnTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

