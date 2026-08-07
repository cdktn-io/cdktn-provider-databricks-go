// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksaccountnetworkpolicies/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference interface {
	cdktn.ComplexObject
	AccountApi() DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountApiOutputReference
	AccountApiInput() interface{}
	AccountDatabricksOne() DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountDatabricksOneOutputReference
	AccountDatabricksOneInput() interface{}
	AccountUi() DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountUiOutputReference
	AccountUiInput() interface{}
	AllDestinations() interface{}
	SetAllDestinations(val interface{})
	AllDestinationsInput() interface{}
	AppsRuntime() DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAppsRuntimeOutputReference
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
	LakebaseRuntime() DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationLakebaseRuntimeOutputReference
	LakebaseRuntimeInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WorkspaceApi() DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceApiOutputReference
	WorkspaceApiInput() interface{}
	WorkspaceUi() DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceUiOutputReference
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
	PutAccountApi(value *DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountApi)
	PutAccountDatabricksOne(value *DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountDatabricksOne)
	PutAccountUi(value *DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountUi)
	PutAppsRuntime(value *DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAppsRuntime)
	PutLakebaseRuntime(value *DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationLakebaseRuntime)
	PutWorkspaceApi(value *DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceApi)
	PutWorkspaceUi(value *DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceUi)
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

// The jsii proxy struct for DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference
type jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) AccountApi() DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountApiOutputReference {
	var returns DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountApiOutputReference
	_jsii_.Get(
		j,
		"accountApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) AccountApiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) AccountDatabricksOne() DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountDatabricksOneOutputReference {
	var returns DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountDatabricksOneOutputReference
	_jsii_.Get(
		j,
		"accountDatabricksOne",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) AccountDatabricksOneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountDatabricksOneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) AccountUi() DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountUiOutputReference {
	var returns DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountUiOutputReference
	_jsii_.Get(
		j,
		"accountUi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) AccountUiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountUiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) AllDestinations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) AllDestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allDestinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) AppsRuntime() DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAppsRuntimeOutputReference {
	var returns DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAppsRuntimeOutputReference
	_jsii_.Get(
		j,
		"appsRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) AppsRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appsRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) LakebaseRuntime() DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationLakebaseRuntimeOutputReference {
	var returns DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationLakebaseRuntimeOutputReference
	_jsii_.Get(
		j,
		"lakebaseRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) LakebaseRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lakebaseRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) WorkspaceApi() DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceApiOutputReference {
	var returns DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceApiOutputReference
	_jsii_.Get(
		j,
		"workspaceApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) WorkspaceApiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) WorkspaceUi() DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceUiOutputReference {
	var returns DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceUiOutputReference
	_jsii_.Get(
		j,
		"workspaceUi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) WorkspaceUiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceUiInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference_Override(d DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference)SetAllDestinations(val interface{}) {
	if err := j.validateSetAllDestinationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allDestinations",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) PutAccountApi(value *DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountApi) {
	if err := d.validatePutAccountApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAccountApi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) PutAccountDatabricksOne(value *DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountDatabricksOne) {
	if err := d.validatePutAccountDatabricksOneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAccountDatabricksOne",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) PutAccountUi(value *DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountUi) {
	if err := d.validatePutAccountUiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAccountUi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) PutAppsRuntime(value *DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAppsRuntime) {
	if err := d.validatePutAppsRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAppsRuntime",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) PutLakebaseRuntime(value *DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationLakebaseRuntime) {
	if err := d.validatePutLakebaseRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLakebaseRuntime",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) PutWorkspaceApi(value *DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceApi) {
	if err := d.validatePutWorkspaceApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWorkspaceApi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) PutWorkspaceUi(value *DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceUi) {
	if err := d.validatePutWorkspaceUiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWorkspaceUi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) ResetAccountApi() {
	_jsii_.InvokeVoid(
		d,
		"resetAccountApi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) ResetAccountDatabricksOne() {
	_jsii_.InvokeVoid(
		d,
		"resetAccountDatabricksOne",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) ResetAccountUi() {
	_jsii_.InvokeVoid(
		d,
		"resetAccountUi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) ResetAllDestinations() {
	_jsii_.InvokeVoid(
		d,
		"resetAllDestinations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) ResetAppsRuntime() {
	_jsii_.InvokeVoid(
		d,
		"resetAppsRuntime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) ResetLakebaseRuntime() {
	_jsii_.InvokeVoid(
		d,
		"resetLakebaseRuntime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) ResetWorkspaceApi() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkspaceApi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) ResetWorkspaceUi() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkspaceUi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

