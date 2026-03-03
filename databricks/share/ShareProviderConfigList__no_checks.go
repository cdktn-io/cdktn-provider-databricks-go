// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package share

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ShareProviderConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ShareProviderConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ShareProviderConfigList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ShareProviderConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ShareProviderConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ShareProviderConfigList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ShareProviderConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewShareProviderConfigListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

