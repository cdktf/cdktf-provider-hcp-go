// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointapplicationtemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WaypointApplicationTemplateConfig struct {
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
	// The name of the Application Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.92.0/docs/resources/waypoint_application_template#name WaypointApplicationTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A brief description of the template, up to 110 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.92.0/docs/resources/waypoint_application_template#summary WaypointApplicationTemplate#summary}
	Summary *string `field:"required" json:"summary" yaml:"summary"`
	// Terraform Cloud Workspace details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.92.0/docs/resources/waypoint_application_template#terraform_cloud_workspace_details WaypointApplicationTemplate#terraform_cloud_workspace_details}
	TerraformCloudWorkspaceDetails *WaypointApplicationTemplateTerraformCloudWorkspaceDetails `field:"required" json:"terraformCloudWorkspaceDetails" yaml:"terraformCloudWorkspaceDetails"`
	// Terraform Cloud No-Code Module details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.92.0/docs/resources/waypoint_application_template#terraform_no_code_module WaypointApplicationTemplate#terraform_no_code_module}
	TerraformNoCodeModule *WaypointApplicationTemplateTerraformNoCodeModule `field:"required" json:"terraformNoCodeModule" yaml:"terraformNoCodeModule"`
	// A description of the template, along with when and why it should be used, up to 500 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.92.0/docs/resources/waypoint_application_template#description WaypointApplicationTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of labels attached to this Application Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.92.0/docs/resources/waypoint_application_template#labels WaypointApplicationTemplate#labels}
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// The ID of the HCP project where the Waypoint Application Template is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.92.0/docs/resources/waypoint_application_template#project_id WaypointApplicationTemplate#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Instructions for using the template (markdown format supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.92.0/docs/resources/waypoint_application_template#readme_markdown_template WaypointApplicationTemplate#readme_markdown_template}
	ReadmeMarkdownTemplate *string `field:"optional" json:"readmeMarkdownTemplate" yaml:"readmeMarkdownTemplate"`
	// List of variable options for the template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.92.0/docs/resources/waypoint_application_template#variable_options WaypointApplicationTemplate#variable_options}
	VariableOptions interface{} `field:"optional" json:"variableOptions" yaml:"variableOptions"`
}

