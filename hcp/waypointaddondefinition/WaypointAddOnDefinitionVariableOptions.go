// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointaddondefinition


type WaypointAddOnDefinitionVariableOptions struct {
	// Variable name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.95.0/docs/resources/waypoint_add_on_definition#name WaypointAddOnDefinition#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.95.0/docs/resources/waypoint_add_on_definition#options WaypointAddOnDefinition#options}
	Options *[]*string `field:"required" json:"options" yaml:"options"`
	// Variable type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.95.0/docs/resources/waypoint_add_on_definition#variable_type WaypointAddOnDefinition#variable_type}
	VariableType *string `field:"required" json:"variableType" yaml:"variableType"`
	// Whether the variable is editable by the user creating an add-on.
	//
	// If options are provided, then the user may only use those options, regardless of this setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.95.0/docs/resources/waypoint_add_on_definition#user_editable WaypointAddOnDefinition#user_editable}
	UserEditable interface{} `field:"optional" json:"userEditable" yaml:"userEditable"`
}

