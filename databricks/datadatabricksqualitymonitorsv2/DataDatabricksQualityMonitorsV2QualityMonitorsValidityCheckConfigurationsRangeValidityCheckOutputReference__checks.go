// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package datadatabricksqualitymonitorsv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (d *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateSetColumnNamesParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheck:
		val := val.(*DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheck)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheck:
		val_ := val.(DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheck)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheck; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateSetLowerBoundParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReference) validateSetUpperBoundParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheckOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

