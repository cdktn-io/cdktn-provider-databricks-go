// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountsettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabricksaccountsettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference interface {
	cdktn.ComplexObject
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
	FontColor() DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseFontColorOutputReference
	FontColorInput() interface{}
	FontFamily() *string
	SetFontFamily(val *string)
	FontFamilyInput() *string
	FontSize() *float64
	SetFontSize(val *float64)
	FontSizeInput() *float64
	FontStyle() *string
	SetFontStyle(val *string)
	FontStyleInput() *string
	FontWeight() *float64
	SetFontWeight(val *float64)
	FontWeightInput() *float64
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
	TextCase() *string
	SetTextCase(val *string)
	TextCaseInput() *string
	TextDecoration() *string
	SetTextDecoration(val *string)
	TextDecorationInput() *string
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
	PutFontColor(value *DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseFontColor)
	ResetFontColor()
	ResetFontFamily()
	ResetFontSize()
	ResetFontStyle()
	ResetFontWeight()
	ResetTextCase()
	ResetTextDecoration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference
type jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) FontColor() DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseFontColorOutputReference {
	var returns DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseFontColorOutputReference
	_jsii_.Get(
		j,
		"fontColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) FontColorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fontColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) FontFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) FontFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) FontSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) FontSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) FontStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) FontStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) FontWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) FontWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) TextCase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) TextCaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textCaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) TextDecoration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textDecoration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) TextDecorationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textDecorationInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountSettingV2.DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference_Override(d DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountSettingV2.DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference)SetFontFamily(val *string) {
	if err := j.validateSetFontFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontFamily",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference)SetFontSize(val *float64) {
	if err := j.validateSetFontSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontSize",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference)SetFontStyle(val *string) {
	if err := j.validateSetFontStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontStyle",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference)SetFontWeight(val *float64) {
	if err := j.validateSetFontWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontWeight",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference)SetTextCase(val *string) {
	if err := j.validateSetTextCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textCase",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference)SetTextDecoration(val *string) {
	if err := j.validateSetTextDecorationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textDecoration",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) PutFontColor(value *DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseFontColor) {
	if err := d.validatePutFontColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFontColor",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) ResetFontColor() {
	_jsii_.InvokeVoid(
		d,
		"resetFontColor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) ResetFontFamily() {
	_jsii_.InvokeVoid(
		d,
		"resetFontFamily",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) ResetFontSize() {
	_jsii_.InvokeVoid(
		d,
		"resetFontSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) ResetFontStyle() {
	_jsii_.InvokeVoid(
		d,
		"resetFontStyle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) ResetFontWeight() {
	_jsii_.InvokeVoid(
		d,
		"resetFontWeight",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) ResetTextCase() {
	_jsii_.InvokeVoid(
		d,
		"resetTextCase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) ResetTextDecoration() {
	_jsii_.InvokeVoid(
		d,
		"resetTextDecoration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

