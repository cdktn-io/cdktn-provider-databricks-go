// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/aisearchendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AiSearchEndpointThroughputInfoOutputReference interface {
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
	InternalValue() *AiSearchEndpointThroughputInfo
	SetInternalValue(val *AiSearchEndpointThroughputInfo)
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

// The jsii proxy struct for AiSearchEndpointThroughputInfoOutputReference
type jsiiProxy_AiSearchEndpointThroughputInfoOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) ChangeRequestMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeRequestMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) ChangeRequestState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeRequestState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) CurrentConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) CurrentConcurrencyUtilizationPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentConcurrencyUtilizationPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) CurrentNumReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentNumReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) InternalValue() *AiSearchEndpointThroughputInfo {
	var returns *AiSearchEndpointThroughputInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) MaximumConcurrencyAllowed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumConcurrencyAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) MaximumConcurrencyAllowedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumConcurrencyAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) MinimalConcurrencyAllowed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimalConcurrencyAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) MinimalConcurrencyAllowedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimalConcurrencyAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) RequestedConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) RequestedConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) RequestedNumReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedNumReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) RequestedNumReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedNumReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAiSearchEndpointThroughputInfoOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AiSearchEndpointThroughputInfoOutputReference {
	_init_.Initialize()

	if err := validateNewAiSearchEndpointThroughputInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiSearchEndpointThroughputInfoOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.aiSearchEndpoint.AiSearchEndpointThroughputInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAiSearchEndpointThroughputInfoOutputReference_Override(a AiSearchEndpointThroughputInfoOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.aiSearchEndpoint.AiSearchEndpointThroughputInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference)SetInternalValue(val *AiSearchEndpointThroughputInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference)SetMaximumConcurrencyAllowed(val *float64) {
	if err := j.validateSetMaximumConcurrencyAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumConcurrencyAllowed",
		val,
	)
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference)SetMinimalConcurrencyAllowed(val *float64) {
	if err := j.validateSetMinimalConcurrencyAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimalConcurrencyAllowed",
		val,
	)
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference)SetRequestedConcurrency(val *float64) {
	if err := j.validateSetRequestedConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestedConcurrency",
		val,
	)
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference)SetRequestedNumReplicas(val *float64) {
	if err := j.validateSetRequestedNumReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestedNumReplicas",
		val,
	)
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) ResetMaximumConcurrencyAllowed() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumConcurrencyAllowed",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) ResetMinimalConcurrencyAllowed() {
	_jsii_.InvokeVoid(
		a,
		"resetMinimalConcurrencyAllowed",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) ResetRequestedConcurrency() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestedConcurrency",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) ResetRequestedNumReplicas() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestedNumReplicas",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchEndpointThroughputInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

