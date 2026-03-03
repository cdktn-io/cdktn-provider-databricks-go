// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/appsspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppsSpaceResourcesOutputReference interface {
	cdktn.ComplexObject
	App() AppsSpaceResourcesAppOutputReference
	AppInput() interface{}
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Database() AppsSpaceResourcesDatabaseOutputReference
	DatabaseInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Experiment() AppsSpaceResourcesExperimentOutputReference
	ExperimentInput() interface{}
	// Experimental.
	Fqn() *string
	GenieSpace() AppsSpaceResourcesGenieSpaceOutputReference
	GenieSpaceInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Job() AppsSpaceResourcesJobOutputReference
	JobInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Secret() AppsSpaceResourcesSecretOutputReference
	SecretInput() interface{}
	ServingEndpoint() AppsSpaceResourcesServingEndpointOutputReference
	ServingEndpointInput() interface{}
	SqlWarehouse() AppsSpaceResourcesSqlWarehouseOutputReference
	SqlWarehouseInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UcSecurable() AppsSpaceResourcesUcSecurableOutputReference
	UcSecurableInput() interface{}
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
	PutApp(value *AppsSpaceResourcesApp)
	PutDatabase(value *AppsSpaceResourcesDatabase)
	PutExperiment(value *AppsSpaceResourcesExperiment)
	PutGenieSpace(value *AppsSpaceResourcesGenieSpace)
	PutJob(value *AppsSpaceResourcesJob)
	PutSecret(value *AppsSpaceResourcesSecret)
	PutServingEndpoint(value *AppsSpaceResourcesServingEndpoint)
	PutSqlWarehouse(value *AppsSpaceResourcesSqlWarehouse)
	PutUcSecurable(value *AppsSpaceResourcesUcSecurable)
	ResetApp()
	ResetDatabase()
	ResetDescription()
	ResetExperiment()
	ResetGenieSpace()
	ResetJob()
	ResetSecret()
	ResetServingEndpoint()
	ResetSqlWarehouse()
	ResetUcSecurable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppsSpaceResourcesOutputReference
type jsiiProxy_AppsSpaceResourcesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) App() AppsSpaceResourcesAppOutputReference {
	var returns AppsSpaceResourcesAppOutputReference
	_jsii_.Get(
		j,
		"app",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) AppInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) Database() AppsSpaceResourcesDatabaseOutputReference {
	var returns AppsSpaceResourcesDatabaseOutputReference
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) DatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) Experiment() AppsSpaceResourcesExperimentOutputReference {
	var returns AppsSpaceResourcesExperimentOutputReference
	_jsii_.Get(
		j,
		"experiment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) ExperimentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"experimentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) GenieSpace() AppsSpaceResourcesGenieSpaceOutputReference {
	var returns AppsSpaceResourcesGenieSpaceOutputReference
	_jsii_.Get(
		j,
		"genieSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) GenieSpaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"genieSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) Job() AppsSpaceResourcesJobOutputReference {
	var returns AppsSpaceResourcesJobOutputReference
	_jsii_.Get(
		j,
		"job",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) JobInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) Secret() AppsSpaceResourcesSecretOutputReference {
	var returns AppsSpaceResourcesSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) ServingEndpoint() AppsSpaceResourcesServingEndpointOutputReference {
	var returns AppsSpaceResourcesServingEndpointOutputReference
	_jsii_.Get(
		j,
		"servingEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) ServingEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"servingEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) SqlWarehouse() AppsSpaceResourcesSqlWarehouseOutputReference {
	var returns AppsSpaceResourcesSqlWarehouseOutputReference
	_jsii_.Get(
		j,
		"sqlWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) SqlWarehouseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlWarehouseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) UcSecurable() AppsSpaceResourcesUcSecurableOutputReference {
	var returns AppsSpaceResourcesUcSecurableOutputReference
	_jsii_.Get(
		j,
		"ucSecurable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference) UcSecurableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ucSecurableInput",
		&returns,
	)
	return returns
}


func NewAppsSpaceResourcesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppsSpaceResourcesOutputReference {
	_init_.Initialize()

	if err := validateNewAppsSpaceResourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsSpaceResourcesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.appsSpace.AppsSpaceResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppsSpaceResourcesOutputReference_Override(a AppsSpaceResourcesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.appsSpace.AppsSpaceResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsSpaceResourcesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) PutApp(value *AppsSpaceResourcesApp) {
	if err := a.validatePutAppParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) PutDatabase(value *AppsSpaceResourcesDatabase) {
	if err := a.validatePutDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatabase",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) PutExperiment(value *AppsSpaceResourcesExperiment) {
	if err := a.validatePutExperimentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExperiment",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) PutGenieSpace(value *AppsSpaceResourcesGenieSpace) {
	if err := a.validatePutGenieSpaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGenieSpace",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) PutJob(value *AppsSpaceResourcesJob) {
	if err := a.validatePutJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJob",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) PutSecret(value *AppsSpaceResourcesSecret) {
	if err := a.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecret",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) PutServingEndpoint(value *AppsSpaceResourcesServingEndpoint) {
	if err := a.validatePutServingEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServingEndpoint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) PutSqlWarehouse(value *AppsSpaceResourcesSqlWarehouse) {
	if err := a.validatePutSqlWarehouseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqlWarehouse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) PutUcSecurable(value *AppsSpaceResourcesUcSecurable) {
	if err := a.validatePutUcSecurableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUcSecurable",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) ResetApp() {
	_jsii_.InvokeVoid(
		a,
		"resetApp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		a,
		"resetDatabase",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) ResetExperiment() {
	_jsii_.InvokeVoid(
		a,
		"resetExperiment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) ResetGenieSpace() {
	_jsii_.InvokeVoid(
		a,
		"resetGenieSpace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) ResetJob() {
	_jsii_.InvokeVoid(
		a,
		"resetJob",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) ResetServingEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetServingEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) ResetSqlWarehouse() {
	_jsii_.InvokeVoid(
		a,
		"resetSqlWarehouse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) ResetUcSecurable() {
	_jsii_.InvokeVoid(
		a,
		"resetUcSecurable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsSpaceResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

