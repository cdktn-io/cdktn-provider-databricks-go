// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/job/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference interface {
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
	EnableConfidenceScores() interface{}
	SetEnableConfidenceScores(val interface{})
	EnableConfidenceScoresInput() interface{}
	EnableRationales() interface{}
	SetEnableRationales(val interface{})
	EnableRationalesInput() interface{}
	// Experimental.
	Fqn() *string
	InlineLabels() JobTaskForEachTaskTaskAiFunctionTaskAiClassifyInlineLabelsOutputReference
	InlineLabelsInput() *JobTaskForEachTaskTaskAiFunctionTaskAiClassifyInlineLabels
	Instructions() *string
	SetInstructions(val *string)
	InstructionsInput() *string
	InternalValue() *JobTaskForEachTaskTaskAiFunctionTaskAiClassify
	SetInternalValue(val *JobTaskForEachTaskTaskAiFunctionTaskAiClassify)
	Multilabel() interface{}
	SetMultilabel(val interface{})
	MultilabelInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutInlineLabels(value *JobTaskForEachTaskTaskAiFunctionTaskAiClassifyInlineLabels)
	ResetEnableConfidenceScores()
	ResetEnableRationales()
	ResetInlineLabels()
	ResetInstructions()
	ResetMultilabel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference
type jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) EnableConfidenceScores() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConfidenceScores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) EnableConfidenceScoresInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConfidenceScoresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) EnableRationales() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRationales",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) EnableRationalesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRationalesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) InlineLabels() JobTaskForEachTaskTaskAiFunctionTaskAiClassifyInlineLabelsOutputReference {
	var returns JobTaskForEachTaskTaskAiFunctionTaskAiClassifyInlineLabelsOutputReference
	_jsii_.Get(
		j,
		"inlineLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) InlineLabelsInput() *JobTaskForEachTaskTaskAiFunctionTaskAiClassifyInlineLabels {
	var returns *JobTaskForEachTaskTaskAiFunctionTaskAiClassifyInlineLabels
	_jsii_.Get(
		j,
		"inlineLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) Instructions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instructions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) InstructionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instructionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) InternalValue() *JobTaskForEachTaskTaskAiFunctionTaskAiClassify {
	var returns *JobTaskForEachTaskTaskAiFunctionTaskAiClassify
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) Multilabel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multilabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) MultilabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multilabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewJobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference_Override(j JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference)SetEnableConfidenceScores(val interface{}) {
	if err := j.validateSetEnableConfidenceScoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConfidenceScores",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference)SetEnableRationales(val interface{}) {
	if err := j.validateSetEnableRationalesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableRationales",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference)SetInstructions(val *string) {
	if err := j.validateSetInstructionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instructions",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference)SetInternalValue(val *JobTaskForEachTaskTaskAiFunctionTaskAiClassify) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference)SetMultilabel(val interface{}) {
	if err := j.validateSetMultilabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multilabel",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) PutInlineLabels(value *JobTaskForEachTaskTaskAiFunctionTaskAiClassifyInlineLabels) {
	if err := j.validatePutInlineLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putInlineLabels",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) ResetEnableConfidenceScores() {
	_jsii_.InvokeVoid(
		j,
		"resetEnableConfidenceScores",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) ResetEnableRationales() {
	_jsii_.InvokeVoid(
		j,
		"resetEnableRationales",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) ResetInlineLabels() {
	_jsii_.InvokeVoid(
		j,
		"resetInlineLabels",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) ResetInstructions() {
	_jsii_.InvokeVoid(
		j,
		"resetInstructions",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) ResetMultilabel() {
	_jsii_.InvokeVoid(
		j,
		"resetMultilabel",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiClassifyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

