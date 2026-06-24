// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sqlglobalconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/sqlglobalconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/sql_global_config databricks_sql_global_config}.
type SqlGlobalConfig interface {
	cdktn.TerraformResource
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
	DataAccessConfig() *map[string]*string
	SetDataAccessConfig(val *map[string]*string)
	DataAccessConfigInput() *map[string]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableServerlessCompute() interface{}
	SetEnableServerlessCompute(val interface{})
	EnableServerlessComputeInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GoogleServiceAccount() *string
	SetGoogleServiceAccount(val *string)
	GoogleServiceAccountInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceProfileArn() *string
	SetInstanceProfileArn(val *string)
	InstanceProfileArnInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProviderConfig() SqlGlobalConfigProviderConfigOutputReference
	ProviderConfigInput() *SqlGlobalConfigProviderConfig
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SecurityPolicy() *string
	SetSecurityPolicy(val *string)
	SecurityPolicyInput() *string
	SqlConfigParams() *map[string]*string
	SetSqlConfigParams(val *map[string]*string)
	SqlConfigParamsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutProviderConfig(value *SqlGlobalConfigProviderConfig)
	ResetDataAccessConfig()
	ResetEnableServerlessCompute()
	ResetGoogleServiceAccount()
	ResetId()
	ResetInstanceProfileArn()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProviderConfig()
	ResetSecurityPolicy()
	ResetSqlConfigParams()
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

// The jsii proxy struct for SqlGlobalConfig
type jsiiProxy_SqlGlobalConfig struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_SqlGlobalConfig) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) DataAccessConfig() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dataAccessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) DataAccessConfigInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dataAccessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) EnableServerlessCompute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableServerlessCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) EnableServerlessComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableServerlessComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) GoogleServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) GoogleServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) InstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) InstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) ProviderConfig() SqlGlobalConfigProviderConfigOutputReference {
	var returns SqlGlobalConfigProviderConfigOutputReference
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) ProviderConfigInput() *SqlGlobalConfigProviderConfig {
	var returns *SqlGlobalConfigProviderConfig
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) SecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) SecurityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) SqlConfigParams() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sqlConfigParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) SqlConfigParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sqlConfigParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlGlobalConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/sql_global_config databricks_sql_global_config} Resource.
func NewSqlGlobalConfig(scope constructs.Construct, id *string, config *SqlGlobalConfigConfig) SqlGlobalConfig {
	_init_.Initialize()

	if err := validateNewSqlGlobalConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqlGlobalConfig{}

	_jsii_.Create(
		"@cdktn/provider-databricks.sqlGlobalConfig.SqlGlobalConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/sql_global_config databricks_sql_global_config} Resource.
func NewSqlGlobalConfig_Override(s SqlGlobalConfig, scope constructs.Construct, id *string, config *SqlGlobalConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.sqlGlobalConfig.SqlGlobalConfig",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SqlGlobalConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SqlGlobalConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SqlGlobalConfig)SetDataAccessConfig(val *map[string]*string) {
	if err := j.validateSetDataAccessConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataAccessConfig",
		val,
	)
}

func (j *jsiiProxy_SqlGlobalConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SqlGlobalConfig)SetEnableServerlessCompute(val interface{}) {
	if err := j.validateSetEnableServerlessComputeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableServerlessCompute",
		val,
	)
}

func (j *jsiiProxy_SqlGlobalConfig)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SqlGlobalConfig)SetGoogleServiceAccount(val *string) {
	if err := j.validateSetGoogleServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleServiceAccount",
		val,
	)
}

func (j *jsiiProxy_SqlGlobalConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SqlGlobalConfig)SetInstanceProfileArn(val *string) {
	if err := j.validateSetInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_SqlGlobalConfig)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SqlGlobalConfig)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SqlGlobalConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SqlGlobalConfig)SetSecurityPolicy(val *string) {
	if err := j.validateSetSecurityPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityPolicy",
		val,
	)
}

func (j *jsiiProxy_SqlGlobalConfig)SetSqlConfigParams(val *map[string]*string) {
	if err := j.validateSetSqlConfigParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlConfigParams",
		val,
	)
}

// Generates CDKTN code for importing a SqlGlobalConfig resource upon running "cdktn plan <stack-name>".
func SqlGlobalConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateSqlGlobalConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.sqlGlobalConfig.SqlGlobalConfig",
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
func SqlGlobalConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSqlGlobalConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.sqlGlobalConfig.SqlGlobalConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SqlGlobalConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSqlGlobalConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.sqlGlobalConfig.SqlGlobalConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SqlGlobalConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSqlGlobalConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.sqlGlobalConfig.SqlGlobalConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SqlGlobalConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-databricks.sqlGlobalConfig.SqlGlobalConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SqlGlobalConfig) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SqlGlobalConfig) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SqlGlobalConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SqlGlobalConfig) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SqlGlobalConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SqlGlobalConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SqlGlobalConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SqlGlobalConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SqlGlobalConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SqlGlobalConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SqlGlobalConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SqlGlobalConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlGlobalConfig) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SqlGlobalConfig) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SqlGlobalConfig) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SqlGlobalConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SqlGlobalConfig) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SqlGlobalConfig) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SqlGlobalConfig) PutProviderConfig(value *SqlGlobalConfigProviderConfig) {
	if err := s.validatePutProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProviderConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SqlGlobalConfig) ResetDataAccessConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetDataAccessConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlGlobalConfig) ResetEnableServerlessCompute() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableServerlessCompute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlGlobalConfig) ResetGoogleServiceAccount() {
	_jsii_.InvokeVoid(
		s,
		"resetGoogleServiceAccount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlGlobalConfig) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlGlobalConfig) ResetInstanceProfileArn() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceProfileArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlGlobalConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlGlobalConfig) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlGlobalConfig) ResetSecurityPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetSecurityPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlGlobalConfig) ResetSqlConfigParams() {
	_jsii_.InvokeVoid(
		s,
		"resetSqlConfigParams",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlGlobalConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlGlobalConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlGlobalConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlGlobalConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlGlobalConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlGlobalConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlGlobalConfig) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

