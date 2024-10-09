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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.97.0/docs/resources/waypoint_add_on_definition#description WaypointAddOnDefinition#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the Add-on Definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.97.0/docs/resources/waypoint_add_on_definition#name WaypointAddOnDefinition#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A short summary of the Add-on Definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.97.0/docs/resources/waypoint_add_on_definition#summary WaypointAddOnDefinition#summary}
	Summary *string `field:"required" json:"summary" yaml:"summary"`
	// Terraform Cloud no-code Module Source, expected to be in one of the following formats: "app.terraform.io/hcp_waypoint_example/ecs-advanced-microservice/aws" or "private/hcp_waypoint_example/ecs-advanced-microservice/aws".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.97.0/docs/resources/waypoint_add_on_definition#terraform_no_code_module_source WaypointAddOnDefinition#terraform_no_code_module_source}
	TerraformNoCodeModuleSource *string `field:"required" json:"terraformNoCodeModuleSource" yaml:"terraformNoCodeModuleSource"`
	// The ID of the Terraform Cloud Project to create workspaces in.
	//
	// The ID is found on the Terraform Cloud Project settings page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.97.0/docs/resources/waypoint_add_on_definition#terraform_project_id WaypointAddOnDefinition#terraform_project_id}
	TerraformProjectId *string `field:"required" json:"terraformProjectId" yaml:"terraformProjectId"`
	// List of labels attached to this Add-on Definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.97.0/docs/resources/waypoint_add_on_definition#labels WaypointAddOnDefinition#labels}
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// The ID of the HCP project where the Waypoint Add-on Definition is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.97.0/docs/resources/waypoint_add_on_definition#project_id WaypointAddOnDefinition#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// The markdown template for the Add-on Definition README (markdown format supported).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.97.0/docs/resources/waypoint_add_on_definition#readme_markdown_template WaypointAddOnDefinition#readme_markdown_template}
	ReadmeMarkdownTemplate *string `field:"optional" json:"readmeMarkdownTemplate" yaml:"readmeMarkdownTemplate"`
	// The ID of the Terraform agent pool to use for running Terraform operations.
	//
	// This is only applicable when the execution mode is set to 'agent'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.97.0/docs/resources/waypoint_add_on_definition#terraform_agent_pool_id WaypointAddOnDefinition#terraform_agent_pool_id}
	TerraformAgentPoolId *string `field:"optional" json:"terraformAgentPoolId" yaml:"terraformAgentPoolId"`
	// Terraform Cloud Workspace details. If not provided, defaults to the HCP Terraform project of the associated application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.97.0/docs/resources/waypoint_add_on_definition#terraform_cloud_workspace_details WaypointAddOnDefinition#terraform_cloud_workspace_details}
	TerraformCloudWorkspaceDetails *WaypointAddOnDefinitionTerraformCloudWorkspaceDetails `field:"optional" json:"terraformCloudWorkspaceDetails" yaml:"terraformCloudWorkspaceDetails"`
	// The execution mode of the HCP Terraform workspaces for add-ons using this add-on definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.97.0/docs/resources/waypoint_add_on_definition#terraform_execution_mode WaypointAddOnDefinition#terraform_execution_mode}
	TerraformExecutionMode *string `field:"optional" json:"terraformExecutionMode" yaml:"terraformExecutionMode"`
	// List of variable options for the Add-on Definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.97.0/docs/resources/waypoint_add_on_definition#variable_options WaypointAddOnDefinition#variable_options}
	VariableOptions interface{} `field:"optional" json:"variableOptions" yaml:"variableOptions"`
}

