// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaisearchindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksaisearchindex/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference interface {
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
	EmbeddingSourceColumns() DataDatabricksAiSearchIndexDeltaSyncIndexSpecEmbeddingSourceColumnsList
	EmbeddingSourceColumnsInput() interface{}
	EmbeddingVectorColumns() DataDatabricksAiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList
	EmbeddingVectorColumnsInput() interface{}
	EmbeddingWritebackTable() *string
	SetEmbeddingWritebackTable(val *string)
	EmbeddingWritebackTableInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksAiSearchIndexDeltaSyncIndexSpec
	SetInternalValue(val *DataDatabricksAiSearchIndexDeltaSyncIndexSpec)
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

// The jsii proxy struct for DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference
type jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) ColumnsToSync() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"columnsToSync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) ColumnsToSyncInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"columnsToSyncInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) EmbeddingSourceColumns() DataDatabricksAiSearchIndexDeltaSyncIndexSpecEmbeddingSourceColumnsList {
	var returns DataDatabricksAiSearchIndexDeltaSyncIndexSpecEmbeddingSourceColumnsList
	_jsii_.Get(
		j,
		"embeddingSourceColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) EmbeddingSourceColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embeddingSourceColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) EmbeddingVectorColumns() DataDatabricksAiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList {
	var returns DataDatabricksAiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList
	_jsii_.Get(
		j,
		"embeddingVectorColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) EmbeddingVectorColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embeddingVectorColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) EmbeddingWritebackTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingWritebackTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) EmbeddingWritebackTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingWritebackTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) InternalValue() *DataDatabricksAiSearchIndexDeltaSyncIndexSpec {
	var returns *DataDatabricksAiSearchIndexDeltaSyncIndexSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) PipelineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) PipelineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) PipelineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) SourceTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) SourceTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiSearchIndex.DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference_Override(d DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiSearchIndex.DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference)SetColumnsToSync(val *[]*string) {
	if err := j.validateSetColumnsToSyncParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnsToSync",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference)SetEmbeddingWritebackTable(val *string) {
	if err := j.validateSetEmbeddingWritebackTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"embeddingWritebackTable",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference)SetInternalValue(val *DataDatabricksAiSearchIndexDeltaSyncIndexSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference)SetPipelineType(val *string) {
	if err := j.validateSetPipelineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference)SetSourceTable(val *string) {
	if err := j.validateSetSourceTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceTable",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) PutEmbeddingSourceColumns(value interface{}) {
	if err := d.validatePutEmbeddingSourceColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEmbeddingSourceColumns",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) PutEmbeddingVectorColumns(value interface{}) {
	if err := d.validatePutEmbeddingVectorColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEmbeddingVectorColumns",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) ResetColumnsToSync() {
	_jsii_.InvokeVoid(
		d,
		"resetColumnsToSync",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) ResetEmbeddingSourceColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetEmbeddingSourceColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) ResetEmbeddingVectorColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetEmbeddingVectorColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) ResetEmbeddingWritebackTable() {
	_jsii_.InvokeVoid(
		d,
		"resetEmbeddingWritebackTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) ResetSourceTable() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexDeltaSyncIndexSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

