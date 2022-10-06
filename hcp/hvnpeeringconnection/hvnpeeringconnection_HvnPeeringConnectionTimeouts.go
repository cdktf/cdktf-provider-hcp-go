package hvnpeeringconnection


type HvnPeeringConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/hvn_peering_connection#create HvnPeeringConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/hvn_peering_connection#default HvnPeeringConnection#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/hvn_peering_connection#delete HvnPeeringConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

