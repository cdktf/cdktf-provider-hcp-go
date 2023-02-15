package vaultcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-hcp-go/hcp/v3/jsii"

	"github.com/cdktf/cdktf-provider-hcp-go/hcp/v3/vaultcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VaultClusterAuditLogConfigOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatadogApiKey() *string
	SetDatadogApiKey(val *string)
	DatadogApiKeyInput() *string
	DatadogRegion() *string
	SetDatadogRegion(val *string)
	DatadogRegionInput() *string
	// Experimental.
	Fqn() *string
	GrafanaEndpoint() *string
	SetGrafanaEndpoint(val *string)
	GrafanaEndpointInput() *string
	GrafanaPassword() *string
	SetGrafanaPassword(val *string)
	GrafanaPasswordInput() *string
	GrafanaUser() *string
	SetGrafanaUser(val *string)
	GrafanaUserInput() *string
	InternalValue() *VaultClusterAuditLogConfig
	SetInternalValue(val *VaultClusterAuditLogConfig)
	SplunkHecendpoint() *string
	SetSplunkHecendpoint(val *string)
	SplunkHecendpointInput() *string
	SplunkToken() *string
	SetSplunkToken(val *string)
	SplunkTokenInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDatadogApiKey()
	ResetDatadogRegion()
	ResetGrafanaEndpoint()
	ResetGrafanaPassword()
	ResetGrafanaUser()
	ResetSplunkHecendpoint()
	ResetSplunkToken()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VaultClusterAuditLogConfigOutputReference
type jsiiProxy_VaultClusterAuditLogConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) DatadogApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) DatadogApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) DatadogRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) DatadogRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GrafanaEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GrafanaEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GrafanaPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GrafanaPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GrafanaUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GrafanaUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) InternalValue() *VaultClusterAuditLogConfig {
	var returns *VaultClusterAuditLogConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) SplunkHecendpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splunkHecendpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) SplunkHecendpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splunkHecendpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) SplunkToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splunkToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) SplunkTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splunkTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVaultClusterAuditLogConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VaultClusterAuditLogConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVaultClusterAuditLogConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VaultClusterAuditLogConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-hcp.vaultCluster.VaultClusterAuditLogConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVaultClusterAuditLogConfigOutputReference_Override(v VaultClusterAuditLogConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-hcp.vaultCluster.VaultClusterAuditLogConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetDatadogApiKey(val *string) {
	if err := j.validateSetDatadogApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datadogApiKey",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetDatadogRegion(val *string) {
	if err := j.validateSetDatadogRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datadogRegion",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetGrafanaEndpoint(val *string) {
	if err := j.validateSetGrafanaEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grafanaEndpoint",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetGrafanaPassword(val *string) {
	if err := j.validateSetGrafanaPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grafanaPassword",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetGrafanaUser(val *string) {
	if err := j.validateSetGrafanaUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grafanaUser",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetInternalValue(val *VaultClusterAuditLogConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetSplunkHecendpoint(val *string) {
	if err := j.validateSetSplunkHecendpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"splunkHecendpoint",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetSplunkToken(val *string) {
	if err := j.validateSetSplunkTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"splunkToken",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VaultClusterAuditLogConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetDatadogApiKey() {
	_jsii_.InvokeVoid(
		v,
		"resetDatadogApiKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetDatadogRegion() {
	_jsii_.InvokeVoid(
		v,
		"resetDatadogRegion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetGrafanaEndpoint() {
	_jsii_.InvokeVoid(
		v,
		"resetGrafanaEndpoint",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetGrafanaPassword() {
	_jsii_.InvokeVoid(
		v,
		"resetGrafanaPassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetGrafanaUser() {
	_jsii_.InvokeVoid(
		v,
		"resetGrafanaUser",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetSplunkHecendpoint() {
	_jsii_.InvokeVoid(
		v,
		"resetSplunkHecendpoint",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ResetSplunkToken() {
	_jsii_.InvokeVoid(
		v,
		"resetSplunkToken",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultClusterAuditLogConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

