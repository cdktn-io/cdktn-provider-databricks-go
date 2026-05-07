// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/accountnetworkpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference interface {
	cdktn.ComplexObject
	AllIpRanges() interface{}
	SetAllIpRanges(val interface{})
	AllIpRangesInput() interface{}
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
	ExcludedIpRanges() AccountNetworkPolicyIngressPublicAccessDenyRulesOriginExcludedIpRangesOutputReference
	ExcludedIpRangesInput() interface{}
	// Experimental.
	Fqn() *string
	IncludedIpRanges() AccountNetworkPolicyIngressPublicAccessDenyRulesOriginIncludedIpRangesOutputReference
	IncludedIpRangesInput() interface{}
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
	PutExcludedIpRanges(value *AccountNetworkPolicyIngressPublicAccessDenyRulesOriginExcludedIpRanges)
	PutIncludedIpRanges(value *AccountNetworkPolicyIngressPublicAccessDenyRulesOriginIncludedIpRanges)
	ResetAllIpRanges()
	ResetExcludedIpRanges()
	ResetIncludedIpRanges()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference
type jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) AllIpRanges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) AllIpRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) ExcludedIpRanges() AccountNetworkPolicyIngressPublicAccessDenyRulesOriginExcludedIpRangesOutputReference {
	var returns AccountNetworkPolicyIngressPublicAccessDenyRulesOriginExcludedIpRangesOutputReference
	_jsii_.Get(
		j,
		"excludedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) ExcludedIpRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) IncludedIpRanges() AccountNetworkPolicyIngressPublicAccessDenyRulesOriginIncludedIpRangesOutputReference {
	var returns AccountNetworkPolicyIngressPublicAccessDenyRulesOriginIncludedIpRangesOutputReference
	_jsii_.Get(
		j,
		"includedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) IncludedIpRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference {
	_init_.Initialize()

	if err := validateNewAccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.accountNetworkPolicy.AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference_Override(a AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.accountNetworkPolicy.AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference)SetAllIpRanges(val interface{}) {
	if err := j.validateSetAllIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allIpRanges",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) PutExcludedIpRanges(value *AccountNetworkPolicyIngressPublicAccessDenyRulesOriginExcludedIpRanges) {
	if err := a.validatePutExcludedIpRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExcludedIpRanges",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) PutIncludedIpRanges(value *AccountNetworkPolicyIngressPublicAccessDenyRulesOriginIncludedIpRanges) {
	if err := a.validatePutIncludedIpRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIncludedIpRanges",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) ResetAllIpRanges() {
	_jsii_.InvokeVoid(
		a,
		"resetAllIpRanges",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) ResetExcludedIpRanges() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludedIpRanges",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) ResetIncludedIpRanges() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludedIpRanges",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressPublicAccessDenyRulesOriginOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

