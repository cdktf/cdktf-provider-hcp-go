package vaultcluster


type VaultClusterAuditLogConfig struct {
	// Datadog api key for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#datadog_api_key VaultCluster#datadog_api_key}
	DatadogApiKey *string `field:"optional" json:"datadogApiKey" yaml:"datadogApiKey"`
	// Datadog region for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#datadog_region VaultCluster#datadog_region}
	DatadogRegion *string `field:"optional" json:"datadogRegion" yaml:"datadogRegion"`
	// Grafana endpoint for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#grafana_endpoint VaultCluster#grafana_endpoint}
	GrafanaEndpoint *string `field:"optional" json:"grafanaEndpoint" yaml:"grafanaEndpoint"`
	// Grafana password for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#grafana_password VaultCluster#grafana_password}
	GrafanaPassword *string `field:"optional" json:"grafanaPassword" yaml:"grafanaPassword"`
	// Grafana user for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#grafana_user VaultCluster#grafana_user}
	GrafanaUser *string `field:"optional" json:"grafanaUser" yaml:"grafanaUser"`
	// Splunk endpoint for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#splunk_hecendpoint VaultCluster#splunk_hecendpoint}
	SplunkHecendpoint *string `field:"optional" json:"splunkHecendpoint" yaml:"splunkHecendpoint"`
	// Splunk token for streaming audit logs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hcp/r/vault_cluster#splunk_token VaultCluster#splunk_token}
	SplunkToken *string `field:"optional" json:"splunkToken" yaml:"splunkToken"`
}

