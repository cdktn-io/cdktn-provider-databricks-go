// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymcpservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/aigatewaymcpservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AiGatewayMcpServiceConfigRateLimitsOutputReference interface {
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

// The jsii proxy struct for AiGatewayMcpServiceConfigRateLimitsOutputReference
type jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) PrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) RenewalPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renewalPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) RenewalPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renewalPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) Requests() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) RequestsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) RequestTagKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTagKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) RequestTagKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTagKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) RequestTagValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTagValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) RequestTagValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTagValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) Tokens() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) TokensInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokensInput",
		&returns,
	)
	return returns
}


func NewAiGatewayMcpServiceConfigRateLimitsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AiGatewayMcpServiceConfigRateLimitsOutputReference {
	_init_.Initialize()

	if err := validateNewAiGatewayMcpServiceConfigRateLimitsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.aiGatewayMcpService.AiGatewayMcpServiceConfigRateLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAiGatewayMcpServiceConfigRateLimitsOutputReference_Override(a AiGatewayMcpServiceConfigRateLimitsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.aiGatewayMcpService.AiGatewayMcpServiceConfigRateLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference)SetPrincipal(val *string) {
	if err := j.validateSetPrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference)SetRenewalPeriod(val *string) {
	if err := j.validateSetRenewalPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renewalPeriod",
		val,
	)
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference)SetRequests(val *float64) {
	if err := j.validateSetRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requests",
		val,
	)
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference)SetRequestTagKey(val *string) {
	if err := j.validateSetRequestTagKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTagKey",
		val,
	)
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference)SetRequestTagValue(val *string) {
	if err := j.validateSetRequestTagValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTagValue",
		val,
	)
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference)SetTokens(val *float64) {
	if err := j.validateSetTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokens",
		val,
	)
}

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) ResetPrincipal() {
	_jsii_.InvokeVoid(
		a,
		"resetPrincipal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) ResetRequests() {
	_jsii_.InvokeVoid(
		a,
		"resetRequests",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) ResetRequestTagKey() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestTagKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) ResetRequestTagValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestTagValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) ResetTokens() {
	_jsii_.InvokeVoid(
		a,
		"resetTokens",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AiGatewayMcpServiceConfigRateLimitsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

