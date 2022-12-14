package consulsnapshot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConsulSnapshotConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the HCP Consul cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/consul_snapshot#cluster_id ConsulSnapshot#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The name of the snapshot.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/consul_snapshot#snapshot_name ConsulSnapshot#snapshot_name}
	SnapshotName *string `field:"required" json:"snapshotName" yaml:"snapshotName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/consul_snapshot#id ConsulSnapshot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/consul_snapshot#timeouts ConsulSnapshot#timeouts}
	Timeouts *ConsulSnapshotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

