// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/job/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference interface {
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
	EnableCitations() interface{}
	SetEnableCitations(val interface{})
	EnableCitationsInput() interface{}
	EnableConfidenceScores() interface{}
	SetEnableConfidenceScores(val interface{})
	EnableConfidenceScoresInput() interface{}
	// Experimental.
	Fqn() *string
	Instructions() *string
	SetInstructions(val *string)
	InstructionsInput() *string
	InternalValue() *JobTaskForEachTaskTaskAiFunctionTaskAiExtract
	SetInternalValue(val *JobTaskForEachTaskTaskAiFunctionTaskAiExtract)
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
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
	ResetEnableCitations()
	ResetEnableConfidenceScores()
	ResetInstructions()
	ResetMode()
	ResetSchema()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference
type jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) EnableCitations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCitations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) EnableCitationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCitationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) EnableConfidenceScores() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConfidenceScores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) EnableConfidenceScoresInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConfidenceScoresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) Instructions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instructions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) InstructionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instructionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) InternalValue() *JobTaskForEachTaskTaskAiFunctionTaskAiExtract {
	var returns *JobTaskForEachTaskTaskAiFunctionTaskAiExtract
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewJobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference_Override(j JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference)SetEnableCitations(val interface{}) {
	if err := j.validateSetEnableCitationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCitations",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference)SetEnableConfidenceScores(val interface{}) {
	if err := j.validateSetEnableConfidenceScoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConfidenceScores",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference)SetInstructions(val *string) {
	if err := j.validateSetInstructionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instructions",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference)SetInternalValue(val *JobTaskForEachTaskTaskAiFunctionTaskAiExtract) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) ResetEnableCitations() {
	_jsii_.InvokeVoid(
		j,
		"resetEnableCitations",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) ResetEnableConfidenceScores() {
	_jsii_.InvokeVoid(
		j,
		"resetEnableConfidenceScores",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) ResetInstructions() {
	_jsii_.InvokeVoid(
		j,
		"resetInstructions",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		j,
		"resetMode",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) ResetSchema() {
	_jsii_.InvokeVoid(
		j,
		"resetSchema",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskAiFunctionTaskAiExtractOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

