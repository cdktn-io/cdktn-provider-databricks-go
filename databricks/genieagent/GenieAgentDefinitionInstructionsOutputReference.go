// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package genieagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/genieagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GenieAgentDefinitionInstructionsOutputReference interface {
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
	ExampleQuestionSqls() GenieAgentDefinitionInstructionsExampleQuestionSqlsList
	ExampleQuestionSqlsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JoinSpecs() GenieAgentDefinitionInstructionsJoinSpecsList
	JoinSpecsInput() interface{}
	SqlFunctions() GenieAgentDefinitionInstructionsSqlFunctionsList
	SqlFunctionsInput() interface{}
	SqlSnippets() GenieAgentDefinitionInstructionsSqlSnippetsOutputReference
	SqlSnippetsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TextInstructions() GenieAgentDefinitionInstructionsTextInstructionsList
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
	PutSqlSnippets(value *GenieAgentDefinitionInstructionsSqlSnippets)
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

// The jsii proxy struct for GenieAgentDefinitionInstructionsOutputReference
type jsiiProxy_GenieAgentDefinitionInstructionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) ExampleQuestionSqls() GenieAgentDefinitionInstructionsExampleQuestionSqlsList {
	var returns GenieAgentDefinitionInstructionsExampleQuestionSqlsList
	_jsii_.Get(
		j,
		"exampleQuestionSqls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) ExampleQuestionSqlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exampleQuestionSqlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) JoinSpecs() GenieAgentDefinitionInstructionsJoinSpecsList {
	var returns GenieAgentDefinitionInstructionsJoinSpecsList
	_jsii_.Get(
		j,
		"joinSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) JoinSpecsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"joinSpecsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) SqlFunctions() GenieAgentDefinitionInstructionsSqlFunctionsList {
	var returns GenieAgentDefinitionInstructionsSqlFunctionsList
	_jsii_.Get(
		j,
		"sqlFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) SqlFunctionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlFunctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) SqlSnippets() GenieAgentDefinitionInstructionsSqlSnippetsOutputReference {
	var returns GenieAgentDefinitionInstructionsSqlSnippetsOutputReference
	_jsii_.Get(
		j,
		"sqlSnippets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) SqlSnippetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlSnippetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) TextInstructions() GenieAgentDefinitionInstructionsTextInstructionsList {
	var returns GenieAgentDefinitionInstructionsTextInstructionsList
	_jsii_.Get(
		j,
		"textInstructions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) TextInstructionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"textInstructionsInput",
		&returns,
	)
	return returns
}


func NewGenieAgentDefinitionInstructionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GenieAgentDefinitionInstructionsOutputReference {
	_init_.Initialize()

	if err := validateNewGenieAgentDefinitionInstructionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GenieAgentDefinitionInstructionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.genieAgent.GenieAgentDefinitionInstructionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGenieAgentDefinitionInstructionsOutputReference_Override(g GenieAgentDefinitionInstructionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.genieAgent.GenieAgentDefinitionInstructionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) PutExampleQuestionSqls(value interface{}) {
	if err := g.validatePutExampleQuestionSqlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExampleQuestionSqls",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) PutJoinSpecs(value interface{}) {
	if err := g.validatePutJoinSpecsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putJoinSpecs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) PutSqlFunctions(value interface{}) {
	if err := g.validatePutSqlFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSqlFunctions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) PutSqlSnippets(value *GenieAgentDefinitionInstructionsSqlSnippets) {
	if err := g.validatePutSqlSnippetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSqlSnippets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) PutTextInstructions(value interface{}) {
	if err := g.validatePutTextInstructionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTextInstructions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) ResetExampleQuestionSqls() {
	_jsii_.InvokeVoid(
		g,
		"resetExampleQuestionSqls",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) ResetJoinSpecs() {
	_jsii_.InvokeVoid(
		g,
		"resetJoinSpecs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) ResetSqlFunctions() {
	_jsii_.InvokeVoid(
		g,
		"resetSqlFunctions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) ResetSqlSnippets() {
	_jsii_.InvokeVoid(
		g,
		"resetSqlSnippets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) ResetTextInstructions() {
	_jsii_.InvokeVoid(
		g,
		"resetTextInstructions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenieAgentDefinitionInstructionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

