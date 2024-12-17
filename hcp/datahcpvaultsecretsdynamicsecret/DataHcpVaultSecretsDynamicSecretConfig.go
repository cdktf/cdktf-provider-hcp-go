// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datahcpvaultsecretsdynamicsecret

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataHcpVaultSecretsDynamicSecretConfig struct {
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
	// The name of the Vault Secrets application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.101.0/docs/data-sources/vault_secrets_dynamic_secret#app_name DataHcpVaultSecretsDynamicSecret#app_name}
	AppName *string `field:"required" json:"appName" yaml:"appName"`
	// The name of the Vault Secrets secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.101.0/docs/data-sources/vault_secrets_dynamic_secret#secret_name DataHcpVaultSecretsDynamicSecret#secret_name}
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

