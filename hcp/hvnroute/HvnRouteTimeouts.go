package hvnroute


type HvnRouteTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/hvn_route#create HvnRoute#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/hvn_route#default HvnRoute#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/hvn_route#delete HvnRoute#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

