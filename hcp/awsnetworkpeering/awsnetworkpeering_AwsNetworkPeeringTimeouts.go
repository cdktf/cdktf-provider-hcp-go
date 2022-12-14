package awsnetworkpeering


type AwsNetworkPeeringTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/aws_network_peering#create AwsNetworkPeering#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/aws_network_peering#default AwsNetworkPeering#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/aws_network_peering#delete AwsNetworkPeering#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

