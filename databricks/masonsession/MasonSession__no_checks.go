// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package masonsession

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MasonSession) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validateImportFromParameters(id *string) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validateMoveToIdParameters(id *string) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validatePutProviderConfigParameters(value *MasonSessionProviderConfig) error {
	return nil
}

func (m *jsiiProxy_MasonSession) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateMasonSession_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateMasonSession_IsConstructParameters(x interface{}) error {
	return nil
}

func validateMasonSession_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateMasonSession_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_MasonSession) validateSetActorIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MasonSession) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MasonSession) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MasonSession) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_MasonSession) validateSetMetadataParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_MasonSession) validateSetParentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MasonSession) validateSetParentSessionIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MasonSession) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_MasonSession) validateSetSessionIdParameters(val *string) error {
	return nil
}

func validateNewMasonSessionParameters(scope constructs.Construct, id *string, config *MasonSessionConfig) error {
	return nil
}

