// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappsspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/datadatabricksappsspaces/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAppsSpacesSpacesResourcesOutputReference interface {
	cdktn.ComplexObject
	App() DataDatabricksAppsSpacesSpacesResourcesAppOutputReference
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
	Database() DataDatabricksAppsSpacesSpacesResourcesDatabaseOutputReference
	DatabaseInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Experiment() DataDatabricksAppsSpacesSpacesResourcesExperimentOutputReference
	ExperimentInput() interface{}
	// Experimental.
	Fqn() *string
	GenieSpace() DataDatabricksAppsSpacesSpacesResourcesGenieSpaceOutputReference
	GenieSpaceInput() interface{}
	InternalValue() *DataDatabricksAppsSpacesSpacesResources
	SetInternalValue(val *DataDatabricksAppsSpacesSpacesResources)
	Job() DataDatabricksAppsSpacesSpacesResourcesJobOutputReference
	JobInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Secret() DataDatabricksAppsSpacesSpacesResourcesSecretOutputReference
	SecretInput() interface{}
	ServingEndpoint() DataDatabricksAppsSpacesSpacesResourcesServingEndpointOutputReference
	ServingEndpointInput() interface{}
	SqlWarehouse() DataDatabricksAppsSpacesSpacesResourcesSqlWarehouseOutputReference
	SqlWarehouseInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UcSecurable() DataDatabricksAppsSpacesSpacesResourcesUcSecurableOutputReference
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
	PutApp(value *DataDatabricksAppsSpacesSpacesResourcesApp)
	PutDatabase(value *DataDatabricksAppsSpacesSpacesResourcesDatabase)
	PutExperiment(value *DataDatabricksAppsSpacesSpacesResourcesExperiment)
	PutGenieSpace(value *DataDatabricksAppsSpacesSpacesResourcesGenieSpace)
	PutJob(value *DataDatabricksAppsSpacesSpacesResourcesJob)
	PutSecret(value *DataDatabricksAppsSpacesSpacesResourcesSecret)
	PutServingEndpoint(value *DataDatabricksAppsSpacesSpacesResourcesServingEndpoint)
	PutSqlWarehouse(value *DataDatabricksAppsSpacesSpacesResourcesSqlWarehouse)
	PutUcSecurable(value *DataDatabricksAppsSpacesSpacesResourcesUcSecurable)
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

// The jsii proxy struct for DataDatabricksAppsSpacesSpacesResourcesOutputReference
type jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) App() DataDatabricksAppsSpacesSpacesResourcesAppOutputReference {
	var returns DataDatabricksAppsSpacesSpacesResourcesAppOutputReference
	_jsii_.Get(
		j,
		"app",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) AppInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) Database() DataDatabricksAppsSpacesSpacesResourcesDatabaseOutputReference {
	var returns DataDatabricksAppsSpacesSpacesResourcesDatabaseOutputReference
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) DatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) Experiment() DataDatabricksAppsSpacesSpacesResourcesExperimentOutputReference {
	var returns DataDatabricksAppsSpacesSpacesResourcesExperimentOutputReference
	_jsii_.Get(
		j,
		"experiment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) ExperimentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"experimentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) GenieSpace() DataDatabricksAppsSpacesSpacesResourcesGenieSpaceOutputReference {
	var returns DataDatabricksAppsSpacesSpacesResourcesGenieSpaceOutputReference
	_jsii_.Get(
		j,
		"genieSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) GenieSpaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"genieSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) InternalValue() *DataDatabricksAppsSpacesSpacesResources {
	var returns *DataDatabricksAppsSpacesSpacesResources
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) Job() DataDatabricksAppsSpacesSpacesResourcesJobOutputReference {
	var returns DataDatabricksAppsSpacesSpacesResourcesJobOutputReference
	_jsii_.Get(
		j,
		"job",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) JobInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) Secret() DataDatabricksAppsSpacesSpacesResourcesSecretOutputReference {
	var returns DataDatabricksAppsSpacesSpacesResourcesSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) ServingEndpoint() DataDatabricksAppsSpacesSpacesResourcesServingEndpointOutputReference {
	var returns DataDatabricksAppsSpacesSpacesResourcesServingEndpointOutputReference
	_jsii_.Get(
		j,
		"servingEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) ServingEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"servingEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) SqlWarehouse() DataDatabricksAppsSpacesSpacesResourcesSqlWarehouseOutputReference {
	var returns DataDatabricksAppsSpacesSpacesResourcesSqlWarehouseOutputReference
	_jsii_.Get(
		j,
		"sqlWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) SqlWarehouseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlWarehouseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) UcSecurable() DataDatabricksAppsSpacesSpacesResourcesUcSecurableOutputReference {
	var returns DataDatabricksAppsSpacesSpacesResourcesUcSecurableOutputReference
	_jsii_.Get(
		j,
		"ucSecurable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) UcSecurableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ucSecurableInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksAppsSpacesSpacesResourcesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksAppsSpacesSpacesResourcesOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAppsSpacesSpacesResourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAppsSpaces.DataDatabricksAppsSpacesSpacesResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksAppsSpacesSpacesResourcesOutputReference_Override(d DataDatabricksAppsSpacesSpacesResourcesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAppsSpaces.DataDatabricksAppsSpacesSpacesResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference)SetInternalValue(val *DataDatabricksAppsSpacesSpacesResources) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) PutApp(value *DataDatabricksAppsSpacesSpacesResourcesApp) {
	if err := d.validatePutAppParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApp",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) PutDatabase(value *DataDatabricksAppsSpacesSpacesResourcesDatabase) {
	if err := d.validatePutDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDatabase",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) PutExperiment(value *DataDatabricksAppsSpacesSpacesResourcesExperiment) {
	if err := d.validatePutExperimentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExperiment",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) PutGenieSpace(value *DataDatabricksAppsSpacesSpacesResourcesGenieSpace) {
	if err := d.validatePutGenieSpaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGenieSpace",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) PutJob(value *DataDatabricksAppsSpacesSpacesResourcesJob) {
	if err := d.validatePutJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJob",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) PutSecret(value *DataDatabricksAppsSpacesSpacesResourcesSecret) {
	if err := d.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecret",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) PutServingEndpoint(value *DataDatabricksAppsSpacesSpacesResourcesServingEndpoint) {
	if err := d.validatePutServingEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServingEndpoint",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) PutSqlWarehouse(value *DataDatabricksAppsSpacesSpacesResourcesSqlWarehouse) {
	if err := d.validatePutSqlWarehouseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSqlWarehouse",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) PutUcSecurable(value *DataDatabricksAppsSpacesSpacesResourcesUcSecurable) {
	if err := d.validatePutUcSecurableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUcSecurable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) ResetApp() {
	_jsii_.InvokeVoid(
		d,
		"resetApp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) ResetExperiment() {
	_jsii_.InvokeVoid(
		d,
		"resetExperiment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) ResetGenieSpace() {
	_jsii_.InvokeVoid(
		d,
		"resetGenieSpace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) ResetJob() {
	_jsii_.InvokeVoid(
		d,
		"resetJob",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		d,
		"resetSecret",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) ResetServingEndpoint() {
	_jsii_.InvokeVoid(
		d,
		"resetServingEndpoint",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) ResetSqlWarehouse() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlWarehouse",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) ResetUcSecurable() {
	_jsii_.InvokeVoid(
		d,
		"resetUcSecurable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAppsSpacesSpacesResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

