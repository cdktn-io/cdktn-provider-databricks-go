// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package mlflowmodel

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MlflowModelTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MlflowModelTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MlflowModelTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MlflowModelTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MlflowModelTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MlflowModelTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MlflowModelTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMlflowModelTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

