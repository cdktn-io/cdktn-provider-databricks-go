// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package job

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobTaskDependsOnList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (j *jsiiProxy_JobTaskDependsOnList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JobTaskDependsOnList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JobTaskDependsOnList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JobTaskDependsOnList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobTaskDependsOnList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JobTaskDependsOnList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJobTaskDependsOnListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

