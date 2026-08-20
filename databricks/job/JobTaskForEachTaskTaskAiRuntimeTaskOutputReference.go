// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/job/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JobTaskForEachTaskTaskAiRuntimeTaskOutputReference interface {
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
	Deployments() JobTaskForEachTaskTaskAiRuntimeTaskDeploymentsList
	DeploymentsInput() interface{}
	DockerImageUrl() *string
	SetDockerImageUrl(val *string)
	DockerImageUrlInput() *string
	Experiment() *string
	SetExperiment(val *string)
	ExperimentInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *JobTaskForEachTaskTaskAiRuntimeTask
	SetInternalValue(val *JobTaskForEachTaskTaskAiRuntimeTask)
	MlflowArtifactLocation() *string
	SetMlflowArtifactLocation(val *string)
	MlflowArtifactLocationInput() *string
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
	ResetDockerImageUrl()
	ResetMlflowArtifactLocation()
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

// The jsii proxy struct for JobTaskForEachTaskTaskAiRuntimeTaskOutputReference
type jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) CodeSourcePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSourcePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) CodeSourcePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSourcePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) Deployments() JobTaskForEachTaskTaskAiRuntimeTaskDeploymentsList {
	var returns JobTaskForEachTaskTaskAiRuntimeTaskDeploymentsList
	_jsii_.Get(
		j,
		"deployments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) DeploymentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) DockerImageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerImageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) DockerImageUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerImageUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) Experiment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"experiment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) ExperimentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"experimentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) InternalValue() *JobTaskForEachTaskTaskAiRuntimeTask {
	var returns *JobTaskForEachTaskTaskAiRuntimeTask
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) MlflowArtifactLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlflowArtifactLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) MlflowArtifactLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlflowArtifactLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) MlflowExperimentDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlflowExperimentDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) MlflowExperimentDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlflowExperimentDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) MlflowRun() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlflowRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) MlflowRunInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlflowRunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobTaskForEachTaskTaskAiRuntimeTaskOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) JobTaskForEachTaskTaskAiRuntimeTaskOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskForEachTaskTaskAiRuntimeTaskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskForEachTaskTaskAiRuntimeTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskForEachTaskTaskAiRuntimeTaskOutputReference_Override(j JobTaskForEachTaskTaskAiRuntimeTaskOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskForEachTaskTaskAiRuntimeTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference)SetCodeSourcePath(val *string) {
	if err := j.validateSetCodeSourcePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeSourcePath",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference)SetDockerImageUrl(val *string) {
	if err := j.validateSetDockerImageUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerImageUrl",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference)SetExperiment(val *string) {
	if err := j.validateSetExperimentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"experiment",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference)SetInternalValue(val *JobTaskForEachTaskTaskAiRuntimeTask) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference)SetMlflowArtifactLocation(val *string) {
	if err := j.validateSetMlflowArtifactLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mlflowArtifactLocation",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference)SetMlflowExperimentDirectory(val *string) {
	if err := j.validateSetMlflowExperimentDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mlflowExperimentDirectory",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference)SetMlflowRun(val *string) {
	if err := j.validateSetMlflowRunParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mlflowRun",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) PutDeployments(value interface{}) {
	if err := j.validatePutDeploymentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDeployments",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) ResetCodeSourcePath() {
	_jsii_.InvokeVoid(
		j,
		"resetCodeSourcePath",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) ResetDockerImageUrl() {
	_jsii_.InvokeVoid(
		j,
		"resetDockerImageUrl",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) ResetMlflowArtifactLocation() {
	_jsii_.InvokeVoid(
		j,
		"resetMlflowArtifactLocation",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) ResetMlflowExperimentDirectory() {
	_jsii_.InvokeVoid(
		j,
		"resetMlflowExperimentDirectory",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) ResetMlflowRun() {
	_jsii_.InvokeVoid(
		j,
		"resetMlflowRun",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiRuntimeTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

