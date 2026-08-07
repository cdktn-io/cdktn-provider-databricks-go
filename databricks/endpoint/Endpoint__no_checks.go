// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package endpoint

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Endpoint) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validateImportFromParameters(id *string) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validateMoveToIdParameters(id *string) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validatePutAwsVpcEndpointInfoParameters(value *EndpointAwsVpcEndpointInfo) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validatePutAzurePrivateEndpointInfoParameters(value *EndpointAzurePrivateEndpointInfo) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validatePutGcpPscEndpointInfoParameters(value *EndpointGcpPscEndpointInfo) error {
	return nil
}

func (e *jsiiProxy_Endpoint) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateEndpoint_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateEndpoint_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEndpoint_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateEndpoint_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Endpoint) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Endpoint) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Endpoint) validateSetDisplayNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Endpoint) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Endpoint) validateSetParentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Endpoint) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Endpoint) validateSetRegionParameters(val *string) error {
	return nil
}

func validateNewEndpointParameters(scope constructs.Construct, id *string, config *EndpointConfig) error {
	return nil
}

