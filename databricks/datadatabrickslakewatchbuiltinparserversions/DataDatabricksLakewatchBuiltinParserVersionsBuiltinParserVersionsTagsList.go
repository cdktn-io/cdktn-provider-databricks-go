// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchbuiltinparserversions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabrickslakewatchbuiltinparserversions/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList interface {
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
	Get(index *float64) DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList
type jsiiProxy_DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList {
	_init_.Initialize()

	if err := validateNewDataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchBuiltinParserVersions.DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList_Override(d DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchBuiltinParserVersions.DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList) Get(index *float64) DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksLakewatchBuiltinParserVersionsBuiltinParserVersionsTagsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

