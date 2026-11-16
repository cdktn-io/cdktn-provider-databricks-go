// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountsettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabricksaccountsettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference interface {
	cdktn.ComplexObject
	Base() DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference
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
	FieldTitle() DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldTitleOutputReference
	FieldTitleInput() interface{}
	FieldValue() DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldValueOutputReference
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
	WidgetDescription() DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetDescriptionOutputReference
	WidgetDescriptionInput() interface{}
	WidgetTitle() DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference
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
	PutBase(value *DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBase)
	PutFieldTitle(value *DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldTitle)
	PutFieldValue(value *DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldValue)
	PutWidgetDescription(value *DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetDescription)
	PutWidgetTitle(value *DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitle)
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

// The jsii proxy struct for DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference
type jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) Base() DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference {
	var returns DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference
	_jsii_.Get(
		j,
		"base",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) BaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"baseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) FieldTitle() DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldTitleOutputReference {
	var returns DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldTitleOutputReference
	_jsii_.Get(
		j,
		"fieldTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) FieldTitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldTitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) FieldValue() DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldValueOutputReference {
	var returns DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldValueOutputReference
	_jsii_.Get(
		j,
		"fieldValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) FieldValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) WidgetDescription() DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetDescriptionOutputReference {
	var returns DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetDescriptionOutputReference
	_jsii_.Get(
		j,
		"widgetDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) WidgetDescriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) WidgetTitle() DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference {
	var returns DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference
	_jsii_.Get(
		j,
		"widgetTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) WidgetTitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetTitleInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountSettingV2.DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference_Override(d DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountSettingV2.DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) PutBase(value *DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBase) {
	if err := d.validatePutBaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBase",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) PutFieldTitle(value *DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldTitle) {
	if err := d.validatePutFieldTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFieldTitle",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) PutFieldValue(value *DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldValue) {
	if err := d.validatePutFieldValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFieldValue",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) PutWidgetDescription(value *DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetDescription) {
	if err := d.validatePutWidgetDescriptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWidgetDescription",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) PutWidgetTitle(value *DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitle) {
	if err := d.validatePutWidgetTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWidgetTitle",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ResetBase() {
	_jsii_.InvokeVoid(
		d,
		"resetBase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ResetFieldTitle() {
	_jsii_.InvokeVoid(
		d,
		"resetFieldTitle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ResetFieldValue() {
	_jsii_.InvokeVoid(
		d,
		"resetFieldValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ResetWidgetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetWidgetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ResetWidgetTitle() {
	_jsii_.InvokeVoid(
		d,
		"resetWidgetTitle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

