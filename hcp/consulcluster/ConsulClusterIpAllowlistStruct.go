package consulcluster


type ConsulClusterIpAllowlistStruct struct {
	// IP address range in CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.61.0/docs/resources/consul_cluster#address ConsulCluster#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// Description to help identify source (maximum 255 chars).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.61.0/docs/resources/consul_cluster#description ConsulCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

