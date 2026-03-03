// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sqlwidget

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqlWidgetParameterList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SqlWidgetParameterList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SqlWidgetParameterList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SqlWidgetParameterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SqlWidgetParameterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SqlWidgetParameterList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SqlWidgetParameterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSqlWidgetParameterListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

