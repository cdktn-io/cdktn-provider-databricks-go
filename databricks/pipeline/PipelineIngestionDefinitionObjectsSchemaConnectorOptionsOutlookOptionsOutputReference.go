// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/pipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference interface {
	cdktn.ComplexObject
	AttachmentMode() *string
	SetAttachmentMode(val *string)
	AttachmentModeInput() *string
	BodyFormat() *string
	SetBodyFormat(val *string)
	BodyFormatInput() *string
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
	FolderFilter() *[]*string
	SetFolderFilter(val *[]*string)
	FolderFilterInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludeFolders() *[]*string
	SetIncludeFolders(val *[]*string)
	IncludeFoldersInput() *[]*string
	IncludeMailboxes() *[]*string
	SetIncludeMailboxes(val *[]*string)
	IncludeMailboxesInput() *[]*string
	IncludeSenders() *[]*string
	SetIncludeSenders(val *[]*string)
	IncludeSendersInput() *[]*string
	IncludeSubjects() *[]*string
	SetIncludeSubjects(val *[]*string)
	IncludeSubjectsInput() *[]*string
	InternalValue() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptions
	SetInternalValue(val *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptions)
	SenderFilter() *[]*string
	SetSenderFilter(val *[]*string)
	SenderFilterInput() *[]*string
	StartDate() *string
	SetStartDate(val *string)
	StartDateInput() *string
	SubjectFilter() *[]*string
	SetSubjectFilter(val *[]*string)
	SubjectFilterInput() *[]*string
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
	ResetAttachmentMode()
	ResetBodyFormat()
	ResetFolderFilter()
	ResetIncludeFolders()
	ResetIncludeMailboxes()
	ResetIncludeSenders()
	ResetIncludeSubjects()
	ResetSenderFilter()
	ResetStartDate()
	ResetSubjectFilter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference
type jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) AttachmentMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) AttachmentModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) BodyFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) BodyFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) FolderFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"folderFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) FolderFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"folderFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) IncludeFolders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) IncludeFoldersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeFoldersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) IncludeMailboxes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeMailboxes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) IncludeMailboxesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeMailboxesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) IncludeSenders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeSenders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) IncludeSendersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeSendersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) IncludeSubjects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeSubjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) IncludeSubjectsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeSubjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) InternalValue() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) SenderFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"senderFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) SenderFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"senderFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) StartDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) StartDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) SubjectFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subjectFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) SubjectFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subjectFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference_Override(p PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference)SetAttachmentMode(val *string) {
	if err := j.validateSetAttachmentModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attachmentMode",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference)SetBodyFormat(val *string) {
	if err := j.validateSetBodyFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bodyFormat",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference)SetFolderFilter(val *[]*string) {
	if err := j.validateSetFolderFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"folderFilter",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference)SetIncludeFolders(val *[]*string) {
	if err := j.validateSetIncludeFoldersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeFolders",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference)SetIncludeMailboxes(val *[]*string) {
	if err := j.validateSetIncludeMailboxesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeMailboxes",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference)SetIncludeSenders(val *[]*string) {
	if err := j.validateSetIncludeSendersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSenders",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference)SetIncludeSubjects(val *[]*string) {
	if err := j.validateSetIncludeSubjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSubjects",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference)SetInternalValue(val *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference)SetSenderFilter(val *[]*string) {
	if err := j.validateSetSenderFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"senderFilter",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference)SetStartDate(val *string) {
	if err := j.validateSetStartDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startDate",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference)SetSubjectFilter(val *[]*string) {
	if err := j.validateSetSubjectFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectFilter",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) ResetAttachmentMode() {
	_jsii_.InvokeVoid(
		p,
		"resetAttachmentMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) ResetBodyFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetBodyFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) ResetFolderFilter() {
	_jsii_.InvokeVoid(
		p,
		"resetFolderFilter",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) ResetIncludeFolders() {
	_jsii_.InvokeVoid(
		p,
		"resetIncludeFolders",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) ResetIncludeMailboxes() {
	_jsii_.InvokeVoid(
		p,
		"resetIncludeMailboxes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) ResetIncludeSenders() {
	_jsii_.InvokeVoid(
		p,
		"resetIncludeSenders",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) ResetIncludeSubjects() {
	_jsii_.InvokeVoid(
		p,
		"resetIncludeSubjects",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) ResetSenderFilter() {
	_jsii_.InvokeVoid(
		p,
		"resetSenderFilter",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) ResetStartDate() {
	_jsii_.InvokeVoid(
		p,
		"resetStartDate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) ResetSubjectFilter() {
	_jsii_.InvokeVoid(
		p,
		"resetSubjectFilter",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

