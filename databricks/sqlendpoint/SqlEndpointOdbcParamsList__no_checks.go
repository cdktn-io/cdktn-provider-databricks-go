// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sqlendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqlEndpointOdbcParamsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SqlEndpointOdbcParamsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SqlEndpointOdbcParamsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SqlEndpointOdbcParamsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SqlEndpointOdbcParamsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SqlEndpointOdbcParamsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSqlEndpointOdbcParamsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

