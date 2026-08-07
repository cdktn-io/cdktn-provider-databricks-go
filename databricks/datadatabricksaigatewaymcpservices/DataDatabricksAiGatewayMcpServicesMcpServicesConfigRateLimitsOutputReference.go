// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymcpservices

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksaigatewaymcpservices/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference interface {
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

// The jsii proxy struct for DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference
type jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) PrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) RenewalPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renewalPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) RenewalPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renewalPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) Requests() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) RequestsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) RequestTagKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTagKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) RequestTagKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTagKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) RequestTagValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTagValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) RequestTagValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTagValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) Tokens() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) TokensInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokensInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiGatewayMcpServices.DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference_Override(d DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiGatewayMcpServices.DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference)SetPrincipal(val *string) {
	if err := j.validateSetPrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference)SetRenewalPeriod(val *string) {
	if err := j.validateSetRenewalPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renewalPeriod",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference)SetRequests(val *float64) {
	if err := j.validateSetRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requests",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference)SetRequestTagKey(val *string) {
	if err := j.validateSetRequestTagKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTagKey",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference)SetRequestTagValue(val *string) {
	if err := j.validateSetRequestTagValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTagValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference)SetTokens(val *float64) {
	if err := j.validateSetTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokens",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) ResetPrincipal() {
	_jsii_.InvokeVoid(
		d,
		"resetPrincipal",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) ResetRequests() {
	_jsii_.InvokeVoid(
		d,
		"resetRequests",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) ResetRequestTagKey() {
	_jsii_.InvokeVoid(
		d,
		"resetRequestTagKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) ResetRequestTagValue() {
	_jsii_.InvokeVoid(
		d,
		"resetRequestTagValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) ResetTokens() {
	_jsii_.InvokeVoid(
		d,
		"resetTokens",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

