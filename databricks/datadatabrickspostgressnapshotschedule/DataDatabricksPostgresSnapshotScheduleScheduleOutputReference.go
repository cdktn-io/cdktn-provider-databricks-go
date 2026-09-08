// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgressnapshotschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabrickspostgressnapshotschedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksPostgresSnapshotScheduleScheduleOutputReference interface {
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
	DailySchedule() DataDatabricksPostgresSnapshotScheduleScheduleDailyScheduleOutputReference
	DailyScheduleInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksPostgresSnapshotScheduleSchedule
	SetInternalValue(val *DataDatabricksPostgresSnapshotScheduleSchedule)
	MonthlySchedule() DataDatabricksPostgresSnapshotScheduleScheduleMonthlyScheduleOutputReference
	MonthlyScheduleInput() interface{}
	Retention() *string
	SetRetention(val *string)
	RetentionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WeeklySchedule() DataDatabricksPostgresSnapshotScheduleScheduleWeeklyScheduleOutputReference
	WeeklyScheduleInput() interface{}
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
	PutDailySchedule(value *DataDatabricksPostgresSnapshotScheduleScheduleDailySchedule)
	PutMonthlySchedule(value *DataDatabricksPostgresSnapshotScheduleScheduleMonthlySchedule)
	PutWeeklySchedule(value *DataDatabricksPostgresSnapshotScheduleScheduleWeeklySchedule)
	ResetDailySchedule()
	ResetMonthlySchedule()
	ResetWeeklySchedule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksPostgresSnapshotScheduleScheduleOutputReference
type jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) DailySchedule() DataDatabricksPostgresSnapshotScheduleScheduleDailyScheduleOutputReference {
	var returns DataDatabricksPostgresSnapshotScheduleScheduleDailyScheduleOutputReference
	_jsii_.Get(
		j,
		"dailySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) DailyScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dailyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) InternalValue() *DataDatabricksPostgresSnapshotScheduleSchedule {
	var returns *DataDatabricksPostgresSnapshotScheduleSchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) MonthlySchedule() DataDatabricksPostgresSnapshotScheduleScheduleMonthlyScheduleOutputReference {
	var returns DataDatabricksPostgresSnapshotScheduleScheduleMonthlyScheduleOutputReference
	_jsii_.Get(
		j,
		"monthlySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) MonthlyScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monthlyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) Retention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) RetentionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) WeeklySchedule() DataDatabricksPostgresSnapshotScheduleScheduleWeeklyScheduleOutputReference {
	var returns DataDatabricksPostgresSnapshotScheduleScheduleWeeklyScheduleOutputReference
	_jsii_.Get(
		j,
		"weeklySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) WeeklyScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"weeklyScheduleInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksPostgresSnapshotScheduleScheduleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksPostgresSnapshotScheduleScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksPostgresSnapshotScheduleScheduleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresSnapshotSchedule.DataDatabricksPostgresSnapshotScheduleScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksPostgresSnapshotScheduleScheduleOutputReference_Override(d DataDatabricksPostgresSnapshotScheduleScheduleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresSnapshotSchedule.DataDatabricksPostgresSnapshotScheduleScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference)SetInternalValue(val *DataDatabricksPostgresSnapshotScheduleSchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference)SetRetention(val *string) {
	if err := j.validateSetRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) PutDailySchedule(value *DataDatabricksPostgresSnapshotScheduleScheduleDailySchedule) {
	if err := d.validatePutDailyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDailySchedule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) PutMonthlySchedule(value *DataDatabricksPostgresSnapshotScheduleScheduleMonthlySchedule) {
	if err := d.validatePutMonthlyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMonthlySchedule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) PutWeeklySchedule(value *DataDatabricksPostgresSnapshotScheduleScheduleWeeklySchedule) {
	if err := d.validatePutWeeklyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWeeklySchedule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) ResetDailySchedule() {
	_jsii_.InvokeVoid(
		d,
		"resetDailySchedule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) ResetMonthlySchedule() {
	_jsii_.InvokeVoid(
		d,
		"resetMonthlySchedule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) ResetWeeklySchedule() {
	_jsii_.InvokeVoid(
		d,
		"resetWeeklySchedule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresSnapshotScheduleScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

