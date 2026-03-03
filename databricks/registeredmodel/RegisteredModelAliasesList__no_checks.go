// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package registeredmodel

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RegisteredModelAliasesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RegisteredModelAliasesList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RegisteredModelAliasesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RegisteredModelAliasesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RegisteredModelAliasesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RegisteredModelAliasesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RegisteredModelAliasesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRegisteredModelAliasesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

