// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaisearchindexes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabricksaisearchindexes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference interface {
	cdktn.ComplexObject
	ColumnsToSync() *[]*string
	SetColumnsToSync(val *[]*string)
	ColumnsToSyncInput() *[]*string
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
	EmbeddingSourceColumns() DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsList
	EmbeddingSourceColumnsInput() interface{}
	EmbeddingVectorColumns() DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingVectorColumnsList
	EmbeddingVectorColumnsInput() interface{}
	EmbeddingWritebackTable() *string
	SetEmbeddingWritebackTable(val *string)
	EmbeddingWritebackTableInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpec
	SetInternalValue(val *DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpec)
	PipelineId() *string
	PipelineType() *string
	SetPipelineType(val *string)
	PipelineTypeInput() *string
	SourceTable() *string
	SetSourceTable(val *string)
	SourceTableInput() *string
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
	ResetColumnsToSync()
	ResetEmbeddingSourceColumns()
	ResetEmbeddingVectorColumns()
	ResetEmbeddingWritebackTable()
	ResetSourceTable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference
type jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) ColumnsToSync() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"columnsToSync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) ColumnsToSyncInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"columnsToSyncInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) EmbeddingSourceColumns() DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsList {
	var returns DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsList
	_jsii_.Get(
		j,
		"embeddingSourceColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) EmbeddingSourceColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embeddingSourceColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) EmbeddingVectorColumns() DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingVectorColumnsList {
	var returns DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingVectorColumnsList
	_jsii_.Get(
		j,
		"embeddingVectorColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) EmbeddingVectorColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embeddingVectorColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) EmbeddingWritebackTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingWritebackTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) EmbeddingWritebackTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingWritebackTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) InternalValue() *DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpec {
	var returns *DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) PipelineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) PipelineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) PipelineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) SourceTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) SourceTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiSearchIndexes.DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference_Override(d DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiSearchIndexes.DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference)SetColumnsToSync(val *[]*string) {
	if err := j.validateSetColumnsToSyncParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnsToSync",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference)SetEmbeddingWritebackTable(val *string) {
	if err := j.validateSetEmbeddingWritebackTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"embeddingWritebackTable",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference)SetInternalValue(val *DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference)SetPipelineType(val *string) {
	if err := j.validateSetPipelineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference)SetSourceTable(val *string) {
	if err := j.validateSetSourceTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceTable",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) PutEmbeddingSourceColumns(value interface{}) {
	if err := d.validatePutEmbeddingSourceColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEmbeddingSourceColumns",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) PutEmbeddingVectorColumns(value interface{}) {
	if err := d.validatePutEmbeddingVectorColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEmbeddingVectorColumns",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) ResetColumnsToSync() {
	_jsii_.InvokeVoid(
		d,
		"resetColumnsToSync",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) ResetEmbeddingSourceColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetEmbeddingSourceColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) ResetEmbeddingVectorColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetEmbeddingVectorColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) ResetEmbeddingWritebackTable() {
	_jsii_.InvokeVoid(
		d,
		"resetEmbeddingWritebackTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) ResetSourceTable() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

