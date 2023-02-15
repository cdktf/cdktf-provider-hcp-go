package hvn


type HvnTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/hvn#create Hvn#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/hvn#default Hvn#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/hvn#delete Hvn#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

