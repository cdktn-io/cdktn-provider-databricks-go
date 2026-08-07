// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaisearchendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksaisearchendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAiSearchEndpointThroughputInfoOutputReference interface {
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
	InternalValue() *DataDatabricksAiSearchEndpointThroughputInfo
	SetInternalValue(val *DataDatabricksAiSearchEndpointThroughputInfo)
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

// The jsii proxy struct for DataDatabricksAiSearchEndpointThroughputInfoOutputReference
type jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) ChangeRequestMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeRequestMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) ChangeRequestState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeRequestState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) CurrentConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) CurrentConcurrencyUtilizationPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentConcurrencyUtilizationPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) CurrentNumReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentNumReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) InternalValue() *DataDatabricksAiSearchEndpointThroughputInfo {
	var returns *DataDatabricksAiSearchEndpointThroughputInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) MaximumConcurrencyAllowed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumConcurrencyAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) MaximumConcurrencyAllowedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumConcurrencyAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) MinimalConcurrencyAllowed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimalConcurrencyAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) MinimalConcurrencyAllowedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimalConcurrencyAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) RequestedConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) RequestedConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) RequestedNumReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedNumReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) RequestedNumReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedNumReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAiSearchEndpointThroughputInfoOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAiSearchEndpointThroughputInfoOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAiSearchEndpointThroughputInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiSearchEndpoint.DataDatabricksAiSearchEndpointThroughputInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAiSearchEndpointThroughputInfoOutputReference_Override(d DataDatabricksAiSearchEndpointThroughputInfoOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiSearchEndpoint.DataDatabricksAiSearchEndpointThroughputInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference)SetInternalValue(val *DataDatabricksAiSearchEndpointThroughputInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference)SetMaximumConcurrencyAllowed(val *float64) {
	if err := j.validateSetMaximumConcurrencyAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumConcurrencyAllowed",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference)SetMinimalConcurrencyAllowed(val *float64) {
	if err := j.validateSetMinimalConcurrencyAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimalConcurrencyAllowed",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference)SetRequestedConcurrency(val *float64) {
	if err := j.validateSetRequestedConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestedConcurrency",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference)SetRequestedNumReplicas(val *float64) {
	if err := j.validateSetRequestedNumReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestedNumReplicas",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) ResetMaximumConcurrencyAllowed() {
	_jsii_.InvokeVoid(
		d,
		"resetMaximumConcurrencyAllowed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) ResetMinimalConcurrencyAllowed() {
	_jsii_.InvokeVoid(
		d,
		"resetMinimalConcurrencyAllowed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) ResetRequestedConcurrency() {
	_jsii_.InvokeVoid(
		d,
		"resetRequestedConcurrency",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) ResetRequestedNumReplicas() {
	_jsii_.InvokeVoid(
		d,
		"resetRequestedNumReplicas",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAiSearchEndpointThroughputInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

