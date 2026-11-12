// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksgenieagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabricksgenieagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference interface {
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
	ExampleQuestionSqls() DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsExampleQuestionSqlsList
	ExampleQuestionSqlsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JoinSpecs() DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsJoinSpecsList
	JoinSpecsInput() interface{}
	SqlFunctions() DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsSqlFunctionsList
	SqlFunctionsInput() interface{}
	SqlSnippets() DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsSqlSnippetsOutputReference
	SqlSnippetsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TextInstructions() DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsTextInstructionsList
	TextInstructionsInput() interface{}
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
	PutExampleQuestionSqls(value interface{})
	PutJoinSpecs(value interface{})
	PutSqlFunctions(value interface{})
	PutSqlSnippets(value *DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsSqlSnippets)
	PutTextInstructions(value interface{})
	ResetExampleQuestionSqls()
	ResetJoinSpecs()
	ResetSqlFunctions()
	ResetSqlSnippets()
	ResetTextInstructions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference
type jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) ExampleQuestionSqls() DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsExampleQuestionSqlsList {
	var returns DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsExampleQuestionSqlsList
	_jsii_.Get(
		j,
		"exampleQuestionSqls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) ExampleQuestionSqlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exampleQuestionSqlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) JoinSpecs() DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsJoinSpecsList {
	var returns DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsJoinSpecsList
	_jsii_.Get(
		j,
		"joinSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) JoinSpecsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"joinSpecsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) SqlFunctions() DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsSqlFunctionsList {
	var returns DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsSqlFunctionsList
	_jsii_.Get(
		j,
		"sqlFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) SqlFunctionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlFunctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) SqlSnippets() DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsSqlSnippetsOutputReference {
	var returns DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsSqlSnippetsOutputReference
	_jsii_.Get(
		j,
		"sqlSnippets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) SqlSnippetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlSnippetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) TextInstructions() DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsTextInstructionsList {
	var returns DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsTextInstructionsList
	_jsii_.Get(
		j,
		"textInstructions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) TextInstructionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"textInstructionsInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksGenieAgents.DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference_Override(d DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksGenieAgents.DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) PutExampleQuestionSqls(value interface{}) {
	if err := d.validatePutExampleQuestionSqlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExampleQuestionSqls",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) PutJoinSpecs(value interface{}) {
	if err := d.validatePutJoinSpecsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJoinSpecs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) PutSqlFunctions(value interface{}) {
	if err := d.validatePutSqlFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSqlFunctions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) PutSqlSnippets(value *DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsSqlSnippets) {
	if err := d.validatePutSqlSnippetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSqlSnippets",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) PutTextInstructions(value interface{}) {
	if err := d.validatePutTextInstructionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTextInstructions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) ResetExampleQuestionSqls() {
	_jsii_.InvokeVoid(
		d,
		"resetExampleQuestionSqls",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) ResetJoinSpecs() {
	_jsii_.InvokeVoid(
		d,
		"resetJoinSpecs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) ResetSqlFunctions() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlFunctions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) ResetSqlSnippets() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlSnippets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) ResetTextInstructions() {
	_jsii_.InvokeVoid(
		d,
		"resetTextInstructions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

