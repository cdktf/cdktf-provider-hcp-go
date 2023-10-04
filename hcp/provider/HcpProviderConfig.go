// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type HcpProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.1/docs#alias HcpProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The OAuth2 Client ID for API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.1/docs#client_id HcpProvider#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The OAuth2 Client Secret for API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.1/docs#client_secret HcpProvider#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The default project in which resources should be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.1/docs#project_id HcpProvider#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

