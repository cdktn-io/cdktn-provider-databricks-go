// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/aisearchindex/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList
type jsiiProxy_AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList {
	_init_.Initialize()

	if err := validateNewAiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList{}

	_jsii_.Create(
		"@cdktn/provider-databricks.aiSearchIndex.AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList_Override(a AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.aiSearchIndex.AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := a.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		a,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList) Get(index *float64) AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AiSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumnsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

