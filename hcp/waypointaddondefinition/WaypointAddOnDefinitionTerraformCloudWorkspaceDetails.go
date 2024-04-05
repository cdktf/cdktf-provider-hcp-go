// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointaddondefinition


type WaypointAddOnDefinitionTerraformCloudWorkspaceDetails struct {
	// Name of the Terraform Cloud Workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.86.0/docs/resources/waypoint_add_on_definition#name WaypointAddOnDefinition#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Tetraform Cloud Project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.86.0/docs/resources/waypoint_add_on_definition#terraform_project_id WaypointAddOnDefinition#terraform_project_id}
	TerraformProjectId *string `field:"required" json:"terraformProjectId" yaml:"terraformProjectId"`
}

