// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickssupervisoragenttools

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabrickssupervisoragenttools/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksSupervisorAgentToolsToolsOutputReference interface {
	cdktn.ComplexObject
	App() DataDatabricksSupervisorAgentToolsToolsAppOutputReference
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
	Description() *string
	// Experimental.
	Fqn() *string
	GenieSpace() DataDatabricksSupervisorAgentToolsToolsGenieSpaceOutputReference
	Id() *string
	InternalValue() *DataDatabricksSupervisorAgentToolsTools
	SetInternalValue(val *DataDatabricksSupervisorAgentToolsTools)
	KnowledgeAssistant() DataDatabricksSupervisorAgentToolsToolsKnowledgeAssistantOutputReference
	Name() *string
	SetName(val *string)
	NameInput() *string
	ProviderConfig() DataDatabricksSupervisorAgentToolsToolsProviderConfigOutputReference
	ProviderConfigInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ToolId() *string
	ToolType() *string
	UcConnection() DataDatabricksSupervisorAgentToolsToolsUcConnectionOutputReference
	UcFunction() DataDatabricksSupervisorAgentToolsToolsUcFunctionOutputReference
	Volume() DataDatabricksSupervisorAgentToolsToolsVolumeOutputReference
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
	PutProviderConfig(value *DataDatabricksSupervisorAgentToolsToolsProviderConfig)
	ResetProviderConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksSupervisorAgentToolsToolsOutputReference
type jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) App() DataDatabricksSupervisorAgentToolsToolsAppOutputReference {
	var returns DataDatabricksSupervisorAgentToolsToolsAppOutputReference
	_jsii_.Get(
		j,
		"app",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) GenieSpace() DataDatabricksSupervisorAgentToolsToolsGenieSpaceOutputReference {
	var returns DataDatabricksSupervisorAgentToolsToolsGenieSpaceOutputReference
	_jsii_.Get(
		j,
		"genieSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) InternalValue() *DataDatabricksSupervisorAgentToolsTools {
	var returns *DataDatabricksSupervisorAgentToolsTools
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) KnowledgeAssistant() DataDatabricksSupervisorAgentToolsToolsKnowledgeAssistantOutputReference {
	var returns DataDatabricksSupervisorAgentToolsToolsKnowledgeAssistantOutputReference
	_jsii_.Get(
		j,
		"knowledgeAssistant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) ProviderConfig() DataDatabricksSupervisorAgentToolsToolsProviderConfigOutputReference {
	var returns DataDatabricksSupervisorAgentToolsToolsProviderConfigOutputReference
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) ToolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) ToolType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) UcConnection() DataDatabricksSupervisorAgentToolsToolsUcConnectionOutputReference {
	var returns DataDatabricksSupervisorAgentToolsToolsUcConnectionOutputReference
	_jsii_.Get(
		j,
		"ucConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) UcFunction() DataDatabricksSupervisorAgentToolsToolsUcFunctionOutputReference {
	var returns DataDatabricksSupervisorAgentToolsToolsUcFunctionOutputReference
	_jsii_.Get(
		j,
		"ucFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) Volume() DataDatabricksSupervisorAgentToolsToolsVolumeOutputReference {
	var returns DataDatabricksSupervisorAgentToolsToolsVolumeOutputReference
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}


func NewDataDatabricksSupervisorAgentToolsToolsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksSupervisorAgentToolsToolsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksSupervisorAgentToolsToolsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksSupervisorAgentTools.DataDatabricksSupervisorAgentToolsToolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksSupervisorAgentToolsToolsOutputReference_Override(d DataDatabricksSupervisorAgentToolsToolsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksSupervisorAgentTools.DataDatabricksSupervisorAgentToolsToolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference)SetInternalValue(val *DataDatabricksSupervisorAgentToolsTools) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) PutProviderConfig(value *DataDatabricksSupervisorAgentToolsToolsProviderConfig) {
	if err := d.validatePutProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProviderConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksSupervisorAgentToolsToolsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

