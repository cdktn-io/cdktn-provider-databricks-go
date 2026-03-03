// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package pipeline

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PipelineIngestionDefinitionSourceConfigurationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PipelineIngestionDefinitionSourceConfigurationsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PipelineIngestionDefinitionSourceConfigurationsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PipelineIngestionDefinitionSourceConfigurationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PipelineIngestionDefinitionSourceConfigurationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PipelineIngestionDefinitionSourceConfigurationsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PipelineIngestionDefinitionSourceConfigurationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPipelineIngestionDefinitionSourceConfigurationsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

