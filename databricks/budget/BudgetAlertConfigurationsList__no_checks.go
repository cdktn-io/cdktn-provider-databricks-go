// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package budget

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BudgetAlertConfigurationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BudgetAlertConfigurationsList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BudgetAlertConfigurationsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BudgetAlertConfigurationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BudgetAlertConfigurationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BudgetAlertConfigurationsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BudgetAlertConfigurationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBudgetAlertConfigurationsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

