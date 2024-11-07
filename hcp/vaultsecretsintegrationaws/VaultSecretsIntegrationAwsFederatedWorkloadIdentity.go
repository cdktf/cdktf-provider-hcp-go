// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegrationaws


type VaultSecretsIntegrationAwsFederatedWorkloadIdentity struct {
	// Audience configured on the AWS IAM identity provider to federate access with HCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.98.1/docs/resources/vault_secrets_integration_aws#audience VaultSecretsIntegrationAws#audience}
	Audience *string `field:"required" json:"audience" yaml:"audience"`
	// AWS IAM role ARN the integration will assume to carry operations for the appropriate capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.98.1/docs/resources/vault_secrets_integration_aws#role_arn VaultSecretsIntegrationAws#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

