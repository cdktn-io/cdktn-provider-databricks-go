// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package databaseinstance

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseInstanceChildInstanceRefsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceChildInstanceRefsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceChildInstanceRefsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabaseInstanceChildInstanceRefsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabaseInstanceChildInstanceRefsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabaseInstanceChildInstanceRefsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabaseInstanceChildInstanceRefsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabaseInstanceChildInstanceRefsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

