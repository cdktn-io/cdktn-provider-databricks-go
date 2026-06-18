// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaisearchindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabricksaisearchindex/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference interface {
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
	EmbeddingSourceColumns() DataDatabricksAiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsList
	EmbeddingSourceColumnsInput() interface{}
	EmbeddingVectorColumns() DataDatabricksAiSearchIndexDirectAccessIndexSpecEmbeddingVectorColumnsList
	EmbeddingVectorColumnsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksAiSearchIndexDirectAccessIndexSpec
	SetInternalValue(val *DataDatabricksAiSearchIndexDirectAccessIndexSpec)
	SchemaJson() *string
	SetSchemaJson(val *string)
	SchemaJsonInput() *string
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
	PutEmbeddingSourceColumns(value interface{})
	PutEmbeddingVectorColumns(value interface{})
	ResetEmbeddingSourceColumns()
	ResetEmbeddingVectorColumns()
	ResetSchemaJson()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference
type jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) EmbeddingSourceColumns() DataDatabricksAiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsList {
	var returns DataDatabricksAiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsList
	_jsii_.Get(
		j,
		"embeddingSourceColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) EmbeddingSourceColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embeddingSourceColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) EmbeddingVectorColumns() DataDatabricksAiSearchIndexDirectAccessIndexSpecEmbeddingVectorColumnsList {
	var returns DataDatabricksAiSearchIndexDirectAccessIndexSpecEmbeddingVectorColumnsList
	_jsii_.Get(
		j,
		"embeddingVectorColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) EmbeddingVectorColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embeddingVectorColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) InternalValue() *DataDatabricksAiSearchIndexDirectAccessIndexSpec {
	var returns *DataDatabricksAiSearchIndexDirectAccessIndexSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) SchemaJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) SchemaJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiSearchIndex.DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference_Override(d DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiSearchIndex.DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference)SetInternalValue(val *DataDatabricksAiSearchIndexDirectAccessIndexSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference)SetSchemaJson(val *string) {
	if err := j.validateSetSchemaJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaJson",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) PutEmbeddingSourceColumns(value interface{}) {
	if err := d.validatePutEmbeddingSourceColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEmbeddingSourceColumns",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) PutEmbeddingVectorColumns(value interface{}) {
	if err := d.validatePutEmbeddingVectorColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEmbeddingVectorColumns",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) ResetEmbeddingSourceColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetEmbeddingSourceColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) ResetEmbeddingVectorColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetEmbeddingVectorColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) ResetSchemaJson() {
	_jsii_.InvokeVoid(
		d,
		"resetSchemaJson",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDirectAccessIndexSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

