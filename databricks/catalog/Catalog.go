// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/catalog/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog databricks_catalog}.
type Catalog interface {
	cdktn.TerraformResource
	BrowseOnly() interface{}
	SetBrowseOnly(val interface{})
	BrowseOnlyInput() interface{}
	CatalogType() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionName() *string
	SetConnectionName(val *string)
	ConnectionNameInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *float64
	CreatedBy() *string
	CustomMaxRetentionHours() *float64
	SetCustomMaxRetentionHours(val *float64)
	CustomMaxRetentionHoursInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EffectivePredictiveOptimizationFlag() CatalogEffectivePredictiveOptimizationFlagOutputReference
	EffectivePredictiveOptimizationFlagInput() *CatalogEffectivePredictiveOptimizationFlag
	EnablePredictiveOptimization() *string
	SetEnablePredictiveOptimization(val *string)
	EnablePredictiveOptimizationInput() *string
	ForceDestroy() interface{}
	SetForceDestroy(val interface{})
	ForceDestroyInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FullName() *string
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
	ManagedEncryptionSettings() CatalogManagedEncryptionSettingsOutputReference
	ManagedEncryptionSettingsInput() *CatalogManagedEncryptionSettings
	MetastoreId() *string
	SetMetastoreId(val *string)
	MetastoreIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Options() *map[string]*string
	SetOptions(val *map[string]*string)
	OptionsInput() *map[string]*string
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	Properties() *map[string]*string
	SetProperties(val *map[string]*string)
	PropertiesInput() *map[string]*string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProviderConfig() CatalogProviderConfigOutputReference
	ProviderConfigInput() *CatalogProviderConfig
	ProviderName() *string
	SetProviderName(val *string)
	ProviderNameInput() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProvisioningInfo() CatalogProvisioningInfoOutputReference
	ProvisioningInfoInput() *CatalogProvisioningInfo
	// Experimental.
	RawOverrides() interface{}
	SecurableType() *string
	ShareName() *string
	SetShareName(val *string)
	ShareNameInput() *string
	StorageLocation() *string
	StorageRoot() *string
	SetStorageRoot(val *string)
	StorageRootInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatedAt() *float64
	UpdatedBy() *string
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
	PutEffectivePredictiveOptimizationFlag(value *CatalogEffectivePredictiveOptimizationFlag)
	PutManagedEncryptionSettings(value *CatalogManagedEncryptionSettings)
	PutProviderConfig(value *CatalogProviderConfig)
	PutProvisioningInfo(value *CatalogProvisioningInfo)
	ResetBrowseOnly()
	ResetComment()
	ResetConnectionName()
	ResetCustomMaxRetentionHours()
	ResetEffectivePredictiveOptimizationFlag()
	ResetEnablePredictiveOptimization()
	ResetForceDestroy()
	ResetId()
	ResetIsolationMode()
	ResetManagedEncryptionSettings()
	ResetMetastoreId()
	ResetName()
	ResetOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwner()
	ResetProperties()
	ResetProviderConfig()
	ResetProviderName()
	ResetProvisioningInfo()
	ResetShareName()
	ResetStorageRoot()
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

// The jsii proxy struct for Catalog
type jsiiProxy_Catalog struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_Catalog) BrowseOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browseOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) BrowseOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browseOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) CatalogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) ConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) ConnectionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) CreatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) CustomMaxRetentionHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customMaxRetentionHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) CustomMaxRetentionHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customMaxRetentionHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) EffectivePredictiveOptimizationFlag() CatalogEffectivePredictiveOptimizationFlagOutputReference {
	var returns CatalogEffectivePredictiveOptimizationFlagOutputReference
	_jsii_.Get(
		j,
		"effectivePredictiveOptimizationFlag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) EffectivePredictiveOptimizationFlagInput() *CatalogEffectivePredictiveOptimizationFlag {
	var returns *CatalogEffectivePredictiveOptimizationFlag
	_jsii_.Get(
		j,
		"effectivePredictiveOptimizationFlagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) EnablePredictiveOptimization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enablePredictiveOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) EnablePredictiveOptimizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enablePredictiveOptimizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) FullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) IsolationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isolationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) IsolationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isolationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) ManagedEncryptionSettings() CatalogManagedEncryptionSettingsOutputReference {
	var returns CatalogManagedEncryptionSettingsOutputReference
	_jsii_.Get(
		j,
		"managedEncryptionSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) ManagedEncryptionSettingsInput() *CatalogManagedEncryptionSettings {
	var returns *CatalogManagedEncryptionSettings
	_jsii_.Get(
		j,
		"managedEncryptionSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) MetastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) MetastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Options() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) OptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) ProviderConfig() CatalogProviderConfigOutputReference {
	var returns CatalogProviderConfigOutputReference
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) ProviderConfigInput() *CatalogProviderConfig {
	var returns *CatalogProviderConfig
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) ProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) ProvisioningInfo() CatalogProvisioningInfoOutputReference {
	var returns CatalogProvisioningInfoOutputReference
	_jsii_.Get(
		j,
		"provisioningInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) ProvisioningInfoInput() *CatalogProvisioningInfo {
	var returns *CatalogProvisioningInfo
	_jsii_.Get(
		j,
		"provisioningInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) SecurableType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securableType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) ShareName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) ShareNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) StorageLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) StorageRoot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) StorageRootInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) UpdatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Catalog) UpdatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedBy",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog databricks_catalog} Resource.
func NewCatalog(scope constructs.Construct, id *string, config *CatalogConfig) Catalog {
	_init_.Initialize()

	if err := validateNewCatalogParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Catalog{}

	_jsii_.Create(
		"@cdktn/provider-databricks.catalog.Catalog",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog databricks_catalog} Resource.
func NewCatalog_Override(c Catalog, scope constructs.Construct, id *string, config *CatalogConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.catalog.Catalog",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_Catalog)SetBrowseOnly(val interface{}) {
	if err := j.validateSetBrowseOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"browseOnly",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetConnectionName(val *string) {
	if err := j.validateSetConnectionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionName",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetCustomMaxRetentionHours(val *float64) {
	if err := j.validateSetCustomMaxRetentionHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customMaxRetentionHours",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetEnablePredictiveOptimization(val *string) {
	if err := j.validateSetEnablePredictiveOptimizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePredictiveOptimization",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetIsolationMode(val *string) {
	if err := j.validateSetIsolationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isolationMode",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetMetastoreId(val *string) {
	if err := j.validateSetMetastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metastoreId",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetOptions(val *map[string]*string) {
	if err := j.validateSetOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetProviderName(val *string) {
	if err := j.validateSetProviderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerName",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetShareName(val *string) {
	if err := j.validateSetShareNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareName",
		val,
	)
}

func (j *jsiiProxy_Catalog)SetStorageRoot(val *string) {
	if err := j.validateSetStorageRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageRoot",
		val,
	)
}

// Generates CDKTN code for importing a Catalog resource upon running "cdktn plan <stack-name>".
func Catalog_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateCatalog_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.catalog.Catalog",
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
func Catalog_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCatalog_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.catalog.Catalog",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Catalog_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCatalog_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.catalog.Catalog",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Catalog_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCatalog_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.catalog.Catalog",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Catalog_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-databricks.catalog.Catalog",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_Catalog) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_Catalog) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_Catalog) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_Catalog) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_Catalog) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_Catalog) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_Catalog) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_Catalog) PutEffectivePredictiveOptimizationFlag(value *CatalogEffectivePredictiveOptimizationFlag) {
	if err := c.validatePutEffectivePredictiveOptimizationFlagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEffectivePredictiveOptimizationFlag",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Catalog) PutManagedEncryptionSettings(value *CatalogManagedEncryptionSettings) {
	if err := c.validatePutManagedEncryptionSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putManagedEncryptionSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Catalog) PutProviderConfig(value *CatalogProviderConfig) {
	if err := c.validatePutProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putProviderConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Catalog) PutProvisioningInfo(value *CatalogProvisioningInfo) {
	if err := c.validatePutProvisioningInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putProvisioningInfo",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Catalog) ResetBrowseOnly() {
	_jsii_.InvokeVoid(
		c,
		"resetBrowseOnly",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetComment() {
	_jsii_.InvokeVoid(
		c,
		"resetComment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetConnectionName() {
	_jsii_.InvokeVoid(
		c,
		"resetConnectionName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetCustomMaxRetentionHours() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomMaxRetentionHours",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetEffectivePredictiveOptimizationFlag() {
	_jsii_.InvokeVoid(
		c,
		"resetEffectivePredictiveOptimizationFlag",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetEnablePredictiveOptimization() {
	_jsii_.InvokeVoid(
		c,
		"resetEnablePredictiveOptimization",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		c,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetIsolationMode() {
	_jsii_.InvokeVoid(
		c,
		"resetIsolationMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetManagedEncryptionSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetManagedEncryptionSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetMetastoreId() {
	_jsii_.InvokeVoid(
		c,
		"resetMetastoreId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetOptions() {
	_jsii_.InvokeVoid(
		c,
		"resetOptions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetOwner() {
	_jsii_.InvokeVoid(
		c,
		"resetOwner",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetProperties() {
	_jsii_.InvokeVoid(
		c,
		"resetProperties",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetProviderName() {
	_jsii_.InvokeVoid(
		c,
		"resetProviderName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetProvisioningInfo() {
	_jsii_.InvokeVoid(
		c,
		"resetProvisioningInfo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetShareName() {
	_jsii_.InvokeVoid(
		c,
		"resetShareName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) ResetStorageRoot() {
	_jsii_.InvokeVoid(
		c,
		"resetStorageRoot",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Catalog) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Catalog) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

