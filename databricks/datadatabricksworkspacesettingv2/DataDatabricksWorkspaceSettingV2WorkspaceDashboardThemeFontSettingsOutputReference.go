// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspacesettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabricksworkspacesettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference interface {
	cdktn.ComplexObject
	Base() DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsBaseOutputReference
	BaseInput() interface{}
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
	FieldTitle() DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference
	FieldTitleInput() interface{}
	FieldValue() DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldValueOutputReference
	FieldValueInput() interface{}
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
	WidgetDescription() DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsWidgetDescriptionOutputReference
	WidgetDescriptionInput() interface{}
	WidgetTitle() DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference
	WidgetTitleInput() interface{}
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
	PutBase(value *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsBase)
	PutFieldTitle(value *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitle)
	PutFieldValue(value *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldValue)
	PutWidgetDescription(value *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsWidgetDescription)
	PutWidgetTitle(value *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitle)
	ResetBase()
	ResetFieldTitle()
	ResetFieldValue()
	ResetWidgetDescription()
	ResetWidgetTitle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference
type jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) Base() DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsBaseOutputReference {
	var returns DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsBaseOutputReference
	_jsii_.Get(
		j,
		"base",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) BaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"baseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) FieldTitle() DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference {
	var returns DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference
	_jsii_.Get(
		j,
		"fieldTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) FieldTitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldTitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) FieldValue() DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldValueOutputReference {
	var returns DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldValueOutputReference
	_jsii_.Get(
		j,
		"fieldValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) FieldValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) WidgetDescription() DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsWidgetDescriptionOutputReference {
	var returns DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsWidgetDescriptionOutputReference
	_jsii_.Get(
		j,
		"widgetDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) WidgetDescriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) WidgetTitle() DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference {
	var returns DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference
	_jsii_.Get(
		j,
		"widgetTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) WidgetTitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetTitleInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksWorkspaceSettingV2.DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference_Override(d DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksWorkspaceSettingV2.DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) PutBase(value *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsBase) {
	if err := d.validatePutBaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBase",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) PutFieldTitle(value *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitle) {
	if err := d.validatePutFieldTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFieldTitle",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) PutFieldValue(value *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldValue) {
	if err := d.validatePutFieldValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFieldValue",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) PutWidgetDescription(value *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsWidgetDescription) {
	if err := d.validatePutWidgetDescriptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWidgetDescription",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) PutWidgetTitle(value *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitle) {
	if err := d.validatePutWidgetTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWidgetTitle",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) ResetBase() {
	_jsii_.InvokeVoid(
		d,
		"resetBase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) ResetFieldTitle() {
	_jsii_.InvokeVoid(
		d,
		"resetFieldTitle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) ResetFieldValue() {
	_jsii_.InvokeVoid(
		d,
		"resetFieldValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) ResetWidgetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetWidgetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) ResetWidgetTitle() {
	_jsii_.InvokeVoid(
		d,
		"resetWidgetTitle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

