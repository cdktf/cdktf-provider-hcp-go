// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type HcpProviderWorkloadIdentity struct {
	// The resource_name of the Workload Identity Provider to exchange the token with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.86.0/docs#resource_name HcpProvider#resource_name}
	ResourceName *string `field:"required" json:"resourceName" yaml:"resourceName"`
	// The path to a file containing a JWT token retrieved from an OpenID Connect (OIDC) or OAuth2 provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.86.0/docs#token_file HcpProvider#token_file}
	TokenFile *string `field:"required" json:"tokenFile" yaml:"tokenFile"`
}

