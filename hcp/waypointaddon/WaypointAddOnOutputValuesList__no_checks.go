// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package waypointaddon

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WaypointAddOnOutputValuesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WaypointAddOnOutputValuesList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WaypointAddOnOutputValuesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WaypointAddOnOutputValuesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WaypointAddOnOutputValuesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WaypointAddOnOutputValuesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWaypointAddOnOutputValuesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

