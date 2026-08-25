// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/workspacesettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/workspace_setting_v2 databricks_workspace_setting_v2}.
type WorkspaceSettingV2 interface {
	cdktn.TerraformResource
	AibiDashboardEmbeddingAccessPolicy() WorkspaceSettingV2AibiDashboardEmbeddingAccessPolicyOutputReference
	AibiDashboardEmbeddingAccessPolicyInput() interface{}
	AibiDashboardEmbeddingApprovedDomains() WorkspaceSettingV2AibiDashboardEmbeddingApprovedDomainsOutputReference
	AibiDashboardEmbeddingApprovedDomainsInput() interface{}
	AllowedAppsUserApiScopes() WorkspaceSettingV2AllowedAppsUserApiScopesOutputReference
	AllowedAppsUserApiScopesInput() interface{}
	AutomaticClusterUpdateWorkspace() WorkspaceSettingV2AutomaticClusterUpdateWorkspaceOutputReference
	AutomaticClusterUpdateWorkspaceInput() interface{}
	BooleanVal() WorkspaceSettingV2BooleanValOutputReference
	BooleanValInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CollaborationPlatformConnectivity() WorkspaceSettingV2CollaborationPlatformConnectivityOutputReference
	CollaborationPlatformConnectivityInput() interface{}
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
	EffectiveAibiDashboardEmbeddingAccessPolicy() WorkspaceSettingV2EffectiveAibiDashboardEmbeddingAccessPolicyOutputReference
	EffectiveAibiDashboardEmbeddingAccessPolicyInput() interface{}
	EffectiveAibiDashboardEmbeddingApprovedDomains() WorkspaceSettingV2EffectiveAibiDashboardEmbeddingApprovedDomainsOutputReference
	EffectiveAibiDashboardEmbeddingApprovedDomainsInput() interface{}
	EffectiveAllowedAppsUserApiScopes() WorkspaceSettingV2EffectiveAllowedAppsUserApiScopesOutputReference
	EffectiveAutomaticClusterUpdateWorkspace() WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference
	EffectiveAutomaticClusterUpdateWorkspaceInput() interface{}
	EffectiveBooleanVal() WorkspaceSettingV2EffectiveBooleanValOutputReference
	EffectiveCollaborationPlatformConnectivity() WorkspaceSettingV2EffectiveCollaborationPlatformConnectivityOutputReference
	EffectiveIntegerVal() WorkspaceSettingV2EffectiveIntegerValOutputReference
	EffectiveOperationalEmailCustomRecipient() WorkspaceSettingV2EffectiveOperationalEmailCustomRecipientOutputReference
	EffectivePersonalCompute() WorkspaceSettingV2EffectivePersonalComputeOutputReference
	EffectivePersonalComputeInput() interface{}
	EffectiveRestrictWorkspaceAdmins() WorkspaceSettingV2EffectiveRestrictWorkspaceAdminsOutputReference
	EffectiveRestrictWorkspaceAdminsInput() interface{}
	EffectiveStringVal() WorkspaceSettingV2EffectiveStringValOutputReference
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IntegerVal() WorkspaceSettingV2IntegerValOutputReference
	IntegerValInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OperationalEmailCustomRecipient() WorkspaceSettingV2OperationalEmailCustomRecipientOutputReference
	OperationalEmailCustomRecipientInput() interface{}
	PersonalCompute() WorkspaceSettingV2PersonalComputeOutputReference
	PersonalComputeInput() interface{}
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProviderConfig() WorkspaceSettingV2ProviderConfigOutputReference
	ProviderConfigInput() interface{}
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RestrictWorkspaceAdmins() WorkspaceSettingV2RestrictWorkspaceAdminsOutputReference
	RestrictWorkspaceAdminsInput() interface{}
	StringVal() WorkspaceSettingV2StringValOutputReference
	StringValInput() interface{}
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
	PutAibiDashboardEmbeddingAccessPolicy(value *WorkspaceSettingV2AibiDashboardEmbeddingAccessPolicy)
	PutAibiDashboardEmbeddingApprovedDomains(value *WorkspaceSettingV2AibiDashboardEmbeddingApprovedDomains)
	PutAllowedAppsUserApiScopes(value *WorkspaceSettingV2AllowedAppsUserApiScopes)
	PutAutomaticClusterUpdateWorkspace(value *WorkspaceSettingV2AutomaticClusterUpdateWorkspace)
	PutBooleanVal(value *WorkspaceSettingV2BooleanVal)
	PutCollaborationPlatformConnectivity(value *WorkspaceSettingV2CollaborationPlatformConnectivity)
	PutEffectiveAibiDashboardEmbeddingAccessPolicy(value *WorkspaceSettingV2EffectiveAibiDashboardEmbeddingAccessPolicy)
	PutEffectiveAibiDashboardEmbeddingApprovedDomains(value *WorkspaceSettingV2EffectiveAibiDashboardEmbeddingApprovedDomains)
	PutEffectiveAutomaticClusterUpdateWorkspace(value *WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspace)
	PutEffectivePersonalCompute(value *WorkspaceSettingV2EffectivePersonalCompute)
	PutEffectiveRestrictWorkspaceAdmins(value *WorkspaceSettingV2EffectiveRestrictWorkspaceAdmins)
	PutIntegerVal(value *WorkspaceSettingV2IntegerVal)
	PutOperationalEmailCustomRecipient(value *WorkspaceSettingV2OperationalEmailCustomRecipient)
	PutPersonalCompute(value *WorkspaceSettingV2PersonalCompute)
	PutProviderConfig(value *WorkspaceSettingV2ProviderConfig)
	PutRestrictWorkspaceAdmins(value *WorkspaceSettingV2RestrictWorkspaceAdmins)
	PutStringVal(value *WorkspaceSettingV2StringVal)
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
	ResetAibiDashboardEmbeddingAccessPolicy()
	ResetAibiDashboardEmbeddingApprovedDomains()
	ResetAllowedAppsUserApiScopes()
	ResetAutomaticClusterUpdateWorkspace()
	ResetBooleanVal()
	ResetCollaborationPlatformConnectivity()
	ResetEffectiveAibiDashboardEmbeddingAccessPolicy()
	ResetEffectiveAibiDashboardEmbeddingApprovedDomains()
	ResetEffectiveAutomaticClusterUpdateWorkspace()
	ResetEffectivePersonalCompute()
	ResetEffectiveRestrictWorkspaceAdmins()
	ResetIntegerVal()
	ResetName()
	ResetOperationalEmailCustomRecipient()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPersonalCompute()
	ResetProviderConfig()
	ResetRestrictWorkspaceAdmins()
	ResetStringVal()
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

// The jsii proxy struct for WorkspaceSettingV2
type jsiiProxy_WorkspaceSettingV2 struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_WorkspaceSettingV2) AibiDashboardEmbeddingAccessPolicy() WorkspaceSettingV2AibiDashboardEmbeddingAccessPolicyOutputReference {
	var returns WorkspaceSettingV2AibiDashboardEmbeddingAccessPolicyOutputReference
	_jsii_.Get(
		j,
		"aibiDashboardEmbeddingAccessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) AibiDashboardEmbeddingAccessPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aibiDashboardEmbeddingAccessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) AibiDashboardEmbeddingApprovedDomains() WorkspaceSettingV2AibiDashboardEmbeddingApprovedDomainsOutputReference {
	var returns WorkspaceSettingV2AibiDashboardEmbeddingApprovedDomainsOutputReference
	_jsii_.Get(
		j,
		"aibiDashboardEmbeddingApprovedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) AibiDashboardEmbeddingApprovedDomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aibiDashboardEmbeddingApprovedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) AllowedAppsUserApiScopes() WorkspaceSettingV2AllowedAppsUserApiScopesOutputReference {
	var returns WorkspaceSettingV2AllowedAppsUserApiScopesOutputReference
	_jsii_.Get(
		j,
		"allowedAppsUserApiScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) AllowedAppsUserApiScopesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedAppsUserApiScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) AutomaticClusterUpdateWorkspace() WorkspaceSettingV2AutomaticClusterUpdateWorkspaceOutputReference {
	var returns WorkspaceSettingV2AutomaticClusterUpdateWorkspaceOutputReference
	_jsii_.Get(
		j,
		"automaticClusterUpdateWorkspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) AutomaticClusterUpdateWorkspaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticClusterUpdateWorkspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) BooleanVal() WorkspaceSettingV2BooleanValOutputReference {
	var returns WorkspaceSettingV2BooleanValOutputReference
	_jsii_.Get(
		j,
		"booleanVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) BooleanValInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"booleanValInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) CollaborationPlatformConnectivity() WorkspaceSettingV2CollaborationPlatformConnectivityOutputReference {
	var returns WorkspaceSettingV2CollaborationPlatformConnectivityOutputReference
	_jsii_.Get(
		j,
		"collaborationPlatformConnectivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) CollaborationPlatformConnectivityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collaborationPlatformConnectivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) EffectiveAibiDashboardEmbeddingAccessPolicy() WorkspaceSettingV2EffectiveAibiDashboardEmbeddingAccessPolicyOutputReference {
	var returns WorkspaceSettingV2EffectiveAibiDashboardEmbeddingAccessPolicyOutputReference
	_jsii_.Get(
		j,
		"effectiveAibiDashboardEmbeddingAccessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) EffectiveAibiDashboardEmbeddingAccessPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveAibiDashboardEmbeddingAccessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) EffectiveAibiDashboardEmbeddingApprovedDomains() WorkspaceSettingV2EffectiveAibiDashboardEmbeddingApprovedDomainsOutputReference {
	var returns WorkspaceSettingV2EffectiveAibiDashboardEmbeddingApprovedDomainsOutputReference
	_jsii_.Get(
		j,
		"effectiveAibiDashboardEmbeddingApprovedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) EffectiveAibiDashboardEmbeddingApprovedDomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveAibiDashboardEmbeddingApprovedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) EffectiveAllowedAppsUserApiScopes() WorkspaceSettingV2EffectiveAllowedAppsUserApiScopesOutputReference {
	var returns WorkspaceSettingV2EffectiveAllowedAppsUserApiScopesOutputReference
	_jsii_.Get(
		j,
		"effectiveAllowedAppsUserApiScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) EffectiveAutomaticClusterUpdateWorkspace() WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference {
	var returns WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference
	_jsii_.Get(
		j,
		"effectiveAutomaticClusterUpdateWorkspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) EffectiveAutomaticClusterUpdateWorkspaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveAutomaticClusterUpdateWorkspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) EffectiveBooleanVal() WorkspaceSettingV2EffectiveBooleanValOutputReference {
	var returns WorkspaceSettingV2EffectiveBooleanValOutputReference
	_jsii_.Get(
		j,
		"effectiveBooleanVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) EffectiveCollaborationPlatformConnectivity() WorkspaceSettingV2EffectiveCollaborationPlatformConnectivityOutputReference {
	var returns WorkspaceSettingV2EffectiveCollaborationPlatformConnectivityOutputReference
	_jsii_.Get(
		j,
		"effectiveCollaborationPlatformConnectivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) EffectiveIntegerVal() WorkspaceSettingV2EffectiveIntegerValOutputReference {
	var returns WorkspaceSettingV2EffectiveIntegerValOutputReference
	_jsii_.Get(
		j,
		"effectiveIntegerVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) EffectiveOperationalEmailCustomRecipient() WorkspaceSettingV2EffectiveOperationalEmailCustomRecipientOutputReference {
	var returns WorkspaceSettingV2EffectiveOperationalEmailCustomRecipientOutputReference
	_jsii_.Get(
		j,
		"effectiveOperationalEmailCustomRecipient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) EffectivePersonalCompute() WorkspaceSettingV2EffectivePersonalComputeOutputReference {
	var returns WorkspaceSettingV2EffectivePersonalComputeOutputReference
	_jsii_.Get(
		j,
		"effectivePersonalCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) EffectivePersonalComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectivePersonalComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) EffectiveRestrictWorkspaceAdmins() WorkspaceSettingV2EffectiveRestrictWorkspaceAdminsOutputReference {
	var returns WorkspaceSettingV2EffectiveRestrictWorkspaceAdminsOutputReference
	_jsii_.Get(
		j,
		"effectiveRestrictWorkspaceAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) EffectiveRestrictWorkspaceAdminsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveRestrictWorkspaceAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) EffectiveStringVal() WorkspaceSettingV2EffectiveStringValOutputReference {
	var returns WorkspaceSettingV2EffectiveStringValOutputReference
	_jsii_.Get(
		j,
		"effectiveStringVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) IntegerVal() WorkspaceSettingV2IntegerValOutputReference {
	var returns WorkspaceSettingV2IntegerValOutputReference
	_jsii_.Get(
		j,
		"integerVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) IntegerValInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"integerValInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) OperationalEmailCustomRecipient() WorkspaceSettingV2OperationalEmailCustomRecipientOutputReference {
	var returns WorkspaceSettingV2OperationalEmailCustomRecipientOutputReference
	_jsii_.Get(
		j,
		"operationalEmailCustomRecipient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) OperationalEmailCustomRecipientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationalEmailCustomRecipientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) PersonalCompute() WorkspaceSettingV2PersonalComputeOutputReference {
	var returns WorkspaceSettingV2PersonalComputeOutputReference
	_jsii_.Get(
		j,
		"personalCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) PersonalComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"personalComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) ProviderConfig() WorkspaceSettingV2ProviderConfigOutputReference {
	var returns WorkspaceSettingV2ProviderConfigOutputReference
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) RestrictWorkspaceAdmins() WorkspaceSettingV2RestrictWorkspaceAdminsOutputReference {
	var returns WorkspaceSettingV2RestrictWorkspaceAdminsOutputReference
	_jsii_.Get(
		j,
		"restrictWorkspaceAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) RestrictWorkspaceAdminsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictWorkspaceAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) StringVal() WorkspaceSettingV2StringValOutputReference {
	var returns WorkspaceSettingV2StringValOutputReference
	_jsii_.Get(
		j,
		"stringVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) StringValInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringValInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/workspace_setting_v2 databricks_workspace_setting_v2} Resource.
func NewWorkspaceSettingV2(scope constructs.Construct, id *string, config *WorkspaceSettingV2Config) WorkspaceSettingV2 {
	_init_.Initialize()

	if err := validateNewWorkspaceSettingV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspaceSettingV2{}

	_jsii_.Create(
		"@cdktn/provider-databricks.workspaceSettingV2.WorkspaceSettingV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/workspace_setting_v2 databricks_workspace_setting_v2} Resource.
func NewWorkspaceSettingV2_Override(w WorkspaceSettingV2, scope constructs.Construct, id *string, config *WorkspaceSettingV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.workspaceSettingV2.WorkspaceSettingV2",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a WorkspaceSettingV2 resource upon running "cdktn plan <stack-name>".
func WorkspaceSettingV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateWorkspaceSettingV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.workspaceSettingV2.WorkspaceSettingV2",
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
func WorkspaceSettingV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkspaceSettingV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.workspaceSettingV2.WorkspaceSettingV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkspaceSettingV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkspaceSettingV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.workspaceSettingV2.WorkspaceSettingV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkspaceSettingV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkspaceSettingV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.workspaceSettingV2.WorkspaceSettingV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WorkspaceSettingV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-databricks.workspaceSettingV2.WorkspaceSettingV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := w.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) PutAibiDashboardEmbeddingAccessPolicy(value *WorkspaceSettingV2AibiDashboardEmbeddingAccessPolicy) {
	if err := w.validatePutAibiDashboardEmbeddingAccessPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAibiDashboardEmbeddingAccessPolicy",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) PutAibiDashboardEmbeddingApprovedDomains(value *WorkspaceSettingV2AibiDashboardEmbeddingApprovedDomains) {
	if err := w.validatePutAibiDashboardEmbeddingApprovedDomainsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAibiDashboardEmbeddingApprovedDomains",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) PutAllowedAppsUserApiScopes(value *WorkspaceSettingV2AllowedAppsUserApiScopes) {
	if err := w.validatePutAllowedAppsUserApiScopesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllowedAppsUserApiScopes",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) PutAutomaticClusterUpdateWorkspace(value *WorkspaceSettingV2AutomaticClusterUpdateWorkspace) {
	if err := w.validatePutAutomaticClusterUpdateWorkspaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAutomaticClusterUpdateWorkspace",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) PutBooleanVal(value *WorkspaceSettingV2BooleanVal) {
	if err := w.validatePutBooleanValParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBooleanVal",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) PutCollaborationPlatformConnectivity(value *WorkspaceSettingV2CollaborationPlatformConnectivity) {
	if err := w.validatePutCollaborationPlatformConnectivityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCollaborationPlatformConnectivity",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) PutEffectiveAibiDashboardEmbeddingAccessPolicy(value *WorkspaceSettingV2EffectiveAibiDashboardEmbeddingAccessPolicy) {
	if err := w.validatePutEffectiveAibiDashboardEmbeddingAccessPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putEffectiveAibiDashboardEmbeddingAccessPolicy",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) PutEffectiveAibiDashboardEmbeddingApprovedDomains(value *WorkspaceSettingV2EffectiveAibiDashboardEmbeddingApprovedDomains) {
	if err := w.validatePutEffectiveAibiDashboardEmbeddingApprovedDomainsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putEffectiveAibiDashboardEmbeddingApprovedDomains",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) PutEffectiveAutomaticClusterUpdateWorkspace(value *WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspace) {
	if err := w.validatePutEffectiveAutomaticClusterUpdateWorkspaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putEffectiveAutomaticClusterUpdateWorkspace",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) PutEffectivePersonalCompute(value *WorkspaceSettingV2EffectivePersonalCompute) {
	if err := w.validatePutEffectivePersonalComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putEffectivePersonalCompute",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) PutEffectiveRestrictWorkspaceAdmins(value *WorkspaceSettingV2EffectiveRestrictWorkspaceAdmins) {
	if err := w.validatePutEffectiveRestrictWorkspaceAdminsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putEffectiveRestrictWorkspaceAdmins",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) PutIntegerVal(value *WorkspaceSettingV2IntegerVal) {
	if err := w.validatePutIntegerValParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putIntegerVal",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) PutOperationalEmailCustomRecipient(value *WorkspaceSettingV2OperationalEmailCustomRecipient) {
	if err := w.validatePutOperationalEmailCustomRecipientParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putOperationalEmailCustomRecipient",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) PutPersonalCompute(value *WorkspaceSettingV2PersonalCompute) {
	if err := w.validatePutPersonalComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putPersonalCompute",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) PutProviderConfig(value *WorkspaceSettingV2ProviderConfig) {
	if err := w.validatePutProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putProviderConfig",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) PutRestrictWorkspaceAdmins(value *WorkspaceSettingV2RestrictWorkspaceAdmins) {
	if err := w.validatePutRestrictWorkspaceAdminsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRestrictWorkspaceAdmins",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) PutStringVal(value *WorkspaceSettingV2StringVal) {
	if err := w.validatePutStringValParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putStringVal",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := w.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetAibiDashboardEmbeddingAccessPolicy() {
	_jsii_.InvokeVoid(
		w,
		"resetAibiDashboardEmbeddingAccessPolicy",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetAibiDashboardEmbeddingApprovedDomains() {
	_jsii_.InvokeVoid(
		w,
		"resetAibiDashboardEmbeddingApprovedDomains",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetAllowedAppsUserApiScopes() {
	_jsii_.InvokeVoid(
		w,
		"resetAllowedAppsUserApiScopes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetAutomaticClusterUpdateWorkspace() {
	_jsii_.InvokeVoid(
		w,
		"resetAutomaticClusterUpdateWorkspace",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetBooleanVal() {
	_jsii_.InvokeVoid(
		w,
		"resetBooleanVal",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetCollaborationPlatformConnectivity() {
	_jsii_.InvokeVoid(
		w,
		"resetCollaborationPlatformConnectivity",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetEffectiveAibiDashboardEmbeddingAccessPolicy() {
	_jsii_.InvokeVoid(
		w,
		"resetEffectiveAibiDashboardEmbeddingAccessPolicy",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetEffectiveAibiDashboardEmbeddingApprovedDomains() {
	_jsii_.InvokeVoid(
		w,
		"resetEffectiveAibiDashboardEmbeddingApprovedDomains",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetEffectiveAutomaticClusterUpdateWorkspace() {
	_jsii_.InvokeVoid(
		w,
		"resetEffectiveAutomaticClusterUpdateWorkspace",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetEffectivePersonalCompute() {
	_jsii_.InvokeVoid(
		w,
		"resetEffectivePersonalCompute",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetEffectiveRestrictWorkspaceAdmins() {
	_jsii_.InvokeVoid(
		w,
		"resetEffectiveRestrictWorkspaceAdmins",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetIntegerVal() {
	_jsii_.InvokeVoid(
		w,
		"resetIntegerVal",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetName() {
	_jsii_.InvokeVoid(
		w,
		"resetName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetOperationalEmailCustomRecipient() {
	_jsii_.InvokeVoid(
		w,
		"resetOperationalEmailCustomRecipient",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetPersonalCompute() {
	_jsii_.InvokeVoid(
		w,
		"resetPersonalCompute",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		w,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetRestrictWorkspaceAdmins() {
	_jsii_.InvokeVoid(
		w,
		"resetRestrictWorkspaceAdmins",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) ResetStringVal() {
	_jsii_.InvokeVoid(
		w,
		"resetStringVal",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		w,
		"with",
		args,
		&returns,
	)

	return returns
}

