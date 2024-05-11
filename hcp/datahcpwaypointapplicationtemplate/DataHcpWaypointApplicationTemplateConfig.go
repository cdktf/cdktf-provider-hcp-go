// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datahcpwaypointapplicationtemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataHcpWaypointApplicationTemplateConfig struct {
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
	// The ID of the Application Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.89.0/docs/data-sources/waypoint_application_template#id DataHcpWaypointApplicationTemplate#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the Application Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.89.0/docs/data-sources/waypoint_application_template#name DataHcpWaypointApplicationTemplate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of the HCP project where the Waypoint Application Template is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.89.0/docs/data-sources/waypoint_application_template#project_id DataHcpWaypointApplicationTemplate#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

