// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/job/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JobTaskPipelineTaskOutputReference interface {
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
	InternalValue() *JobTaskPipelineTask
	SetInternalValue(val *JobTaskPipelineTask)
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	PipelineId() *string
	SetPipelineId(val *string)
	PipelineIdInput() *string
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
	ResetParameters()
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

// The jsii proxy struct for JobTaskPipelineTaskOutputReference
type jsiiProxy_JobTaskPipelineTaskOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) FullRefresh() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fullRefresh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) FullRefreshInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fullRefreshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) FullRefreshSelection() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fullRefreshSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) FullRefreshSelectionInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fullRefreshSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) InternalValue() *JobTaskPipelineTask {
	var returns *JobTaskPipelineTask
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) PipelineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) PipelineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) RefreshFlowSelection() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"refreshFlowSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) RefreshFlowSelectionInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"refreshFlowSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) RefreshSelection() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"refreshSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) RefreshSelectionInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"refreshSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) ResetCheckpointSelection() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resetCheckpointSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) ResetCheckpointSelectionInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resetCheckpointSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobTaskPipelineTaskOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) JobTaskPipelineTaskOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskPipelineTaskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskPipelineTaskOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskPipelineTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskPipelineTaskOutputReference_Override(j JobTaskPipelineTaskOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskPipelineTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference)SetFullRefresh(val interface{}) {
	if err := j.validateSetFullRefreshParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullRefresh",
		val,
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference)SetFullRefreshSelection(val *[]*string) {
	if err := j.validateSetFullRefreshSelectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullRefreshSelection",
		val,
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference)SetInternalValue(val *JobTaskPipelineTask) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference)SetPipelineId(val *string) {
	if err := j.validateSetPipelineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineId",
		val,
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference)SetRefreshFlowSelection(val *[]*string) {
	if err := j.validateSetRefreshFlowSelectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshFlowSelection",
		val,
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference)SetRefreshSelection(val *[]*string) {
	if err := j.validateSetRefreshSelectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshSelection",
		val,
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference)SetResetCheckpointSelection(val *[]*string) {
	if err := j.validateSetResetCheckpointSelectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resetCheckpointSelection",
		val,
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) ResetFullRefresh() {
	_jsii_.InvokeVoid(
		j,
		"resetFullRefresh",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) ResetFullRefreshSelection() {
	_jsii_.InvokeVoid(
		j,
		"resetFullRefreshSelection",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		j,
		"resetParameters",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) ResetRefreshFlowSelection() {
	_jsii_.InvokeVoid(
		j,
		"resetRefreshFlowSelection",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) ResetRefreshSelection() {
	_jsii_.InvokeVoid(
		j,
		"resetRefreshSelection",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) ResetResetCheckpointSelection() {
	_jsii_.InvokeVoid(
		j,
		"resetResetCheckpointSelection",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskPipelineTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

