// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountsettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/accountsettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference interface {
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
	FontColor() AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference
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
	PutFontColor(value *AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleFontColor)
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

// The jsii proxy struct for AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference
type jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) FontColor() AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference {
	var returns AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference
	_jsii_.Get(
		j,
		"fontColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) FontColorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fontColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) FontFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) FontFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) FontSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) FontSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) FontStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) FontStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) FontWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) FontWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) TextCase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) TextCaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textCaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) TextDecoration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textDecoration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) TextDecorationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textDecorationInput",
		&returns,
	)
	return returns
}


func NewAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference {
	_init_.Initialize()

	if err := validateNewAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.accountSettingV2.AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference_Override(a AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.accountSettingV2.AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference)SetFontFamily(val *string) {
	if err := j.validateSetFontFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontFamily",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference)SetFontSize(val *float64) {
	if err := j.validateSetFontSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontSize",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference)SetFontStyle(val *string) {
	if err := j.validateSetFontStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontStyle",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference)SetFontWeight(val *float64) {
	if err := j.validateSetFontWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontWeight",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference)SetTextCase(val *string) {
	if err := j.validateSetTextCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textCase",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference)SetTextDecoration(val *string) {
	if err := j.validateSetTextDecorationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textDecoration",
		val,
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) PutFontColor(value *AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleFontColor) {
	if err := a.validatePutFontColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFontColor",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) ResetFontColor() {
	_jsii_.InvokeVoid(
		a,
		"resetFontColor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) ResetFontFamily() {
	_jsii_.InvokeVoid(
		a,
		"resetFontFamily",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) ResetFontSize() {
	_jsii_.InvokeVoid(
		a,
		"resetFontSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) ResetFontStyle() {
	_jsii_.InvokeVoid(
		a,
		"resetFontStyle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) ResetFontWeight() {
	_jsii_.InvokeVoid(
		a,
		"resetFontWeight",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) ResetTextCase() {
	_jsii_.InvokeVoid(
		a,
		"resetTextCase",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) ResetTextDecoration() {
	_jsii_.InvokeVoid(
		a,
		"resetTextDecoration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

