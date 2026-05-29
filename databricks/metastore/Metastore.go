// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package metastore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/metastore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/metastore databricks_metastore}.
type Metastore interface {
	cdktn.TerraformResource
	Api() *string
	SetApi(val *string)
	ApiInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Cloud() *string
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
	DefaultDataAccessConfigId() *string
	SetDefaultDataAccessConfigId(val *string)
	DefaultDataAccessConfigIdInput() *string
	DeltaSharingOrganizationName() *string
	SetDeltaSharingOrganizationName(val *string)
	DeltaSharingOrganizationNameInput() *string
	DeltaSharingRecipientTokenLifetimeInSeconds() *float64
	SetDeltaSharingRecipientTokenLifetimeInSeconds(val *float64)
	DeltaSharingRecipientTokenLifetimeInSecondsInput() *float64
	DeltaSharingScope() *string
	SetDeltaSharingScope(val *string)
	DeltaSharingScopeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExternalAccessEnabled() interface{}
	SetExternalAccessEnabled(val interface{})
	ExternalAccessEnabledInput() interface{}
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
	GlobalMetastoreId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MetastoreId() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	PrivilegeModelVersion() *string
	SetPrivilegeModelVersion(val *string)
	PrivilegeModelVersionInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProviderConfig() MetastoreProviderConfigOutputReference
	ProviderConfigInput() *MetastoreProviderConfig
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	StorageRoot() *string
	SetStorageRoot(val *string)
	StorageRootCredentialId() *string
	SetStorageRootCredentialId(val *string)
	StorageRootCredentialIdInput() *string
	StorageRootCredentialName() *string
	SetStorageRootCredentialName(val *string)
	StorageRootCredentialNameInput() *string
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
	PutProviderConfig(value *MetastoreProviderConfig)
	ResetApi()
	ResetDefaultDataAccessConfigId()
	ResetDeltaSharingOrganizationName()
	ResetDeltaSharingRecipientTokenLifetimeInSeconds()
	ResetDeltaSharingScope()
	ResetExternalAccessEnabled()
	ResetForceDestroy()
	ResetId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwner()
	ResetPrivilegeModelVersion()
	ResetProviderConfig()
	ResetRegion()
	ResetStorageRoot()
	ResetStorageRootCredentialId()
	ResetStorageRootCredentialName()
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

// The jsii proxy struct for Metastore
type jsiiProxy_Metastore struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_Metastore) Api() *string {
	var returns *string
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) ApiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) Cloud() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloud",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) CreatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) DefaultDataAccessConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDataAccessConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) DefaultDataAccessConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDataAccessConfigIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) DeltaSharingOrganizationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaSharingOrganizationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) DeltaSharingOrganizationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaSharingOrganizationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) DeltaSharingRecipientTokenLifetimeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deltaSharingRecipientTokenLifetimeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) DeltaSharingRecipientTokenLifetimeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deltaSharingRecipientTokenLifetimeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) DeltaSharingScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaSharingScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) DeltaSharingScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaSharingScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) ExternalAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) ExternalAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) GlobalMetastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalMetastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) MetastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) PrivilegeModelVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegeModelVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) PrivilegeModelVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegeModelVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) ProviderConfig() MetastoreProviderConfigOutputReference {
	var returns MetastoreProviderConfigOutputReference
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) ProviderConfigInput() *MetastoreProviderConfig {
	var returns *MetastoreProviderConfig
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) StorageRoot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) StorageRootCredentialId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRootCredentialId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) StorageRootCredentialIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRootCredentialIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) StorageRootCredentialName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRootCredentialName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) StorageRootCredentialNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRootCredentialNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) StorageRootInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) UpdatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Metastore) UpdatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedBy",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/metastore databricks_metastore} Resource.
func NewMetastore(scope constructs.Construct, id *string, config *MetastoreConfig) Metastore {
	_init_.Initialize()

	if err := validateNewMetastoreParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Metastore{}

	_jsii_.Create(
		"@cdktn/provider-databricks.metastore.Metastore",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/metastore databricks_metastore} Resource.
func NewMetastore_Override(m Metastore, scope constructs.Construct, id *string, config *MetastoreConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.metastore.Metastore",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_Metastore)SetApi(val *string) {
	if err := j.validateSetApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetDefaultDataAccessConfigId(val *string) {
	if err := j.validateSetDefaultDataAccessConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultDataAccessConfigId",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetDeltaSharingOrganizationName(val *string) {
	if err := j.validateSetDeltaSharingOrganizationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deltaSharingOrganizationName",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetDeltaSharingRecipientTokenLifetimeInSeconds(val *float64) {
	if err := j.validateSetDeltaSharingRecipientTokenLifetimeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deltaSharingRecipientTokenLifetimeInSeconds",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetDeltaSharingScope(val *string) {
	if err := j.validateSetDeltaSharingScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deltaSharingScope",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetExternalAccessEnabled(val interface{}) {
	if err := j.validateSetExternalAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetPrivilegeModelVersion(val *string) {
	if err := j.validateSetPrivilegeModelVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilegeModelVersion",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetStorageRoot(val *string) {
	if err := j.validateSetStorageRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageRoot",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetStorageRootCredentialId(val *string) {
	if err := j.validateSetStorageRootCredentialIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageRootCredentialId",
		val,
	)
}

func (j *jsiiProxy_Metastore)SetStorageRootCredentialName(val *string) {
	if err := j.validateSetStorageRootCredentialNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageRootCredentialName",
		val,
	)
}

// Generates CDKTN code for importing a Metastore resource upon running "cdktn plan <stack-name>".
func Metastore_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateMetastore_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.metastore.Metastore",
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
func Metastore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMetastore_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.metastore.Metastore",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Metastore_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMetastore_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.metastore.Metastore",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Metastore_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMetastore_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.metastore.Metastore",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Metastore_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-databricks.metastore.Metastore",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_Metastore) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_Metastore) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_Metastore) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metastore) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metastore) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metastore) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metastore) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metastore) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metastore) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metastore) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metastore) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metastore) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metastore) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_Metastore) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metastore) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_Metastore) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_Metastore) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_Metastore) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_Metastore) PutProviderConfig(value *MetastoreProviderConfig) {
	if err := m.validatePutProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putProviderConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Metastore) ResetApi() {
	_jsii_.InvokeVoid(
		m,
		"resetApi",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metastore) ResetDefaultDataAccessConfigId() {
	_jsii_.InvokeVoid(
		m,
		"resetDefaultDataAccessConfigId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metastore) ResetDeltaSharingOrganizationName() {
	_jsii_.InvokeVoid(
		m,
		"resetDeltaSharingOrganizationName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metastore) ResetDeltaSharingRecipientTokenLifetimeInSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetDeltaSharingRecipientTokenLifetimeInSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metastore) ResetDeltaSharingScope() {
	_jsii_.InvokeVoid(
		m,
		"resetDeltaSharingScope",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metastore) ResetExternalAccessEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetExternalAccessEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metastore) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		m,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metastore) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metastore) ResetName() {
	_jsii_.InvokeVoid(
		m,
		"resetName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metastore) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metastore) ResetOwner() {
	_jsii_.InvokeVoid(
		m,
		"resetOwner",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metastore) ResetPrivilegeModelVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetPrivilegeModelVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metastore) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metastore) ResetRegion() {
	_jsii_.InvokeVoid(
		m,
		"resetRegion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metastore) ResetStorageRoot() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageRoot",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metastore) ResetStorageRootCredentialId() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageRootCredentialId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metastore) ResetStorageRootCredentialName() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageRootCredentialName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Metastore) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metastore) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metastore) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metastore) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metastore) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metastore) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Metastore) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		m,
		"with",
		args,
		&returns,
	)

	return returns
}

