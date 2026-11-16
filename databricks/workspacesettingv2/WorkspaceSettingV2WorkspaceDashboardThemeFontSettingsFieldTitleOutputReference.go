// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/workspacesettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference interface {
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
	FontColor() WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleFontColorOutputReference
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
	PutFontColor(value *WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleFontColor)
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

// The jsii proxy struct for WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference
type jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) FontColor() WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleFontColorOutputReference {
	var returns WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleFontColorOutputReference
	_jsii_.Get(
		j,
		"fontColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) FontColorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fontColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) FontFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) FontFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) FontSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) FontSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) FontStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) FontStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) FontWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) FontWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) TextCase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) TextCaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textCaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) TextDecoration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textDecoration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) TextDecorationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textDecorationInput",
		&returns,
	)
	return returns
}


func NewWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.workspaceSettingV2.WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference_Override(w WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.workspaceSettingV2.WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference)SetFontFamily(val *string) {
	if err := j.validateSetFontFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontFamily",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference)SetFontSize(val *float64) {
	if err := j.validateSetFontSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontSize",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference)SetFontStyle(val *string) {
	if err := j.validateSetFontStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontStyle",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference)SetFontWeight(val *float64) {
	if err := j.validateSetFontWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontWeight",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference)SetTextCase(val *string) {
	if err := j.validateSetTextCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textCase",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference)SetTextDecoration(val *string) {
	if err := j.validateSetTextDecorationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textDecoration",
		val,
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) PutFontColor(value *WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleFontColor) {
	if err := w.validatePutFontColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putFontColor",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) ResetFontColor() {
	_jsii_.InvokeVoid(
		w,
		"resetFontColor",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) ResetFontFamily() {
	_jsii_.InvokeVoid(
		w,
		"resetFontFamily",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) ResetFontSize() {
	_jsii_.InvokeVoid(
		w,
		"resetFontSize",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) ResetFontStyle() {
	_jsii_.InvokeVoid(
		w,
		"resetFontStyle",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) ResetFontWeight() {
	_jsii_.InvokeVoid(
		w,
		"resetFontWeight",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) ResetTextCase() {
	_jsii_.InvokeVoid(
		w,
		"resetTextCase",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) ResetTextDecoration() {
	_jsii_.InvokeVoid(
		w,
		"resetTextDecoration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := w.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

