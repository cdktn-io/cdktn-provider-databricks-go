// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package query

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QueryParameterList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (q *jsiiProxy_QueryParameterList) validateGetParameters(index *float64) error {
	return nil
}

func (q *jsiiProxy_QueryParameterList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_QueryParameterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_QueryParameterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QueryParameterList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_QueryParameterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewQueryParameterListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

