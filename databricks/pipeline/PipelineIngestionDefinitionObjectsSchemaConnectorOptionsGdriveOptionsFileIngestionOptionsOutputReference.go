// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/pipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference interface {
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
	CorruptRecordColumn() *string
	SetCorruptRecordColumn(val *string)
	CorruptRecordColumnInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FileFilters() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersList
	FileFiltersInput() interface{}
	Format() *string
	SetFormat(val *string)
	FormatInput() *string
	FormatOptions() *map[string]*string
	SetFormatOptions(val *map[string]*string)
	FormatOptionsInput() *map[string]*string
	// Experimental.
	Fqn() *string
	IgnoreCorruptFiles() interface{}
	SetIgnoreCorruptFiles(val interface{})
	IgnoreCorruptFilesInput() interface{}
	InferColumnTypes() interface{}
	SetInferColumnTypes(val interface{})
	InferColumnTypesInput() interface{}
	InternalValue() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptions
	SetInternalValue(val *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptions)
	ReaderCaseSensitive() interface{}
	SetReaderCaseSensitive(val interface{})
	ReaderCaseSensitiveInput() interface{}
	RescuedDataColumn() *string
	SetRescuedDataColumn(val *string)
	RescuedDataColumnInput() *string
	SchemaEvolutionMode() *string
	SetSchemaEvolutionMode(val *string)
	SchemaEvolutionModeInput() *string
	SchemaHints() *string
	SetSchemaHints(val *string)
	SchemaHintsInput() *string
	SingleVariantColumn() *string
	SetSingleVariantColumn(val *string)
	SingleVariantColumnInput() *string
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
	PutFileFilters(value interface{})
	ResetCorruptRecordColumn()
	ResetFileFilters()
	ResetFormat()
	ResetFormatOptions()
	ResetIgnoreCorruptFiles()
	ResetInferColumnTypes()
	ResetReaderCaseSensitive()
	ResetRescuedDataColumn()
	ResetSchemaEvolutionMode()
	ResetSchemaHints()
	ResetSingleVariantColumn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference
type jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) CorruptRecordColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"corruptRecordColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) CorruptRecordColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"corruptRecordColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) FileFilters() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersList {
	var returns PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersList
	_jsii_.Get(
		j,
		"fileFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) FileFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) FormatOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"formatOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) FormatOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"formatOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) IgnoreCorruptFiles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCorruptFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) IgnoreCorruptFilesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCorruptFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) InferColumnTypes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferColumnTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) InferColumnTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferColumnTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) InternalValue() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) ReaderCaseSensitive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readerCaseSensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) ReaderCaseSensitiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readerCaseSensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) RescuedDataColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rescuedDataColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) RescuedDataColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rescuedDataColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) SchemaEvolutionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaEvolutionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) SchemaEvolutionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaEvolutionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) SchemaHints() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaHints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) SchemaHintsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaHintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) SingleVariantColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleVariantColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) SingleVariantColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleVariantColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference_Override(p PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference)SetCorruptRecordColumn(val *string) {
	if err := j.validateSetCorruptRecordColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"corruptRecordColumn",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference)SetFormatOptions(val *map[string]*string) {
	if err := j.validateSetFormatOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"formatOptions",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference)SetIgnoreCorruptFiles(val interface{}) {
	if err := j.validateSetIgnoreCorruptFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreCorruptFiles",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference)SetInferColumnTypes(val interface{}) {
	if err := j.validateSetInferColumnTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferColumnTypes",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference)SetInternalValue(val *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference)SetReaderCaseSensitive(val interface{}) {
	if err := j.validateSetReaderCaseSensitiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readerCaseSensitive",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference)SetRescuedDataColumn(val *string) {
	if err := j.validateSetRescuedDataColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rescuedDataColumn",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference)SetSchemaEvolutionMode(val *string) {
	if err := j.validateSetSchemaEvolutionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaEvolutionMode",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference)SetSchemaHints(val *string) {
	if err := j.validateSetSchemaHintsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaHints",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference)SetSingleVariantColumn(val *string) {
	if err := j.validateSetSingleVariantColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleVariantColumn",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) PutFileFilters(value interface{}) {
	if err := p.validatePutFileFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFileFilters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) ResetCorruptRecordColumn() {
	_jsii_.InvokeVoid(
		p,
		"resetCorruptRecordColumn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) ResetFileFilters() {
	_jsii_.InvokeVoid(
		p,
		"resetFileFilters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) ResetFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) ResetFormatOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetFormatOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) ResetIgnoreCorruptFiles() {
	_jsii_.InvokeVoid(
		p,
		"resetIgnoreCorruptFiles",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) ResetInferColumnTypes() {
	_jsii_.InvokeVoid(
		p,
		"resetInferColumnTypes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) ResetReaderCaseSensitive() {
	_jsii_.InvokeVoid(
		p,
		"resetReaderCaseSensitive",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) ResetRescuedDataColumn() {
	_jsii_.InvokeVoid(
		p,
		"resetRescuedDataColumn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) ResetSchemaEvolutionMode() {
	_jsii_.InvokeVoid(
		p,
		"resetSchemaEvolutionMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) ResetSchemaHints() {
	_jsii_.InvokeVoid(
		p,
		"resetSchemaHints",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) ResetSingleVariantColumn() {
	_jsii_.InvokeVoid(
		p,
		"resetSingleVariantColumn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

