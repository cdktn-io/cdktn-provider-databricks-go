// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksapp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAppAppOutputReference interface {
	cdktn.ComplexObject
	ActiveDeployment() DataDatabricksAppAppActiveDeploymentOutputReference
	AppStatus() DataDatabricksAppAppAppStatusOutputReference
	BudgetPolicyId() *string
	SetBudgetPolicyId(val *string)
	BudgetPolicyIdInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ComputeMaxInstances() *float64
	SetComputeMaxInstances(val *float64)
	ComputeMaxInstancesInput() *float64
	ComputeMinInstances() *float64
	SetComputeMinInstances(val *float64)
	ComputeMinInstancesInput() *float64
	ComputeSize() *string
	SetComputeSize(val *string)
	ComputeSizeInput() *string
	ComputeStatus() DataDatabricksAppAppComputeStatusOutputReference
	CreateTime() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Creator() *string
	DefaultGitSource() DataDatabricksAppAppDefaultGitSourceOutputReference
	DefaultSourceCodePath() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveBudgetPolicyId() *string
	EffectiveUsagePolicyId() *string
	EffectiveUserApiScopes() *[]*string
	ForwardUserAccessToken() interface{}
	SetForwardUserAccessToken(val interface{})
	ForwardUserAccessTokenInput() interface{}
	// Experimental.
	Fqn() *string
	GitRepository() DataDatabricksAppAppGitRepositoryOutputReference
	GitRepositoryInput() interface{}
	GitSource() DataDatabricksAppAppGitSourceOutputReference
	GitSourceInput() interface{}
	Id() *string
	InternalValue() *DataDatabricksAppApp
	SetInternalValue(val *DataDatabricksAppApp)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Oauth2AppClientId() *string
	Oauth2AppIntegrationId() *string
	PendingDeployment() DataDatabricksAppAppPendingDeploymentOutputReference
	Resources() DataDatabricksAppAppResourcesList
	ResourcesInput() interface{}
	ServicePrincipalClientId() *string
	ServicePrincipalId() *float64
	ServicePrincipalName() *string
	SourceCodePath() *string
	SetSourceCodePath(val *string)
	SourceCodePathInput() *string
	Space() *string
	SetSpace(val *string)
	SpaceInput() *string
	TelemetryExportDestinations() DataDatabricksAppAppTelemetryExportDestinationsList
	TelemetryExportDestinationsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThumbnailUrl() *string
	Updater() *string
	UpdateTime() *string
	Url() *string
	UsagePolicyId() *string
	SetUsagePolicyId(val *string)
	UsagePolicyIdInput() *string
	UserApiScopes() *[]*string
	SetUserApiScopes(val *[]*string)
	UserApiScopesInput() *[]*string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutGitRepository(value *DataDatabricksAppAppGitRepository)
	PutGitSource(value *DataDatabricksAppAppGitSource)
	PutResources(value interface{})
	PutTelemetryExportDestinations(value interface{})
	ResetBudgetPolicyId()
	ResetComputeMaxInstances()
	ResetComputeMinInstances()
	ResetComputeSize()
	ResetDescription()
	ResetForwardUserAccessToken()
	ResetGitRepository()
	ResetGitSource()
	ResetResources()
	ResetSourceCodePath()
	ResetSpace()
	ResetTelemetryExportDestinations()
	ResetUsagePolicyId()
	ResetUserApiScopes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAppAppOutputReference
