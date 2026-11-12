// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchdatasources

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabrickslakewatchdatasources/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksLakewatchDatasourcesDatasourcesOutputReference interface {
	cdktn.ComplexObject
	Annotations() cdktn.StringMap
	Autoloader() DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference
	CatalogName() *string
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
	ComputeMode() *string
	ComputeWorkspaceId() *string
	CreatedBy() DataDatabricksLakewatchDatasourcesDatasourcesCreatedByOutputReference
	CreateTime() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatasourceId() *string
	DatasourceName() *string
	Description() *string
	EffectiveCatalogName() *string
	EffectiveDatasourceName() *string
	EffectiveRunAsPrincipal() *string
	EffectiveSchemaName() *string
	Etag() *string
	ExistingTable() DataDatabricksLakewatchDatasourcesDatasourcesExistingTableOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksLakewatchDatasourcesDatasources
	SetInternalValue(val *DataDatabricksLakewatchDatasourcesDatasources)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Owner() DataDatabricksLakewatchDatasourcesDatasourcesOwnerOutputReference
	Parsers() DataDatabricksLakewatchDatasourcesDatasourcesParsersList
	ProviderConfig() DataDatabricksLakewatchDatasourcesDatasourcesProviderConfigOutputReference
	ProviderConfigInput() interface{}
	RunAsPrincipal() *string
	Schedule() DataDatabricksLakewatchDatasourcesDatasourcesScheduleOutputReference
	SchemaName() *string
	Status() DataDatabricksLakewatchDatasourcesDatasourcesStatusOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UnityCatalogConnection() DataDatabricksLakewatchDatasourcesDatasourcesUnityCatalogConnectionOutputReference
	UpdatedBy() DataDatabricksLakewatchDatasourcesDatasourcesUpdatedByOutputReference
	UpdateTime() *string
	Version() *float64
	WarehouseId() *string
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
	PutProviderConfig(value *DataDatabricksLakewatchDatasourcesDatasourcesProviderConfig)
	ResetProviderConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksLakewatchDatasourcesDatasourcesOutputReference
type jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) Annotations() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) Autoloader() DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference {
	var returns DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference
	_jsii_.Get(
		j,
		"autoloader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) CatalogName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) ComputeMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) ComputeWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) CreatedBy() DataDatabricksLakewatchDatasourcesDatasourcesCreatedByOutputReference {
	var returns DataDatabricksLakewatchDatasourcesDatasourcesCreatedByOutputReference
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) DatasourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) DatasourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) EffectiveCatalogName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveCatalogName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) EffectiveDatasourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveDatasourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) EffectiveRunAsPrincipal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveRunAsPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) EffectiveSchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveSchemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) ExistingTable() DataDatabricksLakewatchDatasourcesDatasourcesExistingTableOutputReference {
	var returns DataDatabricksLakewatchDatasourcesDatasourcesExistingTableOutputReference
	_jsii_.Get(
		j,
		"existingTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) InternalValue() *DataDatabricksLakewatchDatasourcesDatasources {
	var returns *DataDatabricksLakewatchDatasourcesDatasources
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) Owner() DataDatabricksLakewatchDatasourcesDatasourcesOwnerOutputReference {
	var returns DataDatabricksLakewatchDatasourcesDatasourcesOwnerOutputReference
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) Parsers() DataDatabricksLakewatchDatasourcesDatasourcesParsersList {
	var returns DataDatabricksLakewatchDatasourcesDatasourcesParsersList
	_jsii_.Get(
		j,
		"parsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) ProviderConfig() DataDatabricksLakewatchDatasourcesDatasourcesProviderConfigOutputReference {
	var returns DataDatabricksLakewatchDatasourcesDatasourcesProviderConfigOutputReference
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) RunAsPrincipal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) Schedule() DataDatabricksLakewatchDatasourcesDatasourcesScheduleOutputReference {
	var returns DataDatabricksLakewatchDatasourcesDatasourcesScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) SchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) Status() DataDatabricksLakewatchDatasourcesDatasourcesStatusOutputReference {
	var returns DataDatabricksLakewatchDatasourcesDatasourcesStatusOutputReference
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) UnityCatalogConnection() DataDatabricksLakewatchDatasourcesDatasourcesUnityCatalogConnectionOutputReference {
	var returns DataDatabricksLakewatchDatasourcesDatasourcesUnityCatalogConnectionOutputReference
	_jsii_.Get(
		j,
		"unityCatalogConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) UpdatedBy() DataDatabricksLakewatchDatasourcesDatasourcesUpdatedByOutputReference {
	var returns DataDatabricksLakewatchDatasourcesDatasourcesUpdatedByOutputReference
	_jsii_.Get(
		j,
		"updatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) WarehouseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseId",
		&returns,
	)
	return returns
}


func NewDataDatabricksLakewatchDatasourcesDatasourcesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksLakewatchDatasourcesDatasourcesOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksLakewatchDatasourcesDatasourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchDatasources.DataDatabricksLakewatchDatasourcesDatasourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksLakewatchDatasourcesDatasourcesOutputReference_Override(d DataDatabricksLakewatchDatasourcesDatasourcesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchDatasources.DataDatabricksLakewatchDatasourcesDatasourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference)SetInternalValue(val *DataDatabricksLakewatchDatasourcesDatasources) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) PutProviderConfig(value *DataDatabricksLakewatchDatasourcesDatasourcesProviderConfig) {
	if err := d.validatePutProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProviderConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

