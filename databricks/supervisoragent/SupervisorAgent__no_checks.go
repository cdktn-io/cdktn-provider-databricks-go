// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package supervisoragent

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SupervisorAgent) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorAgent) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_SupervisorAgent) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorAgent) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorAgent) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorAgent) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorAgent) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorAgent) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorAgent) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorAgent) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorAgent) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorAgent) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorAgent) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorAgent) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorAgent) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_SupervisorAgent) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorAgent) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_SupervisorAgent) validatePutProviderConfigParameters(value *SupervisorAgentProviderConfig) error {
	return nil
}

func validateSupervisorAgent_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSupervisorAgent_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSupervisorAgent_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSupervisorAgent_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_SupervisorAgent) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SupervisorAgent) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SupervisorAgent) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SupervisorAgent) validateSetDisplayNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SupervisorAgent) validateSetInstructionsParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SupervisorAgent) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_SupervisorAgent) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewSupervisorAgentParameters(scope constructs.Construct, id *string, config *SupervisorAgentConfig) error {
	return nil
}

