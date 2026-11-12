// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/job/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JobTaskAiFunctionTaskAiParseDocumentOutputReference interface {
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
	DescriptionElementTypes() *[]*string
	SetDescriptionElementTypes(val *[]*string)
	DescriptionElementTypesInput() *[]*string
	// Experimental.
	Fqn() *string
	ImageOutputPath() *string
	SetImageOutputPath(val *string)
	ImageOutputPathInput() *string
	InternalValue() *JobTaskAiFunctionTaskAiParseDocument
	SetInternalValue(val *JobTaskAiFunctionTaskAiParseDocument)
	PageRange() *string
	SetPageRange(val *string)
	PageRangeInput() *string
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
	ResetDescriptionElementTypes()
	ResetImageOutputPath()
	ResetPageRange()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTaskAiFunctionTaskAiParseDocumentOutputReference
type jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) DescriptionElementTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"descriptionElementTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) DescriptionElementTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"descriptionElementTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) ImageOutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageOutputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) ImageOutputPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageOutputPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) InternalValue() *JobTaskAiFunctionTaskAiParseDocument {
	var returns *JobTaskAiFunctionTaskAiParseDocument
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) PageRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pageRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) PageRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pageRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewJobTaskAiFunctionTaskAiParseDocumentOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) JobTaskAiFunctionTaskAiParseDocumentOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskAiFunctionTaskAiParseDocumentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskAiFunctionTaskAiParseDocumentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskAiFunctionTaskAiParseDocumentOutputReference_Override(j JobTaskAiFunctionTaskAiParseDocumentOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskAiFunctionTaskAiParseDocumentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference)SetDescriptionElementTypes(val *[]*string) {
	if err := j.validateSetDescriptionElementTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"descriptionElementTypes",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference)SetImageOutputPath(val *string) {
	if err := j.validateSetImageOutputPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageOutputPath",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference)SetInternalValue(val *JobTaskAiFunctionTaskAiParseDocument) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference)SetPageRange(val *string) {
	if err := j.validateSetPageRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pageRange",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) ResetDescriptionElementTypes() {
	_jsii_.InvokeVoid(
		j,
		"resetDescriptionElementTypes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) ResetImageOutputPath() {
	_jsii_.InvokeVoid(
		j,
		"resetImageOutputPath",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) ResetPageRange() {
	_jsii_.InvokeVoid(
		j,
		"resetPageRange",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskAiFunctionTaskAiParseDocumentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

