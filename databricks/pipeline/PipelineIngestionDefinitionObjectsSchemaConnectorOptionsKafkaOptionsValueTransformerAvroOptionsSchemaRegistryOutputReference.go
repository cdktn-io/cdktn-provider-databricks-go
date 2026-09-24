// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/pipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference interface {
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
	ConfluentOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryConfluentOptionsOutputReference
	ConfluentOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryConfluentOptions
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
	InternalValue() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistry
	SetInternalValue(val *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistry)
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
	PutConfluentOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryConfluentOptions)
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

// The jsii proxy struct for PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference
type jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) ConfluentOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryConfluentOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryConfluentOptionsOutputReference
	_jsii_.Get(
		j,
		"confluentOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) ConfluentOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryConfluentOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryConfluentOptions
	_jsii_.Get(
		j,
		"confluentOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) ConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) ConnectionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) InternalValue() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistry {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistry
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) ProtobufMessageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protobufMessageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) ProtobufMessageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protobufMessageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference_Override(p PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference)SetConnectionName(val *string) {
	if err := j.validateSetConnectionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionName",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference)SetInternalValue(val *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistry) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference)SetProtobufMessageName(val *string) {
	if err := j.validateSetProtobufMessageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protobufMessageName",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) PutConfluentOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryConfluentOptions) {
	if err := p.validatePutConfluentOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putConfluentOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) ResetConfluentOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetConfluentOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) ResetConnectionName() {
	_jsii_.InvokeVoid(
		p,
		"resetConnectionName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) ResetProtobufMessageName() {
	_jsii_.InvokeVoid(
		p,
		"resetProtobufMessageName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

