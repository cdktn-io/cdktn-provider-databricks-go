// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/cluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/cluster databricks_cluster}.
type Cluster interface {
	cdktn.TerraformResource
	ApplyPolicyDefaultValues() interface{}
	SetApplyPolicyDefaultValues(val interface{})
	ApplyPolicyDefaultValuesInput() interface{}
	Autoscale() ClusterAutoscaleOutputReference
	AutoscaleInput() *ClusterAutoscale
	AutoterminationMinutes() *float64
	SetAutoterminationMinutes(val *float64)
	AutoterminationMinutesInput() *float64
	AwsAttributes() ClusterAwsAttributesOutputReference
	AwsAttributesInput() *ClusterAwsAttributes
	AzureAttributes() ClusterAzureAttributesOutputReference
	AzureAttributesInput() *ClusterAzureAttributes
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClearCloudAttributesOnRemove() interface{}
	SetClearCloudAttributesOnRemove(val interface{})
	ClearCloudAttributesOnRemoveInput() interface{}
	ClusterId() *string
	ClusterLogConf() ClusterClusterLogConfOutputReference
	ClusterLogConfInput() *ClusterClusterLogConf
	ClusterMountInfo() ClusterClusterMountInfoList
	ClusterMountInfoInput() interface{}
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
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
	CustomTags() *map[string]*string
	SetCustomTags(val *map[string]*string)
	CustomTagsInput() *map[string]*string
	DataSecurityMode() *string
	SetDataSecurityMode(val *string)
	DataSecurityModeInput() *string
	DefaultTags() cdktn.StringMap
	DependencyMode() *string
	SetDependencyMode(val *string)
	DependencyModeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DockerImage() ClusterDockerImageOutputReference
	DockerImageInput() *ClusterDockerImage
	DriverInstancePoolId() *string
	SetDriverInstancePoolId(val *string)
	DriverInstancePoolIdInput() *string
	DriverNodeTypeFlexibility() ClusterDriverNodeTypeFlexibilityOutputReference
	DriverNodeTypeFlexibilityInput() *ClusterDriverNodeTypeFlexibility
	DriverNodeTypeId() *string
	SetDriverNodeTypeId(val *string)
	DriverNodeTypeIdInput() *string
	EnableElasticDisk() interface{}
	SetEnableElasticDisk(val interface{})
	EnableElasticDiskInput() interface{}
	EnableLocalDiskEncryption() interface{}
	SetEnableLocalDiskEncryption(val interface{})
	EnableLocalDiskEncryptionInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcpAttributes() ClusterGcpAttributesOutputReference
	GcpAttributesInput() *ClusterGcpAttributes
	Id() *string
	SetId(val *string)
	IdempotencyToken() *string
	SetIdempotencyToken(val *string)
	IdempotencyTokenInput() *string
	IdInput() *string
	InitScripts() ClusterInitScriptsList
	InitScriptsInput() interface{}
	InstancePoolId() *string
	SetInstancePoolId(val *string)
	InstancePoolIdInput() *string
	IsPinned() interface{}
	SetIsPinned(val interface{})
	IsPinnedInput() interface{}
	IsSingleNode() interface{}
	SetIsSingleNode(val interface{})
	IsSingleNodeInput() interface{}
	Kind() *string
	SetKind(val *string)
	KindInput() *string
	Library() ClusterLibraryList
	LibraryInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	NodeTypeId() *string
	SetNodeTypeId(val *string)
	NodeTypeIdInput() *string
	NoWait() interface{}
	SetNoWait(val interface{})
	NoWaitInput() interface{}
	NumWorkers() *float64
	SetNumWorkers(val *float64)
	NumWorkersInput() *float64
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyIdInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProviderConfig() ClusterProviderConfigOutputReference
	ProviderConfigInput() *ClusterProviderConfig
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RemoteDiskThroughput() *float64
	SetRemoteDiskThroughput(val *float64)
	RemoteDiskThroughputInput() *float64
	RuntimeEngine() *string
	SetRuntimeEngine(val *string)
	RuntimeEngineInput() *string
	SingleUserName() *string
	SetSingleUserName(val *string)
	SingleUserNameInput() *string
	SparkConf() *map[string]*string
	SetSparkConf(val *map[string]*string)
	SparkConfInput() *map[string]*string
	SparkEnvVars() *map[string]*string
	SetSparkEnvVars(val *map[string]*string)
	SparkEnvVarsInput() *map[string]*string
	SparkVersion() *string
	SetSparkVersion(val *string)
	SparkVersionInput() *string
	SshPublicKeys() *[]*string
	SetSshPublicKeys(val *[]*string)
	SshPublicKeysInput() *[]*string
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	TotalInitialRemoteDiskSize() *float64
	SetTotalInitialRemoteDiskSize(val *float64)
	TotalInitialRemoteDiskSizeInput() *float64
	Url() *string
	UseMlRuntime() interface{}
	SetUseMlRuntime(val interface{})
	UseMlRuntimeInput() interface{}
	WorkerNodeTypeFlexibility() ClusterWorkerNodeTypeFlexibilityOutputReference
	WorkerNodeTypeFlexibilityInput() *ClusterWorkerNodeTypeFlexibility
	WorkloadType() ClusterWorkloadTypeOutputReference
	WorkloadTypeInput() *ClusterWorkloadType
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
	PutAutoscale(value *ClusterAutoscale)
	PutAwsAttributes(value *ClusterAwsAttributes)
	PutAzureAttributes(value *ClusterAzureAttributes)
	PutClusterLogConf(value *ClusterClusterLogConf)
	PutClusterMountInfo(value interface{})
	PutDockerImage(value *ClusterDockerImage)
	PutDriverNodeTypeFlexibility(value *ClusterDriverNodeTypeFlexibility)
	PutGcpAttributes(value *ClusterGcpAttributes)
	PutInitScripts(value interface{})
	PutLibrary(value interface{})
	PutProviderConfig(value *ClusterProviderConfig)
	PutTimeouts(value *ClusterTimeouts)
	PutWorkerNodeTypeFlexibility(value *ClusterWorkerNodeTypeFlexibility)
	PutWorkloadType(value *ClusterWorkloadType)
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
	ResetApplyPolicyDefaultValues()
	ResetAutoscale()
	ResetAutoterminationMinutes()
	ResetAwsAttributes()
	ResetAzureAttributes()
	ResetClearCloudAttributesOnRemove()
	ResetClusterLogConf()
	ResetClusterMountInfo()
	ResetClusterName()
	ResetCustomTags()
	ResetDataSecurityMode()
	ResetDependencyMode()
	ResetDockerImage()
	ResetDriverInstancePoolId()
	ResetDriverNodeTypeFlexibility()
	ResetDriverNodeTypeId()
	ResetEnableElasticDisk()
	ResetEnableLocalDiskEncryption()
	ResetGcpAttributes()
	ResetId()
	ResetIdempotencyToken()
	ResetInitScripts()
	ResetInstancePoolId()
	ResetIsPinned()
	ResetIsSingleNode()
	ResetKind()
	ResetLibrary()
	ResetNodeTypeId()
	ResetNoWait()
	ResetNumWorkers()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicyId()
	ResetProviderConfig()
	ResetRemoteDiskThroughput()
	ResetRuntimeEngine()
	ResetSingleUserName()
	ResetSparkConf()
	ResetSparkEnvVars()
	ResetSshPublicKeys()
	ResetTimeouts()
	ResetTotalInitialRemoteDiskSize()
	ResetUseMlRuntime()
	ResetWorkerNodeTypeFlexibility()
	ResetWorkloadType()
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

// The jsii proxy struct for Cluster
type jsiiProxy_Cluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_Cluster) ApplyPolicyDefaultValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyPolicyDefaultValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ApplyPolicyDefaultValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyPolicyDefaultValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Autoscale() ClusterAutoscaleOutputReference {
	var returns ClusterAutoscaleOutputReference
	_jsii_.Get(
		j,
		"autoscale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AutoscaleInput() *ClusterAutoscale {
	var returns *ClusterAutoscale
	_jsii_.Get(
		j,
		"autoscaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AutoterminationMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoterminationMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AutoterminationMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoterminationMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AwsAttributes() ClusterAwsAttributesOutputReference {
	var returns ClusterAwsAttributesOutputReference
	_jsii_.Get(
		j,
		"awsAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AwsAttributesInput() *ClusterAwsAttributes {
	var returns *ClusterAwsAttributes
	_jsii_.Get(
		j,
		"awsAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AzureAttributes() ClusterAzureAttributesOutputReference {
	var returns ClusterAzureAttributesOutputReference
	_jsii_.Get(
		j,
		"azureAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AzureAttributesInput() *ClusterAzureAttributes {
	var returns *ClusterAzureAttributes
	_jsii_.Get(
		j,
		"azureAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClearCloudAttributesOnRemove() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clearCloudAttributesOnRemove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClearCloudAttributesOnRemoveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clearCloudAttributesOnRemoveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterLogConf() ClusterClusterLogConfOutputReference {
	var returns ClusterClusterLogConfOutputReference
	_jsii_.Get(
		j,
		"clusterLogConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterLogConfInput() *ClusterClusterLogConf {
	var returns *ClusterClusterLogConf
	_jsii_.Get(
		j,
		"clusterLogConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterMountInfo() ClusterClusterMountInfoList {
	var returns ClusterClusterMountInfoList
	_jsii_.Get(
		j,
		"clusterMountInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterMountInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterMountInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) CustomTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) CustomTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DataSecurityMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSecurityMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DataSecurityModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSecurityModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DefaultTags() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"defaultTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DependencyMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dependencyMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DependencyModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dependencyModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DockerImage() ClusterDockerImageOutputReference {
	var returns ClusterDockerImageOutputReference
	_jsii_.Get(
		j,
		"dockerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DockerImageInput() *ClusterDockerImage {
	var returns *ClusterDockerImage
	_jsii_.Get(
		j,
		"dockerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DriverInstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInstancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DriverInstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInstancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DriverNodeTypeFlexibility() ClusterDriverNodeTypeFlexibilityOutputReference {
	var returns ClusterDriverNodeTypeFlexibilityOutputReference
	_jsii_.Get(
		j,
		"driverNodeTypeFlexibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DriverNodeTypeFlexibilityInput() *ClusterDriverNodeTypeFlexibility {
	var returns *ClusterDriverNodeTypeFlexibility
	_jsii_.Get(
		j,
		"driverNodeTypeFlexibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DriverNodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DriverNodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) EnableElasticDisk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) EnableElasticDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) EnableLocalDiskEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalDiskEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) EnableLocalDiskEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalDiskEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) GcpAttributes() ClusterGcpAttributesOutputReference {
	var returns ClusterGcpAttributesOutputReference
	_jsii_.Get(
		j,
		"gcpAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) GcpAttributesInput() *ClusterGcpAttributes {
	var returns *ClusterGcpAttributes
	_jsii_.Get(
		j,
		"gcpAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) IdempotencyToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idempotencyToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) IdempotencyTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idempotencyTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) InitScripts() ClusterInitScriptsList {
	var returns ClusterInitScriptsList
	_jsii_.Get(
		j,
		"initScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) InitScriptsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initScriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) InstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) InstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) IsPinned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPinned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) IsPinnedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPinnedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) IsSingleNode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSingleNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) IsSingleNodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSingleNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Library() ClusterLibraryList {
	var returns ClusterLibraryList
	_jsii_.Get(
		j,
		"library",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) LibraryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"libraryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) NodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) NodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) NoWait() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noWait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) NoWaitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noWaitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) NumWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) NumWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderConfig() ClusterProviderConfigOutputReference {
	var returns ClusterProviderConfigOutputReference
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderConfigInput() *ClusterProviderConfig {
	var returns *ClusterProviderConfig
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) RemoteDiskThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"remoteDiskThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) RemoteDiskThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"remoteDiskThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) RuntimeEngine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) RuntimeEngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) SingleUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) SingleUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) SparkConf() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) SparkConfInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) SparkEnvVars() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) SparkEnvVarsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvVarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) SparkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) SparkVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) SshPublicKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Timeouts() ClusterTimeoutsOutputReference {
	var returns ClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) TotalInitialRemoteDiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalInitialRemoteDiskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) TotalInitialRemoteDiskSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalInitialRemoteDiskSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) UseMlRuntime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMlRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) UseMlRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMlRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) WorkerNodeTypeFlexibility() ClusterWorkerNodeTypeFlexibilityOutputReference {
	var returns ClusterWorkerNodeTypeFlexibilityOutputReference
	_jsii_.Get(
		j,
		"workerNodeTypeFlexibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) WorkerNodeTypeFlexibilityInput() *ClusterWorkerNodeTypeFlexibility {
	var returns *ClusterWorkerNodeTypeFlexibility
	_jsii_.Get(
		j,
		"workerNodeTypeFlexibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) WorkloadType() ClusterWorkloadTypeOutputReference {
	var returns ClusterWorkloadTypeOutputReference
	_jsii_.Get(
		j,
		"workloadType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) WorkloadTypeInput() *ClusterWorkloadType {
	var returns *ClusterWorkloadType
	_jsii_.Get(
		j,
		"workloadTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/cluster databricks_cluster} Resource.
func NewCluster(scope constructs.Construct, id *string, config *ClusterConfig) Cluster {
	_init_.Initialize()

	if err := validateNewClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Cluster{}

	_jsii_.Create(
		"@cdktn/provider-databricks.cluster.Cluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/cluster databricks_cluster} Resource.
func NewCluster_Override(c Cluster, scope constructs.Construct, id *string, config *ClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.cluster.Cluster",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_Cluster)SetApplyPolicyDefaultValues(val interface{}) {
	if err := j.validateSetApplyPolicyDefaultValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyPolicyDefaultValues",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetAutoterminationMinutes(val *float64) {
	if err := j.validateSetAutoterminationMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoterminationMinutes",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetClearCloudAttributesOnRemove(val interface{}) {
	if err := j.validateSetClearCloudAttributesOnRemoveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clearCloudAttributesOnRemove",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetCustomTags(val *map[string]*string) {
	if err := j.validateSetCustomTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customTags",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetDataSecurityMode(val *string) {
	if err := j.validateSetDataSecurityModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSecurityMode",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetDependencyMode(val *string) {
	if err := j.validateSetDependencyModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dependencyMode",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetDriverInstancePoolId(val *string) {
	if err := j.validateSetDriverInstancePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driverInstancePoolId",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetDriverNodeTypeId(val *string) {
	if err := j.validateSetDriverNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driverNodeTypeId",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetEnableElasticDisk(val interface{}) {
	if err := j.validateSetEnableElasticDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableElasticDisk",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetEnableLocalDiskEncryption(val interface{}) {
	if err := j.validateSetEnableLocalDiskEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLocalDiskEncryption",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetIdempotencyToken(val *string) {
	if err := j.validateSetIdempotencyTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idempotencyToken",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetInstancePoolId(val *string) {
	if err := j.validateSetInstancePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePoolId",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetIsPinned(val interface{}) {
	if err := j.validateSetIsPinnedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isPinned",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetIsSingleNode(val interface{}) {
	if err := j.validateSetIsSingleNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSingleNode",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetKind(val *string) {
	if err := j.validateSetKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetNodeTypeId(val *string) {
	if err := j.validateSetNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeTypeId",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetNoWait(val interface{}) {
	if err := j.validateSetNoWaitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noWait",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetNumWorkers(val *float64) {
	if err := j.validateSetNumWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numWorkers",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetRemoteDiskThroughput(val *float64) {
	if err := j.validateSetRemoteDiskThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteDiskThroughput",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetRuntimeEngine(val *string) {
	if err := j.validateSetRuntimeEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeEngine",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetSingleUserName(val *string) {
	if err := j.validateSetSingleUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleUserName",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetSparkConf(val *map[string]*string) {
	if err := j.validateSetSparkConfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkConf",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetSparkEnvVars(val *map[string]*string) {
	if err := j.validateSetSparkEnvVarsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkEnvVars",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetSparkVersion(val *string) {
	if err := j.validateSetSparkVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkVersion",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetSshPublicKeys(val *[]*string) {
	if err := j.validateSetSshPublicKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetTotalInitialRemoteDiskSize(val *float64) {
	if err := j.validateSetTotalInitialRemoteDiskSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalInitialRemoteDiskSize",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetUseMlRuntime(val interface{}) {
	if err := j.validateSetUseMlRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useMlRuntime",
		val,
	)
}

// Generates CDKTN code for importing a Cluster resource upon running "cdktn plan <stack-name>".
func Cluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.cluster.Cluster",
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
func Cluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.cluster.Cluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Cluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.cluster.Cluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Cluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.cluster.Cluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Cluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-databricks.cluster.Cluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_Cluster) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_Cluster) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_Cluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_Cluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_Cluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_Cluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_Cluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_Cluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_Cluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_Cluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_Cluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_Cluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_Cluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_Cluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := c.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_Cluster) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_Cluster) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_Cluster) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_Cluster) PutAutoscale(value *ClusterAutoscale) {
	if err := c.validatePutAutoscaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAutoscale",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) PutAwsAttributes(value *ClusterAwsAttributes) {
	if err := c.validatePutAwsAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsAttributes",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) PutAzureAttributes(value *ClusterAzureAttributes) {
	if err := c.validatePutAzureAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAzureAttributes",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) PutClusterLogConf(value *ClusterClusterLogConf) {
	if err := c.validatePutClusterLogConfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putClusterLogConf",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) PutClusterMountInfo(value interface{}) {
	if err := c.validatePutClusterMountInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putClusterMountInfo",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) PutDockerImage(value *ClusterDockerImage) {
	if err := c.validatePutDockerImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDockerImage",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) PutDriverNodeTypeFlexibility(value *ClusterDriverNodeTypeFlexibility) {
	if err := c.validatePutDriverNodeTypeFlexibilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDriverNodeTypeFlexibility",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) PutGcpAttributes(value *ClusterGcpAttributes) {
	if err := c.validatePutGcpAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGcpAttributes",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) PutInitScripts(value interface{}) {
	if err := c.validatePutInitScriptsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putInitScripts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) PutLibrary(value interface{}) {
	if err := c.validatePutLibraryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLibrary",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) PutProviderConfig(value *ClusterProviderConfig) {
	if err := c.validatePutProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putProviderConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) PutTimeouts(value *ClusterTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) PutWorkerNodeTypeFlexibility(value *ClusterWorkerNodeTypeFlexibility) {
	if err := c.validatePutWorkerNodeTypeFlexibilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWorkerNodeTypeFlexibility",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) PutWorkloadType(value *ClusterWorkloadType) {
	if err := c.validatePutWorkloadTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWorkloadType",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := c.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (c *jsiiProxy_Cluster) ResetApplyPolicyDefaultValues() {
	_jsii_.InvokeVoid(
		c,
		"resetApplyPolicyDefaultValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetAutoscale() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoscale",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetAutoterminationMinutes() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoterminationMinutes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetAwsAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetAzureAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetClearCloudAttributesOnRemove() {
	_jsii_.InvokeVoid(
		c,
		"resetClearCloudAttributesOnRemove",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetClusterLogConf() {
	_jsii_.InvokeVoid(
		c,
		"resetClusterLogConf",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetClusterMountInfo() {
	_jsii_.InvokeVoid(
		c,
		"resetClusterMountInfo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetClusterName() {
	_jsii_.InvokeVoid(
		c,
		"resetClusterName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetCustomTags() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetDataSecurityMode() {
	_jsii_.InvokeVoid(
		c,
		"resetDataSecurityMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetDependencyMode() {
	_jsii_.InvokeVoid(
		c,
		"resetDependencyMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetDockerImage() {
	_jsii_.InvokeVoid(
		c,
		"resetDockerImage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetDriverInstancePoolId() {
	_jsii_.InvokeVoid(
		c,
		"resetDriverInstancePoolId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetDriverNodeTypeFlexibility() {
	_jsii_.InvokeVoid(
		c,
		"resetDriverNodeTypeFlexibility",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetDriverNodeTypeId() {
	_jsii_.InvokeVoid(
		c,
		"resetDriverNodeTypeId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetEnableElasticDisk() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableElasticDisk",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetEnableLocalDiskEncryption() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableLocalDiskEncryption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetGcpAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetGcpAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetIdempotencyToken() {
	_jsii_.InvokeVoid(
		c,
		"resetIdempotencyToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetInitScripts() {
	_jsii_.InvokeVoid(
		c,
		"resetInitScripts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetInstancePoolId() {
	_jsii_.InvokeVoid(
		c,
		"resetInstancePoolId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetIsPinned() {
	_jsii_.InvokeVoid(
		c,
		"resetIsPinned",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetIsSingleNode() {
	_jsii_.InvokeVoid(
		c,
		"resetIsSingleNode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetKind() {
	_jsii_.InvokeVoid(
		c,
		"resetKind",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetLibrary() {
	_jsii_.InvokeVoid(
		c,
		"resetLibrary",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetNodeTypeId() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeTypeId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetNoWait() {
	_jsii_.InvokeVoid(
		c,
		"resetNoWait",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetNumWorkers() {
	_jsii_.InvokeVoid(
		c,
		"resetNumWorkers",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetPolicyId() {
	_jsii_.InvokeVoid(
		c,
		"resetPolicyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetRemoteDiskThroughput() {
	_jsii_.InvokeVoid(
		c,
		"resetRemoteDiskThroughput",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetRuntimeEngine() {
	_jsii_.InvokeVoid(
		c,
		"resetRuntimeEngine",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetSingleUserName() {
	_jsii_.InvokeVoid(
		c,
		"resetSingleUserName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetSparkConf() {
	_jsii_.InvokeVoid(
		c,
		"resetSparkConf",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetSparkEnvVars() {
	_jsii_.InvokeVoid(
		c,
		"resetSparkEnvVars",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetSshPublicKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetSshPublicKeys",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetTotalInitialRemoteDiskSize() {
	_jsii_.InvokeVoid(
		c,
		"resetTotalInitialRemoteDiskSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetUseMlRuntime() {
	_jsii_.InvokeVoid(
		c,
		"resetUseMlRuntime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetWorkerNodeTypeFlexibility() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkerNodeTypeFlexibility",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetWorkloadType() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkloadType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

