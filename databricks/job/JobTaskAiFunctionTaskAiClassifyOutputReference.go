// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/job/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JobTaskAiFunctionTaskAiClassifyOutputReference interface {
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
	InlineLabels() JobTaskAiFunctionTaskAiClassifyInlineLabelsOutputReference
	InlineLabelsInput() *JobTaskAiFunctionTaskAiClassifyInlineLabels
	Instructions() *string
	SetInstructions(val *string)
	InstructionsInput() *string
	InternalValue() *JobTaskAiFunctionTaskAiClassify
	SetInternalValue(val *JobTaskAiFunctionTaskAiClassify)
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
	PutInlineLabels(value *JobTaskAiFunctionTaskAiClassifyInlineLabels)
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

// The jsii proxy struct for JobTaskAiFunctionTaskAiClassifyOutputReference
type jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) EnableConfidenceScores() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConfidenceScores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) EnableConfidenceScoresInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConfidenceScoresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) EnableRationales() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRationales",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) EnableRationalesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRationalesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) InlineLabels() JobTaskAiFunctionTaskAiClassifyInlineLabelsOutputReference {
	var returns JobTaskAiFunctionTaskAiClassifyInlineLabelsOutputReference
	_jsii_.Get(
		j,
		"inlineLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) InlineLabelsInput() *JobTaskAiFunctionTaskAiClassifyInlineLabels {
	var returns *JobTaskAiFunctionTaskAiClassifyInlineLabels
	_jsii_.Get(
		j,
		"inlineLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) Instructions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instructions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) InstructionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instructionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) InternalValue() *JobTaskAiFunctionTaskAiClassify {
	var returns *JobTaskAiFunctionTaskAiClassify
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) Multilabel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multilabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) MultilabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multilabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewJobTaskAiFunctionTaskAiClassifyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) JobTaskAiFunctionTaskAiClassifyOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskAiFunctionTaskAiClassifyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskAiFunctionTaskAiClassifyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskAiFunctionTaskAiClassifyOutputReference_Override(j JobTaskAiFunctionTaskAiClassifyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskAiFunctionTaskAiClassifyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference)SetEnableConfidenceScores(val interface{}) {
	if err := j.validateSetEnableConfidenceScoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConfidenceScores",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference)SetEnableRationales(val interface{}) {
	if err := j.validateSetEnableRationalesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableRationales",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference)SetInstructions(val *string) {
	if err := j.validateSetInstructionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instructions",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference)SetInternalValue(val *JobTaskAiFunctionTaskAiClassify) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference)SetMultilabel(val interface{}) {
	if err := j.validateSetMultilabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multilabel",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) PutInlineLabels(value *JobTaskAiFunctionTaskAiClassifyInlineLabels) {
	if err := j.validatePutInlineLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putInlineLabels",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) ResetEnableConfidenceScores() {
	_jsii_.InvokeVoid(
		j,
		"resetEnableConfidenceScores",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) ResetEnableRationales() {
	_jsii_.InvokeVoid(
		j,
		"resetEnableRationales",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) ResetInlineLabels() {
	_jsii_.InvokeVoid(
		j,
		"resetInlineLabels",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) ResetInstructions() {
	_jsii_.InvokeVoid(
		j,
		"resetInstructions",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) ResetMultilabel() {
	_jsii_.InvokeVoid(
		j,
		"resetMultilabel",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiClassifyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

