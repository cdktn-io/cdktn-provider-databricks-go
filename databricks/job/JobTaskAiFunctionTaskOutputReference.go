// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/job/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JobTaskAiFunctionTaskOutputReference interface {
	cdktn.ComplexObject
	AiClassify() JobTaskAiFunctionTaskAiClassifyOutputReference
	AiClassifyInput() *JobTaskAiFunctionTaskAiClassify
	AiExtract() JobTaskAiFunctionTaskAiExtractOutputReference
	AiExtractInput() *JobTaskAiFunctionTaskAiExtract
	AiParseDocument() JobTaskAiFunctionTaskAiParseDocumentOutputReference
	AiParseDocumentInput() *JobTaskAiFunctionTaskAiParseDocument
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
	DeltaTable() JobTaskAiFunctionTaskDeltaTableOutputReference
	DeltaTableInput() *JobTaskAiFunctionTaskDeltaTable
	// Experimental.
	Fqn() *string
	InternalValue() *JobTaskAiFunctionTask
	SetInternalValue(val *JobTaskAiFunctionTask)
	Output() JobTaskAiFunctionTaskOutputOutputReference
	OutputInput() *JobTaskAiFunctionTaskOutput
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Volume() JobTaskAiFunctionTaskVolumeOutputReference
	VolumeInput() *JobTaskAiFunctionTaskVolume
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
	PutAiClassify(value *JobTaskAiFunctionTaskAiClassify)
	PutAiExtract(value *JobTaskAiFunctionTaskAiExtract)
	PutAiParseDocument(value *JobTaskAiFunctionTaskAiParseDocument)
	PutDeltaTable(value *JobTaskAiFunctionTaskDeltaTable)
	PutOutput(value *JobTaskAiFunctionTaskOutput)
	PutVolume(value *JobTaskAiFunctionTaskVolume)
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

// The jsii proxy struct for JobTaskAiFunctionTaskOutputReference
type jsiiProxy_JobTaskAiFunctionTaskOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) AiClassify() JobTaskAiFunctionTaskAiClassifyOutputReference {
	var returns JobTaskAiFunctionTaskAiClassifyOutputReference
	_jsii_.Get(
		j,
		"aiClassify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) AiClassifyInput() *JobTaskAiFunctionTaskAiClassify {
	var returns *JobTaskAiFunctionTaskAiClassify
	_jsii_.Get(
		j,
		"aiClassifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) AiExtract() JobTaskAiFunctionTaskAiExtractOutputReference {
	var returns JobTaskAiFunctionTaskAiExtractOutputReference
	_jsii_.Get(
		j,
		"aiExtract",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) AiExtractInput() *JobTaskAiFunctionTaskAiExtract {
	var returns *JobTaskAiFunctionTaskAiExtract
	_jsii_.Get(
		j,
		"aiExtractInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) AiParseDocument() JobTaskAiFunctionTaskAiParseDocumentOutputReference {
	var returns JobTaskAiFunctionTaskAiParseDocumentOutputReference
	_jsii_.Get(
		j,
		"aiParseDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) AiParseDocumentInput() *JobTaskAiFunctionTaskAiParseDocument {
	var returns *JobTaskAiFunctionTaskAiParseDocument
	_jsii_.Get(
		j,
		"aiParseDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) DeltaTable() JobTaskAiFunctionTaskDeltaTableOutputReference {
	var returns JobTaskAiFunctionTaskDeltaTableOutputReference
	_jsii_.Get(
		j,
		"deltaTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) DeltaTableInput() *JobTaskAiFunctionTaskDeltaTable {
	var returns *JobTaskAiFunctionTaskDeltaTable
	_jsii_.Get(
		j,
		"deltaTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) InternalValue() *JobTaskAiFunctionTask {
	var returns *JobTaskAiFunctionTask
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) Output() JobTaskAiFunctionTaskOutputOutputReference {
	var returns JobTaskAiFunctionTaskOutputOutputReference
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) OutputInput() *JobTaskAiFunctionTaskOutput {
	var returns *JobTaskAiFunctionTaskOutput
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) Volume() JobTaskAiFunctionTaskVolumeOutputReference {
	var returns JobTaskAiFunctionTaskVolumeOutputReference
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) VolumeInput() *JobTaskAiFunctionTaskVolume {
	var returns *JobTaskAiFunctionTaskVolume
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}


func NewJobTaskAiFunctionTaskOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) JobTaskAiFunctionTaskOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskAiFunctionTaskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskAiFunctionTaskOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskAiFunctionTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskAiFunctionTaskOutputReference_Override(j JobTaskAiFunctionTaskOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskAiFunctionTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference)SetInternalValue(val *JobTaskAiFunctionTask) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) PutAiClassify(value *JobTaskAiFunctionTaskAiClassify) {
	if err := j.validatePutAiClassifyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAiClassify",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) PutAiExtract(value *JobTaskAiFunctionTaskAiExtract) {
	if err := j.validatePutAiExtractParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAiExtract",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) PutAiParseDocument(value *JobTaskAiFunctionTaskAiParseDocument) {
	if err := j.validatePutAiParseDocumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAiParseDocument",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) PutDeltaTable(value *JobTaskAiFunctionTaskDeltaTable) {
	if err := j.validatePutDeltaTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDeltaTable",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) PutOutput(value *JobTaskAiFunctionTaskOutput) {
	if err := j.validatePutOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putOutput",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) PutVolume(value *JobTaskAiFunctionTaskVolume) {
	if err := j.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putVolume",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) ResetAiClassify() {
	_jsii_.InvokeVoid(
		j,
		"resetAiClassify",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) ResetAiExtract() {
	_jsii_.InvokeVoid(
		j,
		"resetAiExtract",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) ResetAiParseDocument() {
	_jsii_.InvokeVoid(
		j,
		"resetAiParseDocument",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) ResetDeltaTable() {
	_jsii_.InvokeVoid(
		j,
		"resetDeltaTable",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) ResetVolume() {
	_jsii_.InvokeVoid(
		j,
		"resetVolume",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

