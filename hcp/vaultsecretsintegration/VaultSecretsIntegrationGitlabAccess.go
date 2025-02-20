// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegration


type VaultSecretsIntegrationGitlabAccess struct {
	// Access token used to authenticate against the target GitLab account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.103.0/docs/resources/vault_secrets_integration#token VaultSecretsIntegration#token}
	Token *string `field:"required" json:"token" yaml:"token"`
}

