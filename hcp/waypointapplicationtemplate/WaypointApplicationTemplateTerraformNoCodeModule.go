// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointapplicationtemplate


type WaypointApplicationTemplateTerraformNoCodeModule struct {
	// No-Code Module Source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.91.0/docs/resources/waypoint_application_template#source WaypointApplicationTemplate#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// No-Code Module Version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.91.0/docs/resources/waypoint_application_template#version WaypointApplicationTemplate#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}

