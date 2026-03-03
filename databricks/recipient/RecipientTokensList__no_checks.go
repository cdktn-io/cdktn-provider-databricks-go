// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package recipient

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RecipientTokensList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RecipientTokensList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RecipientTokensList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RecipientTokensList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RecipientTokensList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RecipientTokensList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RecipientTokensList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRecipientTokensListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

