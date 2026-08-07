// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/accountnetworkpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
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
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList
type jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList {
	_init_.Initialize()

	if err := validateNewAccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList{}

	_jsii_.Create(
		"@cdktn/provider-databricks.accountNetworkPolicy.AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList_Override(a AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.accountNetworkPolicy.AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := a.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		a,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList) Get(index *float64) AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

