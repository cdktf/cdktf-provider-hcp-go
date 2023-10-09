// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datahcpconsulcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataHcpConsulClusterConfig struct {
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
	// The ID of the HCP Consul cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.2/docs/data-sources/consul_cluster#cluster_id DataHcpConsulCluster#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.2/docs/data-sources/consul_cluster#id DataHcpConsulCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the HCP project where the HCP Consul cluster is located.
	//
	// If not specified, the project specified in the HCP Provider config block will be used, if configured.
	// If a project is not configured in the HCP Provider config block, the oldest project in the organization will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.2/docs/data-sources/consul_cluster#project_id DataHcpConsulCluster#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.2/docs/data-sources/consul_cluster#timeouts DataHcpConsulCluster#timeouts}
	Timeouts *DataHcpConsulClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

