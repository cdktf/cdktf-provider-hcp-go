package hvnroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HvnRouteConfig struct {
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
	// The destination CIDR of the HVN route.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/hvn_route#destination_cidr HvnRoute#destination_cidr}
	DestinationCidr *string `field:"required" json:"destinationCidr" yaml:"destinationCidr"`
	// The `self_link` of the HashiCorp Virtual Network (HVN).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/hvn_route#hvn_link HvnRoute#hvn_link}
	HvnLink *string `field:"required" json:"hvnLink" yaml:"hvnLink"`
	// The ID of the HVN route.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/hvn_route#hvn_route_id HvnRoute#hvn_route_id}
	HvnRouteId *string `field:"required" json:"hvnRouteId" yaml:"hvnRouteId"`
	// A unique URL identifying the target of the HVN route. Examples of the target: [`aws_network_peering`](aws_network_peering.md), [`aws_transit_gateway_attachment`](aws_transit_gateway_attachment.md).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/hvn_route#target_link HvnRoute#target_link}
	TargetLink *string `field:"required" json:"targetLink" yaml:"targetLink"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/hvn_route#id HvnRoute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/hvn_route#timeouts HvnRoute#timeouts}
	Timeouts *HvnRouteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

