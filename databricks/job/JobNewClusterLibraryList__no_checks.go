// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package job

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobNewClusterLibraryList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (j *jsiiProxy_JobNewClusterLibraryList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JobNewClusterLibraryList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JobNewClusterLibraryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JobNewClusterLibraryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobNewClusterLibraryList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JobNewClusterLibraryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJobNewClusterLibraryListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

