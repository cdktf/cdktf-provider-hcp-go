// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegrationtwilio


type VaultSecretsIntegrationTwilioStaticCredentialDetails struct {
	// Account SID for the target Twilio account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.105.0/docs/resources/vault_secrets_integration_twilio#account_sid VaultSecretsIntegrationTwilio#account_sid}
	AccountSid *string `field:"required" json:"accountSid" yaml:"accountSid"`
	// Api key secret used with the api key SID to authenticate against the target Twilio account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.105.0/docs/resources/vault_secrets_integration_twilio#api_key_secret VaultSecretsIntegrationTwilio#api_key_secret}
	ApiKeySecret *string `field:"required" json:"apiKeySecret" yaml:"apiKeySecret"`
	// Api key SID to authenticate against the target Twilio account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.105.0/docs/resources/vault_secrets_integration_twilio#api_key_sid VaultSecretsIntegrationTwilio#api_key_sid}
	ApiKeySid *string `field:"required" json:"apiKeySid" yaml:"apiKeySid"`
}

