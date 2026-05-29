// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/job/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JobTaskRunJobTaskPipelineParamsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	FullRefresh() interface{}
	SetFullRefresh(val interface{})
	FullRefreshInput() interface{}
	FullRefreshSelection() *[]*string
	SetFullRefreshSelection(val *[]*string)
	FullRefreshSelectionInput() *[]*string
	InternalValue() *JobTaskRunJobTaskPipelineParams
	SetInternalValue(val *JobTaskRunJobTaskPipelineParams)
	RefreshFlowSelection() *[]*string
	SetRefreshFlowSelection(val *[]*string)
	RefreshFlowSelectionInput() *[]*string
	RefreshSelection() *[]*string
	SetRefreshSelection(val *[]*string)
	RefreshSelectionInput() *[]*string
	ResetCheckpointSelection() *[]*string
	SetResetCheckpointSelection(val *[]*string)
	ResetCheckpointSelectionInput() *[]*string
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
	ResetFullRefresh()
	ResetFullRefreshSelection()
	ResetRefreshFlowSelection()
	ResetRefreshSelection()
	ResetResetCheckpointSelection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTaskRunJobTaskPipelineParamsOutputReference
type jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) FullRefresh() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fullRefresh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) FullRefreshInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fullRefreshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) FullRefreshSelection() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fullRefreshSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) FullRefreshSelectionInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fullRefreshSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) InternalValue() *JobTaskRunJobTaskPipelineParams {
	var returns *JobTaskRunJobTaskPipelineParams
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) RefreshFlowSelection() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"refreshFlowSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) RefreshFlowSelectionInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"refreshFlowSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) RefreshSelection() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"refreshSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) RefreshSelectionInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"refreshSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) ResetCheckpointSelection() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resetCheckpointSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) ResetCheckpointSelectionInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resetCheckpointSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobTaskRunJobTaskPipelineParamsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) JobTaskRunJobTaskPipelineParamsOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskRunJobTaskPipelineParamsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskRunJobTaskPipelineParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskRunJobTaskPipelineParamsOutputReference_Override(j JobTaskRunJobTaskPipelineParamsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskRunJobTaskPipelineParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference)SetFullRefresh(val interface{}) {
	if err := j.validateSetFullRefreshParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullRefresh",
		val,
	)
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference)SetFullRefreshSelection(val *[]*string) {
	if err := j.validateSetFullRefreshSelectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullRefreshSelection",
		val,
	)
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference)SetInternalValue(val *JobTaskRunJobTaskPipelineParams) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference)SetRefreshFlowSelection(val *[]*string) {
	if err := j.validateSetRefreshFlowSelectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshFlowSelection",
		val,
	)
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference)SetRefreshSelection(val *[]*string) {
	if err := j.validateSetRefreshSelectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshSelection",
		val,
	)
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference)SetResetCheckpointSelection(val *[]*string) {
	if err := j.validateSetResetCheckpointSelectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resetCheckpointSelection",
		val,
	)
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) ResetFullRefresh() {
	_jsii_.InvokeVoid(
		j,
		"resetFullRefresh",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) ResetFullRefreshSelection() {
	_jsii_.InvokeVoid(
		j,
		"resetFullRefreshSelection",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) ResetRefreshFlowSelection() {
	_jsii_.InvokeVoid(
		j,
		"resetRefreshFlowSelection",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) ResetRefreshSelection() {
	_jsii_.InvokeVoid(
		j,
		"resetRefreshSelection",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) ResetResetCheckpointSelection() {
	_jsii_.InvokeVoid(
		j,
		"resetResetCheckpointSelection",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := j.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskRunJobTaskPipelineParamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

