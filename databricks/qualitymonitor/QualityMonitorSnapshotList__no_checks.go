// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package qualitymonitor

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QualityMonitorSnapshotList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (q *jsiiProxy_QualityMonitorSnapshotList) validateGetParameters(index *float64) error {
	return nil
}

func (q *jsiiProxy_QualityMonitorSnapshotList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorSnapshotList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorSnapshotList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorSnapshotList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_QualityMonitorSnapshotList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewQualityMonitorSnapshotListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

