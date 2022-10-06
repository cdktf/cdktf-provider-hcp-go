//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HcpProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (h *jsiiProxy_HcpProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateHcpProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewHcpProviderParameters(scope constructs.Construct, id *string, config *HcpProviderConfig) error {
	return nil
}

