// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package library

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LibraryCranList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LibraryCranList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LibraryCranList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LibraryCranList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LibraryCranList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LibraryCranList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LibraryCranList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLibraryCranListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

