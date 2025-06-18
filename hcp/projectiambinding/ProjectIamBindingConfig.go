// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectiambinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectIamBindingConfig struct {
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
	// The principal to bind to the given role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.107.0/docs/resources/project_iam_binding#principal_id ProjectIamBinding#principal_id}
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// The role name to bind to the given principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.107.0/docs/resources/project_iam_binding#role ProjectIamBinding#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// The ID of the HCP project to apply the IAM Policy to.
	//
	// If unspecified, the project configured on the provider is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.107.0/docs/resources/project_iam_binding#project_id ProjectIamBinding#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

