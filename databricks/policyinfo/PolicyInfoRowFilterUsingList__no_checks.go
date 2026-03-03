// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package policyinfo

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PolicyInfoRowFilterUsingList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PolicyInfoRowFilterUsingList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PolicyInfoRowFilterUsingList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PolicyInfoRowFilterUsingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PolicyInfoRowFilterUsingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PolicyInfoRowFilterUsingList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PolicyInfoRowFilterUsingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPolicyInfoRowFilterUsingListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

