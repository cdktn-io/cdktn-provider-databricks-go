// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/pipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference interface {
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
	ConfluentOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryConfluentOptionsOutputReference
	ConfluentOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryConfluentOptions
	ConnectionName() *string
	SetConnectionName(val *string)
	ConnectionNameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistry
	SetInternalValue(val *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistry)
	ProtobufMessageName() *string
	SetProtobufMessageName(val *string)
	ProtobufMessageNameInput() *string
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
	PutConfluentOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryConfluentOptions)
	ResetConfluentOptions()
	ResetConnectionName()
	ResetProtobufMessageName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference
type jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) ConfluentOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryConfluentOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryConfluentOptionsOutputReference
	_jsii_.Get(
		j,
		"confluentOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) ConfluentOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryConfluentOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryConfluentOptions
	_jsii_.Get(
		j,
		"confluentOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) ConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) ConnectionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) InternalValue() *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistry {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistry
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) ProtobufMessageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protobufMessageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) ProtobufMessageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protobufMessageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference_Override(p PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference)SetConnectionName(val *string) {
	if err := j.validateSetConnectionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionName",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference)SetInternalValue(val *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistry) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference)SetProtobufMessageName(val *string) {
	if err := j.validateSetProtobufMessageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protobufMessageName",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) PutConfluentOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryConfluentOptions) {
	if err := p.validatePutConfluentOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putConfluentOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) ResetConfluentOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetConfluentOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) ResetConnectionName() {
	_jsii_.InvokeVoid(
		p,
		"resetConnectionName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) ResetProtobufMessageName() {
	_jsii_.InvokeVoid(
		p,
		"resetProtobufMessageName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerAvroOptionsSchemaRegistryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

