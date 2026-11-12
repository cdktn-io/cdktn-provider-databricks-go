// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchalertconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabrickslakewatchalertconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference interface {
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

// The jsii proxy struct for DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference
type jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) ApiTokenUcSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiTokenUcSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) ApiTokenUcSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiTokenUcSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) DefaultIssueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultIssueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) DefaultIssueTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultIssueTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) InstanceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) InstanceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) ProjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) ProjectKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) UserEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) UserEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userEmailInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchAlertConfig.DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference_Override(d DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchAlertConfig.DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference)SetApiTokenUcSecret(val *string) {
	if err := j.validateSetApiTokenUcSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiTokenUcSecret",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference)SetDefaultIssueType(val *string) {
	if err := j.validateSetDefaultIssueTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultIssueType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference)SetInstanceUrl(val *string) {
	if err := j.validateSetInstanceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceUrl",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference)SetProjectKey(val *string) {
	if err := j.validateSetProjectKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectKey",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference)SetUserEmail(val *string) {
	if err := j.validateSetUserEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userEmail",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSecurityAlertExportConfigJiraOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

