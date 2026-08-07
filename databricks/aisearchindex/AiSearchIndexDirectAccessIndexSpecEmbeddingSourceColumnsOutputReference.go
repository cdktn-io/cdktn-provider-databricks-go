// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/aisearchindex/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference interface {
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

// The jsii proxy struct for AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference
type jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) EmbeddingModelEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingModelEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) EmbeddingModelEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingModelEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) ModelEndpointNameForQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelEndpointNameForQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) ModelEndpointNameForQueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelEndpointNameForQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference {
	_init_.Initialize()

	if err := validateNewAiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.aiSearchIndex.AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference_Override(a AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.aiSearchIndex.AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference)SetEmbeddingModelEndpoint(val *string) {
	if err := j.validateSetEmbeddingModelEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"embeddingModelEndpoint",
		val,
	)
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference)SetModelEndpointNameForQuery(val *string) {
	if err := j.validateSetModelEndpointNameForQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelEndpointNameForQuery",
		val,
	)
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) ResetEmbeddingModelEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetEmbeddingModelEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) ResetModelEndpointNameForQuery() {
	_jsii_.InvokeVoid(
		a,
		"resetModelEndpointNameForQuery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumnsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

