// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package share

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ShareObjectPartitionValueList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ShareObjectPartitionValueList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ShareObjectPartitionValueList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ShareObjectPartitionValueList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ShareObjectPartitionValueList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ShareObjectPartitionValueList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ShareObjectPartitionValueList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewShareObjectPartitionValueListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

