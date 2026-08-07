// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksknowledgeassistantknowledgesources

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksknowledgeassistantknowledgesources/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference interface {
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
	DocUriCol() *string
	SetDocUriCol(val *string)
	DocUriColInput() *string
	// Experimental.
	Fqn() *string
	IndexName() *string
	SetIndexName(val *string)
	IndexNameInput() *string
	InternalValue() *DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndex
	SetInternalValue(val *DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndex)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TextCol() *string
	SetTextCol(val *string)
	TextColInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference
type jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) DocUriCol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"docUriCol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) DocUriColInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"docUriColInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) IndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) IndexNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) InternalValue() *DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndex {
	var returns *DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndex
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) TextCol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textCol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) TextColInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textColInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksKnowledgeAssistantKnowledgeSources.DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference_Override(d DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksKnowledgeAssistantKnowledgeSources.DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference)SetDocUriCol(val *string) {
	if err := j.validateSetDocUriColParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"docUriCol",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference)SetIndexName(val *string) {
	if err := j.validateSetIndexNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference)SetInternalValue(val *DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndex) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference)SetTextCol(val *string) {
	if err := j.validateSetTextColParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textCol",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndexOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

