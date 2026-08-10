// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package app

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/app/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/app databricks_app}.
type App interface {
	cdktn.TerraformResource
	ActiveDeployment() AppActiveDeploymentOutputReference
	AppStatus() AppAppStatusOutputReference
	BudgetPolicyId() *string
	SetBudgetPolicyId(val *string)
	BudgetPolicyIdInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ComputeMaxInstances() *float64
	SetComputeMaxInstances(val *float64)
	ComputeMaxInstancesInput() *float64
	ComputeMinInstances() *float64
	SetComputeMinInstances(val *float64)
	ComputeMinInstancesInput() *float64
	ComputeSize() *string
	SetComputeSize(val *string)
	ComputeSizeInput() *string
	ComputeStatus() AppComputeStatusOutputReference
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
	CreateTime() *string
	Creator() *string
	DefaultSourceCodePath() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveBudgetPolicyId() *string
	EffectiveUsagePolicyId() *string
	EffectiveUserApiScopes() *[]*string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GitRepository() AppGitRepositoryOutputReference
	GitRepositoryInput() interface{}
	Id() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NoCompute() interface{}
	SetNoCompute(val interface{})
	NoComputeInput() interface{}
	// The tree node.
	Node() constructs.Node
	Oauth2AppClientId() *string
	Oauth2AppIntegrationId() *string
	PendingDeployment() AppPendingDeploymentOutputReference
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProviderConfig() AppProviderConfigOutputReference
	ProviderConfigInput() interface{}
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Resources() AppResourcesList
	ResourcesInput() interface{}
	ServicePrincipalClientId() *string
	ServicePrincipalId() *float64
	ServicePrincipalName() *string
	Space() *string
	SetSpace(val *string)
	SpaceInput() *string
	TelemetryExportDestinations() AppTelemetryExportDestinationsList
	TelemetryExportDestinationsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThumbnailUrl() *string
	Updater() *string
	UpdateTime() *string
	Url() *string
	UsagePolicyId() *string
	SetUsagePolicyId(val *string)
	UsagePolicyIdInput() *string
	UserApiScopes() *[]*string
	SetUserApiScopes(val *[]*string)
	UserApiScopesInput() *[]*string
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
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
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
	PutGitRepository(value *AppGitRepository)
	PutProviderConfig(value *AppProviderConfig)
	PutResources(value interface{})
	PutTelemetryExportDestinations(value interface{})
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetBudgetPolicyId()
	ResetComputeMaxInstances()
	ResetComputeMinInstances()
	ResetComputeSize()
	ResetDescription()
	ResetGitRepository()
	ResetNoCompute()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProviderConfig()
	ResetResources()
	ResetSpace()
	ResetTelemetryExportDestinations()
	ResetUsagePolicyId()
	ResetUserApiScopes()
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

// The jsii proxy struct for App
type jsiiProxy_App struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_App) ActiveDeployment() AppActiveDeploymentOutputReference {
	var returns AppActiveDeploymentOutputReference
	_jsii_.Get(
		j,
		"activeDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) AppStatus() AppAppStatusOutputReference {
	var returns AppAppStatusOutputReference
	_jsii_.Get(
		j,
		"appStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) BudgetPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) BudgetPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ComputeMaxInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeMaxInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ComputeMaxInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeMaxInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ComputeMinInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeMinInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ComputeMinInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeMinInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ComputeSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ComputeSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ComputeStatus() AppComputeStatusOutputReference {
	var returns AppComputeStatusOutputReference
	_jsii_.Get(
		j,
		"computeStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Creator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) DefaultSourceCodePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSourceCodePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) EffectiveBudgetPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveBudgetPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) EffectiveUsagePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveUsagePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) EffectiveUserApiScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"effectiveUserApiScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) GitRepository() AppGitRepositoryOutputReference {
	var returns AppGitRepositoryOutputReference
	_jsii_.Get(
		j,
		"gitRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) GitRepositoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gitRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) NoCompute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) NoComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Oauth2AppClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2AppClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Oauth2AppIntegrationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2AppIntegrationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) PendingDeployment() AppPendingDeploymentOutputReference {
	var returns AppPendingDeploymentOutputReference
	_jsii_.Get(
		j,
		"pendingDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ProviderConfig() AppProviderConfigOutputReference {
	var returns AppProviderConfigOutputReference
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Resources() AppResourcesList {
	var returns AppResourcesList
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ServicePrincipalClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ServicePrincipalId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"servicePrincipalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ServicePrincipalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Space() *string {
	var returns *string
	_jsii_.Get(
		j,
		"space",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) SpaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) TelemetryExportDestinations() AppTelemetryExportDestinationsList {
	var returns AppTelemetryExportDestinationsList
	_jsii_.Get(
		j,
		"telemetryExportDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) TelemetryExportDestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"telemetryExportDestinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ThumbnailUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbnailUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Updater() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updater",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) UsagePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) UsagePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) UserApiScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userApiScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) UserApiScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userApiScopesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/app databricks_app} Resource.
func NewApp(scope constructs.Construct, id *string, config *AppConfig) App {
	_init_.Initialize()

	if err := validateNewAppParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_App{}

	_jsii_.Create(
		"@cdktn/provider-databricks.app.App",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/app databricks_app} Resource.
func NewApp_Override(a App, scope constructs.Construct, id *string, config *AppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.app.App",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_App)SetBudgetPolicyId(val *string) {
	if err := j.validateSetBudgetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"budgetPolicyId",
		val,
	)
}

func (j *jsiiProxy_App)SetComputeMaxInstances(val *float64) {
	if err := j.validateSetComputeMaxInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeMaxInstances",
		val,
	)
}

func (j *jsiiProxy_App)SetComputeMinInstances(val *float64) {
	if err := j.validateSetComputeMinInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeMinInstances",
		val,
	)
}

func (j *jsiiProxy_App)SetComputeSize(val *string) {
	if err := j.validateSetComputeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeSize",
		val,
	)
}

func (j *jsiiProxy_App)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_App)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_App)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_App)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_App)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_App)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_App)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_App)SetNoCompute(val interface{}) {
	if err := j.validateSetNoComputeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noCompute",
		val,
	)
}

