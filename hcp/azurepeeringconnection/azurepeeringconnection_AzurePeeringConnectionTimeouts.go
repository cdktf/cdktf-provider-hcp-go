package azurepeeringconnection


type AzurePeeringConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/azure_peering_connection#create AzurePeeringConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/azure_peering_connection#default AzurePeeringConnection#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/azure_peering_connection#delete AzurePeeringConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

