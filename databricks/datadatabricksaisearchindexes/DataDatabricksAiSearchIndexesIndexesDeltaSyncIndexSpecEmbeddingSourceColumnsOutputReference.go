// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaisearchindexes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabricksaisearchindexes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference interface {
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
	EmbeddingModelEndpoint() *string
	SetEmbeddingModelEndpoint(val *string)
	EmbeddingModelEndpointInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ModelEndpointNameForQuery() *string
	SetModelEndpointNameForQuery(val *string)
	ModelEndpointNameForQueryInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	ResetEmbeddingModelEndpoint()
	ResetModelEndpointNameForQuery()
	ResetName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference
type jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) EmbeddingModelEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingModelEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) EmbeddingModelEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingModelEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) ModelEndpointNameForQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelEndpointNameForQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) ModelEndpointNameForQueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelEndpointNameForQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiSearchIndexes.DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference_Override(d DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiSearchIndexes.DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference)SetEmbeddingModelEndpoint(val *string) {
	if err := j.validateSetEmbeddingModelEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"embeddingModelEndpoint",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference)SetModelEndpointNameForQuery(val *string) {
	if err := j.validateSetModelEndpointNameForQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelEndpointNameForQuery",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) ResetEmbeddingModelEndpoint() {
	_jsii_.InvokeVoid(
		d,
		"resetEmbeddingModelEndpoint",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) ResetModelEndpointNameForQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetModelEndpointNameForQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpecEmbeddingSourceColumnsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

