// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sandbox

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Sandbox) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validatePutProviderConfigParameters(value *SandboxProviderConfig) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validatePutSpecParameters(value *SandboxSpec) error {
	return nil
}

func (s *jsiiProxy_Sandbox) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateSandbox_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSandbox_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSandbox_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSandbox_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Sandbox) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Sandbox) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Sandbox) validateSetDisplayNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Sandbox) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Sandbox) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Sandbox) validateSetSandboxIdParameters(val *string) error {
	return nil
}

func validateNewSandboxParameters(scope constructs.Construct, id *string, config *SandboxConfig) error {
	return nil
}