type jsiiProxy_DataDatabricksAppAppOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ActiveDeployment() DataDatabricksAppAppActiveDeploymentOutputReference {
	var returns DataDatabricksAppAppActiveDeploymentOutputReference
	_jsii_.Get(
		j,
		"activeDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) AppStatus() DataDatabricksAppAppAppStatusOutputReference {
	var returns DataDatabricksAppAppAppStatusOutputReference
	_jsii_.Get(
		j,
		"appStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) BudgetPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) BudgetPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ComputeMaxInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeMaxInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ComputeMaxInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeMaxInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ComputeMinInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeMinInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ComputeMinInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeMinInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ComputeSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ComputeSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ComputeStatus() DataDatabricksAppAppComputeStatusOutputReference {
	var returns DataDatabricksAppAppComputeStatusOutputReference
	_jsii_.Get(
		j,
		"computeStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) Creator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) DefaultGitSource() DataDatabricksAppAppDefaultGitSourceOutputReference {
	var returns DataDatabricksAppAppDefaultGitSourceOutputReference
	_jsii_.Get(
		j,
		"defaultGitSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) DefaultSourceCodePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSourceCodePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) EffectiveBudgetPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveBudgetPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) EffectiveUsagePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveUsagePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) EffectiveUserApiScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"effectiveUserApiScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ForwardUserAccessToken() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardUserAccessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ForwardUserAccessTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardUserAccessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) GitRepository() DataDatabricksAppAppGitRepositoryOutputReference {
	var returns DataDatabricksAppAppGitRepositoryOutputReference
	_jsii_.Get(
		j,
		"gitRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) GitRepositoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gitRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) GitSource() DataDatabricksAppAppGitSourceOutputReference {
	var returns DataDatabricksAppAppGitSourceOutputReference
	_jsii_.Get(
		j,
		"gitSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) GitSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gitSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) InternalValue() *DataDatabricksAppApp {
	var returns *DataDatabricksAppApp
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) Oauth2AppClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2AppClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) Oauth2AppIntegrationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2AppIntegrationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) PendingDeployment() DataDatabricksAppAppPendingDeploymentOutputReference {
	var returns DataDatabricksAppAppPendingDeploymentOutputReference
	_jsii_.Get(
		j,
		"pendingDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) Resources() DataDatabricksAppAppResourcesList {
	var returns DataDatabricksAppAppResourcesList
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ServicePrincipalClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ServicePrincipalId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"servicePrincipalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ServicePrincipalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) SourceCodePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCodePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) SourceCodePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCodePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) Space() *string {
	var returns *string
	_jsii_.Get(
		j,
		"space",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) SpaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) TelemetryExportDestinations() DataDatabricksAppAppTelemetryExportDestinationsList {
	var returns DataDatabricksAppAppTelemetryExportDestinationsList
	_jsii_.Get(
		j,
		"telemetryExportDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) TelemetryExportDestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"telemetryExportDestinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ThumbnailUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbnailUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) Updater() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updater",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) UsagePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) UsagePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usagePolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) UserApiScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userApiScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) UserApiScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userApiScopesInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksAppAppOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAppAppOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAppAppOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAppAppOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksApp.DataDatabricksAppAppOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAppAppOutputReference_Override(d DataDatabricksAppAppOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksApp.DataDatabricksAppAppOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetBudgetPolicyId(val *string) {
	if err := j.validateSetBudgetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"budgetPolicyId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetComputeMaxInstances(val *float64) {
	if err := j.validateSetComputeMaxInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeMaxInstances",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetComputeMinInstances(val *float64) {
	if err := j.validateSetComputeMinInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeMinInstances",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetComputeSize(val *string) {
	if err := j.validateSetComputeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeSize",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetForwardUserAccessToken(val interface{}) {
	if err := j.validateSetForwardUserAccessTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardUserAccessToken",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetInternalValue(val *DataDatabricksAppApp) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetSourceCodePath(val *string) {
	if err := j.validateSetSourceCodePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCodePath",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetSpace(val *string) {
	if err := j.validateSetSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"space",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetUsagePolicyId(val *string) {
	if err := j.validateSetUsagePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usagePolicyId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetUserApiScopes(val *[]*string) {
	if err := j.validateSetUserApiScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userApiScopes",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) PutGitRepository(value *DataDatabricksAppAppGitRepository) {
	if err := d.validatePutGitRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGitRepository",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) PutGitSource(value *DataDatabricksAppAppGitSource) {
	if err := d.validatePutGitSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGitSource",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) PutResources(value interface{}) {
	if err := d.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putResources",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) PutTelemetryExportDestinations(value interface{}) {
	if err := d.validatePutTelemetryExportDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTelemetryExportDestinations",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetBudgetPolicyId() {
	_jsii_.InvokeVoid(
		d,
		"resetBudgetPolicyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetComputeMaxInstances() {
	_jsii_.InvokeVoid(
		d,
		"resetComputeMaxInstances",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetComputeMinInstances() {
	_jsii_.InvokeVoid(
		d,
		"resetComputeMinInstances",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetComputeSize() {
	_jsii_.InvokeVoid(
		d,
		"resetComputeSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetForwardUserAccessToken() {
	_jsii_.InvokeVoid(
		d,
		"resetForwardUserAccessToken",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetGitRepository() {
	_jsii_.InvokeVoid(
		d,
		"resetGitRepository",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetGitSource() {
	_jsii_.InvokeVoid(
		d,
		"resetGitSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		d,
		"resetResources",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetSourceCodePath() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceCodePath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetSpace() {
	_jsii_.InvokeVoid(
		d,
		"resetSpace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetTelemetryExportDestinations() {
	_jsii_.InvokeVoid(
		d,
		"resetTelemetryExportDestinations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetUsagePolicyId() {
	_jsii_.InvokeVoid(
		d,
		"resetUsagePolicyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetUserApiScopes() {
	_jsii_.InvokeVoid(
		d,
		"resetUserApiScopes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

