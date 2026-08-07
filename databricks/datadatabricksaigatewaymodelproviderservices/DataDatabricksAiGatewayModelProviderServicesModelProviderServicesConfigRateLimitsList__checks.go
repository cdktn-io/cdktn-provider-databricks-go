// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package datadatabricksaigatewaymodelproviderservices

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigRateLimitsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigRateLimitsList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigRateLimitsList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigRateLimitsList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigRateLimits:
		val := val.(*[]*DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigRateLimits)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigRateLimits:
		val_ := val.([]*DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigRateLimits)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *[]*DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigRateLimits; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigRateLimitsList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigRateLimitsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigRateLimitsList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigRateLimitsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

