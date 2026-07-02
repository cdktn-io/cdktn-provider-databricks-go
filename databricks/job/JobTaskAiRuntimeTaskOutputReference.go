// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/job/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JobTaskAiRuntimeTaskOutputReference interface {
	cdktn.ComplexObject
	CodeSourcePath() *string
	SetCodeSourcePath(val *string)
	CodeSourcePathInput() *string
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
	Deployments() JobTaskAiRuntimeTaskDeploymentsList
	DeploymentsInput() interface{}
	Experiment() *string
	SetExperiment(val *string)
	ExperimentInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *JobTaskAiRuntimeTask
	SetInternalValue(val *JobTaskAiRuntimeTask)
	MlflowExperimentDirectory() *string
	SetMlflowExperimentDirectory(val *string)
	MlflowExperimentDirectoryInput() *string
	MlflowRun() *string
	SetMlflowRun(val *string)
	MlflowRunInput() *string
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
	PutDeployments(value interface{})
	ResetCodeSourcePath()
	ResetMlflowExperimentDirectory()
	ResetMlflowRun()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTaskAiRuntimeTaskOutputReference
type jsiiProxy_JobTaskAiRuntimeTaskOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) CodeSourcePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSourcePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) CodeSourcePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSourcePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) Deployments() JobTaskAiRuntimeTaskDeploymentsList {
	var returns JobTaskAiRuntimeTaskDeploymentsList
	_jsii_.Get(
		j,
		"deployments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) DeploymentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) Experiment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"experiment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) ExperimentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"experimentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) InternalValue() *JobTaskAiRuntimeTask {
	var returns *JobTaskAiRuntimeTask
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) MlflowExperimentDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlflowExperimentDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) MlflowExperimentDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlflowExperimentDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) MlflowRun() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlflowRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) MlflowRunInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlflowRunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobTaskAiRuntimeTaskOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) JobTaskAiRuntimeTaskOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskAiRuntimeTaskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskAiRuntimeTaskOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskAiRuntimeTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskAiRuntimeTaskOutputReference_Override(j JobTaskAiRuntimeTaskOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskAiRuntimeTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference)SetCodeSourcePath(val *string) {
	if err := j.validateSetCodeSourcePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeSourcePath",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference)SetExperiment(val *string) {
	if err := j.validateSetExperimentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"experiment",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference)SetInternalValue(val *JobTaskAiRuntimeTask) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference)SetMlflowExperimentDirectory(val *string) {
	if err := j.validateSetMlflowExperimentDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mlflowExperimentDirectory",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference)SetMlflowRun(val *string) {
	if err := j.validateSetMlflowRunParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mlflowRun",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) PutDeployments(value interface{}) {
	if err := j.validatePutDeploymentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDeployments",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) ResetCodeSourcePath() {
	_jsii_.InvokeVoid(
		j,
		"resetCodeSourcePath",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) ResetMlflowExperimentDirectory() {
	_jsii_.InvokeVoid(
		j,
		"resetMlflowExperimentDirectory",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) ResetMlflowRun() {
	_jsii_.InvokeVoid(
		j,
		"resetMlflowRun",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskAiRuntimeTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

