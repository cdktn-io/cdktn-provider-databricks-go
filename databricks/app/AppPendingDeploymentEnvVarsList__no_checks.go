// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package app

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppPendingDeploymentEnvVarsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppPendingDeploymentEnvVarsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppPendingDeploymentEnvVarsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppPendingDeploymentEnvVarsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppPendingDeploymentEnvVarsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppPendingDeploymentEnvVarsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppPendingDeploymentEnvVarsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppPendingDeploymentEnvVarsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

