package vaultcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VaultClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#cluster_id VaultCluster#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The ID of the HVN this HCP Vault cluster is associated to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#hvn_id VaultCluster#hvn_id}
	HvnId *string `field:"required" json:"hvnId" yaml:"hvnId"`
	// audit_log_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#audit_log_config VaultCluster#audit_log_config}
	AuditLogConfig *VaultClusterAuditLogConfig `field:"optional" json:"auditLogConfig" yaml:"auditLogConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#id VaultCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// major_version_upgrade_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#major_version_upgrade_config VaultCluster#major_version_upgrade_config}
	MajorVersionUpgradeConfig *VaultClusterMajorVersionUpgradeConfig `field:"optional" json:"majorVersionUpgradeConfig" yaml:"majorVersionUpgradeConfig"`
	// metrics_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#metrics_config VaultCluster#metrics_config}
	MetricsConfig *VaultClusterMetricsConfig `field:"optional" json:"metricsConfig" yaml:"metricsConfig"`
	// The minimum Vault version to use when creating the cluster.
	//
	// If not specified, it is defaulted to the version that is currently recommended by HCP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#min_vault_version VaultCluster#min_vault_version}
	MinVaultVersion *string `field:"optional" json:"minVaultVersion" yaml:"minVaultVersion"`
	// The performance replication [paths filter](https://developer.hashicorp.com/vault/tutorials/cloud-ops/vault-replication-terraform). Applies to performance replication secondaries only and operates in "deny" mode only.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#paths_filter VaultCluster#paths_filter}
	PathsFilter *[]*string `field:"optional" json:"pathsFilter" yaml:"pathsFilter"`
	// The `self_link` of the HCP Vault Plus tier cluster which is the primary in the performance replication setup with this HCP Vault Plus tier cluster.
	//
	// If not specified, it is a standalone Plus tier HCP Vault cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#primary_link VaultCluster#primary_link}
	PrimaryLink *string `field:"optional" json:"primaryLink" yaml:"primaryLink"`
	// Denotes that the cluster has a public endpoint. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#public_endpoint VaultCluster#public_endpoint}
	PublicEndpoint interface{} `field:"optional" json:"publicEndpoint" yaml:"publicEndpoint"`
	// Tier of the HCP Vault cluster.
	//
	// Valid options for tiers - `dev`, `starter_small`, `standard_small`, `standard_medium`, `standard_large`, `plus_small`, `plus_medium`, `plus_large`. See [pricing information](https://cloud.hashicorp.com/pricing/vault). Changing a cluster's size or tier is only available to admins. See [Scale a cluster](https://registry.terraform.io/providers/hashicorp/hcp/latest/docs/guides/vault-scaling).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#tier VaultCluster#tier}
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#timeouts VaultCluster#timeouts}
	Timeouts *VaultClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

