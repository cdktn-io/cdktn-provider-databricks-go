// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabricksapps/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAppsAppEffectiveResourcesOutputReference interface {
	cdktn.ComplexObject
	App() DataDatabricksAppsAppEffectiveResourcesAppOutputReference
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
	Database() DataDatabricksAppsAppEffectiveResourcesDatabaseOutputReference
	DatabaseInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Experiment() DataDatabricksAppsAppEffectiveResourcesExperimentOutputReference
	ExperimentInput() interface{}
	// Experimental.
	Fqn() *string
	GenieSpace() DataDatabricksAppsAppEffectiveResourcesGenieSpaceOutputReference
	GenieSpaceInput() interface{}
	InternalValue() *DataDatabricksAppsAppEffectiveResources
	SetInternalValue(val *DataDatabricksAppsAppEffectiveResources)
	Job() DataDatabricksAppsAppEffectiveResourcesJobOutputReference
	JobInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Postgres() DataDatabricksAppsAppEffectiveResourcesPostgresOutputReference
	PostgresInput() interface{}
	Secret() DataDatabricksAppsAppEffectiveResourcesSecretOutputReference
	SecretInput() interface{}
	ServingEndpoint() DataDatabricksAppsAppEffectiveResourcesServingEndpointOutputReference
	ServingEndpointInput() interface{}
	SqlWarehouse() DataDatabricksAppsAppEffectiveResourcesSqlWarehouseOutputReference
	SqlWarehouseInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UcSecurable() DataDatabricksAppsAppEffectiveResourcesUcSecurableOutputReference
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
	PutApp(value *DataDatabricksAppsAppEffectiveResourcesApp)
	PutDatabase(value *DataDatabricksAppsAppEffectiveResourcesDatabase)
	PutExperiment(value *DataDatabricksAppsAppEffectiveResourcesExperiment)
	PutGenieSpace(value *DataDatabricksAppsAppEffectiveResourcesGenieSpace)
	PutJob(value *DataDatabricksAppsAppEffectiveResourcesJob)
	PutPostgres(value *DataDatabricksAppsAppEffectiveResourcesPostgres)
	PutSecret(value *DataDatabricksAppsAppEffectiveResourcesSecret)
	PutServingEndpoint(value *DataDatabricksAppsAppEffectiveResourcesServingEndpoint)
	PutSqlWarehouse(value *DataDatabricksAppsAppEffectiveResourcesSqlWarehouse)
	PutUcSecurable(value *DataDatabricksAppsAppEffectiveResourcesUcSecurable)
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

