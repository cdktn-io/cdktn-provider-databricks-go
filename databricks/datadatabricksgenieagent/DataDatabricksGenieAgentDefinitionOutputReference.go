// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksgenieagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabricksgenieagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksGenieAgentDefinitionOutputReference interface {
	cdktn.ComplexObject
	Benchmarks() DataDatabricksGenieAgentDefinitionBenchmarksOutputReference
	BenchmarksInput() interface{}
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
	Config() DataDatabricksGenieAgentDefinitionConfigOutputReference
	ConfigInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataSources() DataDatabricksGenieAgentDefinitionDataSourcesOutputReference
	DataSourcesInput() interface{}
	// Experimental.
	Fqn() *string
	Instructions() DataDatabricksGenieAgentDefinitionInstructionsOutputReference
	InstructionsInput() interface{}
	InternalValue() *DataDatabricksGenieAgentDefinition
	SetInternalValue(val *DataDatabricksGenieAgentDefinition)
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
	PutBenchmarks(value *DataDatabricksGenieAgentDefinitionBenchmarks)
	PutConfig(value *DataDatabricksGenieAgentDefinitionConfig)
	PutDataSources(value *DataDatabricksGenieAgentDefinitionDataSources)
	PutInstructions(value *DataDatabricksGenieAgentDefinitionInstructions)
	ResetBenchmarks()
	ResetConfig()
	ResetDataSources()
	ResetInstructions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksGenieAgentDefinitionOutputReference
type jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) Benchmarks() DataDatabricksGenieAgentDefinitionBenchmarksOutputReference {
	var returns DataDatabricksGenieAgentDefinitionBenchmarksOutputReference
	_jsii_.Get(
		j,
		"benchmarks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) BenchmarksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"benchmarksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) Config() DataDatabricksGenieAgentDefinitionConfigOutputReference {
	var returns DataDatabricksGenieAgentDefinitionConfigOutputReference
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) ConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) DataSources() DataDatabricksGenieAgentDefinitionDataSourcesOutputReference {
	var returns DataDatabricksGenieAgentDefinitionDataSourcesOutputReference
	_jsii_.Get(
		j,
		"dataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) DataSourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) Instructions() DataDatabricksGenieAgentDefinitionInstructionsOutputReference {
	var returns DataDatabricksGenieAgentDefinitionInstructionsOutputReference
	_jsii_.Get(
		j,
		"instructions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) InstructionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instructionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) InternalValue() *DataDatabricksGenieAgentDefinition {
	var returns *DataDatabricksGenieAgentDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksGenieAgentDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksGenieAgentDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksGenieAgentDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksGenieAgent.DataDatabricksGenieAgentDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksGenieAgentDefinitionOutputReference_Override(d DataDatabricksGenieAgentDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksGenieAgent.DataDatabricksGenieAgentDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference)SetInternalValue(val *DataDatabricksGenieAgentDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) PutBenchmarks(value *DataDatabricksGenieAgentDefinitionBenchmarks) {
	if err := d.validatePutBenchmarksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBenchmarks",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) PutConfig(value *DataDatabricksGenieAgentDefinitionConfig) {
	if err := d.validatePutConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) PutDataSources(value *DataDatabricksGenieAgentDefinitionDataSources) {
	if err := d.validatePutDataSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDataSources",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) PutInstructions(value *DataDatabricksGenieAgentDefinitionInstructions) {
	if err := d.validatePutInstructionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInstructions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) ResetBenchmarks() {
	_jsii_.InvokeVoid(
		d,
		"resetBenchmarks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) ResetConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) ResetDataSources() {
	_jsii_.InvokeVoid(
		d,
		"resetDataSources",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) ResetInstructions() {
	_jsii_.InvokeVoid(
		d,
		"resetInstructions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksGenieAgentDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

