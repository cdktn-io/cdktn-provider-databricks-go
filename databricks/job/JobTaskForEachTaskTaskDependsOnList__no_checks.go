// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package job

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobTaskForEachTaskTaskDependsOnList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDependsOnList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDependsOnList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDependsOnList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDependsOnList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDependsOnList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDependsOnList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJobTaskForEachTaskTaskDependsOnListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

