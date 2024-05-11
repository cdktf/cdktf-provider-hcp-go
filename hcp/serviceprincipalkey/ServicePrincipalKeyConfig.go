// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipalkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicePrincipalKeyConfig struct {
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
	// The service principal's resource name for which a key should be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.89.0/docs/resources/service_principal_key#service_principal ServicePrincipalKey#service_principal}
	ServicePrincipal *string `field:"required" json:"servicePrincipal" yaml:"servicePrincipal"`
	// A map of arbitrary string key/value pairs that will force recreation of the key when they change, enabling key based on external conditions such as a rotating timestamp.
	//
	// Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.89.0/docs/resources/service_principal_key#rotate_triggers ServicePrincipalKey#rotate_triggers}
	RotateTriggers *map[string]*string `field:"optional" json:"rotateTriggers" yaml:"rotateTriggers"`
}

