// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointaddondefinition


type WaypointAddOnDefinitionTerraformNoCodeModule struct {
	// Terraform Cloud no-code Module Source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.91.1/docs/resources/waypoint_add_on_definition#source WaypointAddOnDefinition#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// Terraform Cloud no-code Module Version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.91.1/docs/resources/waypoint_add_on_definition#version WaypointAddOnDefinition#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}

