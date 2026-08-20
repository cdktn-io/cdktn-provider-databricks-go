// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/job/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JobTriggersOutputReference interface {
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
	Continuous() JobTriggersContinuousOutputReference
	ContinuousInput() *JobTriggersContinuous
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FileArrival() JobTriggersFileArrivalOutputReference
	FileArrivalInput() *JobTriggersFileArrival
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Model() JobTriggersModelOutputReference
	ModelInput() *JobTriggersModel
	PauseStatus() *string
	SetPauseStatus(val *string)
	PauseStatusInput() *string
	Periodic() JobTriggersPeriodicOutputReference
	PeriodicInput() *JobTriggersPeriodic
	Schedule() JobTriggersScheduleOutputReference
	ScheduleInput() *JobTriggersSchedule
	SqlCondition() JobTriggersSqlConditionOutputReference
	SqlConditionInput() *JobTriggersSqlCondition
	TableUpdate() JobTriggersTableUpdateOutputReference
	TableUpdateInput() *JobTriggersTableUpdate
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
	PutContinuous(value *JobTriggersContinuous)
	PutFileArrival(value *JobTriggersFileArrival)
	PutModel(value *JobTriggersModel)
	PutPeriodic(value *JobTriggersPeriodic)
	PutSchedule(value *JobTriggersSchedule)
	PutSqlCondition(value *JobTriggersSqlCondition)
	PutTableUpdate(value *JobTriggersTableUpdate)
	ResetContinuous()
	ResetFileArrival()
	ResetModel()
	ResetPauseStatus()
	ResetPeriodic()
	ResetSchedule()
	ResetSqlCondition()
	ResetTableUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTriggersOutputReference
type jsiiProxy_JobTriggersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_JobTriggersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) Continuous() JobTriggersContinuousOutputReference {
	var returns JobTriggersContinuousOutputReference
	_jsii_.Get(
		j,
		"continuous",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) ContinuousInput() *JobTriggersContinuous {
	var returns *JobTriggersContinuous
	_jsii_.Get(
		j,
		"continuousInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) FileArrival() JobTriggersFileArrivalOutputReference {
	var returns JobTriggersFileArrivalOutputReference
	_jsii_.Get(
		j,
		"fileArrival",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) FileArrivalInput() *JobTriggersFileArrival {
	var returns *JobTriggersFileArrival
	_jsii_.Get(
		j,
		"fileArrivalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) Model() JobTriggersModelOutputReference {
	var returns JobTriggersModelOutputReference
	_jsii_.Get(
		j,
		"model",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) ModelInput() *JobTriggersModel {
	var returns *JobTriggersModel
	_jsii_.Get(
		j,
		"modelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) PauseStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pauseStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) PauseStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pauseStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) Periodic() JobTriggersPeriodicOutputReference {
	var returns JobTriggersPeriodicOutputReference
	_jsii_.Get(
		j,
		"periodic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) PeriodicInput() *JobTriggersPeriodic {
	var returns *JobTriggersPeriodic
	_jsii_.Get(
		j,
		"periodicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) Schedule() JobTriggersScheduleOutputReference {
	var returns JobTriggersScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) ScheduleInput() *JobTriggersSchedule {
	var returns *JobTriggersSchedule
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) SqlCondition() JobTriggersSqlConditionOutputReference {
	var returns JobTriggersSqlConditionOutputReference
	_jsii_.Get(
		j,
		"sqlCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) SqlConditionInput() *JobTriggersSqlCondition {
	var returns *JobTriggersSqlCondition
	_jsii_.Get(
		j,
		"sqlConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) TableUpdate() JobTriggersTableUpdateOutputReference {
	var returns JobTriggersTableUpdateOutputReference
	_jsii_.Get(
		j,
		"tableUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) TableUpdateInput() *JobTriggersTableUpdate {
	var returns *JobTriggersTableUpdate
	_jsii_.Get(
		j,
		"tableUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobTriggersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JobTriggersOutputReference {
	_init_.Initialize()

	if err := validateNewJobTriggersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTriggersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTriggersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJobTriggersOutputReference_Override(j JobTriggersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTriggersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JobTriggersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTriggersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTriggersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTriggersOutputReference)SetPauseStatus(val *string) {
	if err := j.validateSetPauseStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pauseStatus",
		val,
	)
}

func (j *jsiiProxy_JobTriggersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTriggersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTriggersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) PutContinuous(value *JobTriggersContinuous) {
	if err := j.validatePutContinuousParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putContinuous",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTriggersOutputReference) PutFileArrival(value *JobTriggersFileArrival) {
	if err := j.validatePutFileArrivalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putFileArrival",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTriggersOutputReference) PutModel(value *JobTriggersModel) {
	if err := j.validatePutModelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putModel",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTriggersOutputReference) PutPeriodic(value *JobTriggersPeriodic) {
	if err := j.validatePutPeriodicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPeriodic",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTriggersOutputReference) PutSchedule(value *JobTriggersSchedule) {
	if err := j.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSchedule",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTriggersOutputReference) PutSqlCondition(value *JobTriggersSqlCondition) {
	if err := j.validatePutSqlConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSqlCondition",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTriggersOutputReference) PutTableUpdate(value *JobTriggersTableUpdate) {
	if err := j.validatePutTableUpdateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putTableUpdate",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTriggersOutputReference) ResetContinuous() {
	_jsii_.InvokeVoid(
		j,
		"resetContinuous",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggersOutputReference) ResetFileArrival() {
	_jsii_.InvokeVoid(
		j,
		"resetFileArrival",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggersOutputReference) ResetModel() {
	_jsii_.InvokeVoid(
		j,
		"resetModel",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggersOutputReference) ResetPauseStatus() {
	_jsii_.InvokeVoid(
		j,
		"resetPauseStatus",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggersOutputReference) ResetPeriodic() {
	_jsii_.InvokeVoid(
		j,
		"resetPeriodic",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggersOutputReference) ResetSchedule() {
	_jsii_.InvokeVoid(
		j,
		"resetSchedule",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggersOutputReference) ResetSqlCondition() {
	_jsii_.InvokeVoid(
		j,
		"resetSqlCondition",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggersOutputReference) ResetTableUpdate() {
	_jsii_.InvokeVoid(
		j,
		"resetTableUpdate",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := j.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

