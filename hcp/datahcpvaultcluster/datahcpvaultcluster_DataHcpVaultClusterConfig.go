package datahcpvaultcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataHcpVaultClusterConfig struct {
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
	// The ID of the HCP Vault cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/vault_cluster#cluster_id DataHcpVaultCluster#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// audit_log_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/vault_cluster#audit_log_config DataHcpVaultCluster#audit_log_config}
	AuditLogConfig interface{} `field:"optional" json:"auditLogConfig" yaml:"auditLogConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/vault_cluster#id DataHcpVaultCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// metrics_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/vault_cluster#metrics_config DataHcpVaultCluster#metrics_config}
	MetricsConfig interface{} `field:"optional" json:"metricsConfig" yaml:"metricsConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/d/vault_cluster#timeouts DataHcpVaultCluster#timeouts}
	Timeouts *DataHcpVaultClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

