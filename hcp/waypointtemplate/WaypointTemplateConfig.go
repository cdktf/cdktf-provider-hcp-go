// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointtemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WaypointTemplateConfig struct {
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
	// The name of the Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.1/docs/resources/waypoint_template#name WaypointTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A brief description of the template, up to 110 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.1/docs/resources/waypoint_template#summary WaypointTemplate#summary}
	Summary *string `field:"required" json:"summary" yaml:"summary"`
	// Terraform Cloud Workspace details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.1/docs/resources/waypoint_template#terraform_cloud_workspace_details WaypointTemplate#terraform_cloud_workspace_details}
	TerraformCloudWorkspaceDetails *WaypointTemplateTerraformCloudWorkspaceDetails `field:"required" json:"terraformCloudWorkspaceDetails" yaml:"terraformCloudWorkspaceDetails"`
	// Terraform Cloud No-Code Module details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.1/docs/resources/waypoint_template#terraform_no_code_module WaypointTemplate#terraform_no_code_module}
	TerraformNoCodeModule *WaypointTemplateTerraformNoCodeModule `field:"required" json:"terraformNoCodeModule" yaml:"terraformNoCodeModule"`
	// A description of the template, along with when and why it should be used, up to 500 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.1/docs/resources/waypoint_template#description WaypointTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of labels attached to this Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.1/docs/resources/waypoint_template#labels WaypointTemplate#labels}
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// The ID of the HCP project where the Waypoint Template is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.1/docs/resources/waypoint_template#project_id WaypointTemplate#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Instructions for using the template (markdown format supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.1/docs/resources/waypoint_template#readme_markdown_template WaypointTemplate#readme_markdown_template}
	ReadmeMarkdownTemplate *string `field:"optional" json:"readmeMarkdownTemplate" yaml:"readmeMarkdownTemplate"`
	// List of variable options for the template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.1/docs/resources/waypoint_template#variable_options WaypointTemplate#variable_options}
	VariableOptions interface{} `field:"optional" json:"variableOptions" yaml:"variableOptions"`
}

