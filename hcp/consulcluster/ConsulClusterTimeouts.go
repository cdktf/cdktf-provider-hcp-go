package consulcluster


type ConsulClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/consul_cluster#create ConsulCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/consul_cluster#default ConsulCluster#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/consul_cluster#delete ConsulCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/consul_cluster#update ConsulCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

