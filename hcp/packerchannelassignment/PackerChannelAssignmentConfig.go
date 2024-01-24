// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package packerchannelassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PackerChannelAssignmentConfig struct {
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
	// The slug of the HCP Packer Registry bucket where the channel is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.81.0/docs/resources/packer_channel_assignment#bucket_name PackerChannelAssignment#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The name of the HCP Packer channel being managed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.81.0/docs/resources/packer_channel_assignment#channel_name PackerChannelAssignment#channel_name}
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.81.0/docs/resources/packer_channel_assignment#id PackerChannelAssignment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The fingerprint of the iteration assigned to the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.81.0/docs/resources/packer_channel_assignment#iteration_fingerprint PackerChannelAssignment#iteration_fingerprint}
	IterationFingerprint *string `field:"optional" json:"iterationFingerprint" yaml:"iterationFingerprint"`
	// The ID of the iteration assigned to the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.81.0/docs/resources/packer_channel_assignment#iteration_id PackerChannelAssignment#iteration_id}
	IterationId *string `field:"optional" json:"iterationId" yaml:"iterationId"`
	// The incremental version of the iteration assigned to the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.81.0/docs/resources/packer_channel_assignment#iteration_version PackerChannelAssignment#iteration_version}
	IterationVersion *float64 `field:"optional" json:"iterationVersion" yaml:"iterationVersion"`
	// The ID of the HCP project where the channel is located.
	//
	// If not specified, the project specified in the HCP Provider config block will be used, if configured.
	// If a project is not configured in the HCP Provider config block, the oldest project in the organization will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.81.0/docs/resources/packer_channel_assignment#project_id PackerChannelAssignment#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.81.0/docs/resources/packer_channel_assignment#timeouts PackerChannelAssignment#timeouts}
	Timeouts *PackerChannelAssignmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

