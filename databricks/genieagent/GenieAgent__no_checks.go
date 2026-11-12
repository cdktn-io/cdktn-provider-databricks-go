// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package genieagent

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GenieAgent) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validateImportFromParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validateMoveToIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validatePutDefinitionParameters(value *GenieAgentDefinition) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validatePutProviderConfigParameters(value *GenieAgentProviderConfig) error {
	return nil
}

func (g *jsiiProxy_GenieAgent) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateGenieAgent_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateGenieAgent_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGenieAgent_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateGenieAgent_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_GenieAgent) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GenieAgent) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GenieAgent) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GenieAgent) validateSetDisplayNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GenieAgent) validateSetGenieAgentIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GenieAgent) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_GenieAgent) validateSetParentPathParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GenieAgent) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_GenieAgent) validateSetWarehouseIdParameters(val *string) error {
	return nil
}

func validateNewGenieAgentParameters(scope constructs.Construct, id *string, config *GenieAgentConfig) error {
	return nil
}

