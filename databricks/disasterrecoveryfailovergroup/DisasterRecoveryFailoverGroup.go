// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package disasterrecoveryfailovergroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/disasterrecoveryfailovergroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/disaster_recovery_failover_group databricks_disaster_recovery_failover_group}.
type DisasterRecoveryFailoverGroup interface {
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EffectivePrimaryRegion() *string
	Etag() *string
	SetEtag(val *string)
	EtagInput() *string
	FailoverGroupId() *string
	SetFailoverGroupId(val *string)
	FailoverGroupIdInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	InitialPrimaryRegion() *string
	SetInitialPrimaryRegion(val *string)
	InitialPrimaryRegionInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	Parent() *string
	SetParent(val *string)
	ParentInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Regions() *[]*string
	SetRegions(val *[]*string)
	RegionsInput() *[]*string
	ReplicationPoint() *string
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UnityCatalogAssets() DisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference
	UnityCatalogAssetsInput() interface{}
	UpdateTime() *string
	WorkspaceSets() DisasterRecoveryFailoverGroupWorkspaceSetsList
	WorkspaceSetsInput() interface{}
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
	PutUnityCatalogAssets(value *DisasterRecoveryFailoverGroupUnityCatalogAssets)
	PutWorkspaceSets(value interface{})
	ResetEtag()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetUnityCatalogAssets()
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

// The jsii proxy struct for DisasterRecoveryFailoverGroup
type jsiiProxy_DisasterRecoveryFailoverGroup struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) EffectivePrimaryRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectivePrimaryRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) EtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) FailoverGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) FailoverGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) InitialPrimaryRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialPrimaryRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) InitialPrimaryRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialPrimaryRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) Parent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) ParentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) ReplicationPoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) UnityCatalogAssets() DisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference {
	var returns DisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference
	_jsii_.Get(
		j,
		"unityCatalogAssets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) UnityCatalogAssetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unityCatalogAssetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) WorkspaceSets() DisasterRecoveryFailoverGroupWorkspaceSetsList {
	var returns DisasterRecoveryFailoverGroupWorkspaceSetsList
	_jsii_.Get(
		j,
		"workspaceSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup) WorkspaceSetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceSetsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/disaster_recovery_failover_group databricks_disaster_recovery_failover_group} Resource.
func NewDisasterRecoveryFailoverGroup(scope constructs.Construct, id *string, config *DisasterRecoveryFailoverGroupConfig) DisasterRecoveryFailoverGroup {
	_init_.Initialize()

	if err := validateNewDisasterRecoveryFailoverGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DisasterRecoveryFailoverGroup{}

	_jsii_.Create(
		"@cdktn/provider-databricks.disasterRecoveryFailoverGroup.DisasterRecoveryFailoverGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/disaster_recovery_failover_group databricks_disaster_recovery_failover_group} Resource.
func NewDisasterRecoveryFailoverGroup_Override(d DisasterRecoveryFailoverGroup, scope constructs.Construct, id *string, config *DisasterRecoveryFailoverGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.disasterRecoveryFailoverGroup.DisasterRecoveryFailoverGroup",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup)SetEtag(val *string) {
	if err := j.validateSetEtagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etag",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup)SetFailoverGroupId(val *string) {
	if err := j.validateSetFailoverGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failoverGroupId",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup)SetInitialPrimaryRegion(val *string) {
	if err := j.validateSetInitialPrimaryRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialPrimaryRegion",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup)SetParent(val *string) {
	if err := j.validateSetParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parent",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroup)SetRegions(val *[]*string) {
	if err := j.validateSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

// Generates CDKTN code for importing a DisasterRecoveryFailoverGroup resource upon running "cdktn plan <stack-name>".
func DisasterRecoveryFailoverGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDisasterRecoveryFailoverGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.disasterRecoveryFailoverGroup.DisasterRecoveryFailoverGroup",
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
func DisasterRecoveryFailoverGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDisasterRecoveryFailoverGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.disasterRecoveryFailoverGroup.DisasterRecoveryFailoverGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DisasterRecoveryFailoverGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDisasterRecoveryFailoverGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.disasterRecoveryFailoverGroup.DisasterRecoveryFailoverGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DisasterRecoveryFailoverGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDisasterRecoveryFailoverGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.disasterRecoveryFailoverGroup.DisasterRecoveryFailoverGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DisasterRecoveryFailoverGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-databricks.disasterRecoveryFailoverGroup.DisasterRecoveryFailoverGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) PutUnityCatalogAssets(value *DisasterRecoveryFailoverGroupUnityCatalogAssets) {
	if err := d.validatePutUnityCatalogAssetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUnityCatalogAssets",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) PutWorkspaceSets(value interface{}) {
	if err := d.validatePutWorkspaceSetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWorkspaceSets",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) ResetEtag() {
	_jsii_.InvokeVoid(
		d,
		"resetEtag",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) ResetUnityCatalogAssets() {
	_jsii_.InvokeVoid(
		d,
		"resetUnityCatalogAssets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

