// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package datadatabricksaccountnetworkpolicy

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference) validatePutIdentitiesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationIdentities:
		value := value.(*[]*DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationIdentities)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationIdentities:
		value_ := value.([]*DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationIdentities)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationIdentities; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference) validateSetIdentityTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthentication:
		val := val.(*DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthentication)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthentication:
		val_ := val.(DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthentication)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthentication; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

