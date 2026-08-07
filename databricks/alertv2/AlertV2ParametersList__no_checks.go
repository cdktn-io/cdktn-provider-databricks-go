// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package alertv2

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertV2ParametersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AlertV2ParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlertV2ParametersList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlertV2ParametersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlertV2ParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlertV2ParametersList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlertV2ParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlertV2ParametersListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

