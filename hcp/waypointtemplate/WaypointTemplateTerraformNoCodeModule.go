// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointtemplate


type WaypointTemplateTerraformNoCodeModule struct {
	// No-Code Module Source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.0/docs/resources/waypoint_template#source WaypointTemplate#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// No-Code Module Version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.94.0/docs/resources/waypoint_template#version WaypointTemplate#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}

