// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package accountsettingv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) validatePutEndParameters(value *AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialEnd) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) validatePutStartParameters(value *AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialStart) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequential:
		val := val.(*AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequential)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequential:
		val_ := val.(AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequential)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequential; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

