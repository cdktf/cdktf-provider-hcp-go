// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointapplicationtemplate


type WaypointApplicationTemplateVariableOptions struct {
	// Variable name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.92.0/docs/resources/waypoint_application_template#name WaypointApplicationTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Variable type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.92.0/docs/resources/waypoint_application_template#variable_type WaypointApplicationTemplate#variable_type}
	VariableType *string `field:"required" json:"variableType" yaml:"variableType"`
	// List of options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.92.0/docs/resources/waypoint_application_template#options WaypointApplicationTemplate#options}
	Options *[]*string `field:"optional" json:"options" yaml:"options"`
	// Whether the variable is editable by the user creating an application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.92.0/docs/resources/waypoint_application_template#user_editable WaypointApplicationTemplate#user_editable}
	UserEditable interface{} `field:"optional" json:"userEditable" yaml:"userEditable"`
}

