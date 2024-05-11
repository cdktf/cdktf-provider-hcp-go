// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package awsnetworkpeering


type AwsNetworkPeeringTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.89.0/docs/resources/aws_network_peering#create AwsNetworkPeering#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.89.0/docs/resources/aws_network_peering#default AwsNetworkPeering#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.89.0/docs/resources/aws_network_peering#delete AwsNetworkPeering#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

