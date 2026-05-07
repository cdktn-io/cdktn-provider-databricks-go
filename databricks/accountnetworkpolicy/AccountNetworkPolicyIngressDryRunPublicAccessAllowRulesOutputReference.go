// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/accountnetworkpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference interface {
	cdktn.ComplexObject
	Authentication() AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesAuthenticationOutputReference
	AuthenticationInput() interface{}
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
	Destination() AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesDestinationOutputReference
	DestinationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	Origin() AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOriginOutputReference
	OriginInput() interface{}
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
	PutAuthentication(value *AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesAuthentication)
	PutDestination(value *AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesDestination)
	PutOrigin(value *AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOrigin)
	ResetAuthentication()
	ResetDestination()
	ResetLabel()
	ResetOrigin()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference
type jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) Authentication() AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesAuthenticationOutputReference {
	var returns AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesAuthenticationOutputReference
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) AuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) Destination() AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesDestinationOutputReference {
	var returns AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesDestinationOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) DestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) Origin() AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOriginOutputReference {
	var returns AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOriginOutputReference
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) OriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference {
	_init_.Initialize()

	if err := validateNewAccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.accountNetworkPolicy.AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference_Override(a AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.accountNetworkPolicy.AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) PutAuthentication(value *AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesAuthentication) {
	if err := a.validatePutAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthentication",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) PutDestination(value *AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesDestination) {
	if err := a.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) PutOrigin(value *AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOrigin) {
	if err := a.validatePutOriginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOrigin",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) ResetAuthentication() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) ResetDestination() {
	_jsii_.InvokeVoid(
		a,
		"resetDestination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		a,
		"resetLabel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) ResetOrigin() {
	_jsii_.InvokeVoid(
		a,
		"resetOrigin",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

