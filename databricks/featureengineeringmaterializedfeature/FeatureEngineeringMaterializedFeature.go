// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringmaterializedfeature

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/featureengineeringmaterializedfeature/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_materialized_feature databricks_feature_engineering_materialized_feature}.
type FeatureEngineeringMaterializedFeature interface {
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
	CronSchedule() *string
	SetCronSchedule(val *string)
	CronScheduleInput() *string
	CronScheduleTrigger() FeatureEngineeringMaterializedFeatureCronScheduleTriggerOutputReference
	CronScheduleTriggerInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FeatureName() *string
	SetFeatureName(val *string)
	FeatureNameInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IsOnline() cdktn.IResolvable
	LastMaterializationTime() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaterializedFeatureId() *string
	// The tree node.
	Node() constructs.Node
	OfflineStoreConfig() FeatureEngineeringMaterializedFeatureOfflineStoreConfigOutputReference
	OfflineStoreConfigInput() interface{}
	OnlineStoreConfig() FeatureEngineeringMaterializedFeatureOnlineStoreConfigOutputReference
	OnlineStoreConfigInput() interface{}
	PipelineScheduleState() *string
	SetPipelineScheduleState(val *string)
	PipelineScheduleStateInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProviderConfig() FeatureEngineeringMaterializedFeatureProviderConfigOutputReference
	ProviderConfigInput() interface{}
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	StreamingMode() FeatureEngineeringMaterializedFeatureStreamingModeOutputReference
	StreamingModeInput() interface{}
	TableName() *string
	TableTrigger() FeatureEngineeringMaterializedFeatureTableTriggerOutputReference
	TableTriggerInput() interface{}
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
	PutCronScheduleTrigger(value *FeatureEngineeringMaterializedFeatureCronScheduleTrigger)
	PutOfflineStoreConfig(value *FeatureEngineeringMaterializedFeatureOfflineStoreConfig)
	PutOnlineStoreConfig(value *FeatureEngineeringMaterializedFeatureOnlineStoreConfig)
	PutProviderConfig(value *FeatureEngineeringMaterializedFeatureProviderConfig)
	PutStreamingMode(value *FeatureEngineeringMaterializedFeatureStreamingMode)
	PutTableTrigger(value *FeatureEngineeringMaterializedFeatureTableTrigger)
	ResetCronSchedule()
	ResetCronScheduleTrigger()
	ResetOfflineStoreConfig()
	ResetOnlineStoreConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPipelineScheduleState()
	ResetProviderConfig()
	ResetStreamingMode()
	ResetTableTrigger()
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

// The jsii proxy struct for FeatureEngineeringMaterializedFeature
type jsiiProxy_FeatureEngineeringMaterializedFeature struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) CronSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cronSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) CronScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cronScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) CronScheduleTrigger() FeatureEngineeringMaterializedFeatureCronScheduleTriggerOutputReference {
	var returns FeatureEngineeringMaterializedFeatureCronScheduleTriggerOutputReference
	_jsii_.Get(
		j,
		"cronScheduleTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) CronScheduleTriggerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cronScheduleTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) FeatureName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) FeatureNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) IsOnline() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isOnline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) LastMaterializationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastMaterializationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) MaterializedFeatureId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"materializedFeatureId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) OfflineStoreConfig() FeatureEngineeringMaterializedFeatureOfflineStoreConfigOutputReference {
	var returns FeatureEngineeringMaterializedFeatureOfflineStoreConfigOutputReference
	_jsii_.Get(
		j,
		"offlineStoreConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) OfflineStoreConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"offlineStoreConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) OnlineStoreConfig() FeatureEngineeringMaterializedFeatureOnlineStoreConfigOutputReference {
	var returns FeatureEngineeringMaterializedFeatureOnlineStoreConfigOutputReference
	_jsii_.Get(
		j,
		"onlineStoreConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) OnlineStoreConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlineStoreConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) PipelineScheduleState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineScheduleState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) PipelineScheduleStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineScheduleStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) ProviderConfig() FeatureEngineeringMaterializedFeatureProviderConfigOutputReference {
	var returns FeatureEngineeringMaterializedFeatureProviderConfigOutputReference
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) StreamingMode() FeatureEngineeringMaterializedFeatureStreamingModeOutputReference {
	var returns FeatureEngineeringMaterializedFeatureStreamingModeOutputReference
	_jsii_.Get(
		j,
		"streamingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) StreamingModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) TableTrigger() FeatureEngineeringMaterializedFeatureTableTriggerOutputReference {
	var returns FeatureEngineeringMaterializedFeatureTableTriggerOutputReference
	_jsii_.Get(
		j,
		"tableTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) TableTriggerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_materialized_feature databricks_feature_engineering_materialized_feature} Resource.
func NewFeatureEngineeringMaterializedFeature(scope constructs.Construct, id *string, config *FeatureEngineeringMaterializedFeatureConfig) FeatureEngineeringMaterializedFeature {
	_init_.Initialize()

	if err := validateNewFeatureEngineeringMaterializedFeatureParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FeatureEngineeringMaterializedFeature{}

	_jsii_.Create(
		"@cdktn/provider-databricks.featureEngineeringMaterializedFeature.FeatureEngineeringMaterializedFeature",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_materialized_feature databricks_feature_engineering_materialized_feature} Resource.
func NewFeatureEngineeringMaterializedFeature_Override(f FeatureEngineeringMaterializedFeature, scope constructs.Construct, id *string, config *FeatureEngineeringMaterializedFeatureConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.featureEngineeringMaterializedFeature.FeatureEngineeringMaterializedFeature",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature)SetCronSchedule(val *string) {
	if err := j.validateSetCronScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cronSchedule",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature)SetFeatureName(val *string) {
	if err := j.validateSetFeatureNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"featureName",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature)SetPipelineScheduleState(val *string) {
	if err := j.validateSetPipelineScheduleStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineScheduleState",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringMaterializedFeature)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a FeatureEngineeringMaterializedFeature resource upon running "cdktn plan <stack-name>".
func FeatureEngineeringMaterializedFeature_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateFeatureEngineeringMaterializedFeature_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.featureEngineeringMaterializedFeature.FeatureEngineeringMaterializedFeature",
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
func FeatureEngineeringMaterializedFeature_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFeatureEngineeringMaterializedFeature_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.featureEngineeringMaterializedFeature.FeatureEngineeringMaterializedFeature",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FeatureEngineeringMaterializedFeature_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFeatureEngineeringMaterializedFeature_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.featureEngineeringMaterializedFeature.FeatureEngineeringMaterializedFeature",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FeatureEngineeringMaterializedFeature_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFeatureEngineeringMaterializedFeature_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.featureEngineeringMaterializedFeature.FeatureEngineeringMaterializedFeature",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FeatureEngineeringMaterializedFeature_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-databricks.featureEngineeringMaterializedFeature.FeatureEngineeringMaterializedFeature",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) MoveFromId(id *string) {
	if err := f.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveFromId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) MoveToId(id *string) {
	if err := f.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveToId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) PutCronScheduleTrigger(value *FeatureEngineeringMaterializedFeatureCronScheduleTrigger) {
	if err := f.validatePutCronScheduleTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putCronScheduleTrigger",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) PutOfflineStoreConfig(value *FeatureEngineeringMaterializedFeatureOfflineStoreConfig) {
	if err := f.validatePutOfflineStoreConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putOfflineStoreConfig",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) PutOnlineStoreConfig(value *FeatureEngineeringMaterializedFeatureOnlineStoreConfig) {
	if err := f.validatePutOnlineStoreConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putOnlineStoreConfig",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) PutProviderConfig(value *FeatureEngineeringMaterializedFeatureProviderConfig) {
	if err := f.validatePutProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putProviderConfig",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) PutStreamingMode(value *FeatureEngineeringMaterializedFeatureStreamingMode) {
	if err := f.validatePutStreamingModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putStreamingMode",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) PutTableTrigger(value *FeatureEngineeringMaterializedFeatureTableTrigger) {
	if err := f.validatePutTableTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTableTrigger",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) ResetCronSchedule() {
	_jsii_.InvokeVoid(
		f,
		"resetCronSchedule",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) ResetCronScheduleTrigger() {
	_jsii_.InvokeVoid(
		f,
		"resetCronScheduleTrigger",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) ResetOfflineStoreConfig() {
	_jsii_.InvokeVoid(
		f,
		"resetOfflineStoreConfig",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) ResetOnlineStoreConfig() {
	_jsii_.InvokeVoid(
		f,
		"resetOnlineStoreConfig",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) ResetPipelineScheduleState() {
	_jsii_.InvokeVoid(
		f,
		"resetPipelineScheduleState",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		f,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) ResetStreamingMode() {
	_jsii_.InvokeVoid(
		f,
		"resetStreamingMode",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) ResetTableTrigger() {
	_jsii_.InvokeVoid(
		f,
		"resetTableTrigger",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringMaterializedFeature) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		f,
		"with",
		args,
		&returns,
	)

	return returns
}

