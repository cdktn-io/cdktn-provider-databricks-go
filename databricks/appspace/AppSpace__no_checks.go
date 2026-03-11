// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appspace

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSpace) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (a *jsiiProxy_AppSpace) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_AppSpace) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppSpace) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppSpace) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppSpace) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppSpace) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppSpace) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppSpace) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppSpace) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppSpace) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppSpace) validateImportFromParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_AppSpace) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppSpace) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_AppSpace) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (a *jsiiProxy_AppSpace) validateMoveToIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_AppSpace) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (a *jsiiProxy_AppSpace) validatePutProviderConfigParameters(value *AppSpaceProviderConfig) error {
	return nil
}

func (a *jsiiProxy_AppSpace) validatePutResourcesParameters(value interface{}) error {
	return nil
}

func validateAppSpace_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateAppSpace_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAppSpace_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateAppSpace_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSpace) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSpace) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSpace) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppSpace) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_AppSpace) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppSpace) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_AppSpace) validateSetUsagePolicyIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppSpace) validateSetUserApiScopesParameters(val *[]*string) error {
	return nil
}

func validateNewAppSpaceParameters(scope constructs.Construct, id *string, config *AppSpaceConfig) error {
	return nil
}

