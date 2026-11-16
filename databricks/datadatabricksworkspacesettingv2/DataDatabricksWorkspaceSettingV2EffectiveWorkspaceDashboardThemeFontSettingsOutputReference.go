// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspacesettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabricksworkspacesettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference interface {
	cdktn.ComplexObject
	Base() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference
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
	FieldTitle() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldTitleOutputReference
	FieldTitleInput() interface{}
	FieldValue() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldValueOutputReference
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
	WidgetDescription() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetDescriptionOutputReference
	WidgetDescriptionInput() interface{}
	WidgetTitle() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference
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
	PutBase(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBase)
	PutFieldTitle(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldTitle)
	PutFieldValue(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldValue)
	PutWidgetDescription(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetDescription)
	PutWidgetTitle(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitle)
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

// The jsii proxy struct for DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference
type jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) Base() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference
	_jsii_.Get(
		j,
		"base",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) BaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"baseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) FieldTitle() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldTitleOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldTitleOutputReference
	_jsii_.Get(
		j,
		"fieldTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) FieldTitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldTitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) FieldValue() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldValueOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldValueOutputReference
	_jsii_.Get(
		j,
		"fieldValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) FieldValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) WidgetDescription() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetDescriptionOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetDescriptionOutputReference
	_jsii_.Get(
		j,
		"widgetDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) WidgetDescriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) WidgetTitle() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference
	_jsii_.Get(
		j,
		"widgetTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) WidgetTitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetTitleInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksWorkspaceSettingV2.DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference_Override(d DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksWorkspaceSettingV2.DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) PutBase(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBase) {
	if err := d.validatePutBaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBase",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) PutFieldTitle(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldTitle) {
	if err := d.validatePutFieldTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFieldTitle",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) PutFieldValue(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldValue) {
	if err := d.validatePutFieldValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFieldValue",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) PutWidgetDescription(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetDescription) {
	if err := d.validatePutWidgetDescriptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWidgetDescription",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) PutWidgetTitle(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitle) {
	if err := d.validatePutWidgetTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWidgetTitle",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ResetBase() {
	_jsii_.InvokeVoid(
		d,
		"resetBase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ResetFieldTitle() {
	_jsii_.InvokeVoid(
		d,
		"resetFieldTitle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ResetFieldValue() {
	_jsii_.InvokeVoid(
		d,
		"resetFieldValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ResetWidgetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetWidgetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ResetWidgetTitle() {
	_jsii_.InvokeVoid(
		d,
		"resetWidgetTitle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

