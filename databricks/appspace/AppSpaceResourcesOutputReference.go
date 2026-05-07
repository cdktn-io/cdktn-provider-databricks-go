// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/appspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppSpaceResourcesOutputReference interface {
	cdktn.ComplexObject
	App() AppSpaceResourcesAppOutputReference
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
	Database() AppSpaceResourcesDatabaseOutputReference
	DatabaseInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Experiment() AppSpaceResourcesExperimentOutputReference
	ExperimentInput() interface{}
	// Experimental.
	Fqn() *string
	GenieSpace() AppSpaceResourcesGenieSpaceOutputReference
	GenieSpaceInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Job() AppSpaceResourcesJobOutputReference
	JobInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Postgres() AppSpaceResourcesPostgresOutputReference
	PostgresInput() interface{}
	Secret() AppSpaceResourcesSecretOutputReference
	SecretInput() interface{}
	ServingEndpoint() AppSpaceResourcesServingEndpointOutputReference
	ServingEndpointInput() interface{}
	SqlWarehouse() AppSpaceResourcesSqlWarehouseOutputReference
	SqlWarehouseInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UcSecurable() AppSpaceResourcesUcSecurableOutputReference
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
	PutApp(value *AppSpaceResourcesApp)
	PutDatabase(value *AppSpaceResourcesDatabase)
	PutExperiment(value *AppSpaceResourcesExperiment)
	PutGenieSpace(value *AppSpaceResourcesGenieSpace)
	PutJob(value *AppSpaceResourcesJob)
	PutPostgres(value *AppSpaceResourcesPostgres)
	PutSecret(value *AppSpaceResourcesSecret)
	PutServingEndpoint(value *AppSpaceResourcesServingEndpoint)
	PutSqlWarehouse(value *AppSpaceResourcesSqlWarehouse)
	PutUcSecurable(value *AppSpaceResourcesUcSecurable)
	ResetApp()
	ResetDatabase()
	ResetDescription()
	ResetExperiment()
	ResetGenieSpace()
	ResetJob()
	ResetPostgres()
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

// The jsii proxy struct for AppSpaceResourcesOutputReference
type jsiiProxy_AppSpaceResourcesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) App() AppSpaceResourcesAppOutputReference {
	var returns AppSpaceResourcesAppOutputReference
	_jsii_.Get(
		j,
		"app",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) AppInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) Database() AppSpaceResourcesDatabaseOutputReference {
	var returns AppSpaceResourcesDatabaseOutputReference
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) DatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) Experiment() AppSpaceResourcesExperimentOutputReference {
	var returns AppSpaceResourcesExperimentOutputReference
	_jsii_.Get(
		j,
		"experiment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) ExperimentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"experimentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) GenieSpace() AppSpaceResourcesGenieSpaceOutputReference {
	var returns AppSpaceResourcesGenieSpaceOutputReference
	_jsii_.Get(
		j,
		"genieSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) GenieSpaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"genieSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) Job() AppSpaceResourcesJobOutputReference {
	var returns AppSpaceResourcesJobOutputReference
	_jsii_.Get(
		j,
		"job",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) JobInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) Postgres() AppSpaceResourcesPostgresOutputReference {
	var returns AppSpaceResourcesPostgresOutputReference
	_jsii_.Get(
		j,
		"postgres",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) PostgresInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postgresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) Secret() AppSpaceResourcesSecretOutputReference {
	var returns AppSpaceResourcesSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) ServingEndpoint() AppSpaceResourcesServingEndpointOutputReference {
	var returns AppSpaceResourcesServingEndpointOutputReference
	_jsii_.Get(
		j,
		"servingEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) ServingEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"servingEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) SqlWarehouse() AppSpaceResourcesSqlWarehouseOutputReference {
	var returns AppSpaceResourcesSqlWarehouseOutputReference
	_jsii_.Get(
		j,
		"sqlWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) SqlWarehouseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlWarehouseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) UcSecurable() AppSpaceResourcesUcSecurableOutputReference {
	var returns AppSpaceResourcesUcSecurableOutputReference
	_jsii_.Get(
		j,
		"ucSecurable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference) UcSecurableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ucSecurableInput",
		&returns,
	)
	return returns
}


func NewAppSpaceResourcesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppSpaceResourcesOutputReference {
	_init_.Initialize()

	if err := validateNewAppSpaceResourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSpaceResourcesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.appSpace.AppSpaceResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppSpaceResourcesOutputReference_Override(a AppSpaceResourcesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.appSpace.AppSpaceResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppSpaceResourcesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppSpaceResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AppSpaceResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppSpaceResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppSpaceResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppSpaceResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppSpaceResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppSpaceResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppSpaceResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppSpaceResourcesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AppSpaceResourcesOutputReference) PutApp(value *AppSpaceResourcesApp) {
	if err := a.validatePutAppParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) PutDatabase(value *AppSpaceResourcesDatabase) {
	if err := a.validatePutDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatabase",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) PutExperiment(value *AppSpaceResourcesExperiment) {
	if err := a.validatePutExperimentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExperiment",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) PutGenieSpace(value *AppSpaceResourcesGenieSpace) {
	if err := a.validatePutGenieSpaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGenieSpace",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) PutJob(value *AppSpaceResourcesJob) {
	if err := a.validatePutJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJob",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) PutPostgres(value *AppSpaceResourcesPostgres) {
	if err := a.validatePutPostgresParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPostgres",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) PutSecret(value *AppSpaceResourcesSecret) {
	if err := a.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecret",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) PutServingEndpoint(value *AppSpaceResourcesServingEndpoint) {
	if err := a.validatePutServingEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServingEndpoint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) PutSqlWarehouse(value *AppSpaceResourcesSqlWarehouse) {
	if err := a.validatePutSqlWarehouseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqlWarehouse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) PutUcSecurable(value *AppSpaceResourcesUcSecurable) {
	if err := a.validatePutUcSecurableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUcSecurable",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) ResetApp() {
	_jsii_.InvokeVoid(
		a,
		"resetApp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		a,
		"resetDatabase",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) ResetExperiment() {
	_jsii_.InvokeVoid(
		a,
		"resetExperiment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) ResetGenieSpace() {
	_jsii_.InvokeVoid(
		a,
		"resetGenieSpace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) ResetJob() {
	_jsii_.InvokeVoid(
		a,
		"resetJob",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) ResetPostgres() {
	_jsii_.InvokeVoid(
		a,
		"resetPostgres",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) ResetServingEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetServingEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) ResetSqlWarehouse() {
	_jsii_.InvokeVoid(
		a,
		"resetSqlWarehouse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) ResetUcSecurable() {
	_jsii_.InvokeVoid(
		a,
		"resetUcSecurable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpaceResourcesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppSpaceResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

