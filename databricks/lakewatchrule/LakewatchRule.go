// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakewatchrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/lakewatchrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule databricks_lakewatch_rule}.
type LakewatchRule interface {
	cdktn.TerraformResource
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	Batch() LakewatchRuleBatchOutputReference
	BatchInput() interface{}
	CatalogName() *string
	SetCatalogName(val *string)
	CatalogNameInput() *string
	Category() *string
	SetCategory(val *string)
	CategoryInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ComputeWorkspaceId() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContentResourceId() *string
	SetContentResourceId(val *string)
	ContentResourceIdInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedBy() LakewatchRuleCreatedByOutputReference
	CreateTime() *string
	Deduplication() *string
	SetDeduplication(val *string)
	DeduplicationInput() *string
	DeduplicationTimeWindow() *string
	SetDeduplicationTimeWindow(val *string)
	DeduplicationTimeWindowInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveCatalogName() *string
	EffectiveRuleName() *string
	EffectiveRunAsPrincipal() *string
	EffectiveSchemaName() *string
	Etag() *string
	Fidelity() *string
	SetFidelity(val *string)
	FidelityInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Mitre() LakewatchRuleMitreList
	MitreInput() interface{}
	Name() *string
	// The tree node.
	Node() constructs.Node
	Objective() *string
	SetObjective(val *string)
	ObjectiveInput() *string
	OriginalContentPack() LakewatchRuleOriginalContentPackOutputReference
	Output() LakewatchRuleOutputOutputReference
	OutputInput() interface{}
	Owner() LakewatchRuleOwnerOutputReference
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProviderConfig() LakewatchRuleProviderConfigOutputReference
	ProviderConfigInput() interface{}
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	References() *[]*string
	SetReferences(val *[]*string)
	ReferencesInput() *[]*string
	RuleId() *string
	RuleName() *string
	SetRuleName(val *string)
	RuleNameInput() *string
	RunAsPrincipal() *string
	SetRunAsPrincipal(val *string)
	RunAsPrincipalInput() *string
	Runbook() *string
	SetRunbook(val *string)
	RunbookInput() *string
	Schedule() LakewatchRuleScheduleOutputReference
	ScheduleInput() interface{}
	SchemaName() *string
	SetSchemaName(val *string)
	SchemaNameInput() *string
	Severity() *string
	SetSeverity(val *string)
	SeverityInput() *string
	Status() LakewatchRuleStatusOutputReference
	SuppressUntilTime() *string
	SetSuppressUntilTime(val *string)
	SuppressUntilTimeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timerange() LakewatchRuleTimerangeOutputReference
	TimerangeInput() interface{}
	UpdatedBy() LakewatchRuleUpdatedByOutputReference
	UpdateTime() *string
	Version() *float64
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
	PutBatch(value *LakewatchRuleBatch)
	PutMitre(value interface{})
	PutOutput(value *LakewatchRuleOutput)
	PutProviderConfig(value *LakewatchRuleProviderConfig)
	PutSchedule(value *LakewatchRuleSchedule)
	PutTimerange(value *LakewatchRuleTimerange)
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
	ResetAnnotations()
	ResetBatch()
	ResetCategory()
	ResetContentResourceId()
	ResetDeduplication()
	ResetDeduplicationTimeWindow()
	ResetDescription()
	ResetMitre()
	ResetObjective()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProviderConfig()
	ResetReferences()
	ResetRunAsPrincipal()
	ResetRunbook()
	ResetSuppressUntilTime()
	ResetTimerange()
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

// The jsii proxy struct for LakewatchRule
type jsiiProxy_LakewatchRule struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_LakewatchRule) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Batch() LakewatchRuleBatchOutputReference {
	var returns LakewatchRuleBatchOutputReference
	_jsii_.Get(
		j,
		"batch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) BatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) CatalogName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) CatalogNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Category() *string {
	var returns *string
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) CategoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"categoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) ComputeWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) ContentResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) ContentResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) CreatedBy() LakewatchRuleCreatedByOutputReference {
	var returns LakewatchRuleCreatedByOutputReference
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Deduplication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deduplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) DeduplicationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deduplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) DeduplicationTimeWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deduplicationTimeWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) DeduplicationTimeWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deduplicationTimeWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) EffectiveCatalogName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveCatalogName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) EffectiveRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) EffectiveRunAsPrincipal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveRunAsPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) EffectiveSchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveSchemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Fidelity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fidelity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) FidelityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fidelityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Mitre() LakewatchRuleMitreList {
	var returns LakewatchRuleMitreList
	_jsii_.Get(
		j,
		"mitre",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) MitreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mitreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Objective() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objective",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) ObjectiveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) OriginalContentPack() LakewatchRuleOriginalContentPackOutputReference {
	var returns LakewatchRuleOriginalContentPackOutputReference
	_jsii_.Get(
		j,
		"originalContentPack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Output() LakewatchRuleOutputOutputReference {
	var returns LakewatchRuleOutputOutputReference
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) OutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Owner() LakewatchRuleOwnerOutputReference {
	var returns LakewatchRuleOwnerOutputReference
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) ProviderConfig() LakewatchRuleProviderConfigOutputReference {
	var returns LakewatchRuleProviderConfigOutputReference
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) References() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"references",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) ReferencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"referencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) RuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) RuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) RunAsPrincipal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) RunAsPrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsPrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Runbook() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runbook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) RunbookInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runbookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Schedule() LakewatchRuleScheduleOutputReference {
	var returns LakewatchRuleScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) SchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) SchemaNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Severity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) SeverityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Status() LakewatchRuleStatusOutputReference {
	var returns LakewatchRuleStatusOutputReference
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) SuppressUntilTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suppressUntilTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) SuppressUntilTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suppressUntilTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Timerange() LakewatchRuleTimerangeOutputReference {
	var returns LakewatchRuleTimerangeOutputReference
	_jsii_.Get(
		j,
		"timerange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) TimerangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timerangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) UpdatedBy() LakewatchRuleUpdatedByOutputReference {
	var returns LakewatchRuleUpdatedByOutputReference
	_jsii_.Get(
		j,
		"updatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRule) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule databricks_lakewatch_rule} Resource.
func NewLakewatchRule(scope constructs.Construct, id *string, config *LakewatchRuleConfig) LakewatchRule {
	_init_.Initialize()

	if err := validateNewLakewatchRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LakewatchRule{}

	_jsii_.Create(
		"@cdktn/provider-databricks.lakewatchRule.LakewatchRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule databricks_lakewatch_rule} Resource.
func NewLakewatchRule_Override(l LakewatchRule, scope constructs.Construct, id *string, config *LakewatchRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.lakewatchRule.LakewatchRule",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LakewatchRule)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetCatalogName(val *string) {
	if err := j.validateSetCatalogNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogName",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetCategory(val *string) {
	if err := j.validateSetCategoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"category",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetContentResourceId(val *string) {
	if err := j.validateSetContentResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentResourceId",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetDeduplication(val *string) {
	if err := j.validateSetDeduplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deduplication",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetDeduplicationTimeWindow(val *string) {
	if err := j.validateSetDeduplicationTimeWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deduplicationTimeWindow",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetFidelity(val *string) {
	if err := j.validateSetFidelityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fidelity",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetObjective(val *string) {
	if err := j.validateSetObjectiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objective",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetReferences(val *[]*string) {
	if err := j.validateSetReferencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"references",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetRuleName(val *string) {
	if err := j.validateSetRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleName",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetRunAsPrincipal(val *string) {
	if err := j.validateSetRunAsPrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsPrincipal",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetRunbook(val *string) {
	if err := j.validateSetRunbookParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runbook",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetSchemaName(val *string) {
	if err := j.validateSetSchemaNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaName",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetSeverity(val *string) {
	if err := j.validateSetSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_LakewatchRule)SetSuppressUntilTime(val *string) {
	if err := j.validateSetSuppressUntilTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suppressUntilTime",
		val,
	)
}

// Generates CDKTN code for importing a LakewatchRule resource upon running "cdktn plan <stack-name>".
func LakewatchRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateLakewatchRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.lakewatchRule.LakewatchRule",
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
func LakewatchRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLakewatchRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.lakewatchRule.LakewatchRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LakewatchRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLakewatchRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.lakewatchRule.LakewatchRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LakewatchRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLakewatchRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.lakewatchRule.LakewatchRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LakewatchRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-databricks.lakewatchRule.LakewatchRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LakewatchRule) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LakewatchRule) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LakewatchRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRule) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRule) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRule) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LakewatchRule) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRule) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := l.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRule) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LakewatchRule) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LakewatchRule) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LakewatchRule) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LakewatchRule) PutBatch(value *LakewatchRuleBatch) {
	if err := l.validatePutBatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putBatch",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakewatchRule) PutMitre(value interface{}) {
	if err := l.validatePutMitreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMitre",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakewatchRule) PutOutput(value *LakewatchRuleOutput) {
	if err := l.validatePutOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putOutput",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakewatchRule) PutProviderConfig(value *LakewatchRuleProviderConfig) {
	if err := l.validatePutProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putProviderConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakewatchRule) PutSchedule(value *LakewatchRuleSchedule) {
	if err := l.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSchedule",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakewatchRule) PutTimerange(value *LakewatchRuleTimerange) {
	if err := l.validatePutTimerangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimerange",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakewatchRule) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := l.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (l *jsiiProxy_LakewatchRule) ResetAnnotations() {
	_jsii_.InvokeVoid(
		l,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRule) ResetBatch() {
	_jsii_.InvokeVoid(
		l,
		"resetBatch",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRule) ResetCategory() {
	_jsii_.InvokeVoid(
		l,
		"resetCategory",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRule) ResetContentResourceId() {
	_jsii_.InvokeVoid(
		l,
		"resetContentResourceId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRule) ResetDeduplication() {
	_jsii_.InvokeVoid(
		l,
		"resetDeduplication",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRule) ResetDeduplicationTimeWindow() {
	_jsii_.InvokeVoid(
		l,
		"resetDeduplicationTimeWindow",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRule) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRule) ResetMitre() {
	_jsii_.InvokeVoid(
		l,
		"resetMitre",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRule) ResetObjective() {
	_jsii_.InvokeVoid(
		l,
		"resetObjective",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRule) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRule) ResetReferences() {
	_jsii_.InvokeVoid(
		l,
		"resetReferences",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRule) ResetRunAsPrincipal() {
	_jsii_.InvokeVoid(
		l,
		"resetRunAsPrincipal",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRule) ResetRunbook() {
	_jsii_.InvokeVoid(
		l,
		"resetRunbook",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRule) ResetSuppressUntilTime() {
	_jsii_.InvokeVoid(
		l,
		"resetSuppressUntilTime",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRule) ResetTimerange() {
	_jsii_.InvokeVoid(
		l,
		"resetTimerange",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRule) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		l,
		"with",
		args,
		&returns,
	)

	return returns
}

