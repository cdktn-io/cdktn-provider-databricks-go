// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package datadatabricksworkspacesettingv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference) validateSetDarkParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMid:
		val := val.(*DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMid)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMid:
		val_ := val.(DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMid)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMid; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference) validateSetLightParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

