// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hvnroute


type HvnRouteTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.86.0/docs/resources/hvn_route#create HvnRoute#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.86.0/docs/resources/hvn_route#default HvnRoute#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.86.0/docs/resources/hvn_route#delete HvnRoute#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

