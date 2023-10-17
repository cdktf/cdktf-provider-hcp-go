// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultcluster

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-hcp.vaultCluster.VaultCluster",
		reflect.TypeOf((*VaultCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "auditLogConfig", GoGetter: "AuditLogConfig"},
			_jsii_.MemberProperty{JsiiProperty: "auditLogConfigInput", GoGetter: "AuditLogConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cloudProvider", GoGetter: "CloudProvider"},
			_jsii_.MemberProperty{JsiiProperty: "clusterId", GoGetter: "ClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIdInput", GoGetter: "ClusterIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "hvnId", GoGetter: "HvnId"},
			_jsii_.MemberProperty{JsiiProperty: "hvnIdInput", GoGetter: "HvnIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "majorVersionUpgradeConfig", GoGetter: "MajorVersionUpgradeConfig"},
			_jsii_.MemberProperty{JsiiProperty: "majorVersionUpgradeConfigInput", GoGetter: "MajorVersionUpgradeConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "metricsConfig", GoGetter: "MetricsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "metricsConfigInput", GoGetter: "MetricsConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "minVaultVersion", GoGetter: "MinVaultVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minVaultVersionInput", GoGetter: "MinVaultVersionInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "organizationId", GoGetter: "OrganizationId"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pathsFilter", GoGetter: "PathsFilter"},
			_jsii_.MemberProperty{JsiiProperty: "pathsFilterInput", GoGetter: "PathsFilterInput"},
			_jsii_.MemberProperty{JsiiProperty: "primaryLink", GoGetter: "PrimaryLink"},
			_jsii_.MemberProperty{JsiiProperty: "primaryLinkInput", GoGetter: "PrimaryLinkInput"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "projectIdInput", GoGetter: "ProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "proxyEndpoint", GoGetter: "ProxyEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "proxyEndpointInput", GoGetter: "ProxyEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "publicEndpoint", GoGetter: "PublicEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "publicEndpointInput", GoGetter: "PublicEndpointInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAuditLogConfig", GoMethod: "PutAuditLogConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putMajorVersionUpgradeConfig", GoMethod: "PutMajorVersionUpgradeConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putMetricsConfig", GoMethod: "PutMetricsConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditLogConfig", GoMethod: "ResetAuditLogConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMajorVersionUpgradeConfig", GoMethod: "ResetMajorVersionUpgradeConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetricsConfig", GoMethod: "ResetMetricsConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinVaultVersion", GoMethod: "ResetMinVaultVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPathsFilter", GoMethod: "ResetPathsFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimaryLink", GoMethod: "ResetPrimaryLink"},
			_jsii_.MemberMethod{JsiiMethod: "resetProjectId", GoMethod: "ResetProjectId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProxyEndpoint", GoMethod: "ResetProxyEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetPublicEndpoint", GoMethod: "ResetPublicEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetTier", GoMethod: "ResetTier"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "selfLink", GoGetter: "SelfLink"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "tier", GoGetter: "Tier"},
			_jsii_.MemberProperty{JsiiProperty: "tierInput", GoGetter: "TierInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "vaultPrivateEndpointUrl", GoGetter: "VaultPrivateEndpointUrl"},
			_jsii_.MemberProperty{JsiiProperty: "vaultProxyEndpointUrl", GoGetter: "VaultProxyEndpointUrl"},
			_jsii_.MemberProperty{JsiiProperty: "vaultPublicEndpointUrl", GoGetter: "VaultPublicEndpointUrl"},
			_jsii_.MemberProperty{JsiiProperty: "vaultVersion", GoGetter: "VaultVersion"},
		},
		func() interface{} {
			j := jsiiProxy_VaultCluster{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-hcp.vaultCluster.VaultClusterAuditLogConfig",
		reflect.TypeOf((*VaultClusterAuditLogConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-hcp.vaultCluster.VaultClusterAuditLogConfigOutputReference",
		reflect.TypeOf((*VaultClusterAuditLogConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchAccessKeyId", GoGetter: "CloudwatchAccessKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchAccessKeyIdInput", GoGetter: "CloudwatchAccessKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchGroupName", GoGetter: "CloudwatchGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchRegion", GoGetter: "CloudwatchRegion"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchRegionInput", GoGetter: "CloudwatchRegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchSecretAccessKey", GoGetter: "CloudwatchSecretAccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchSecretAccessKeyInput", GoGetter: "CloudwatchSecretAccessKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchStreamName", GoGetter: "CloudwatchStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "datadogApiKey", GoGetter: "DatadogApiKey"},
			_jsii_.MemberProperty{JsiiProperty: "datadogApiKeyInput", GoGetter: "DatadogApiKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "datadogRegion", GoGetter: "DatadogRegion"},
			_jsii_.MemberProperty{JsiiProperty: "datadogRegionInput", GoGetter: "DatadogRegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchDataset", GoGetter: "ElasticsearchDataset"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchEndpoint", GoGetter: "ElasticsearchEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchEndpointInput", GoGetter: "ElasticsearchEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchPassword", GoGetter: "ElasticsearchPassword"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchPasswordInput", GoGetter: "ElasticsearchPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchUser", GoGetter: "ElasticsearchUser"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchUserInput", GoGetter: "ElasticsearchUserInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grafanaEndpoint", GoGetter: "GrafanaEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "grafanaEndpointInput", GoGetter: "GrafanaEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "grafanaPassword", GoGetter: "GrafanaPassword"},
			_jsii_.MemberProperty{JsiiProperty: "grafanaPasswordInput", GoGetter: "GrafanaPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "grafanaUser", GoGetter: "GrafanaUser"},
			_jsii_.MemberProperty{JsiiProperty: "grafanaUserInput", GoGetter: "GrafanaUserInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchAccessKeyId", GoMethod: "ResetCloudwatchAccessKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchRegion", GoMethod: "ResetCloudwatchRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchSecretAccessKey", GoMethod: "ResetCloudwatchSecretAccessKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatadogApiKey", GoMethod: "ResetDatadogApiKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatadogRegion", GoMethod: "ResetDatadogRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetElasticsearchEndpoint", GoMethod: "ResetElasticsearchEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetElasticsearchPassword", GoMethod: "ResetElasticsearchPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetElasticsearchUser", GoMethod: "ResetElasticsearchUser"},
			_jsii_.MemberMethod{JsiiMethod: "resetGrafanaEndpoint", GoMethod: "ResetGrafanaEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetGrafanaPassword", GoMethod: "ResetGrafanaPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetGrafanaUser", GoMethod: "ResetGrafanaUser"},
			_jsii_.MemberMethod{JsiiMethod: "resetSplunkHecendpoint", GoMethod: "ResetSplunkHecendpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetSplunkToken", GoMethod: "ResetSplunkToken"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "splunkHecendpoint", GoGetter: "SplunkHecendpoint"},
			_jsii_.MemberProperty{JsiiProperty: "splunkHecendpointInput", GoGetter: "SplunkHecendpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "splunkToken", GoGetter: "SplunkToken"},
			_jsii_.MemberProperty{JsiiProperty: "splunkTokenInput", GoGetter: "SplunkTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VaultClusterAuditLogConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-hcp.vaultCluster.VaultClusterConfig",
		reflect.TypeOf((*VaultClusterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-hcp.vaultCluster.VaultClusterMajorVersionUpgradeConfig",
		reflect.TypeOf((*VaultClusterMajorVersionUpgradeConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-hcp.vaultCluster.VaultClusterMajorVersionUpgradeConfigOutputReference",
		reflect.TypeOf((*VaultClusterMajorVersionUpgradeConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindowDay", GoGetter: "MaintenanceWindowDay"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindowDayInput", GoGetter: "MaintenanceWindowDayInput"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindowTime", GoGetter: "MaintenanceWindowTime"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindowTimeInput", GoGetter: "MaintenanceWindowTimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaintenanceWindowDay", GoMethod: "ResetMaintenanceWindowDay"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaintenanceWindowTime", GoMethod: "ResetMaintenanceWindowTime"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "upgradeType", GoGetter: "UpgradeType"},
			_jsii_.MemberProperty{JsiiProperty: "upgradeTypeInput", GoGetter: "UpgradeTypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_VaultClusterMajorVersionUpgradeConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-hcp.vaultCluster.VaultClusterMetricsConfig",
		reflect.TypeOf((*VaultClusterMetricsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-hcp.vaultCluster.VaultClusterMetricsConfigOutputReference",
		reflect.TypeOf((*VaultClusterMetricsConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchAccessKeyId", GoGetter: "CloudwatchAccessKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchAccessKeyIdInput", GoGetter: "CloudwatchAccessKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchNamespace", GoGetter: "CloudwatchNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchRegion", GoGetter: "CloudwatchRegion"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchRegionInput", GoGetter: "CloudwatchRegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchSecretAccessKey", GoGetter: "CloudwatchSecretAccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchSecretAccessKeyInput", GoGetter: "CloudwatchSecretAccessKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "datadogApiKey", GoGetter: "DatadogApiKey"},
			_jsii_.MemberProperty{JsiiProperty: "datadogApiKeyInput", GoGetter: "DatadogApiKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "datadogRegion", GoGetter: "DatadogRegion"},
			_jsii_.MemberProperty{JsiiProperty: "datadogRegionInput", GoGetter: "DatadogRegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchDataset", GoGetter: "ElasticsearchDataset"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchEndpoint", GoGetter: "ElasticsearchEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchEndpointInput", GoGetter: "ElasticsearchEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchPassword", GoGetter: "ElasticsearchPassword"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchPasswordInput", GoGetter: "ElasticsearchPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchUser", GoGetter: "ElasticsearchUser"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchUserInput", GoGetter: "ElasticsearchUserInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grafanaEndpoint", GoGetter: "GrafanaEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "grafanaEndpointInput", GoGetter: "GrafanaEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "grafanaPassword", GoGetter: "GrafanaPassword"},
			_jsii_.MemberProperty{JsiiProperty: "grafanaPasswordInput", GoGetter: "GrafanaPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "grafanaUser", GoGetter: "GrafanaUser"},
			_jsii_.MemberProperty{JsiiProperty: "grafanaUserInput", GoGetter: "GrafanaUserInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchAccessKeyId", GoMethod: "ResetCloudwatchAccessKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchRegion", GoMethod: "ResetCloudwatchRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchSecretAccessKey", GoMethod: "ResetCloudwatchSecretAccessKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatadogApiKey", GoMethod: "ResetDatadogApiKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatadogRegion", GoMethod: "ResetDatadogRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetElasticsearchEndpoint", GoMethod: "ResetElasticsearchEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetElasticsearchPassword", GoMethod: "ResetElasticsearchPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetElasticsearchUser", GoMethod: "ResetElasticsearchUser"},
			_jsii_.MemberMethod{JsiiMethod: "resetGrafanaEndpoint", GoMethod: "ResetGrafanaEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetGrafanaPassword", GoMethod: "ResetGrafanaPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetGrafanaUser", GoMethod: "ResetGrafanaUser"},
			_jsii_.MemberMethod{JsiiMethod: "resetSplunkHecendpoint", GoMethod: "ResetSplunkHecendpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetSplunkToken", GoMethod: "ResetSplunkToken"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "splunkHecendpoint", GoGetter: "SplunkHecendpoint"},
			_jsii_.MemberProperty{JsiiProperty: "splunkHecendpointInput", GoGetter: "SplunkHecendpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "splunkToken", GoGetter: "SplunkToken"},
			_jsii_.MemberProperty{JsiiProperty: "splunkTokenInput", GoGetter: "SplunkTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VaultClusterMetricsConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-hcp.vaultCluster.VaultClusterTimeouts",
		reflect.TypeOf((*VaultClusterTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-hcp.vaultCluster.VaultClusterTimeoutsOutputReference",
		reflect.TypeOf((*VaultClusterTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "default", GoGetter: "Default"},
			_jsii_.MemberProperty{JsiiProperty: "defaultInput", GoGetter: "DefaultInput"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefault", GoMethod: "ResetDefault"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_VaultClusterTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
