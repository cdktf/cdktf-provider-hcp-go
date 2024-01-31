// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package notificationswebhook

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NotificationsWebhookSubscriptionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NotificationsWebhookSubscriptionsList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NotificationsWebhookSubscriptionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NotificationsWebhookSubscriptionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NotificationsWebhookSubscriptionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NotificationsWebhookSubscriptionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NotificationsWebhookSubscriptionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNotificationsWebhookSubscriptionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

