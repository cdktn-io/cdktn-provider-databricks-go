// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/workspacesettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference interface {
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
	End() WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingEndOutputReference
	EndInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Mid() WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference
	MidInput() interface{}
	Start() WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingStartOutputReference
	StartInput() interface{}
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
	PutEnd(value *WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingEnd)
	PutMid(value *WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMid)
	PutStart(value *WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingStart)
	ResetEnd()
	ResetMid()
	ResetStart()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference
type jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) End() WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingEndOutputReference {
	var returns WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingEndOutputReference
	_jsii_.Get(
		j,
		"end",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) EndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) Mid() WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference {
	var returns WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMidOutputReference
	_jsii_.Get(
		j,
		"mid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) MidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"midInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) Start() WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingStartOutputReference {
	var returns WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingStartOutputReference
	_jsii_.Get(
		j,
		"start",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) StartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.workspaceSettingV2.WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference_Override(w WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.workspaceSettingV2.WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) PutEnd(value *WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingEnd) {
	if err := w.validatePutEndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putEnd",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) PutMid(value *WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingMid) {
	if err := w.validatePutMidParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMid",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) PutStart(value *WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingStart) {
	if err := w.validatePutStartParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putStart",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) ResetEnd() {
	_jsii_.InvokeVoid(
		w,
		"resetEnd",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) ResetMid() {
	_jsii_.InvokeVoid(
		w,
		"resetMid",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) ResetStart() {
	_jsii_.InvokeVoid(
		w,
		"resetStart",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

