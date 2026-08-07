// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksaccountnetworkpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference interface {
	cdktn.ComplexObject
	AccountApi() DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAccountApiOutputReference
	AccountApiInput() interface{}
	AccountDatabricksOne() DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAccountDatabricksOneOutputReference
	AccountDatabricksOneInput() interface{}
	AccountUi() DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAccountUiOutputReference
	AccountUiInput() interface{}
	AllDestinations() interface{}
	SetAllDestinations(val interface{})
	AllDestinationsInput() interface{}
	AppsRuntime() DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAppsRuntimeOutputReference
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
	LakebaseRuntime() DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationLakebaseRuntimeOutputReference
	LakebaseRuntimeInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WorkspaceApi() DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationWorkspaceApiOutputReference
	WorkspaceApiInput() interface{}
	WorkspaceUi() DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationWorkspaceUiOutputReference
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
	PutAccountApi(value *DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAccountApi)
	PutAccountDatabricksOne(value *DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAccountDatabricksOne)
	PutAccountUi(value *DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAccountUi)
	PutAppsRuntime(value *DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAppsRuntime)
	PutLakebaseRuntime(value *DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationLakebaseRuntime)
	PutWorkspaceApi(value *DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationWorkspaceApi)
	PutWorkspaceUi(value *DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationWorkspaceUi)
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

// The jsii proxy struct for DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference
type jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) AccountApi() DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAccountApiOutputReference {
	var returns DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAccountApiOutputReference
	_jsii_.Get(
		j,
		"accountApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) AccountApiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) AccountDatabricksOne() DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAccountDatabricksOneOutputReference {
	var returns DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAccountDatabricksOneOutputReference
	_jsii_.Get(
		j,
		"accountDatabricksOne",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) AccountDatabricksOneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountDatabricksOneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) AccountUi() DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAccountUiOutputReference {
	var returns DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAccountUiOutputReference
	_jsii_.Get(
		j,
		"accountUi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) AccountUiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountUiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) AllDestinations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) AllDestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allDestinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) AppsRuntime() DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAppsRuntimeOutputReference {
	var returns DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAppsRuntimeOutputReference
	_jsii_.Get(
		j,
		"appsRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) AppsRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appsRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) LakebaseRuntime() DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationLakebaseRuntimeOutputReference {
	var returns DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationLakebaseRuntimeOutputReference
	_jsii_.Get(
		j,
		"lakebaseRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) LakebaseRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lakebaseRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) WorkspaceApi() DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationWorkspaceApiOutputReference {
	var returns DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationWorkspaceApiOutputReference
	_jsii_.Get(
		j,
		"workspaceApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) WorkspaceApiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) WorkspaceUi() DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationWorkspaceUiOutputReference {
	var returns DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationWorkspaceUiOutputReference
	_jsii_.Get(
		j,
		"workspaceUi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) WorkspaceUiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceUiInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicy.DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference_Override(d DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicy.DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference)SetAllDestinations(val interface{}) {
	if err := j.validateSetAllDestinationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allDestinations",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) PutAccountApi(value *DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAccountApi) {
	if err := d.validatePutAccountApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAccountApi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) PutAccountDatabricksOne(value *DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAccountDatabricksOne) {
	if err := d.validatePutAccountDatabricksOneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAccountDatabricksOne",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) PutAccountUi(value *DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAccountUi) {
	if err := d.validatePutAccountUiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAccountUi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) PutAppsRuntime(value *DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationAppsRuntime) {
	if err := d.validatePutAppsRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAppsRuntime",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) PutLakebaseRuntime(value *DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationLakebaseRuntime) {
	if err := d.validatePutLakebaseRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLakebaseRuntime",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) PutWorkspaceApi(value *DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationWorkspaceApi) {
	if err := d.validatePutWorkspaceApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWorkspaceApi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) PutWorkspaceUi(value *DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationWorkspaceUi) {
	if err := d.validatePutWorkspaceUiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWorkspaceUi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) ResetAccountApi() {
	_jsii_.InvokeVoid(
		d,
		"resetAccountApi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) ResetAccountDatabricksOne() {
	_jsii_.InvokeVoid(
		d,
		"resetAccountDatabricksOne",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) ResetAccountUi() {
	_jsii_.InvokeVoid(
		d,
		"resetAccountUi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) ResetAllDestinations() {
	_jsii_.InvokeVoid(
		d,
		"resetAllDestinations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) ResetAppsRuntime() {
	_jsii_.InvokeVoid(
		d,
		"resetAppsRuntime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) ResetLakebaseRuntime() {
	_jsii_.InvokeVoid(
		d,
		"resetLakebaseRuntime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) ResetWorkspaceApi() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkspaceApi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) ResetWorkspaceUi() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkspaceUi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressPrivateAccessDenyRulesDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

