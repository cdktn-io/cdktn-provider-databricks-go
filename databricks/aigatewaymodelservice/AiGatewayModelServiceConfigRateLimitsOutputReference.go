// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/aigatewaymodelservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AiGatewayModelServiceConfigRateLimitsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	Principal() *string
	SetPrincipal(val *string)
	PrincipalInput() *string
	RenewalPeriod() *string
	SetRenewalPeriod(val *string)
	RenewalPeriodInput() *string
	Requests() *float64
	SetRequests(val *float64)
	RequestsInput() *float64
	RequestTagKey() *string
	SetRequestTagKey(val *string)
	RequestTagKeyInput() *string
	RequestTagValue() *string
	SetRequestTagValue(val *string)
	RequestTagValueInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Tokens() *float64
	SetTokens(val *float64)
	TokensInput() *float64
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
	ResetPrincipal()
	ResetRequests()
	ResetRequestTagKey()
	ResetRequestTagValue()
	ResetTokens()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AiGatewayModelServiceConfigRateLimitsOutputReference
type jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) PrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) RenewalPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renewalPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) RenewalPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renewalPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) Requests() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) RequestsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) RequestTagKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTagKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) RequestTagKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTagKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) RequestTagValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTagValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) RequestTagValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTagValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) Tokens() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) TokensInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokensInput",
		&returns,
	)
	return returns
}


func NewAiGatewayModelServiceConfigRateLimitsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AiGatewayModelServiceConfigRateLimitsOutputReference {
	_init_.Initialize()

	if err := validateNewAiGatewayModelServiceConfigRateLimitsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.aiGatewayModelService.AiGatewayModelServiceConfigRateLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAiGatewayModelServiceConfigRateLimitsOutputReference_Override(a AiGatewayModelServiceConfigRateLimitsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.aiGatewayModelService.AiGatewayModelServiceConfigRateLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference)SetPrincipal(val *string) {
	if err := j.validateSetPrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference)SetRenewalPeriod(val *string) {
	if err := j.validateSetRenewalPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renewalPeriod",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference)SetRequests(val *float64) {
	if err := j.validateSetRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requests",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference)SetRequestTagKey(val *string) {
	if err := j.validateSetRequestTagKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTagKey",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference)SetRequestTagValue(val *string) {
	if err := j.validateSetRequestTagValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTagValue",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference)SetTokens(val *float64) {
	if err := j.validateSetTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokens",
		val,
	)
}

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) ResetPrincipal() {
	_jsii_.InvokeVoid(
		a,
		"resetPrincipal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) ResetRequests() {
	_jsii_.InvokeVoid(
		a,
		"resetRequests",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) ResetRequestTagKey() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestTagKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) ResetRequestTagValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestTagValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) ResetTokens() {
	_jsii_.InvokeVoid(
		a,
		"resetTokens",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AiGatewayModelServiceConfigRateLimitsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