func (j *jsiiProxy_App)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_App)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_App)SetSpace(val *string) {
	if err := j.validateSetSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"space",
		val,
	)
}

func (j *jsiiProxy_App)SetUsagePolicyId(val *string) {
	if err := j.validateSetUsagePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usagePolicyId",
		val,
	)
}

func (j *jsiiProxy_App)SetUserApiScopes(val *[]*string) {
	if err := j.validateSetUserApiScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userApiScopes",
		val,
	)
}

// Generates CDKTN code for importing a App resource upon running "cdktn plan <stack-name>".
func App_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateApp_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.app.App",
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
func App_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.app.App",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func App_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.app.App",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func App_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.app.App",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func App_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-databricks.app.App",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_App) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_App) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_App) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_App) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := a.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_App) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_App) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_App) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_App) PutGitRepository(value *AppGitRepository) {
	if err := a.validatePutGitRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGitRepository",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_App) PutProviderConfig(value *AppProviderConfig) {
	if err := a.validatePutProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProviderConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_App) PutResources(value interface{}) {
	if err := a.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResources",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_App) PutTelemetryExportDestinations(value interface{}) {
	if err := a.validatePutTelemetryExportDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTelemetryExportDestinations",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_App) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_App) ResetBudgetPolicyId() {
	_jsii_.InvokeVoid(
		a,
		"resetBudgetPolicyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_App) ResetComputeMaxInstances() {
	_jsii_.InvokeVoid(
		a,
		"resetComputeMaxInstances",
		nil, // no parameters
	)
}

func (a *jsiiProxy_App) ResetComputeMinInstances() {
	_jsii_.InvokeVoid(
		a,
		"resetComputeMinInstances",
		nil, // no parameters
	)
}

func (a *jsiiProxy_App) ResetComputeSize() {
	_jsii_.InvokeVoid(
		a,
		"resetComputeSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_App) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_App) ResetGitRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetGitRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_App) ResetNoCompute() {
	_jsii_.InvokeVoid(
		a,
		"resetNoCompute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_App) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_App) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_App) ResetResources() {
	_jsii_.InvokeVoid(
		a,
		"resetResources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_App) ResetSpace() {
	_jsii_.InvokeVoid(
		a,
		"resetSpace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_App) ResetTelemetryExportDestinations() {
	_jsii_.InvokeVoid(
		a,
		"resetTelemetryExportDestinations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_App) ResetUsagePolicyId() {
	_jsii_.InvokeVoid(
		a,
		"resetUsagePolicyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_App) ResetUserApiScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetUserApiScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_App) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

