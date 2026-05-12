// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package supervisoragenttool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/supervisoragenttool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/supervisor_agent_tool databricks_supervisor_agent_tool}.
type SupervisorAgentTool interface {
	cdktn.TerraformResource
	App() SupervisorAgentToolAppOutputReference
	AppInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GenieSpace() SupervisorAgentToolGenieSpaceOutputReference
	GenieSpaceInput() interface{}
	Id() *string
	KnowledgeAssistant() SupervisorAgentToolKnowledgeAssistantOutputReference
	KnowledgeAssistantInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	Parent() *string
	SetParent(val *string)
	ParentInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProviderConfig() SupervisorAgentToolProviderConfigOutputReference
	ProviderConfigInput() interface{}
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ToolId() *string
	SetToolId(val *string)
	ToolIdInput() *string
	ToolType() *string
	SetToolType(val *string)
	ToolTypeInput() *string
	UcConnection() SupervisorAgentToolUcConnectionOutputReference
	UcConnectionInput() interface{}
	UcFunction() SupervisorAgentToolUcFunctionOutputReference
	UcFunctionInput() interface{}
	Volume() SupervisorAgentToolVolumeOutputReference
	VolumeInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutApp(value *SupervisorAgentToolApp)
	PutGenieSpace(value *SupervisorAgentToolGenieSpace)
	PutKnowledgeAssistant(value *SupervisorAgentToolKnowledgeAssistant)
	PutProviderConfig(value *SupervisorAgentToolProviderConfig)
	PutUcConnection(value *SupervisorAgentToolUcConnection)
	PutUcFunction(value *SupervisorAgentToolUcFunction)
	PutVolume(value *SupervisorAgentToolVolume)
	ResetApp()
	ResetDescription()
	ResetGenieSpace()
	ResetKnowledgeAssistant()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProviderConfig()
	ResetUcConnection()
	ResetUcFunction()
	ResetVolume()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for SupervisorAgentTool
type jsiiProxy_SupervisorAgentTool struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_SupervisorAgentTool) App() SupervisorAgentToolAppOutputReference {
	var returns SupervisorAgentToolAppOutputReference
	_jsii_.Get(
		j,
		"app",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) AppInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) GenieSpace() SupervisorAgentToolGenieSpaceOutputReference {
	var returns SupervisorAgentToolGenieSpaceOutputReference
	_jsii_.Get(
		j,
		"genieSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) GenieSpaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"genieSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) KnowledgeAssistant() SupervisorAgentToolKnowledgeAssistantOutputReference {
	var returns SupervisorAgentToolKnowledgeAssistantOutputReference
	_jsii_.Get(
		j,
		"knowledgeAssistant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) KnowledgeAssistantInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"knowledgeAssistantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) Parent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) ParentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) ProviderConfig() SupervisorAgentToolProviderConfigOutputReference {
	var returns SupervisorAgentToolProviderConfigOutputReference
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) ToolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) ToolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) ToolType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) ToolTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) UcConnection() SupervisorAgentToolUcConnectionOutputReference {
	var returns SupervisorAgentToolUcConnectionOutputReference
	_jsii_.Get(
		j,
		"ucConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) UcConnectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ucConnectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) UcFunction() SupervisorAgentToolUcFunctionOutputReference {
	var returns SupervisorAgentToolUcFunctionOutputReference
	_jsii_.Get(
		j,
		"ucFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) UcFunctionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ucFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) Volume() SupervisorAgentToolVolumeOutputReference {
	var returns SupervisorAgentToolVolumeOutputReference
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupervisorAgentTool) VolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/supervisor_agent_tool databricks_supervisor_agent_tool} Resource.
func NewSupervisorAgentTool(scope constructs.Construct, id *string, config *SupervisorAgentToolConfig) SupervisorAgentTool {
	_init_.Initialize()

	if err := validateNewSupervisorAgentToolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SupervisorAgentTool{}

	_jsii_.Create(
		"@cdktn/provider-databricks.supervisorAgentTool.SupervisorAgentTool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/supervisor_agent_tool databricks_supervisor_agent_tool} Resource.
func NewSupervisorAgentTool_Override(s SupervisorAgentTool, scope constructs.Construct, id *string, config *SupervisorAgentToolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.supervisorAgentTool.SupervisorAgentTool",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SupervisorAgentTool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SupervisorAgentTool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SupervisorAgentTool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SupervisorAgentTool)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SupervisorAgentTool)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SupervisorAgentTool)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SupervisorAgentTool)SetParent(val *string) {
	if err := j.validateSetParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parent",
		val,
	)
}

func (j *jsiiProxy_SupervisorAgentTool)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SupervisorAgentTool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SupervisorAgentTool)SetToolId(val *string) {
	if err := j.validateSetToolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toolId",
		val,
	)
}

