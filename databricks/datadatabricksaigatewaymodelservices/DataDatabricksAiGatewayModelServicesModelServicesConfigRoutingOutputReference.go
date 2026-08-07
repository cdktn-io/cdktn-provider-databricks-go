// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelservices

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksaigatewaymodelservices/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference interface {
	cdktn.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Destinations() DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingDestinationsList
	DestinationsInput() interface{}
	Fallback() DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackOutputReference
	FallbackInput() interface{}
	FirstTokenTimeout() *string
	SetFirstTokenTimeout(val *string)
	FirstTokenTimeoutInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrafficSplitting() DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingTrafficSplittingOutputReference
	TrafficSplittingInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutDestinations(value interface{})
	PutFallback(value *DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallback)
	PutTrafficSplitting(value *DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingTrafficSplitting)
	ResetDestinations()
	ResetFallback()
	ResetFirstTokenTimeout()
	ResetTrafficSplitting()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference
type jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) Destinations() DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingDestinationsList {
	var returns DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingDestinationsList
	_jsii_.Get(
		j,
		"destinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) DestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) Fallback() DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackOutputReference {
	var returns DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackOutputReference
	_jsii_.Get(
		j,
		"fallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) FallbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) FirstTokenTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstTokenTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) FirstTokenTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstTokenTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) TrafficSplitting() DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingTrafficSplittingOutputReference {
	var returns DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingTrafficSplittingOutputReference
	_jsii_.Get(
		j,
		"trafficSplitting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) TrafficSplittingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficSplittingInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiGatewayModelServices.DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference_Override(d DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiGatewayModelServices.DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference)SetFirstTokenTimeout(val *string) {
	if err := j.validateSetFirstTokenTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstTokenTimeout",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) PutDestinations(value interface{}) {
	if err := d.validatePutDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDestinations",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) PutFallback(value *DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallback) {
	if err := d.validatePutFallbackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFallback",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) PutTrafficSplitting(value *DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingTrafficSplitting) {
	if err := d.validatePutTrafficSplittingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTrafficSplitting",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) ResetDestinations() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) ResetFallback() {
	_jsii_.InvokeVoid(
		d,
		"resetFallback",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) ResetFirstTokenTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetFirstTokenTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) ResetTrafficSplitting() {
	_jsii_.InvokeVoid(
		d,
		"resetTrafficSplitting",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

