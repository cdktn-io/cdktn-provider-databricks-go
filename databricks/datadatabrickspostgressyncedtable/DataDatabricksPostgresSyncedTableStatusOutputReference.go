// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgressyncedtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabrickspostgressyncedtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksPostgresSyncedTableStatusOutputReference interface {
	cdktn.ComplexObject
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
	DetailedState() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksPostgresSyncedTableStatus
	SetInternalValue(val *DataDatabricksPostgresSyncedTableStatus)
	LastProcessedCommitVersion() *float64
	LastSync() DataDatabricksPostgresSyncedTableStatusLastSyncOutputReference
	LastSyncTime() *string
	Message() *string
	OngoingSyncProgress() DataDatabricksPostgresSyncedTableStatusOngoingSyncProgressOutputReference
	PipelineId() *string
	Project() *string
	ProvisioningPhase() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UnityCatalogProvisioningState() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksPostgresSyncedTableStatusOutputReference
type jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) DetailedState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detailedState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) InternalValue() *DataDatabricksPostgresSyncedTableStatus {
	var returns *DataDatabricksPostgresSyncedTableStatus
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) LastProcessedCommitVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastProcessedCommitVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) LastSync() DataDatabricksPostgresSyncedTableStatusLastSyncOutputReference {
	var returns DataDatabricksPostgresSyncedTableStatusLastSyncOutputReference
	_jsii_.Get(
		j,
		"lastSync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) LastSyncTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastSyncTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) Message() *string {
	var returns *string
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) OngoingSyncProgress() DataDatabricksPostgresSyncedTableStatusOngoingSyncProgressOutputReference {
	var returns DataDatabricksPostgresSyncedTableStatusOngoingSyncProgressOutputReference
	_jsii_.Get(
		j,
		"ongoingSyncProgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) PipelineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) ProvisioningPhase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningPhase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) UnityCatalogProvisioningState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unityCatalogProvisioningState",
		&returns,
	)
	return returns
}


func NewDataDatabricksPostgresSyncedTableStatusOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksPostgresSyncedTableStatusOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksPostgresSyncedTableStatusOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresSyncedTable.DataDatabricksPostgresSyncedTableStatusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksPostgresSyncedTableStatusOutputReference_Override(d DataDatabricksPostgresSyncedTableStatusOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresSyncedTable.DataDatabricksPostgresSyncedTableStatusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference)SetInternalValue(val *DataDatabricksPostgresSyncedTableStatus) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresSyncedTableStatusOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

