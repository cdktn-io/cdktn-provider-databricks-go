// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/workspacesettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference interface {
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
	End() WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingEndOutputReference
	EndInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Mid() WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference
	MidInput() interface{}
	Start() WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingStartOutputReference
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
	PutEnd(value *WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingEnd)
	PutMid(value *WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMid)
	PutStart(value *WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingStart)
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

// The jsii proxy struct for WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference
type jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) End() WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingEndOutputReference {
	var returns WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingEndOutputReference
	_jsii_.Get(
		j,
		"end",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) EndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) Mid() WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference {
	var returns WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMidOutputReference
	_jsii_.Get(
		j,
		"mid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) MidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"midInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) Start() WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingStartOutputReference {
	var returns WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingStartOutputReference
	_jsii_.Get(
		j,
		"start",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) StartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.workspaceSettingV2.WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference_Override(w WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.workspaceSettingV2.WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) PutEnd(value *WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingEnd) {
	if err := w.validatePutEndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putEnd",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) PutMid(value *WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingMid) {
	if err := w.validatePutMidParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMid",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) PutStart(value *WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingStart) {
	if err := w.validatePutStartParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putStart",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) ResetEnd() {
	_jsii_.InvokeVoid(
		w,
		"resetEnd",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) ResetMid() {
	_jsii_.InvokeVoid(
		w,
		"resetMid",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) ResetStart() {
	_jsii_.InvokeVoid(
		w,
		"resetStart",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

