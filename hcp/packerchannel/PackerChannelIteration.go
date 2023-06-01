package packerchannel


type PackerChannelIteration struct {
	// The fingerprint of the iteration assigned to the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.58.0/docs/resources/packer_channel#fingerprint PackerChannel#fingerprint}
	Fingerprint *string `field:"optional" json:"fingerprint" yaml:"fingerprint"`
	// The ID of the iteration assigned to the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.58.0/docs/resources/packer_channel#id PackerChannel#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The incremental_version of the iteration assigned to the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.58.0/docs/resources/packer_channel#incremental_version PackerChannel#incremental_version}
	IncrementalVersion *float64 `field:"optional" json:"incrementalVersion" yaml:"incrementalVersion"`
}

