// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaisearchendpoints

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabricksaisearchendpoints/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference interface {
	cdktn.ComplexObject
	ChangeRequestMessage() *string
	ChangeRequestState() *string
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
	CurrentConcurrency() *float64
	CurrentConcurrencyUtilizationPercentage() *float64
	CurrentNumReplicas() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksAiSearchEndpointsEndpointsThroughputInfo
	SetInternalValue(val *DataDatabricksAiSearchEndpointsEndpointsThroughputInfo)
	MaximumConcurrencyAllowed() *float64
	SetMaximumConcurrencyAllowed(val *float64)
	MaximumConcurrencyAllowedInput() *float64
	MinimalConcurrencyAllowed() *float64
	SetMinimalConcurrencyAllowed(val *float64)
	MinimalConcurrencyAllowedInput() *float64
	RequestedConcurrency() *float64
	SetRequestedConcurrency(val *float64)
	RequestedConcurrencyInput() *float64
	RequestedNumReplicas() *float64
	SetRequestedNumReplicas(val *float64)
	RequestedNumReplicasInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	ResetMaximumConcurrencyAllowed()
	ResetMinimalConcurrencyAllowed()
	ResetRequestedConcurrency()
	ResetRequestedNumReplicas()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference
type jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) ChangeRequestMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeRequestMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) ChangeRequestState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeRequestState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) CurrentConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) CurrentConcurrencyUtilizationPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentConcurrencyUtilizationPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) CurrentNumReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentNumReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) InternalValue() *DataDatabricksAiSearchEndpointsEndpointsThroughputInfo {
	var returns *DataDatabricksAiSearchEndpointsEndpointsThroughputInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) MaximumConcurrencyAllowed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumConcurrencyAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) MaximumConcurrencyAllowedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumConcurrencyAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) MinimalConcurrencyAllowed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimalConcurrencyAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) MinimalConcurrencyAllowedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimalConcurrencyAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) RequestedConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) RequestedConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) RequestedNumReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedNumReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) RequestedNumReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedNumReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiSearchEndpoints.DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference_Override(d DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiSearchEndpoints.DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference)SetInternalValue(val *DataDatabricksAiSearchEndpointsEndpointsThroughputInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference)SetMaximumConcurrencyAllowed(val *float64) {
	if err := j.validateSetMaximumConcurrencyAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumConcurrencyAllowed",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference)SetMinimalConcurrencyAllowed(val *float64) {
	if err := j.validateSetMinimalConcurrencyAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimalConcurrencyAllowed",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference)SetRequestedConcurrency(val *float64) {
	if err := j.validateSetRequestedConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestedConcurrency",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference)SetRequestedNumReplicas(val *float64) {
	if err := j.validateSetRequestedNumReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestedNumReplicas",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) ResetMaximumConcurrencyAllowed() {
	_jsii_.InvokeVoid(
		d,
		"resetMaximumConcurrencyAllowed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) ResetMinimalConcurrencyAllowed() {
	_jsii_.InvokeVoid(
		d,
		"resetMinimalConcurrencyAllowed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) ResetRequestedConcurrency() {
	_jsii_.InvokeVoid(
		d,
		"resetRequestedConcurrency",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) ResetRequestedNumReplicas() {
	_jsii_.InvokeVoid(
		d,
		"resetRequestedNumReplicas",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointsEndpointsThroughputInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

