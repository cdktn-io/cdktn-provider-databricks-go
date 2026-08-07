// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/accountnetworkpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference interface {
	cdktn.ComplexObject
	AccountApi() AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountApiOutputReference
	AccountApiInput() interface{}
	AccountDatabricksOne() AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountDatabricksOneOutputReference
	AccountDatabricksOneInput() interface{}
	AccountUi() AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountUiOutputReference
	AccountUiInput() interface{}
	AllDestinations() interface{}
	SetAllDestinations(val interface{})
	AllDestinationsInput() interface{}
	AppsRuntime() AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAppsRuntimeOutputReference
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
	LakebaseRuntime() AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationLakebaseRuntimeOutputReference
	LakebaseRuntimeInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WorkspaceApi() AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceApiOutputReference
	WorkspaceApiInput() interface{}
	WorkspaceUi() AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceUiOutputReference
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
	PutAccountApi(value *AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountApi)
	PutAccountDatabricksOne(value *AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountDatabricksOne)
	PutAccountUi(value *AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountUi)
	PutAppsRuntime(value *AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAppsRuntime)
	PutLakebaseRuntime(value *AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationLakebaseRuntime)
	PutWorkspaceApi(value *AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceApi)
	PutWorkspaceUi(value *AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceUi)
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

// The jsii proxy struct for AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference
type jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) AccountApi() AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountApiOutputReference {
	var returns AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountApiOutputReference
	_jsii_.Get(
		j,
		"accountApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) AccountApiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) AccountDatabricksOne() AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountDatabricksOneOutputReference {
	var returns AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountDatabricksOneOutputReference
	_jsii_.Get(
		j,
		"accountDatabricksOne",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) AccountDatabricksOneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountDatabricksOneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) AccountUi() AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountUiOutputReference {
	var returns AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountUiOutputReference
	_jsii_.Get(
		j,
		"accountUi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) AccountUiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountUiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) AllDestinations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) AllDestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allDestinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) AppsRuntime() AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAppsRuntimeOutputReference {
	var returns AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAppsRuntimeOutputReference
	_jsii_.Get(
		j,
		"appsRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) AppsRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appsRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) LakebaseRuntime() AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationLakebaseRuntimeOutputReference {
	var returns AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationLakebaseRuntimeOutputReference
	_jsii_.Get(
		j,
		"lakebaseRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) LakebaseRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lakebaseRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) WorkspaceApi() AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceApiOutputReference {
	var returns AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceApiOutputReference
	_jsii_.Get(
		j,
		"workspaceApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) WorkspaceApiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) WorkspaceUi() AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceUiOutputReference {
	var returns AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceUiOutputReference
	_jsii_.Get(
		j,
		"workspaceUi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) WorkspaceUiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceUiInput",
		&returns,
	)
	return returns
}


func NewAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.accountNetworkPolicy.AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference_Override(a AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.accountNetworkPolicy.AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference)SetAllDestinations(val interface{}) {
	if err := j.validateSetAllDestinationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allDestinations",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) PutAccountApi(value *AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountApi) {
	if err := a.validatePutAccountApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccountApi",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) PutAccountDatabricksOne(value *AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountDatabricksOne) {
	if err := a.validatePutAccountDatabricksOneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccountDatabricksOne",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) PutAccountUi(value *AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountUi) {
	if err := a.validatePutAccountUiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccountUi",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) PutAppsRuntime(value *AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAppsRuntime) {
	if err := a.validatePutAppsRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAppsRuntime",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) PutLakebaseRuntime(value *AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationLakebaseRuntime) {
	if err := a.validatePutLakebaseRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLakebaseRuntime",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) PutWorkspaceApi(value *AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceApi) {
	if err := a.validatePutWorkspaceApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorkspaceApi",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) PutWorkspaceUi(value *AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceUi) {
	if err := a.validatePutWorkspaceUiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorkspaceUi",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) ResetAccountApi() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountApi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) ResetAccountDatabricksOne() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountDatabricksOne",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) ResetAccountUi() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountUi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) ResetAllDestinations() {
	_jsii_.InvokeVoid(
		a,
		"resetAllDestinations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) ResetAppsRuntime() {
	_jsii_.InvokeVoid(
		a,
		"resetAppsRuntime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) ResetLakebaseRuntime() {
	_jsii_.InvokeVoid(
		a,
		"resetLakebaseRuntime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) ResetWorkspaceApi() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkspaceApi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) ResetWorkspaceUi() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkspaceUi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

