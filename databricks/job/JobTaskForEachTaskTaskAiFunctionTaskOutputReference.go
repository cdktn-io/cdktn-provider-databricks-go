// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/job/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JobTaskForEachTaskTaskAiFunctionTaskOutputReference interface {
	cdktn.ComplexObject
	AiClassify() JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference
	AiClassifyInput() *JobTaskForEachTaskTaskAiFunctionTaskAiClassify
	AiExtract() JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference
	AiExtractInput() *JobTaskForEachTaskTaskAiFunctionTaskAiExtract
	AiParseDocument() JobTaskForEachTaskTaskAiFunctionTaskAiParseDocumentOutputReference
	AiParseDocumentInput() *JobTaskForEachTaskTaskAiFunctionTaskAiParseDocument
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
	DeltaTable() JobTaskForEachTaskTaskAiFunctionTaskDeltaTableOutputReference
	DeltaTableInput() *JobTaskForEachTaskTaskAiFunctionTaskDeltaTable
	// Experimental.
	Fqn() *string
	InternalValue() *JobTaskForEachTaskTaskAiFunctionTask
	SetInternalValue(val *JobTaskForEachTaskTaskAiFunctionTask)
	Output() JobTaskForEachTaskTaskAiFunctionTaskOutputOutputReference
	OutputInput() *JobTaskForEachTaskTaskAiFunctionTaskOutput
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Volume() JobTaskForEachTaskTaskAiFunctionTaskVolumeOutputReference
	VolumeInput() *JobTaskForEachTaskTaskAiFunctionTaskVolume
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
	PutAiClassify(value *JobTaskForEachTaskTaskAiFunctionTaskAiClassify)
	PutAiExtract(value *JobTaskForEachTaskTaskAiFunctionTaskAiExtract)
	PutAiParseDocument(value *JobTaskForEachTaskTaskAiFunctionTaskAiParseDocument)
	PutDeltaTable(value *JobTaskForEachTaskTaskAiFunctionTaskDeltaTable)
	PutOutput(value *JobTaskForEachTaskTaskAiFunctionTaskOutput)
	PutVolume(value *JobTaskForEachTaskTaskAiFunctionTaskVolume)
	ResetAiClassify()
	ResetAiExtract()
	ResetAiParseDocument()
	ResetDeltaTable()
	ResetVolume()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTaskForEachTaskTaskAiFunctionTaskOutputReference
type jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) AiClassify() JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference {
	var returns JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference
	_jsii_.Get(
		j,
		"aiClassify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) AiClassifyInput() *JobTaskForEachTaskTaskAiFunctionTaskAiClassify {
	var returns *JobTaskForEachTaskTaskAiFunctionTaskAiClassify
	_jsii_.Get(
		j,
		"aiClassifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) AiExtract() JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference {
	var returns JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference
	_jsii_.Get(
		j,
		"aiExtract",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) AiExtractInput() *JobTaskForEachTaskTaskAiFunctionTaskAiExtract {
	var returns *JobTaskForEachTaskTaskAiFunctionTaskAiExtract
	_jsii_.Get(
		j,
		"aiExtractInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) AiParseDocument() JobTaskForEachTaskTaskAiFunctionTaskAiParseDocumentOutputReference {
	var returns JobTaskForEachTaskTaskAiFunctionTaskAiParseDocumentOutputReference
	_jsii_.Get(
		j,
		"aiParseDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) AiParseDocumentInput() *JobTaskForEachTaskTaskAiFunctionTaskAiParseDocument {
	var returns *JobTaskForEachTaskTaskAiFunctionTaskAiParseDocument
	_jsii_.Get(
		j,
		"aiParseDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) DeltaTable() JobTaskForEachTaskTaskAiFunctionTaskDeltaTableOutputReference {
	var returns JobTaskForEachTaskTaskAiFunctionTaskDeltaTableOutputReference
	_jsii_.Get(
		j,
		"deltaTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) DeltaTableInput() *JobTaskForEachTaskTaskAiFunctionTaskDeltaTable {
	var returns *JobTaskForEachTaskTaskAiFunctionTaskDeltaTable
	_jsii_.Get(
		j,
		"deltaTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) InternalValue() *JobTaskForEachTaskTaskAiFunctionTask {
	var returns *JobTaskForEachTaskTaskAiFunctionTask
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) Output() JobTaskForEachTaskTaskAiFunctionTaskOutputOutputReference {
	var returns JobTaskForEachTaskTaskAiFunctionTaskOutputOutputReference
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) OutputInput() *JobTaskForEachTaskTaskAiFunctionTaskOutput {
	var returns *JobTaskForEachTaskTaskAiFunctionTaskOutput
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) Volume() JobTaskForEachTaskTaskAiFunctionTaskVolumeOutputReference {
	var returns JobTaskForEachTaskTaskAiFunctionTaskVolumeOutputReference
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) VolumeInput() *JobTaskForEachTaskTaskAiFunctionTaskVolume {
	var returns *JobTaskForEachTaskTaskAiFunctionTaskVolume
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}


func NewJobTaskForEachTaskTaskAiFunctionTaskOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) JobTaskForEachTaskTaskAiFunctionTaskOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskForEachTaskTaskAiFunctionTaskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskForEachTaskTaskAiFunctionTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskForEachTaskTaskAiFunctionTaskOutputReference_Override(j JobTaskForEachTaskTaskAiFunctionTaskOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskForEachTaskTaskAiFunctionTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference)SetInternalValue(val *JobTaskForEachTaskTaskAiFunctionTask) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) PutAiClassify(value *JobTaskForEachTaskTaskAiFunctionTaskAiClassify) {
	if err := j.validatePutAiClassifyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAiClassify",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) PutAiExtract(value *JobTaskForEachTaskTaskAiFunctionTaskAiExtract) {
	if err := j.validatePutAiExtractParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAiExtract",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) PutAiParseDocument(value *JobTaskForEachTaskTaskAiFunctionTaskAiParseDocument) {
	if err := j.validatePutAiParseDocumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAiParseDocument",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) PutDeltaTable(value *JobTaskForEachTaskTaskAiFunctionTaskDeltaTable) {
	if err := j.validatePutDeltaTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDeltaTable",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) PutOutput(value *JobTaskForEachTaskTaskAiFunctionTaskOutput) {
	if err := j.validatePutOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putOutput",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) PutVolume(value *JobTaskForEachTaskTaskAiFunctionTaskVolume) {
	if err := j.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putVolume",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) ResetAiClassify() {
	_jsii_.InvokeVoid(
		j,
		"resetAiClassify",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) ResetAiExtract() {
	_jsii_.InvokeVoid(
		j,
		"resetAiExtract",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) ResetAiParseDocument() {
	_jsii_.InvokeVoid(
		j,
		"resetAiParseDocument",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) ResetDeltaTable() {
	_jsii_.InvokeVoid(
		j,
		"resetDeltaTable",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) ResetVolume() {
	_jsii_.InvokeVoid(
		j,
		"resetVolume",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

