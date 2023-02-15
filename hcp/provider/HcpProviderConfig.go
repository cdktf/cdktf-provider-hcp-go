package provider


type HcpProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp#alias HcpProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The OAuth2 Client ID for API operations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp#client_id HcpProvider#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The OAuth2 Client Secret for API operations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp#client_secret HcpProvider#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

