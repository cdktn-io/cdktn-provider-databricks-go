// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package externallocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/externallocation/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/external_location databricks_external_location}.
type ExternalLocation interface {
	cdktn.TerraformResource
	BrowseOnly() cdktn.IResolvable
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
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
	CreatedAt() *float64
	CreatedBy() *string
	CredentialId() *string
	CredentialName() *string
	SetCredentialName(val *string)
	CredentialNameInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EffectiveEnableFileEvents() cdktn.IResolvable
	EffectiveFileEventQueue() ExternalLocationEffectiveFileEventQueueOutputReference
	EffectiveFileEventQueueInput() *ExternalLocationEffectiveFileEventQueue
	EnableFileEvents() interface{}
	SetEnableFileEvents(val interface{})
	EnableFileEventsInput() interface{}
	EncryptionDetails() ExternalLocationEncryptionDetailsOutputReference
	EncryptionDetailsInput() *ExternalLocationEncryptionDetails
	Fallback() interface{}
	SetFallback(val interface{})
	FallbackInput() interface{}
	FileEventQueue() ExternalLocationFileEventQueueOutputReference
	FileEventQueueInput() *ExternalLocationFileEventQueue
	ForceDestroy() interface{}
	SetForceDestroy(val interface{})
	ForceDestroyInput() interface{}
	ForceUpdate() interface{}
	SetForceUpdate(val interface{})
	ForceUpdateInput() interface{}
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
	IsolationMode() *string
	SetIsolationMode(val *string)
	IsolationModeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MetastoreId() *string
	SetMetastoreId(val *string)
	MetastoreIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProviderConfig() ExternalLocationProviderConfigOutputReference
	ProviderConfigInput() *ExternalLocationProviderConfig
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	ReadOnlyInput() interface{}
	SkipValidation() interface{}
	SetSkipValidation(val interface{})
	SkipValidationInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatedAt() *float64
	UpdatedBy() *string
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	PutEffectiveFileEventQueue(value *ExternalLocationEffectiveFileEventQueue)
	PutEncryptionDetails(value *ExternalLocationEncryptionDetails)
	PutFileEventQueue(value *ExternalLocationFileEventQueue)
	PutProviderConfig(value *ExternalLocationProviderConfig)
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
	ResetComment()
	ResetEffectiveFileEventQueue()
	ResetEnableFileEvents()
	ResetEncryptionDetails()
	ResetFallback()
	ResetFileEventQueue()
	ResetForceDestroy()
	ResetForceUpdate()
	ResetId()
	ResetIsolationMode()
	ResetMetastoreId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwner()
	ResetProviderConfig()
	ResetReadOnly()
	ResetSkipValidation()
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

// The jsii proxy struct for ExternalLocation
type jsiiProxy_ExternalLocation struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ExternalLocation) BrowseOnly() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"browseOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) CreatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) CredentialId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) CredentialName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) CredentialNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) EffectiveEnableFileEvents() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"effectiveEnableFileEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) EffectiveFileEventQueue() ExternalLocationEffectiveFileEventQueueOutputReference {
	var returns ExternalLocationEffectiveFileEventQueueOutputReference
	_jsii_.Get(
		j,
		"effectiveFileEventQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) EffectiveFileEventQueueInput() *ExternalLocationEffectiveFileEventQueue {
	var returns *ExternalLocationEffectiveFileEventQueue
	_jsii_.Get(
		j,
		"effectiveFileEventQueueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) EnableFileEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFileEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) EnableFileEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFileEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) EncryptionDetails() ExternalLocationEncryptionDetailsOutputReference {
	var returns ExternalLocationEncryptionDetailsOutputReference
	_jsii_.Get(
		j,
		"encryptionDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) EncryptionDetailsInput() *ExternalLocationEncryptionDetails {
	var returns *ExternalLocationEncryptionDetails
	_jsii_.Get(
		j,
		"encryptionDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) Fallback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) FallbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) FileEventQueue() ExternalLocationFileEventQueueOutputReference {
	var returns ExternalLocationFileEventQueueOutputReference
	_jsii_.Get(
		j,
		"fileEventQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) FileEventQueueInput() *ExternalLocationFileEventQueue {
	var returns *ExternalLocationFileEventQueue
	_jsii_.Get(
		j,
		"fileEventQueueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) ForceUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) ForceUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) IsolationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isolationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) IsolationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isolationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) MetastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) MetastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) ProviderConfig() ExternalLocationProviderConfigOutputReference {
	var returns ExternalLocationProviderConfigOutputReference
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) ProviderConfigInput() *ExternalLocationProviderConfig {
	var returns *ExternalLocationProviderConfig
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) SkipValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) SkipValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) UpdatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) UpdatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocation) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/external_location databricks_external_location} Resource.
func NewExternalLocation(scope constructs.Construct, id *string, config *ExternalLocationConfig) ExternalLocation {
	_init_.Initialize()

	if err := validateNewExternalLocationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExternalLocation{}

	_jsii_.Create(
		"@cdktn/provider-databricks.externalLocation.ExternalLocation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/external_location databricks_external_location} Resource.
func NewExternalLocation_Override(e ExternalLocation, scope constructs.Construct, id *string, config *ExternalLocationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.externalLocation.ExternalLocation",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ExternalLocation)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetCredentialName(val *string) {
	if err := j.validateSetCredentialNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialName",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetEnableFileEvents(val interface{}) {
	if err := j.validateSetEnableFileEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableFileEvents",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetFallback(val interface{}) {
	if err := j.validateSetFallbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fallback",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetForceUpdate(val interface{}) {
	if err := j.validateSetForceUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceUpdate",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetIsolationMode(val *string) {
	if err := j.validateSetIsolationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isolationMode",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetMetastoreId(val *string) {
	if err := j.validateSetMetastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metastoreId",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetSkipValidation(val interface{}) {
	if err := j.validateSetSkipValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipValidation",
		val,
	)
}

func (j *jsiiProxy_ExternalLocation)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

// Generates CDKTN code for importing a ExternalLocation resource upon running "cdktn plan <stack-name>".
func ExternalLocation_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateExternalLocation_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.externalLocation.ExternalLocation",
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
func ExternalLocation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalLocation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.externalLocation.ExternalLocation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ExternalLocation_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalLocation_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.externalLocation.ExternalLocation",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ExternalLocation_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalLocation_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.externalLocation.ExternalLocation",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ExternalLocation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-databricks.externalLocation.ExternalLocation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ExternalLocation) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_ExternalLocation) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ExternalLocation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocation) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocation) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocation) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocation) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocation) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocation) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_ExternalLocation) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocation) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := e.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocation) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ExternalLocation) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_ExternalLocation) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ExternalLocation) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ExternalLocation) PutEffectiveFileEventQueue(value *ExternalLocationEffectiveFileEventQueue) {
	if err := e.validatePutEffectiveFileEventQueueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEffectiveFileEventQueue",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalLocation) PutEncryptionDetails(value *ExternalLocationEncryptionDetails) {
	if err := e.validatePutEncryptionDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEncryptionDetails",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalLocation) PutFileEventQueue(value *ExternalLocationFileEventQueue) {
	if err := e.validatePutFileEventQueueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putFileEventQueue",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalLocation) PutProviderConfig(value *ExternalLocationProviderConfig) {
	if err := e.validatePutProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putProviderConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalLocation) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_ExternalLocation) ResetComment() {
	_jsii_.InvokeVoid(
		e,
		"resetComment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocation) ResetEffectiveFileEventQueue() {
	_jsii_.InvokeVoid(
		e,
		"resetEffectiveFileEventQueue",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocation) ResetEnableFileEvents() {
	_jsii_.InvokeVoid(
		e,
		"resetEnableFileEvents",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocation) ResetEncryptionDetails() {
	_jsii_.InvokeVoid(
		e,
		"resetEncryptionDetails",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocation) ResetFallback() {
	_jsii_.InvokeVoid(
		e,
		"resetFallback",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocation) ResetFileEventQueue() {
	_jsii_.InvokeVoid(
		e,
		"resetFileEventQueue",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocation) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		e,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocation) ResetForceUpdate() {
	_jsii_.InvokeVoid(
		e,
		"resetForceUpdate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocation) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocation) ResetIsolationMode() {
	_jsii_.InvokeVoid(
		e,
		"resetIsolationMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocation) ResetMetastoreId() {
	_jsii_.InvokeVoid(
		e,
		"resetMetastoreId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocation) ResetOwner() {
	_jsii_.InvokeVoid(
		e,
		"resetOwner",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocation) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocation) ResetReadOnly() {
	_jsii_.InvokeVoid(
		e,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocation) ResetSkipValidation() {
	_jsii_.InvokeVoid(
		e,
		"resetSkipValidation",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocation) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocation) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocation) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		e,
		"with",
		args,
		&returns,
	)

	return returns
}

