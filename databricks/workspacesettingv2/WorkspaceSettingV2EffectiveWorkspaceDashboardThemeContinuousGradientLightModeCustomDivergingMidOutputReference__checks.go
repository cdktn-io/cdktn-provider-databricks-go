// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package workspacesettingv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference) validateSetDarkParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMid:
		val := val.(*WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMid)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMid:
		val_ := val.(WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMid)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMid; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference) validateSetLightParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewWorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

