// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/workspacesettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference interface {
	cdktn.ComplexObject
	Color() *string
	SetColor(val *string)
	ColorInput() *string
	ColorScheme() *string
	SetColorScheme(val *string)
	ColorSchemeInput() *string
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
	CustomDiverging() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference
	CustomDivergingInput() interface{}
	CustomSequential() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomSequentialOutputReference
	CustomSequentialInput() interface{}
	// Experimental.
	Fqn() *string
	FromCategoricalIndex() *float64
	SetFromCategoricalIndex(val *float64)
	FromCategoricalIndexInput() *float64
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
	PutCustomDiverging(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDiverging)
	PutCustomSequential(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomSequential)
	ResetColor()
	ResetColorScheme()
	ResetCustomDiverging()
	ResetCustomSequential()
	ResetFromCategoricalIndex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference
type jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) Color() *string {
	var returns *string
	_jsii_.Get(
		j,
		"color",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ColorScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ColorSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) CustomDiverging() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference {
	var returns WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference
	_jsii_.Get(
		j,
		"customDiverging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) CustomDivergingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customDivergingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) CustomSequential() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomSequentialOutputReference {
	var returns WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomSequentialOutputReference
	_jsii_.Get(
		j,
		"customSequential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) CustomSequentialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customSequentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) FromCategoricalIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromCategoricalIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) FromCategoricalIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromCategoricalIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.workspaceSettingV2.WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference_Override(w WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.workspaceSettingV2.WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference)SetColor(val *string) {
	if err := j.validateSetColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"color",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference)SetColorScheme(val *string) {
	if err := j.validateSetColorSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"colorScheme",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference)SetFromCategoricalIndex(val *float64) {
	if err := j.validateSetFromCategoricalIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fromCategoricalIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) PutCustomDiverging(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDiverging) {
	if err := w.validatePutCustomDivergingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCustomDiverging",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) PutCustomSequential(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomSequential) {
	if err := w.validatePutCustomSequentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCustomSequential",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ResetColor() {
	_jsii_.InvokeVoid(
		w,
		"resetColor",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ResetColorScheme() {
	_jsii_.InvokeVoid(
		w,
		"resetColorScheme",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ResetCustomDiverging() {
	_jsii_.InvokeVoid(
		w,
		"resetCustomDiverging",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ResetCustomSequential() {
	_jsii_.InvokeVoid(
		w,
		"resetCustomSequential",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ResetFromCategoricalIndex() {
	_jsii_.InvokeVoid(
		w,
		"resetFromCategoricalIndex",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