func (j *jsiiProxy_SupervisorAgentTool)SetToolType(val *string) {
	if err := j.validateSetToolTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toolType",
		val,
	)
}

// Generates CDKTN code for importing a SupervisorAgentTool resource upon running "cdktn plan <stack-name>".
func SupervisorAgentTool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateSupervisorAgentTool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.supervisorAgentTool.SupervisorAgentTool",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SupervisorAgentTool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSupervisorAgentTool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.supervisorAgentTool.SupervisorAgentTool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SupervisorAgentTool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSupervisorAgentTool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.supervisorAgentTool.SupervisorAgentTool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SupervisorAgentTool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSupervisorAgentTool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.supervisorAgentTool.SupervisorAgentTool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SupervisorAgentTool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-databricks.supervisorAgentTool.SupervisorAgentTool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SupervisorAgentTool) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SupervisorAgentTool) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SupervisorAgentTool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupervisorAgentTool) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupervisorAgentTool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupervisorAgentTool) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupervisorAgentTool) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupervisorAgentTool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupervisorAgentTool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupervisorAgentTool) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupervisorAgentTool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupervisorAgentTool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupervisorAgentTool) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SupervisorAgentTool) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupervisorAgentTool) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SupervisorAgentTool) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SupervisorAgentTool) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SupervisorAgentTool) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SupervisorAgentTool) PutApp(value *SupervisorAgentToolApp) {
	if err := s.validatePutAppParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putApp",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SupervisorAgentTool) PutGenieSpace(value *SupervisorAgentToolGenieSpace) {
	if err := s.validatePutGenieSpaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGenieSpace",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SupervisorAgentTool) PutKnowledgeAssistant(value *SupervisorAgentToolKnowledgeAssistant) {
	if err := s.validatePutKnowledgeAssistantParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putKnowledgeAssistant",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SupervisorAgentTool) PutProviderConfig(value *SupervisorAgentToolProviderConfig) {
	if err := s.validatePutProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProviderConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SupervisorAgentTool) PutUcConnection(value *SupervisorAgentToolUcConnection) {
	if err := s.validatePutUcConnectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putUcConnection",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SupervisorAgentTool) PutUcFunction(value *SupervisorAgentToolUcFunction) {
	if err := s.validatePutUcFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putUcFunction",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SupervisorAgentTool) PutVolume(value *SupervisorAgentToolVolume) {
	if err := s.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVolume",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SupervisorAgentTool) ResetApp() {
	_jsii_.InvokeVoid(
		s,
		"resetApp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SupervisorAgentTool) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SupervisorAgentTool) ResetGenieSpace() {
	_jsii_.InvokeVoid(
		s,
		"resetGenieSpace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SupervisorAgentTool) ResetKnowledgeAssistant() {
	_jsii_.InvokeVoid(
		s,
		"resetKnowledgeAssistant",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SupervisorAgentTool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SupervisorAgentTool) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SupervisorAgentTool) ResetUcConnection() {
	_jsii_.InvokeVoid(
		s,
		"resetUcConnection",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SupervisorAgentTool) ResetUcFunction() {
	_jsii_.InvokeVoid(
		s,
		"resetUcFunction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SupervisorAgentTool) ResetVolume() {
	_jsii_.InvokeVoid(
		s,
		"resetVolume",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SupervisorAgentTool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupervisorAgentTool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupervisorAgentTool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupervisorAgentTool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupervisorAgentTool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupervisorAgentTool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SupervisorAgentTool) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		s,
		"with",
		args,
		&returns,
	)

	return returns
}

