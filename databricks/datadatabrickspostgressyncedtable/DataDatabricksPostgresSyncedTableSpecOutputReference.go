// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgressyncedtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/datadatabrickspostgressyncedtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksPostgresSyncedTableSpecOutputReference interface {
	cdktn.ComplexObject
	Branch() *string
	SetBranch(val *string)
	BranchInput() *string
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
	CreateDatabaseObjectsIfMissing() interface{}
	SetCreateDatabaseObjectsIfMissing(val interface{})
	CreateDatabaseObjectsIfMissingInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExistingPipelineId() *string
	SetExistingPipelineId(val *string)
	ExistingPipelineIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksPostgresSyncedTableSpec
	SetInternalValue(val *DataDatabricksPostgresSyncedTableSpec)
	NewPipelineSpec() DataDatabricksPostgresSyncedTableSpecNewPipelineSpecOutputReference
	NewPipelineSpecInput() interface{}
	PostgresDatabase() *string
	SetPostgresDatabase(val *string)
	PostgresDatabaseInput() *string
	PrimaryKeyColumns() *[]*string
	SetPrimaryKeyColumns(val *[]*string)
	PrimaryKeyColumnsInput() *[]*string
	SchedulingPolicy() *string
	SetSchedulingPolicy(val *string)
	SchedulingPolicyInput() *string
	SourceTableFullName() *string
	SetSourceTableFullName(val *string)
	SourceTableFullNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeseriesKey() *string
	SetTimeseriesKey(val *string)
	TimeseriesKeyInput() *string
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
	PutNewPipelineSpec(value *DataDatabricksPostgresSyncedTableSpecNewPipelineSpec)
	ResetBranch()
	ResetCreateDatabaseObjectsIfMissing()
	ResetExistingPipelineId()
	ResetNewPipelineSpec()
	ResetPostgresDatabase()
	ResetPrimaryKeyColumns()
	ResetSchedulingPolicy()
	ResetSourceTableFullName()
	ResetTimeseriesKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksPostgresSyncedTableSpecOutputReference
type jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) Branch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) BranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) CreateDatabaseObjectsIfMissing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDatabaseObjectsIfMissing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) CreateDatabaseObjectsIfMissingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDatabaseObjectsIfMissingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) ExistingPipelineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingPipelineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) ExistingPipelineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingPipelineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) InternalValue() *DataDatabricksPostgresSyncedTableSpec {
	var returns *DataDatabricksPostgresSyncedTableSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) NewPipelineSpec() DataDatabricksPostgresSyncedTableSpecNewPipelineSpecOutputReference {
	var returns DataDatabricksPostgresSyncedTableSpecNewPipelineSpecOutputReference
	_jsii_.Get(
		j,
		"newPipelineSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) NewPipelineSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newPipelineSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) PostgresDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) PostgresDatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) PrimaryKeyColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"primaryKeyColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) PrimaryKeyColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"primaryKeyColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) SchedulingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) SchedulingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) SourceTableFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTableFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) SourceTableFullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTableFullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) TimeseriesKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeseriesKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) TimeseriesKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeseriesKeyInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksPostgresSyncedTableSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksPostgresSyncedTableSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksPostgresSyncedTableSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresSyncedTable.DataDatabricksPostgresSyncedTableSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksPostgresSyncedTableSpecOutputReference_Override(d DataDatabricksPostgresSyncedTableSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresSyncedTable.DataDatabricksPostgresSyncedTableSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference)SetBranch(val *string) {
	if err := j.validateSetBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branch",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference)SetCreateDatabaseObjectsIfMissing(val interface{}) {
	if err := j.validateSetCreateDatabaseObjectsIfMissingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDatabaseObjectsIfMissing",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference)SetExistingPipelineId(val *string) {
	if err := j.validateSetExistingPipelineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"existingPipelineId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference)SetInternalValue(val *DataDatabricksPostgresSyncedTableSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference)SetPostgresDatabase(val *string) {
	if err := j.validateSetPostgresDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postgresDatabase",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference)SetPrimaryKeyColumns(val *[]*string) {
	if err := j.validateSetPrimaryKeyColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryKeyColumns",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference)SetSchedulingPolicy(val *string) {
	if err := j.validateSetSchedulingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulingPolicy",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference)SetSourceTableFullName(val *string) {
	if err := j.validateSetSourceTableFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceTableFullName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference)SetTimeseriesKey(val *string) {
	if err := j.validateSetTimeseriesKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeseriesKey",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) PutNewPipelineSpec(value *DataDatabricksPostgresSyncedTableSpecNewPipelineSpec) {
	if err := d.validatePutNewPipelineSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNewPipelineSpec",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) ResetBranch() {
	_jsii_.InvokeVoid(
		d,
		"resetBranch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) ResetCreateDatabaseObjectsIfMissing() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateDatabaseObjectsIfMissing",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) ResetExistingPipelineId() {
	_jsii_.InvokeVoid(
		d,
		"resetExistingPipelineId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) ResetNewPipelineSpec() {
	_jsii_.InvokeVoid(
		d,
		"resetNewPipelineSpec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) ResetPostgresDatabase() {
	_jsii_.InvokeVoid(
		d,
		"resetPostgresDatabase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) ResetPrimaryKeyColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetPrimaryKeyColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) ResetSchedulingPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetSchedulingPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) ResetSourceTableFullName() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceTableFullName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) ResetTimeseriesKey() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeseriesKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

