// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appsspace

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppsSpaceResourcesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppsSpaceResourcesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppsSpaceResourcesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppsSpaceResourcesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppsSpaceResourcesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppsSpaceResourcesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppsSpaceResourcesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppsSpaceResourcesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

