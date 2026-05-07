// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/pipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ModifiedAfter() *string
	SetModifiedAfter(val *string)
	ModifiedAfterInput() *string
	ModifiedBefore() *string
	SetModifiedBefore(val *string)
	ModifiedBeforeInput() *string
	PathFilter() *string
	SetPathFilter(val *string)
	PathFilterInput() *string
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
	ResetModifiedAfter()
	ResetModifiedBefore()
	ResetPathFilter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference
type jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) ModifiedAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) ModifiedAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) ModifiedBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) ModifiedBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) PathFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) PathFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference_Override(p PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference)SetModifiedAfter(val *string) {
	if err := j.validateSetModifiedAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modifiedAfter",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference)SetModifiedBefore(val *string) {
	if err := j.validateSetModifiedBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modifiedBefore",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference)SetPathFilter(val *string) {
	if err := j.validateSetPathFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathFilter",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) ResetModifiedAfter() {
	_jsii_.InvokeVoid(
		p,
		"resetModifiedAfter",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) ResetModifiedBefore() {
	_jsii_.InvokeVoid(
		p,
		"resetModifiedBefore",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) ResetPathFilter() {
	_jsii_.InvokeVoid(
		p,
		"resetPathFilter",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

