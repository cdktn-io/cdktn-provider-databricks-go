// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aibidashboardembeddingapproveddomainssetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/aibidashboardembeddingapproveddomainssetting/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/aibi_dashboard_embedding_approved_domains_setting databricks_aibi_dashboard_embedding_approved_domains_setting}.
type AibiDashboardEmbeddingApprovedDomainsSetting interface {
	cdktn.TerraformResource
	AibiDashboardEmbeddingApprovedDomains() AibiDashboardEmbeddingApprovedDomainsSettingAibiDashboardEmbeddingApprovedDomainsOutputReference
	AibiDashboardEmbeddingApprovedDomainsInput() *AibiDashboardEmbeddingApprovedDomainsSettingAibiDashboardEmbeddingApprovedDomains
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Etag() *string
	SetEtag(val *string)
	EtagInput() *string
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
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProviderConfig() AibiDashboardEmbeddingApprovedDomainsSettingProviderConfigOutputReference
	ProviderConfigInput() *AibiDashboardEmbeddingApprovedDomainsSettingProviderConfig
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SettingName() *string
	SetSettingName(val *string)
	SettingNameInput() *string
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
	PutAibiDashboardEmbeddingApprovedDomains(value *AibiDashboardEmbeddingApprovedDomainsSettingAibiDashboardEmbeddingApprovedDomains)
	PutProviderConfig(value *AibiDashboardEmbeddingApprovedDomainsSettingProviderConfig)
	ResetEtag()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProviderConfig()
	ResetSettingName()
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

// The jsii proxy struct for AibiDashboardEmbeddingApprovedDomainsSetting
type jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) AibiDashboardEmbeddingApprovedDomains() AibiDashboardEmbeddingApprovedDomainsSettingAibiDashboardEmbeddingApprovedDomainsOutputReference {
	var returns AibiDashboardEmbeddingApprovedDomainsSettingAibiDashboardEmbeddingApprovedDomainsOutputReference
	_jsii_.Get(
		j,
		"aibiDashboardEmbeddingApprovedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) AibiDashboardEmbeddingApprovedDomainsInput() *AibiDashboardEmbeddingApprovedDomainsSettingAibiDashboardEmbeddingApprovedDomains {
	var returns *AibiDashboardEmbeddingApprovedDomainsSettingAibiDashboardEmbeddingApprovedDomains
	_jsii_.Get(
		j,
		"aibiDashboardEmbeddingApprovedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) EtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) ProviderConfig() AibiDashboardEmbeddingApprovedDomainsSettingProviderConfigOutputReference {
	var returns AibiDashboardEmbeddingApprovedDomainsSettingProviderConfigOutputReference
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) ProviderConfigInput() *AibiDashboardEmbeddingApprovedDomainsSettingProviderConfig {
	var returns *AibiDashboardEmbeddingApprovedDomainsSettingProviderConfig
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) SettingName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"settingName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) SettingNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"settingNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/aibi_dashboard_embedding_approved_domains_setting databricks_aibi_dashboard_embedding_approved_domains_setting} Resource.
func NewAibiDashboardEmbeddingApprovedDomainsSetting(scope constructs.Construct, id *string, config *AibiDashboardEmbeddingApprovedDomainsSettingConfig) AibiDashboardEmbeddingApprovedDomainsSetting {
	_init_.Initialize()

	if err := validateNewAibiDashboardEmbeddingApprovedDomainsSettingParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting{}

	_jsii_.Create(
		"@cdktn/provider-databricks.aibiDashboardEmbeddingApprovedDomainsSetting.AibiDashboardEmbeddingApprovedDomainsSetting",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/aibi_dashboard_embedding_approved_domains_setting databricks_aibi_dashboard_embedding_approved_domains_setting} Resource.
func NewAibiDashboardEmbeddingApprovedDomainsSetting_Override(a AibiDashboardEmbeddingApprovedDomainsSetting, scope constructs.Construct, id *string, config *AibiDashboardEmbeddingApprovedDomainsSettingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.aibiDashboardEmbeddingApprovedDomainsSetting.AibiDashboardEmbeddingApprovedDomainsSetting",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting)SetEtag(val *string) {
	if err := j.validateSetEtagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etag",
		val,
	)
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting)SetSettingName(val *string) {
	if err := j.validateSetSettingNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"settingName",
		val,
	)
}

// Generates CDKTN code for importing a AibiDashboardEmbeddingApprovedDomainsSetting resource upon running "cdktn plan <stack-name>".
func AibiDashboardEmbeddingApprovedDomainsSetting_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAibiDashboardEmbeddingApprovedDomainsSetting_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.aibiDashboardEmbeddingApprovedDomainsSetting.AibiDashboardEmbeddingApprovedDomainsSetting",
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
func AibiDashboardEmbeddingApprovedDomainsSetting_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAibiDashboardEmbeddingApprovedDomainsSetting_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.aibiDashboardEmbeddingApprovedDomainsSetting.AibiDashboardEmbeddingApprovedDomainsSetting",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AibiDashboardEmbeddingApprovedDomainsSetting_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAibiDashboardEmbeddingApprovedDomainsSetting_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.aibiDashboardEmbeddingApprovedDomainsSetting.AibiDashboardEmbeddingApprovedDomainsSetting",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AibiDashboardEmbeddingApprovedDomainsSetting_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAibiDashboardEmbeddingApprovedDomainsSetting_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-databricks.aibiDashboardEmbeddingApprovedDomainsSetting.AibiDashboardEmbeddingApprovedDomainsSetting",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AibiDashboardEmbeddingApprovedDomainsSetting_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-databricks.aibiDashboardEmbeddingApprovedDomainsSetting.AibiDashboardEmbeddingApprovedDomainsSetting",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) PutAibiDashboardEmbeddingApprovedDomains(value *AibiDashboardEmbeddingApprovedDomainsSettingAibiDashboardEmbeddingApprovedDomains) {
	if err := a.validatePutAibiDashboardEmbeddingApprovedDomainsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAibiDashboardEmbeddingApprovedDomains",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) PutProviderConfig(value *AibiDashboardEmbeddingApprovedDomainsSettingProviderConfig) {
	if err := a.validatePutProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProviderConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) ResetEtag() {
	_jsii_.InvokeVoid(
		a,
		"resetEtag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) ResetSettingName() {
	_jsii_.InvokeVoid(
		a,
		"resetSettingName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AibiDashboardEmbeddingApprovedDomainsSetting) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

