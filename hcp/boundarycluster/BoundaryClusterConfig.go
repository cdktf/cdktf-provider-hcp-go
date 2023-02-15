package boundarycluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BoundaryClusterConfig struct {
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
	// The ID of the Boundary cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/boundary_cluster#cluster_id BoundaryCluster#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The password of the initial admin user.
	//
	// This must be at least 8 characters in length. Note that this may show up in logs, and it will be stored in the state file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/boundary_cluster#password BoundaryCluster#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The username of the initial admin user.
	//
	// This must be at least 3 characters in length, alphanumeric, hyphen, or period.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/boundary_cluster#username BoundaryCluster#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/boundary_cluster#id BoundaryCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/boundary_cluster#timeouts BoundaryCluster#timeouts}
	Timeouts *BoundaryClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

