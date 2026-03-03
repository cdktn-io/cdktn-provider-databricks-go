// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package modelserving

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ModelServingRateLimitsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_ModelServingRateLimitsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ModelServingRateLimitsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ModelServingRateLimitsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ModelServingRateLimitsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ModelServingRateLimitsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ModelServingRateLimitsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewModelServingRateLimitsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

