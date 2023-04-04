package consulcluster


type ConsulClusterIpAllowlist struct {
	// IP address range in CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/consul_cluster#address ConsulCluster#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// Description to help identify source (maximum 255 chars).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/consul_cluster#description ConsulCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

