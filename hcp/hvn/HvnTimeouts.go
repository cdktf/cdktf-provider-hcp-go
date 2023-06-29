package hvn


type HvnTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.63.0/docs/resources/hvn#create Hvn#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.63.0/docs/resources/hvn#default Hvn#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.63.0/docs/resources/hvn#delete Hvn#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

