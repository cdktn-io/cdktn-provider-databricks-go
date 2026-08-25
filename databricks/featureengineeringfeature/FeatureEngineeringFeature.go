// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/featureengineeringfeature/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/feature_engineering_feature databricks_feature_engineering_feature}.
type FeatureEngineeringFeature interface {
	cdktn.TerraformResource
	CatalogName() *string
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
	CreatedAt() *string
	CreatedBy() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Entities() FeatureEngineeringFeatureEntitiesList
	EntitiesInput() interface{}
	FilterCondition() *string
	SetFilterCondition(val *string)
	FilterConditionInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FullName() *string
	SetFullName(val *string)
	FullNameInput() *string
	Function() FeatureEngineeringFeatureFunctionOutputReference
	FunctionInput() interface{}
	Inputs() *[]*string
	SetInputs(val *[]*string)
	InputsInput() *[]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LineageContext() FeatureEngineeringFeatureLineageContextOutputReference
	LineageContextInput() interface{}
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProviderConfig() FeatureEngineeringFeatureProviderConfigOutputReference
	ProviderConfigInput() interface{}
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SchemaName() *string
	Source() FeatureEngineeringFeatureSourceOutputReference
	SourceInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeseriesColumn() FeatureEngineeringFeatureTimeseriesColumnOutputReference
	TimeseriesColumnInput() interface{}
	TimeWindow() FeatureEngineeringFeatureTimeWindowOutputReference
	TimeWindowInput() interface{}
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
	PutEntities(value interface{})
	PutFunction(value *FeatureEngineeringFeatureFunction)
	PutLineageContext(value *FeatureEngineeringFeatureLineageContext)
	PutProviderConfig(value *FeatureEngineeringFeatureProviderConfig)
	PutSource(value *FeatureEngineeringFeatureSource)
	PutTimeseriesColumn(value *FeatureEngineeringFeatureTimeseriesColumn)
	PutTimeWindow(value *FeatureEngineeringFeatureTimeWindow)
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
	ResetDescription()
	ResetEntities()
	ResetFilterCondition()
	ResetInputs()
	ResetLineageContext()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProviderConfig()
	ResetTimeseriesColumn()
	ResetTimeWindow()
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

// The jsii proxy struct for FeatureEngineeringFeature
type jsiiProxy_FeatureEngineeringFeature struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_FeatureEngineeringFeature) CatalogName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) Entities() FeatureEngineeringFeatureEntitiesList {
	var returns FeatureEngineeringFeatureEntitiesList
	_jsii_.Get(
		j,
		"entities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) EntitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) FilterCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) FilterConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) FullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) FullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) Function() FeatureEngineeringFeatureFunctionOutputReference {
	var returns FeatureEngineeringFeatureFunctionOutputReference
	_jsii_.Get(
		j,
		"function",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) FunctionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) Inputs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) InputsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) LineageContext() FeatureEngineeringFeatureLineageContextOutputReference {
	var returns FeatureEngineeringFeatureLineageContextOutputReference
	_jsii_.Get(
		j,
		"lineageContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) LineageContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lineageContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) ProviderConfig() FeatureEngineeringFeatureProviderConfigOutputReference {
	var returns FeatureEngineeringFeatureProviderConfigOutputReference
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) SchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) Source() FeatureEngineeringFeatureSourceOutputReference {
	var returns FeatureEngineeringFeatureSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) SourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) TimeseriesColumn() FeatureEngineeringFeatureTimeseriesColumnOutputReference {
	var returns FeatureEngineeringFeatureTimeseriesColumnOutputReference
	_jsii_.Get(
		j,
		"timeseriesColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) TimeseriesColumnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeseriesColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) TimeWindow() FeatureEngineeringFeatureTimeWindowOutputReference {
	var returns FeatureEngineeringFeatureTimeWindowOutputReference
	_jsii_.Get(
		j,
		"timeWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeature) TimeWindowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeWindowInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/feature_engineering_feature databricks_feature_engineering_feature} Resource.
func NewFeatureEngineeringFeature(scope constructs.Construct, id *string, config *FeatureEngineeringFeatureConfig) FeatureEngineeringFeature {
	_init_.Initialize()

	if err := validateNewFeatureEngineeringFeatureParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FeatureEngineeringFeature{}

	_jsii_.Create(
		"@cdktn/provider-databricks.featureEngineeringFeature.FeatureEngineeringFeature",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/feature_engineering_feature databricks_feature_engineering_feature} Resource.
func NewFeatureEngineeringFeature_Override(f FeatureEngineeringFeature, scope constructs.Construct, id *string, config *FeatureEngineeringFeatureConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.featureEngineeringFeature.FeatureEngineeringFeature",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeature)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeature)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeature)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeature)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeature)SetFilterCondition(val *string) {
	if err := j.validateSetFilterConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterCondition",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeature)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeature)SetFullName(val *string) {
	if err := j.validateSetFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullName",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeature)SetInputs(val *[]*string) {
	if err := j.validateSetInputsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputs",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeature)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeature)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeature)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a FeatureEngineeringFeature resource upon running "cdktn plan <stack-name>".
func FeatureEngineeringFeature_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateFeatureEngineeringFeature_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.featureEngineeringFeature.FeatureEngineeringFeature",
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
func FeatureEngineeringFeature_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFeatureEngineeringFeature_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.featureEngineeringFeature.FeatureEngineeringFeature",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FeatureEngineeringFeature_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFeatureEngineeringFeature_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.featureEngineeringFeature.FeatureEngineeringFeature",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FeatureEngineeringFeature_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFeatureEngineeringFeature_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.featureEngineeringFeature.FeatureEngineeringFeature",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FeatureEngineeringFeature_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-databricks.featureEngineeringFeature.FeatureEngineeringFeature",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeature) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FeatureEngineeringFeature) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (f *jsiiProxy_FeatureEngineeringFeature) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FeatureEngineeringFeature) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FeatureEngineeringFeature) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FeatureEngineeringFeature) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FeatureEngineeringFeature) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FeatureEngineeringFeature) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FeatureEngineeringFeature) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FeatureEngineeringFeature) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeature) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (f *jsiiProxy_FeatureEngineeringFeature) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := f.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeature) MoveFromId(id *string) {
	if err := f.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveFromId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) MoveToId(id *string) {
	if err := f.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveToId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) PutEntities(value interface{}) {
	if err := f.validatePutEntitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putEntities",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) PutFunction(value *FeatureEngineeringFeatureFunction) {
	if err := f.validatePutFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putFunction",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) PutLineageContext(value *FeatureEngineeringFeatureLineageContext) {
	if err := f.validatePutLineageContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putLineageContext",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) PutProviderConfig(value *FeatureEngineeringFeatureProviderConfig) {
	if err := f.validatePutProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putProviderConfig",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) PutSource(value *FeatureEngineeringFeatureSource) {
	if err := f.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putSource",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) PutTimeseriesColumn(value *FeatureEngineeringFeatureTimeseriesColumn) {
	if err := f.validatePutTimeseriesColumnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTimeseriesColumn",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) PutTimeWindow(value *FeatureEngineeringFeatureTimeWindow) {
	if err := f.validatePutTimeWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTimeWindow",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := f.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) ResetDescription() {
	_jsii_.InvokeVoid(
		f,
		"resetDescription",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) ResetEntities() {
	_jsii_.InvokeVoid(
		f,
		"resetEntities",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) ResetFilterCondition() {
	_jsii_.InvokeVoid(
		f,
		"resetFilterCondition",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) ResetInputs() {
	_jsii_.InvokeVoid(
		f,
		"resetInputs",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) ResetLineageContext() {
	_jsii_.InvokeVoid(
		f,
		"resetLineageContext",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		f,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) ResetTimeseriesColumn() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeseriesColumn",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) ResetTimeWindow() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeWindow",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeature) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeature) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeature) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeature) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeature) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeature) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeature) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

