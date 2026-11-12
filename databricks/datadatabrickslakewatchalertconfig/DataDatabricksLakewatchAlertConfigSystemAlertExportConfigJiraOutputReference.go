// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchalertconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabrickslakewatchalertconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference interface {
	cdktn.ComplexObject
	ApiTokenUcSecret() *string
	SetApiTokenUcSecret(val *string)
	ApiTokenUcSecretInput() *string
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
	DefaultIssueType() *string
	SetDefaultIssueType(val *string)
	DefaultIssueTypeInput() *string
	// Experimental.
	Fqn() *string
	InstanceUrl() *string
	SetInstanceUrl(val *string)
	InstanceUrlInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProjectKey() *string
	SetProjectKey(val *string)
	ProjectKeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserEmail() *string
	SetUserEmail(val *string)
	UserEmailInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference
type jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) ApiTokenUcSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiTokenUcSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) ApiTokenUcSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiTokenUcSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) DefaultIssueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultIssueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) DefaultIssueTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultIssueTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) InstanceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) InstanceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) ProjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) ProjectKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) UserEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) UserEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userEmailInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchAlertConfig.DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference_Override(d DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchAlertConfig.DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference)SetApiTokenUcSecret(val *string) {
	if err := j.validateSetApiTokenUcSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiTokenUcSecret",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference)SetDefaultIssueType(val *string) {
	if err := j.validateSetDefaultIssueTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultIssueType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference)SetInstanceUrl(val *string) {
	if err := j.validateSetInstanceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceUrl",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference)SetProjectKey(val *string) {
	if err := j.validateSetProjectKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectKey",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference)SetUserEmail(val *string) {
	if err := j.validateSetUserEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userEmail",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

