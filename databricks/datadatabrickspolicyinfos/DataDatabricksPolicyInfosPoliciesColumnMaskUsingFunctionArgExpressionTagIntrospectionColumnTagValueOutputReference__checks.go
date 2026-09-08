// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package datadatabrickspolicyinfos

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference) validateSetColumnAliasParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValue:
		val := val.(*DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValue)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValue:
		val_ := val.(DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValue)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValue; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference) validateSetTagKeyParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

