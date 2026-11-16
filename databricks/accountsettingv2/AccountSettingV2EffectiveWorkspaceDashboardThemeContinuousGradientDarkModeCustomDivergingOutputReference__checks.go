// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package accountsettingv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validatePutEndParameters(value *AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingEnd) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validatePutMidParameters(value *AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMid) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validatePutStartParameters(value *AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingStart) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDiverging:
		val := val.(*AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDiverging)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDiverging:
		val_ := val.(AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDiverging)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDiverging; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

