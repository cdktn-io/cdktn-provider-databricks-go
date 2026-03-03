// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package tagpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TagPolicyValuesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TagPolicyValuesList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TagPolicyValuesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TagPolicyValuesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TagPolicyValuesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TagPolicyValuesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TagPolicyValuesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTagPolicyValuesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

