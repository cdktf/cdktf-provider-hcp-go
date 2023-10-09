// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultcluster


type VaultClusterMetricsConfig struct {
	// CloudWatch access key ID for streaming metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.2/docs/resources/vault_cluster#cloudwatch_access_key_id VaultCluster#cloudwatch_access_key_id}
	CloudwatchAccessKeyId *string `field:"optional" json:"cloudwatchAccessKeyId" yaml:"cloudwatchAccessKeyId"`
	// CloudWatch region for streaming metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.2/docs/resources/vault_cluster#cloudwatch_region VaultCluster#cloudwatch_region}
	CloudwatchRegion *string `field:"optional" json:"cloudwatchRegion" yaml:"cloudwatchRegion"`
	// CloudWatch secret access key for streaming metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.2/docs/resources/vault_cluster#cloudwatch_secret_access_key VaultCluster#cloudwatch_secret_access_key}
	CloudwatchSecretAccessKey *string `field:"optional" json:"cloudwatchSecretAccessKey" yaml:"cloudwatchSecretAccessKey"`
	// Datadog api key for streaming metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.2/docs/resources/vault_cluster#datadog_api_key VaultCluster#datadog_api_key}
	DatadogApiKey *string `field:"optional" json:"datadogApiKey" yaml:"datadogApiKey"`
	// Datadog region for streaming metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.2/docs/resources/vault_cluster#datadog_region VaultCluster#datadog_region}
	DatadogRegion *string `field:"optional" json:"datadogRegion" yaml:"datadogRegion"`
	// ElasticSearch endpoint for streaming metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.2/docs/resources/vault_cluster#elasticsearch_endpoint VaultCluster#elasticsearch_endpoint}
	ElasticsearchEndpoint *string `field:"optional" json:"elasticsearchEndpoint" yaml:"elasticsearchEndpoint"`
	// ElasticSearch password for streaming metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.2/docs/resources/vault_cluster#elasticsearch_password VaultCluster#elasticsearch_password}
	ElasticsearchPassword *string `field:"optional" json:"elasticsearchPassword" yaml:"elasticsearchPassword"`
	// ElasticSearch user for streaming metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.2/docs/resources/vault_cluster#elasticsearch_user VaultCluster#elasticsearch_user}
	ElasticsearchUser *string `field:"optional" json:"elasticsearchUser" yaml:"elasticsearchUser"`
	// Grafana endpoint for streaming metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.2/docs/resources/vault_cluster#grafana_endpoint VaultCluster#grafana_endpoint}
	GrafanaEndpoint *string `field:"optional" json:"grafanaEndpoint" yaml:"grafanaEndpoint"`
	// Grafana password for streaming metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.2/docs/resources/vault_cluster#grafana_password VaultCluster#grafana_password}
	GrafanaPassword *string `field:"optional" json:"grafanaPassword" yaml:"grafanaPassword"`
	// Grafana user for streaming metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.2/docs/resources/vault_cluster#grafana_user VaultCluster#grafana_user}
	GrafanaUser *string `field:"optional" json:"grafanaUser" yaml:"grafanaUser"`
	// Splunk endpoint for streaming metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.2/docs/resources/vault_cluster#splunk_hecendpoint VaultCluster#splunk_hecendpoint}
	SplunkHecendpoint *string `field:"optional" json:"splunkHecendpoint" yaml:"splunkHecendpoint"`
	// Splunk token for streaming metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.72.2/docs/resources/vault_cluster#splunk_token VaultCluster#splunk_token}
	SplunkToken *string `field:"optional" json:"splunkToken" yaml:"splunkToken"`
}

