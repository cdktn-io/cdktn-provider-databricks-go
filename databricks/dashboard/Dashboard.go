// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/dashboard/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/dashboard databricks_dashboard}.
type Dashboard interface {
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
	CreateTime() *string
	SetCreateTime(val *string)
	CreateTimeInput() *string
	DashboardChangeDetected() interface{}
	SetDashboardChangeDetected(val interface{})
	DashboardChangeDetectedInput() interface{}
	DashboardId() *string
	SetDashboardId(val *string)
	DashboardIdInput() *string
	DatasetCatalog() *string
	SetDatasetCatalog(val *string)
	DatasetCatalogInput() *string
	DatasetSchema() *string
	SetDatasetSchema(val *string)
	DatasetSchemaInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EmbedCredentials() interface{}
	SetEmbedCredentials(val interface{})
	EmbedCredentialsInput() interface{}
	Etag() *string
	SetEtag(val *string)
	EtagInput() *string
	FilePath() *string
	SetFilePath(val *string)
	FilePathInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LifecycleState() *string
	SetLifecycleState(val *string)
	LifecycleStateInput() *string
	Md5() *string
	SetMd5(val *string)
	Md5Input() *string
	// The tree node.
	Node() constructs.Node
	ParentPath() *string
	SetParentPath(val *string)
	ParentPathInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProviderConfig() DashboardProviderConfigOutputReference
	ProviderConfigInput() *DashboardProviderConfig
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SerializedDashboard() *string
	SetSerializedDashboard(val *string)
	SerializedDashboardInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdateTime() *string
	SetUpdateTime(val *string)
	UpdateTimeInput() *string
	WarehouseId() *string
	SetWarehouseId(val *string)
	WarehouseIdInput() *string
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
	PutProviderConfig(value *DashboardProviderConfig)
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
	ResetCreateTime()
	ResetDashboardChangeDetected()
	ResetDashboardId()
	ResetDatasetCatalog()
	ResetDatasetSchema()
	ResetEmbedCredentials()
	ResetEtag()
	ResetFilePath()
	ResetId()
	ResetLifecycleState()
	ResetMd5()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPath()
	ResetProviderConfig()
	ResetSerializedDashboard()
	ResetUpdateTime()
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

// The jsii proxy struct for Dashboard
type jsiiProxy_Dashboard struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_Dashboard) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) CreateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) DashboardChangeDetected() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dashboardChangeDetected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) DashboardChangeDetectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dashboardChangeDetectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) DashboardId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) DashboardIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) DatasetCatalog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetCatalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) DatasetCatalogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetCatalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) DatasetSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) DatasetSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) EmbedCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embedCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) EmbedCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embedCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) EtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) FilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) FilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) LifecycleState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) LifecycleStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Md5() *string {
	var returns *string
	_jsii_.Get(
		j,
		"md5",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Md5Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"md5Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) ParentPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) ParentPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) ProviderConfig() DashboardProviderConfigOutputReference {
	var returns DashboardProviderConfigOutputReference
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) ProviderConfigInput() *DashboardProviderConfig {
	var returns *DashboardProviderConfig
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) SerializedDashboard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serializedDashboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) SerializedDashboardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serializedDashboardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) UpdateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) WarehouseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) WarehouseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/dashboard databricks_dashboard} Resource.
func NewDashboard(scope constructs.Construct, id *string, config *DashboardConfig) Dashboard {
	_init_.Initialize()

	if err := validateNewDashboardParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Dashboard{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dashboard.Dashboard",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/dashboard databricks_dashboard} Resource.
func NewDashboard_Override(d Dashboard, scope constructs.Construct, id *string, config *DashboardConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dashboard.Dashboard",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_Dashboard)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetCreateTime(val *string) {
	if err := j.validateSetCreateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createTime",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetDashboardChangeDetected(val interface{}) {
	if err := j.validateSetDashboardChangeDetectedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dashboardChangeDetected",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetDashboardId(val *string) {
	if err := j.validateSetDashboardIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dashboardId",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetDatasetCatalog(val *string) {
	if err := j.validateSetDatasetCatalogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetCatalog",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetDatasetSchema(val *string) {
	if err := j.validateSetDatasetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetSchema",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetEmbedCredentials(val interface{}) {
	if err := j.validateSetEmbedCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"embedCredentials",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetEtag(val *string) {
	if err := j.validateSetEtagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etag",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetFilePath(val *string) {
	if err := j.validateSetFilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filePath",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetLifecycleState(val *string) {
	if err := j.validateSetLifecycleStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleState",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetMd5(val *string) {
	if err := j.validateSetMd5Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"md5",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetParentPath(val *string) {
	if err := j.validateSetParentPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentPath",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetSerializedDashboard(val *string) {
	if err := j.validateSetSerializedDashboardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serializedDashboard",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetUpdateTime(val *string) {
	if err := j.validateSetUpdateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateTime",
		val,
	)
}

func (j *jsiiProxy_Dashboard)SetWarehouseId(val *string) {
	if err := j.validateSetWarehouseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warehouseId",
		val,
	)
}

// Generates CDKTN code for importing a Dashboard resource upon running "cdktn plan <stack-name>".
func Dashboard_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDashboard_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.dashboard.Dashboard",
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
func Dashboard_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDashboard_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.dashboard.Dashboard",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Dashboard_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDashboard_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.dashboard.Dashboard",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Dashboard_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDashboard_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.dashboard.Dashboard",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Dashboard_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-databricks.dashboard.Dashboard",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_Dashboard) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_Dashboard) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_Dashboard) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_Dashboard) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_Dashboard) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_Dashboard) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_Dashboard) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_Dashboard) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_Dashboard) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_Dashboard) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_Dashboard) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_Dashboard) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_Dashboard) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_Dashboard) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := d.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_Dashboard) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_Dashboard) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_Dashboard) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_Dashboard) PutProviderConfig(value *DashboardProviderConfig) {
	if err := d.validatePutProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProviderConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_Dashboard) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_Dashboard) ResetCreateTime() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetDashboardChangeDetected() {
	_jsii_.InvokeVoid(
		d,
		"resetDashboardChangeDetected",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetDashboardId() {
	_jsii_.InvokeVoid(
		d,
		"resetDashboardId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetDatasetCatalog() {
	_jsii_.InvokeVoid(
		d,
		"resetDatasetCatalog",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetDatasetSchema() {
	_jsii_.InvokeVoid(
		d,
		"resetDatasetSchema",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetEmbedCredentials() {
	_jsii_.InvokeVoid(
		d,
		"resetEmbedCredentials",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetEtag() {
	_jsii_.InvokeVoid(
		d,
		"resetEtag",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetFilePath() {
	_jsii_.InvokeVoid(
		d,
		"resetFilePath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetLifecycleState() {
	_jsii_.InvokeVoid(
		d,
		"resetLifecycleState",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetMd5() {
	_jsii_.InvokeVoid(
		d,
		"resetMd5",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetPath() {
	_jsii_.InvokeVoid(
		d,
		"resetPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetSerializedDashboard() {
	_jsii_.InvokeVoid(
		d,
		"resetSerializedDashboard",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetUpdateTime() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdateTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

