// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointapplicationtemplate


type WaypointApplicationTemplateTerraformCloudWorkspaceDetails struct {
	// Name of the Terraform Cloud Workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.91.1/docs/resources/waypoint_application_template#name WaypointApplicationTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Terraform Cloud Project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.91.1/docs/resources/waypoint_application_template#terraform_project_id WaypointApplicationTemplate#terraform_project_id}
	TerraformProjectId *string `field:"required" json:"terraformProjectId" yaml:"terraformProjectId"`
}

