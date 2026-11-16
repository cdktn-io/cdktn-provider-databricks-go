// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package accountsettingv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validatePutEndParameters(value *AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingEnd) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validatePutMidParameters(value *AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMid) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validatePutStartParameters(value *AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingStart) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDiverging:
		val := val.(*AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDiverging)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDiverging:
		val_ := val.(AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDiverging)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDiverging; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

