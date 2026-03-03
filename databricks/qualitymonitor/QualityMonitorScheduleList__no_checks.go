// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package qualitymonitor

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QualityMonitorScheduleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (q *jsiiProxy_QualityMonitorScheduleList) validateGetParameters(index *float64) error {
	return nil
}

func (q *jsiiProxy_QualityMonitorScheduleList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorScheduleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorScheduleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorScheduleList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorScheduleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewQualityMonitorScheduleListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

