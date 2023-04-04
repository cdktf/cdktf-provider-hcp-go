//go:build no_runtime_type_checking

package consulcluster

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConsulClusterIpAllowlistList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConsulClusterIpAllowlistList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConsulClusterIpAllowlistList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConsulClusterIpAllowlistList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConsulClusterIpAllowlistList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConsulClusterIpAllowlistList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConsulClusterIpAllowlistListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

