// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/aigatewaymodelservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference interface {
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
	DestinationType() *string
	SetDestinationType(val *string)
	DestinationTypeInput() *string
	ExternalModelConfig() AiGatewayModelServiceConfigRoutingFallbackDestinationsExternalModelConfigOutputReference
	ExternalModelConfigInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsDeleted() cdktn.IResolvable
	Name() *string
	SetName(val *string)
	NameInput() *string
	PayPerTokenConfig() AiGatewayModelServiceConfigRoutingFallbackDestinationsPayPerTokenConfigOutputReference
	PayPerTokenConfigInput() interface{}
	ProvisionedThroughputConfig() AiGatewayModelServiceConfigRoutingFallbackDestinationsProvisionedThroughputConfigOutputReference
	ProvisionedThroughputConfigInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrafficPercentage() *float64
	SetTrafficPercentage(val *float64)
	TrafficPercentageInput() *float64
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
	PutExternalModelConfig(value *AiGatewayModelServiceConfigRoutingFallbackDestinationsExternalModelConfig)
	PutPayPerTokenConfig(value *AiGatewayModelServiceConfigRoutingFallbackDestinationsPayPerTokenConfig)
	PutProvisionedThroughputConfig(value *AiGatewayModelServiceConfigRoutingFallbackDestinationsProvisionedThroughputConfig)
	ResetExternalModelConfig()
	ResetPayPerTokenConfig()
	ResetProvisionedThroughputConfig()
	ResetTrafficPercentage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference
type jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) DestinationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) DestinationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ExternalModelConfig() AiGatewayModelServiceConfigRoutingFallbackDestinationsExternalModelConfigOutputReference {
	var returns AiGatewayModelServiceConfigRoutingFallbackDestinationsExternalModelConfigOutputReference
	_jsii_.Get(
		j,
		"externalModelConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ExternalModelConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalModelConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) IsDeleted() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isDeleted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) PayPerTokenConfig() AiGatewayModelServiceConfigRoutingFallbackDestinationsPayPerTokenConfigOutputReference {
	var returns AiGatewayModelServiceConfigRoutingFallbackDestinationsPayPerTokenConfigOutputReference
	_jsii_.Get(
		j,
		"payPerTokenConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) PayPerTokenConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"payPerTokenConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ProvisionedThroughputConfig() AiGatewayModelServiceConfigRoutingFallbackDestinationsProvisionedThroughputConfigOutputReference {
	var returns AiGatewayModelServiceConfigRoutingFallbackDestinationsProvisionedThroughputConfigOutputReference
	_jsii_.Get(
		j,
		"provisionedThroughputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ProvisionedThroughputConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedThroughputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) TrafficPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trafficPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) TrafficPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trafficPercentageInput",
		&returns,
	)
	return returns
}


func NewAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference {
	_init_.Initialize()

	if err := validateNewAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.aiGatewayModelService.AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference_Override(a AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.aiGatewayModelService.AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference)SetDestinationType(val *string) {
	if err := j.validateSetDestinationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationType",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference)SetTrafficPercentage(val *float64) {
	if err := j.validateSetTrafficPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficPercentage",
		val,
	)
}

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) PutExternalModelConfig(value *AiGatewayModelServiceConfigRoutingFallbackDestinationsExternalModelConfig) {
	if err := a.validatePutExternalModelConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExternalModelConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) PutPayPerTokenConfig(value *AiGatewayModelServiceConfigRoutingFallbackDestinationsPayPerTokenConfig) {
	if err := a.validatePutPayPerTokenConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPayPerTokenConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) PutProvisionedThroughputConfig(value *AiGatewayModelServiceConfigRoutingFallbackDestinationsProvisionedThroughputConfig) {
	if err := a.validatePutProvisionedThroughputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProvisionedThroughputConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ResetExternalModelConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalModelConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ResetPayPerTokenConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetPayPerTokenConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ResetProvisionedThroughputConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetProvisionedThroughputConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ResetTrafficPercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetTrafficPercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

