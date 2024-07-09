// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointaddondefinition

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WaypointAddOnDefinitionConfig struct {
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
	// A longer description of the Add-on Definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.1/docs/resources/waypoint_add_on_definition#description WaypointAddOnDefinition#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the Add-on Definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.1/docs/resources/waypoint_add_on_definition#name WaypointAddOnDefinition#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A short summary of the Add-on Definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.1/docs/resources/waypoint_add_on_definition#summary WaypointAddOnDefinition#summary}
	Summary *string `field:"required" json:"summary" yaml:"summary"`
	// Terraform Cloud Workspace details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.1/docs/resources/waypoint_add_on_definition#terraform_cloud_workspace_details WaypointAddOnDefinition#terraform_cloud_workspace_details}
	TerraformCloudWorkspaceDetails *WaypointAddOnDefinitionTerraformCloudWorkspaceDetails `field:"required" json:"terraformCloudWorkspaceDetails" yaml:"terraformCloudWorkspaceDetails"`
	// Terraform Cloud no-code Module details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.1/docs/resources/waypoint_add_on_definition#terraform_no_code_module WaypointAddOnDefinition#terraform_no_code_module}
	TerraformNoCodeModule *WaypointAddOnDefinitionTerraformNoCodeModule `field:"required" json:"terraformNoCodeModule" yaml:"terraformNoCodeModule"`
	// List of labels attached to this Add-on Definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.1/docs/resources/waypoint_add_on_definition#labels WaypointAddOnDefinition#labels}
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// The ID of the HCP project where the Waypoint Add-on Definition is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.1/docs/resources/waypoint_add_on_definition#project_id WaypointAddOnDefinition#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// The markdown template for the Add-on Definition README.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.1/docs/resources/waypoint_add_on_definition#readme_markdown_template WaypointAddOnDefinition#readme_markdown_template}
	ReadmeMarkdownTemplate *string `field:"optional" json:"readmeMarkdownTemplate" yaml:"readmeMarkdownTemplate"`
	// List of variable options for the Add-on Definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.1/docs/resources/waypoint_add_on_definition#variable_options WaypointAddOnDefinition#variable_options}
	VariableOptions interface{} `field:"optional" json:"variableOptions" yaml:"variableOptions"`
}

