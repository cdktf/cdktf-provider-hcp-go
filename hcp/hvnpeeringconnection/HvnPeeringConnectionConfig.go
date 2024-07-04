// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hvnpeeringconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HvnPeeringConnectionConfig struct {
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
	// The unique URL of one of the HVNs being peered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.0/docs/resources/hvn_peering_connection#hvn_1 HvnPeeringConnection#hvn_1}
	Hvn1 *string `field:"required" json:"hvn1" yaml:"hvn1"`
	// The unique URL of one of the HVNs being peered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.0/docs/resources/hvn_peering_connection#hvn_2 HvnPeeringConnection#hvn_2}
	Hvn2 *string `field:"required" json:"hvn2" yaml:"hvn2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.0/docs/resources/hvn_peering_connection#id HvnPeeringConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the HCP project where HVN peering connection is located.
	//
	// Always matches hvn_1's project ID. Setting this attribute is deprecated, but it will remain usable in read-only form.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.0/docs/resources/hvn_peering_connection#project_id HvnPeeringConnection#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.0/docs/resources/hvn_peering_connection#timeouts HvnPeeringConnection#timeouts}
	Timeouts *HvnPeeringConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