// The jsii proxy struct for DataDatabricksAppsAppEffectiveResourcesOutputReference
type jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) App() DataDatabricksAppsAppEffectiveResourcesAppOutputReference {
	var returns DataDatabricksAppsAppEffectiveResourcesAppOutputReference
	_jsii_.Get(
		j,
		"app",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) AppInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) Database() DataDatabricksAppsAppEffectiveResourcesDatabaseOutputReference {
	var returns DataDatabricksAppsAppEffectiveResourcesDatabaseOutputReference
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) DatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) Experiment() DataDatabricksAppsAppEffectiveResourcesExperimentOutputReference {
	var returns DataDatabricksAppsAppEffectiveResourcesExperimentOutputReference
	_jsii_.Get(
		j,
		"experiment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) ExperimentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"experimentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) GenieSpace() DataDatabricksAppsAppEffectiveResourcesGenieSpaceOutputReference {
	var returns DataDatabricksAppsAppEffectiveResourcesGenieSpaceOutputReference
	_jsii_.Get(
		j,
		"genieSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) GenieSpaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"genieSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) InternalValue() *DataDatabricksAppsAppEffectiveResources {
	var returns *DataDatabricksAppsAppEffectiveResources
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) Job() DataDatabricksAppsAppEffectiveResourcesJobOutputReference {
	var returns DataDatabricksAppsAppEffectiveResourcesJobOutputReference
	_jsii_.Get(
		j,
		"job",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) JobInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) Postgres() DataDatabricksAppsAppEffectiveResourcesPostgresOutputReference {
	var returns DataDatabricksAppsAppEffectiveResourcesPostgresOutputReference
	_jsii_.Get(
		j,
		"postgres",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) PostgresInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postgresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) Secret() DataDatabricksAppsAppEffectiveResourcesSecretOutputReference {
	var returns DataDatabricksAppsAppEffectiveResourcesSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) ServingEndpoint() DataDatabricksAppsAppEffectiveResourcesServingEndpointOutputReference {
	var returns DataDatabricksAppsAppEffectiveResourcesServingEndpointOutputReference
	_jsii_.Get(
		j,
		"servingEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) ServingEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"servingEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) SqlWarehouse() DataDatabricksAppsAppEffectiveResourcesSqlWarehouseOutputReference {
	var returns DataDatabricksAppsAppEffectiveResourcesSqlWarehouseOutputReference
	_jsii_.Get(
		j,
		"sqlWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) SqlWarehouseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlWarehouseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) UcSecurable() DataDatabricksAppsAppEffectiveResourcesUcSecurableOutputReference {
	var returns DataDatabricksAppsAppEffectiveResourcesUcSecurableOutputReference
	_jsii_.Get(
		j,
		"ucSecurable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) UcSecurableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ucSecurableInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksAppsAppEffectiveResourcesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksAppsAppEffectiveResourcesOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAppsAppEffectiveResourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksApps.DataDatabricksAppsAppEffectiveResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksAppsAppEffectiveResourcesOutputReference_Override(d DataDatabricksAppsAppEffectiveResourcesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksApps.DataDatabricksAppsAppEffectiveResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference)SetInternalValue(val *DataDatabricksAppsAppEffectiveResources) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) PutApp(value *DataDatabricksAppsAppEffectiveResourcesApp) {
	if err := d.validatePutAppParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApp",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) PutDatabase(value *DataDatabricksAppsAppEffectiveResourcesDatabase) {
	if err := d.validatePutDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDatabase",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) PutExperiment(value *DataDatabricksAppsAppEffectiveResourcesExperiment) {
	if err := d.validatePutExperimentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExperiment",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) PutGenieSpace(value *DataDatabricksAppsAppEffectiveResourcesGenieSpace) {
	if err := d.validatePutGenieSpaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGenieSpace",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) PutJob(value *DataDatabricksAppsAppEffectiveResourcesJob) {
	if err := d.validatePutJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJob",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) PutPostgres(value *DataDatabricksAppsAppEffectiveResourcesPostgres) {
	if err := d.validatePutPostgresParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPostgres",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) PutSecret(value *DataDatabricksAppsAppEffectiveResourcesSecret) {
	if err := d.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecret",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) PutServingEndpoint(value *DataDatabricksAppsAppEffectiveResourcesServingEndpoint) {
	if err := d.validatePutServingEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServingEndpoint",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) PutSqlWarehouse(value *DataDatabricksAppsAppEffectiveResourcesSqlWarehouse) {
	if err := d.validatePutSqlWarehouseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSqlWarehouse",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) PutUcSecurable(value *DataDatabricksAppsAppEffectiveResourcesUcSecurable) {
	if err := d.validatePutUcSecurableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUcSecurable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) ResetApp() {
	_jsii_.InvokeVoid(
		d,
		"resetApp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) ResetExperiment() {
	_jsii_.InvokeVoid(
		d,
		"resetExperiment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) ResetGenieSpace() {
	_jsii_.InvokeVoid(
		d,
		"resetGenieSpace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) ResetJob() {
	_jsii_.InvokeVoid(
		d,
		"resetJob",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) ResetPostgres() {
	_jsii_.InvokeVoid(
		d,
		"resetPostgres",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		d,
		"resetSecret",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) ResetServingEndpoint() {
	_jsii_.InvokeVoid(
		d,
		"resetServingEndpoint",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) ResetSqlWarehouse() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlWarehouse",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) ResetUcSecurable() {
	_jsii_.InvokeVoid(
		d,
		"resetUcSecurable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAppsAppEffectiveResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

