// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/aisearchindex/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AiSearchIndexDirectAccessIndexSpecOutputReference interface {
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
	EmbeddingSourceColumns() AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsList
	EmbeddingSourceColumnsInput() interface{}
	EmbeddingVectorColumns() AiSearchIndexDirectAccessIndexSpecEmbeddingVectorColumnsList
	EmbeddingVectorColumnsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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

// The jsii proxy struct for AiSearchIndexDirectAccessIndexSpecOutputReference
type jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) EmbeddingSourceColumns() AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsList {
	var returns AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsList
	_jsii_.Get(
		j,
		"embeddingSourceColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) EmbeddingSourceColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embeddingSourceColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) EmbeddingVectorColumns() AiSearchIndexDirectAccessIndexSpecEmbeddingVectorColumnsList {
	var returns AiSearchIndexDirectAccessIndexSpecEmbeddingVectorColumnsList
	_jsii_.Get(
		j,
		"embeddingVectorColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) EmbeddingVectorColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embeddingVectorColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) SchemaJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) SchemaJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAiSearchIndexDirectAccessIndexSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AiSearchIndexDirectAccessIndexSpecOutputReference {
	_init_.Initialize()

	if err := validateNewAiSearchIndexDirectAccessIndexSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.aiSearchIndex.AiSearchIndexDirectAccessIndexSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAiSearchIndexDirectAccessIndexSpecOutputReference_Override(a AiSearchIndexDirectAccessIndexSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.aiSearchIndex.AiSearchIndexDirectAccessIndexSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference)SetSchemaJson(val *string) {
	if err := j.validateSetSchemaJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaJson",
		val,
	)
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) PutEmbeddingSourceColumns(value interface{}) {
	if err := a.validatePutEmbeddingSourceColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEmbeddingSourceColumns",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) PutEmbeddingVectorColumns(value interface{}) {
	if err := a.validatePutEmbeddingVectorColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEmbeddingVectorColumns",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) ResetEmbeddingSourceColumns() {
	_jsii_.InvokeVoid(
		a,
		"resetEmbeddingSourceColumns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) ResetEmbeddingVectorColumns() {
	_jsii_.InvokeVoid(
		a,
		"resetEmbeddingVectorColumns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) ResetSchemaJson() {
	_jsii_.InvokeVoid(
		a,
		"resetSchemaJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

