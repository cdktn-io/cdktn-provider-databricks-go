// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabrickspostgresproject/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksPostgresProjectSpecOutputReference interface {
	cdktn.ComplexObject
	BudgetPolicyId() *string
	SetBudgetPolicyId(val *string)
	BudgetPolicyIdInput() *string
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
	CustomTags() DataDatabricksPostgresProjectSpecCustomTagsList
	CustomTagsInput() interface{}
	DefaultBranch() *string
	SetDefaultBranch(val *string)
	DefaultBranchInput() *string
	DefaultEndpointSettings() DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference
	DefaultEndpointSettingsInput() interface{}
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EnablePgNativeLogin() interface{}
	SetEnablePgNativeLogin(val interface{})
	EnablePgNativeLoginInput() interface{}
	// Experimental.
	Fqn() *string
	HistoryRetentionDuration() *string
	SetHistoryRetentionDuration(val *string)
	HistoryRetentionDurationInput() *string
	InternalValue() *DataDatabricksPostgresProjectSpec
	SetInternalValue(val *DataDatabricksPostgresProjectSpec)
	PgVersion() *float64
	SetPgVersion(val *float64)
	PgVersionInput() *float64
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
	PutCustomTags(value interface{})
	PutDefaultEndpointSettings(value *DataDatabricksPostgresProjectSpecDefaultEndpointSettings)
	ResetBudgetPolicyId()
	ResetCustomTags()
	ResetDefaultBranch()
	ResetDefaultEndpointSettings()
	ResetDisplayName()
	ResetEnablePgNativeLogin()
	ResetHistoryRetentionDuration()
	ResetPgVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksPostgresProjectSpecOutputReference
type jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) BudgetPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) BudgetPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) CustomTags() DataDatabricksPostgresProjectSpecCustomTagsList {
	var returns DataDatabricksPostgresProjectSpecCustomTagsList
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) CustomTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) DefaultBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) DefaultBranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) DefaultEndpointSettings() DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference {
	var returns DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference
	_jsii_.Get(
		j,
		"defaultEndpointSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) DefaultEndpointSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultEndpointSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) EnablePgNativeLogin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePgNativeLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) EnablePgNativeLoginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePgNativeLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) HistoryRetentionDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"historyRetentionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) HistoryRetentionDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"historyRetentionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) InternalValue() *DataDatabricksPostgresProjectSpec {
	var returns *DataDatabricksPostgresProjectSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) PgVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pgVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) PgVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pgVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksPostgresProjectSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksPostgresProjectSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksPostgresProjectSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresProject.DataDatabricksPostgresProjectSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksPostgresProjectSpecOutputReference_Override(d DataDatabricksPostgresProjectSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresProject.DataDatabricksPostgresProjectSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference)SetBudgetPolicyId(val *string) {
	if err := j.validateSetBudgetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"budgetPolicyId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference)SetDefaultBranch(val *string) {
	if err := j.validateSetDefaultBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultBranch",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference)SetEnablePgNativeLogin(val interface{}) {
	if err := j.validateSetEnablePgNativeLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePgNativeLogin",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference)SetHistoryRetentionDuration(val *string) {
	if err := j.validateSetHistoryRetentionDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"historyRetentionDuration",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference)SetInternalValue(val *DataDatabricksPostgresProjectSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference)SetPgVersion(val *float64) {
	if err := j.validateSetPgVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pgVersion",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) PutCustomTags(value interface{}) {
	if err := d.validatePutCustomTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) PutDefaultEndpointSettings(value *DataDatabricksPostgresProjectSpecDefaultEndpointSettings) {
	if err := d.validatePutDefaultEndpointSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDefaultEndpointSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) ResetBudgetPolicyId() {
	_jsii_.InvokeVoid(
		d,
		"resetBudgetPolicyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) ResetCustomTags() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) ResetDefaultBranch() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultBranch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) ResetDefaultEndpointSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultEndpointSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) ResetEnablePgNativeLogin() {
	_jsii_.InvokeVoid(
		d,
		"resetEnablePgNativeLogin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) ResetHistoryRetentionDuration() {
	_jsii_.InvokeVoid(
		d,
		"resetHistoryRetentionDuration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) ResetPgVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetPgVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

