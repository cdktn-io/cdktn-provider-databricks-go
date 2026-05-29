// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/accountnetworkpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference interface {
	cdktn.ComplexObject
	AccountApi() AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAccountApiOutputReference
	AccountApiInput() interface{}
	AccountDatabricksOne() AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAccountDatabricksOneOutputReference
	AccountDatabricksOneInput() interface{}
	AccountUi() AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAccountUiOutputReference
	AccountUiInput() interface{}
	AllDestinations() interface{}
	SetAllDestinations(val interface{})
	AllDestinationsInput() interface{}
	AppsRuntime() AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAppsRuntimeOutputReference
	AppsRuntimeInput() interface{}
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
	LakebaseRuntime() AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationLakebaseRuntimeOutputReference
	LakebaseRuntimeInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WorkspaceApi() AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceApiOutputReference
	WorkspaceApiInput() interface{}
	WorkspaceUi() AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceUiOutputReference
	WorkspaceUiInput() interface{}
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
	PutAccountApi(value *AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAccountApi)
	PutAccountDatabricksOne(value *AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAccountDatabricksOne)
	PutAccountUi(value *AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAccountUi)
	PutAppsRuntime(value *AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAppsRuntime)
	PutLakebaseRuntime(value *AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationLakebaseRuntime)
	PutWorkspaceApi(value *AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceApi)
	PutWorkspaceUi(value *AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceUi)
	ResetAccountApi()
	ResetAccountDatabricksOne()
	ResetAccountUi()
	ResetAllDestinations()
	ResetAppsRuntime()
	ResetLakebaseRuntime()
	ResetWorkspaceApi()
	ResetWorkspaceUi()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference
type jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) AccountApi() AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAccountApiOutputReference {
	var returns AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAccountApiOutputReference
	_jsii_.Get(
		j,
		"accountApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) AccountApiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) AccountDatabricksOne() AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAccountDatabricksOneOutputReference {
	var returns AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAccountDatabricksOneOutputReference
	_jsii_.Get(
		j,
		"accountDatabricksOne",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) AccountDatabricksOneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountDatabricksOneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) AccountUi() AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAccountUiOutputReference {
	var returns AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAccountUiOutputReference
	_jsii_.Get(
		j,
		"accountUi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) AccountUiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountUiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) AllDestinations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) AllDestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allDestinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) AppsRuntime() AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAppsRuntimeOutputReference {
	var returns AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAppsRuntimeOutputReference
	_jsii_.Get(
		j,
		"appsRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) AppsRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appsRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) LakebaseRuntime() AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationLakebaseRuntimeOutputReference {
	var returns AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationLakebaseRuntimeOutputReference
	_jsii_.Get(
		j,
		"lakebaseRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) LakebaseRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lakebaseRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) WorkspaceApi() AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceApiOutputReference {
	var returns AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceApiOutputReference
	_jsii_.Get(
		j,
		"workspaceApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) WorkspaceApiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) WorkspaceUi() AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceUiOutputReference {
	var returns AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceUiOutputReference
	_jsii_.Get(
		j,
		"workspaceUi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) WorkspaceUiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceUiInput",
		&returns,
	)
	return returns
}


func NewAccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewAccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.accountNetworkPolicy.AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference_Override(a AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.accountNetworkPolicy.AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference)SetAllDestinations(val interface{}) {
	if err := j.validateSetAllDestinationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allDestinations",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) PutAccountApi(value *AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAccountApi) {
	if err := a.validatePutAccountApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccountApi",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) PutAccountDatabricksOne(value *AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAccountDatabricksOne) {
	if err := a.validatePutAccountDatabricksOneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccountDatabricksOne",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) PutAccountUi(value *AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAccountUi) {
	if err := a.validatePutAccountUiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccountUi",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) PutAppsRuntime(value *AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAppsRuntime) {
	if err := a.validatePutAppsRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAppsRuntime",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) PutLakebaseRuntime(value *AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationLakebaseRuntime) {
	if err := a.validatePutLakebaseRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLakebaseRuntime",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) PutWorkspaceApi(value *AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceApi) {
	if err := a.validatePutWorkspaceApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorkspaceApi",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) PutWorkspaceUi(value *AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceUi) {
	if err := a.validatePutWorkspaceUiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorkspaceUi",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) ResetAccountApi() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountApi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) ResetAccountDatabricksOne() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountDatabricksOne",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) ResetAccountUi() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountUi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) ResetAllDestinations() {
	_jsii_.InvokeVoid(
		a,
		"resetAllDestinations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) ResetAppsRuntime() {
	_jsii_.InvokeVoid(
		a,
		"resetAppsRuntime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) ResetLakebaseRuntime() {
	_jsii_.InvokeVoid(
		a,
		"resetLakebaseRuntime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) ResetWorkspaceApi() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkspaceApi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) ResetWorkspaceUi() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkspaceUi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

