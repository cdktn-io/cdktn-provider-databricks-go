// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lakewatchdatasource

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LakewatchDatasourceParsersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LakewatchDatasourceParsersList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LakewatchDatasourceParsersList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LakewatchDatasourceParsersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LakewatchDatasourceParsersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LakewatchDatasourceParsersList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LakewatchDatasourceParsersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLakewatchDatasourceParsersListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

