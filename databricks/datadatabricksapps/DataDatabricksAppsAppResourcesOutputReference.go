// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksapps/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAppsAppResourcesOutputReference interface {
	cdktn.ComplexObject
	App() DataDatabricksAppsAppResourcesAppOutputReference
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
	Database() DataDatabricksAppsAppResourcesDatabaseOutputReference
	DatabaseInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Experiment() DataDatabricksAppsAppResourcesExperimentOutputReference
	ExperimentInput() interface{}
	// Experimental.
	Fqn() *string
	GenieSpace() DataDatabricksAppsAppResourcesGenieSpaceOutputReference
	GenieSpaceInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Job() DataDatabricksAppsAppResourcesJobOutputReference
	JobInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Postgres() DataDatabricksAppsAppResourcesPostgresOutputReference
	PostgresInput() interface{}
	Secret() DataDatabricksAppsAppResourcesSecretOutputReference
	SecretInput() interface{}
	ServingEndpoint() DataDatabricksAppsAppResourcesServingEndpointOutputReference
	ServingEndpointInput() interface{}
	SqlWarehouse() DataDatabricksAppsAppResourcesSqlWarehouseOutputReference
	SqlWarehouseInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UcSecurable() DataDatabricksAppsAppResourcesUcSecurableOutputReference
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
	PutApp(value *DataDatabricksAppsAppResourcesApp)
	PutDatabase(value *DataDatabricksAppsAppResourcesDatabase)
	PutExperiment(value *DataDatabricksAppsAppResourcesExperiment)
	PutGenieSpace(value *DataDatabricksAppsAppResourcesGenieSpace)
	PutJob(value *DataDatabricksAppsAppResourcesJob)
	PutPostgres(value *DataDatabricksAppsAppResourcesPostgres)
	PutSecret(value *DataDatabricksAppsAppResourcesSecret)
	PutServingEndpoint(value *DataDatabricksAppsAppResourcesServingEndpoint)
	PutSqlWarehouse(value *DataDatabricksAppsAppResourcesSqlWarehouse)
	PutUcSecurable(value *DataDatabricksAppsAppResourcesUcSecurable)
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

// The jsii proxy struct for DataDatabricksAppsAppResourcesOutputReference
type jsiiProxy_DataDatabricksAppsAppResourcesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) App() DataDatabricksAppsAppResourcesAppOutputReference {
	var returns DataDatabricksAppsAppResourcesAppOutputReference
	_jsii_.Get(
		j,
		"app",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) AppInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) Database() DataDatabricksAppsAppResourcesDatabaseOutputReference {
	var returns DataDatabricksAppsAppResourcesDatabaseOutputReference
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) DatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) Experiment() DataDatabricksAppsAppResourcesExperimentOutputReference {
	var returns DataDatabricksAppsAppResourcesExperimentOutputReference
	_jsii_.Get(
		j,
		"experiment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) ExperimentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"experimentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) GenieSpace() DataDatabricksAppsAppResourcesGenieSpaceOutputReference {
	var returns DataDatabricksAppsAppResourcesGenieSpaceOutputReference
	_jsii_.Get(
		j,
		"genieSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) GenieSpaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"genieSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) Job() DataDatabricksAppsAppResourcesJobOutputReference {
	var returns DataDatabricksAppsAppResourcesJobOutputReference
	_jsii_.Get(
		j,
		"job",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) JobInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) Postgres() DataDatabricksAppsAppResourcesPostgresOutputReference {
	var returns DataDatabricksAppsAppResourcesPostgresOutputReference
	_jsii_.Get(
		j,
		"postgres",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) PostgresInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postgresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) Secret() DataDatabricksAppsAppResourcesSecretOutputReference {
	var returns DataDatabricksAppsAppResourcesSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) ServingEndpoint() DataDatabricksAppsAppResourcesServingEndpointOutputReference {
	var returns DataDatabricksAppsAppResourcesServingEndpointOutputReference
	_jsii_.Get(
		j,
		"servingEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) ServingEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"servingEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) SqlWarehouse() DataDatabricksAppsAppResourcesSqlWarehouseOutputReference {
	var returns DataDatabricksAppsAppResourcesSqlWarehouseOutputReference
	_jsii_.Get(
		j,
		"sqlWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) SqlWarehouseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlWarehouseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) UcSecurable() DataDatabricksAppsAppResourcesUcSecurableOutputReference {
	var returns DataDatabricksAppsAppResourcesUcSecurableOutputReference
	_jsii_.Get(
		j,
		"ucSecurable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) UcSecurableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ucSecurableInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksAppsAppResourcesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksAppsAppResourcesOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAppsAppResourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAppsAppResourcesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksApps.DataDatabricksAppsAppResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksAppsAppResourcesOutputReference_Override(d DataDatabricksAppsAppResourcesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksApps.DataDatabricksAppsAppResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) PutApp(value *DataDatabricksAppsAppResourcesApp) {
	if err := d.validatePutAppParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApp",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) PutDatabase(value *DataDatabricksAppsAppResourcesDatabase) {
	if err := d.validatePutDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDatabase",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) PutExperiment(value *DataDatabricksAppsAppResourcesExperiment) {
	if err := d.validatePutExperimentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExperiment",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) PutGenieSpace(value *DataDatabricksAppsAppResourcesGenieSpace) {
	if err := d.validatePutGenieSpaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGenieSpace",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) PutJob(value *DataDatabricksAppsAppResourcesJob) {
	if err := d.validatePutJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJob",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) PutPostgres(value *DataDatabricksAppsAppResourcesPostgres) {
	if err := d.validatePutPostgresParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPostgres",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) PutSecret(value *DataDatabricksAppsAppResourcesSecret) {
	if err := d.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecret",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) PutServingEndpoint(value *DataDatabricksAppsAppResourcesServingEndpoint) {
	if err := d.validatePutServingEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServingEndpoint",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) PutSqlWarehouse(value *DataDatabricksAppsAppResourcesSqlWarehouse) {
	if err := d.validatePutSqlWarehouseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSqlWarehouse",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) PutUcSecurable(value *DataDatabricksAppsAppResourcesUcSecurable) {
	if err := d.validatePutUcSecurableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUcSecurable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) ResetApp() {
	_jsii_.InvokeVoid(
		d,
		"resetApp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) ResetExperiment() {
	_jsii_.InvokeVoid(
		d,
		"resetExperiment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) ResetGenieSpace() {
	_jsii_.InvokeVoid(
		d,
		"resetGenieSpace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) ResetJob() {
	_jsii_.InvokeVoid(
		d,
		"resetJob",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) ResetPostgres() {
	_jsii_.InvokeVoid(
		d,
		"resetPostgres",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		d,
		"resetSecret",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) ResetServingEndpoint() {
	_jsii_.InvokeVoid(
		d,
		"resetServingEndpoint",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) ResetSqlWarehouse() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlWarehouse",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) ResetUcSecurable() {
	_jsii_.InvokeVoid(
		d,
		"resetUcSecurable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAppsAppResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

