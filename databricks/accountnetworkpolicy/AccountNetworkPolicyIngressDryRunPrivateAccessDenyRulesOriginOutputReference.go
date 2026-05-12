// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/accountnetworkpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference interface {
	cdktn.ComplexObject
	AllPrivateAccess() interface{}
	SetAllPrivateAccess(val interface{})
	AllPrivateAccessInput() interface{}
	AllRegisteredEndpoints() interface{}
	SetAllRegisteredEndpoints(val interface{})
	AllRegisteredEndpointsInput() interface{}
	AzureWorkspacePrivateLink() interface{}
	SetAzureWorkspacePrivateLink(val interface{})
	AzureWorkspacePrivateLinkInput() interface{}
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
	Endpoints() AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginEndpointsOutputReference
	EndpointsInput() interface{}
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
	PutEndpoints(value *AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginEndpoints)
	ResetAllPrivateAccess()
	ResetAllRegisteredEndpoints()
	ResetAzureWorkspacePrivateLink()
	ResetEndpoints()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference
type jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) AllPrivateAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allPrivateAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) AllPrivateAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allPrivateAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) AllRegisteredEndpoints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allRegisteredEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) AllRegisteredEndpointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allRegisteredEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) AzureWorkspacePrivateLink() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureWorkspacePrivateLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) AzureWorkspacePrivateLinkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureWorkspacePrivateLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) Endpoints() AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginEndpointsOutputReference {
	var returns AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginEndpointsOutputReference
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) EndpointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference {
	_init_.Initialize()

	if err := validateNewAccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.accountNetworkPolicy.AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference_Override(a AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.accountNetworkPolicy.AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference)SetAllPrivateAccess(val interface{}) {
	if err := j.validateSetAllPrivateAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allPrivateAccess",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference)SetAllRegisteredEndpoints(val interface{}) {
	if err := j.validateSetAllRegisteredEndpointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allRegisteredEndpoints",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference)SetAzureWorkspacePrivateLink(val interface{}) {
	if err := j.validateSetAzureWorkspacePrivateLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureWorkspacePrivateLink",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) PutEndpoints(value *AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginEndpoints) {
	if err := a.validatePutEndpointsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEndpoints",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) ResetAllPrivateAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetAllPrivateAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) ResetAllRegisteredEndpoints() {
	_jsii_.InvokeVoid(
		a,
		"resetAllRegisteredEndpoints",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) ResetAzureWorkspacePrivateLink() {
	_jsii_.InvokeVoid(
		a,
		"resetAzureWorkspacePrivateLink",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) ResetEndpoints() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpoints",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesOriginOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

