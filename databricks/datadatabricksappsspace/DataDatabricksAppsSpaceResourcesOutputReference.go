// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappsspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/datadatabricksappsspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAppsSpaceResourcesOutputReference interface {
	cdktn.ComplexObject
	App() DataDatabricksAppsSpaceResourcesAppOutputReference
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
	Database() DataDatabricksAppsSpaceResourcesDatabaseOutputReference
	DatabaseInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Experiment() DataDatabricksAppsSpaceResourcesExperimentOutputReference
	ExperimentInput() interface{}
	// Experimental.
	Fqn() *string
	GenieSpace() DataDatabricksAppsSpaceResourcesGenieSpaceOutputReference
	GenieSpaceInput() interface{}
	InternalValue() *DataDatabricksAppsSpaceResources
	SetInternalValue(val *DataDatabricksAppsSpaceResources)
	Job() DataDatabricksAppsSpaceResourcesJobOutputReference
	JobInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Secret() DataDatabricksAppsSpaceResourcesSecretOutputReference
	SecretInput() interface{}
	ServingEndpoint() DataDatabricksAppsSpaceResourcesServingEndpointOutputReference
	ServingEndpointInput() interface{}
	SqlWarehouse() DataDatabricksAppsSpaceResourcesSqlWarehouseOutputReference
	SqlWarehouseInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UcSecurable() DataDatabricksAppsSpaceResourcesUcSecurableOutputReference
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
	PutApp(value *DataDatabricksAppsSpaceResourcesApp)
	PutDatabase(value *DataDatabricksAppsSpaceResourcesDatabase)
	PutExperiment(value *DataDatabricksAppsSpaceResourcesExperiment)
	PutGenieSpace(value *DataDatabricksAppsSpaceResourcesGenieSpace)
	PutJob(value *DataDatabricksAppsSpaceResourcesJob)
	PutSecret(value *DataDatabricksAppsSpaceResourcesSecret)
	PutServingEndpoint(value *DataDatabricksAppsSpaceResourcesServingEndpoint)
	PutSqlWarehouse(value *DataDatabricksAppsSpaceResourcesSqlWarehouse)
	PutUcSecurable(value *DataDatabricksAppsSpaceResourcesUcSecurable)
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

// The jsii proxy struct for DataDatabricksAppsSpaceResourcesOutputReference
type jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) App() DataDatabricksAppsSpaceResourcesAppOutputReference {
	var returns DataDatabricksAppsSpaceResourcesAppOutputReference
	_jsii_.Get(
		j,
		"app",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) AppInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) Database() DataDatabricksAppsSpaceResourcesDatabaseOutputReference {
	var returns DataDatabricksAppsSpaceResourcesDatabaseOutputReference
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) DatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) Experiment() DataDatabricksAppsSpaceResourcesExperimentOutputReference {
	var returns DataDatabricksAppsSpaceResourcesExperimentOutputReference
	_jsii_.Get(
		j,
		"experiment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) ExperimentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"experimentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) GenieSpace() DataDatabricksAppsSpaceResourcesGenieSpaceOutputReference {
	var returns DataDatabricksAppsSpaceResourcesGenieSpaceOutputReference
	_jsii_.Get(
		j,
		"genieSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) GenieSpaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"genieSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) InternalValue() *DataDatabricksAppsSpaceResources {
	var returns *DataDatabricksAppsSpaceResources
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) Job() DataDatabricksAppsSpaceResourcesJobOutputReference {
	var returns DataDatabricksAppsSpaceResourcesJobOutputReference
	_jsii_.Get(
		j,
		"job",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) JobInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) Secret() DataDatabricksAppsSpaceResourcesSecretOutputReference {
	var returns DataDatabricksAppsSpaceResourcesSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) ServingEndpoint() DataDatabricksAppsSpaceResourcesServingEndpointOutputReference {
	var returns DataDatabricksAppsSpaceResourcesServingEndpointOutputReference
	_jsii_.Get(
		j,
		"servingEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) ServingEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"servingEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) SqlWarehouse() DataDatabricksAppsSpaceResourcesSqlWarehouseOutputReference {
	var returns DataDatabricksAppsSpaceResourcesSqlWarehouseOutputReference
	_jsii_.Get(
		j,
		"sqlWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) SqlWarehouseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlWarehouseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) UcSecurable() DataDatabricksAppsSpaceResourcesUcSecurableOutputReference {
	var returns DataDatabricksAppsSpaceResourcesUcSecurableOutputReference
	_jsii_.Get(
		j,
		"ucSecurable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) UcSecurableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ucSecurableInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksAppsSpaceResourcesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksAppsSpaceResourcesOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAppsSpaceResourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAppsSpace.DataDatabricksAppsSpaceResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksAppsSpaceResourcesOutputReference_Override(d DataDatabricksAppsSpaceResourcesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAppsSpace.DataDatabricksAppsSpaceResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference)SetInternalValue(val *DataDatabricksAppsSpaceResources) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) PutApp(value *DataDatabricksAppsSpaceResourcesApp) {
	if err := d.validatePutAppParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApp",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) PutDatabase(value *DataDatabricksAppsSpaceResourcesDatabase) {
	if err := d.validatePutDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDatabase",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) PutExperiment(value *DataDatabricksAppsSpaceResourcesExperiment) {
	if err := d.validatePutExperimentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExperiment",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) PutGenieSpace(value *DataDatabricksAppsSpaceResourcesGenieSpace) {
	if err := d.validatePutGenieSpaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGenieSpace",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) PutJob(value *DataDatabricksAppsSpaceResourcesJob) {
	if err := d.validatePutJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJob",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) PutSecret(value *DataDatabricksAppsSpaceResourcesSecret) {
	if err := d.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecret",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) PutServingEndpoint(value *DataDatabricksAppsSpaceResourcesServingEndpoint) {
	if err := d.validatePutServingEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServingEndpoint",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) PutSqlWarehouse(value *DataDatabricksAppsSpaceResourcesSqlWarehouse) {
	if err := d.validatePutSqlWarehouseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSqlWarehouse",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) PutUcSecurable(value *DataDatabricksAppsSpaceResourcesUcSecurable) {
	if err := d.validatePutUcSecurableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUcSecurable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) ResetApp() {
	_jsii_.InvokeVoid(
		d,
		"resetApp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) ResetExperiment() {
	_jsii_.InvokeVoid(
		d,
		"resetExperiment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) ResetGenieSpace() {
	_jsii_.InvokeVoid(
		d,
		"resetGenieSpace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) ResetJob() {
	_jsii_.InvokeVoid(
		d,
		"resetJob",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		d,
		"resetSecret",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) ResetServingEndpoint() {
	_jsii_.InvokeVoid(
		d,
		"resetServingEndpoint",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) ResetSqlWarehouse() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlWarehouse",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) ResetUcSecurable() {
	_jsii_.InvokeVoid(
		d,
		"resetUcSecurable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAppsSpaceResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

