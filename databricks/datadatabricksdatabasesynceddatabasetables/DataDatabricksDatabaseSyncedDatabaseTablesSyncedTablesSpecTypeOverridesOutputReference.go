// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasesynceddatabasetables

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksdatabasesynceddatabasetables/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference interface {
	cdktn.ComplexObject
	ColumnName() *string
	SetColumnName(val *string)
	ColumnNameInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PgType() *string
	SetPgType(val *string)
	PgTypeInput() *string
	Size() *float64
	SetSize(val *float64)
	SizeInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	ResetSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference
type jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) ColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) ColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) PgType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pgType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) PgTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pgTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) SizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference_Override(d DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference)SetColumnName(val *string) {
	if err := j.validateSetColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference)SetPgType(val *string) {
	if err := j.validateSetPgTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pgType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference)SetSize(val *float64) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) ResetSize() {
	_jsii_.InvokeVoid(
		d,
		"resetSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

