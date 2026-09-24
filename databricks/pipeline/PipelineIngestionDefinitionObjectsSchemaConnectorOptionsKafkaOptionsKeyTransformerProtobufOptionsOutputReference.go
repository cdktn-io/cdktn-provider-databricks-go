// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/pipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference interface {
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
	DescFilePath() *string
	SetDescFilePath(val *string)
	DescFilePathInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptions
	SetInternalValue(val *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptions)
	MessageName() *string
	SetMessageName(val *string)
	MessageNameInput() *string
	ParseMode() *string
	SetParseMode(val *string)
	ParseModeInput() *string
	RecursiveFieldsMaxDepth() *float64
	SetRecursiveFieldsMaxDepth(val *float64)
	RecursiveFieldsMaxDepthInput() *float64
	SchemaRegistry() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsSchemaRegistryOutputReference
	SchemaRegistryInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsSchemaRegistry
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
	PutSchemaRegistry(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsSchemaRegistry)
	ResetDescFilePath()
	ResetMessageName()
	ResetParseMode()
	ResetRecursiveFieldsMaxDepth()
	ResetSchemaRegistry()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference
type jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) DescFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) DescFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) InternalValue() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) MessageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) MessageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) ParseMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parseMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) ParseModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parseModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) RecursiveFieldsMaxDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recursiveFieldsMaxDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) RecursiveFieldsMaxDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recursiveFieldsMaxDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) SchemaRegistry() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsSchemaRegistryOutputReference {
	var returns PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsSchemaRegistryOutputReference
	_jsii_.Get(
		j,
		"schemaRegistry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) SchemaRegistryInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsSchemaRegistry {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsSchemaRegistry
	_jsii_.Get(
		j,
		"schemaRegistryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference_Override(p PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference)SetDescFilePath(val *string) {
	if err := j.validateSetDescFilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"descFilePath",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference)SetInternalValue(val *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference)SetMessageName(val *string) {
	if err := j.validateSetMessageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageName",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference)SetParseMode(val *string) {
	if err := j.validateSetParseModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parseMode",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference)SetRecursiveFieldsMaxDepth(val *float64) {
	if err := j.validateSetRecursiveFieldsMaxDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recursiveFieldsMaxDepth",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) PutSchemaRegistry(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsSchemaRegistry) {
	if err := p.validatePutSchemaRegistryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSchemaRegistry",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) ResetDescFilePath() {
	_jsii_.InvokeVoid(
		p,
		"resetDescFilePath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) ResetMessageName() {
	_jsii_.InvokeVoid(
		p,
		"resetMessageName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) ResetParseMode() {
	_jsii_.InvokeVoid(
		p,
		"resetParseMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) ResetRecursiveFieldsMaxDepth() {
	_jsii_.InvokeVoid(
		p,
		"resetRecursiveFieldsMaxDepth",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) ResetSchemaRegistry() {
	_jsii_.InvokeVoid(
		p,
		"resetSchemaRegistry",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsKeyTransformerProtobufOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

