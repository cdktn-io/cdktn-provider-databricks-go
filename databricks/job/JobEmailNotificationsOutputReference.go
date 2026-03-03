// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/job/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JobEmailNotificationsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *JobEmailNotifications
	SetInternalValue(val *JobEmailNotifications)
	NoAlertForSkippedRuns() interface{}
	SetNoAlertForSkippedRuns(val interface{})
	NoAlertForSkippedRunsInput() interface{}
	OnDurationWarningThresholdExceeded() *[]*string
	SetOnDurationWarningThresholdExceeded(val *[]*string)
	OnDurationWarningThresholdExceededInput() *[]*string
	OnFailure() *[]*string
	SetOnFailure(val *[]*string)
	OnFailureInput() *[]*string
	OnStart() *[]*string
	SetOnStart(val *[]*string)
	OnStartInput() *[]*string
	OnStreamingBacklogExceeded() *[]*string
	SetOnStreamingBacklogExceeded(val *[]*string)
	OnStreamingBacklogExceededInput() *[]*string
	OnSuccess() *[]*string
	SetOnSuccess(val *[]*string)
	OnSuccessInput() *[]*string
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
	ResetNoAlertForSkippedRuns()
	ResetOnDurationWarningThresholdExceeded()
	ResetOnFailure()
	ResetOnStart()
	ResetOnStreamingBacklogExceeded()
	ResetOnSuccess()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobEmailNotificationsOutputReference
type jsiiProxy_JobEmailNotificationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) InternalValue() *JobEmailNotifications {
	var returns *JobEmailNotifications
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) NoAlertForSkippedRuns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAlertForSkippedRuns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) NoAlertForSkippedRunsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAlertForSkippedRunsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) OnDurationWarningThresholdExceeded() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onDurationWarningThresholdExceeded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) OnDurationWarningThresholdExceededInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onDurationWarningThresholdExceededInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) OnFailure() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) OnFailureInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) OnStart() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) OnStartInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) OnStreamingBacklogExceeded() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onStreamingBacklogExceeded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) OnStreamingBacklogExceededInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onStreamingBacklogExceededInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) OnSuccess() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) OnSuccessInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onSuccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobEmailNotificationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) JobEmailNotificationsOutputReference {
	_init_.Initialize()

	if err := validateNewJobEmailNotificationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobEmailNotificationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobEmailNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobEmailNotificationsOutputReference_Override(j JobEmailNotificationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobEmailNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference)SetInternalValue(val *JobEmailNotifications) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference)SetNoAlertForSkippedRuns(val interface{}) {
	if err := j.validateSetNoAlertForSkippedRunsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noAlertForSkippedRuns",
		val,
	)
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference)SetOnDurationWarningThresholdExceeded(val *[]*string) {
	if err := j.validateSetOnDurationWarningThresholdExceededParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDurationWarningThresholdExceeded",
		val,
	)
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference)SetOnFailure(val *[]*string) {
	if err := j.validateSetOnFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onFailure",
		val,
	)
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference)SetOnStart(val *[]*string) {
	if err := j.validateSetOnStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onStart",
		val,
	)
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference)SetOnStreamingBacklogExceeded(val *[]*string) {
	if err := j.validateSetOnStreamingBacklogExceededParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onStreamingBacklogExceeded",
		val,
	)
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference)SetOnSuccess(val *[]*string) {
	if err := j.validateSetOnSuccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onSuccess",
		val,
	)
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobEmailNotificationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobEmailNotificationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobEmailNotificationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobEmailNotificationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobEmailNotificationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobEmailNotificationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobEmailNotificationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobEmailNotificationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobEmailNotificationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobEmailNotificationsOutputReference) ResetNoAlertForSkippedRuns() {
	_jsii_.InvokeVoid(
		j,
		"resetNoAlertForSkippedRuns",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) ResetOnDurationWarningThresholdExceeded() {
	_jsii_.InvokeVoid(
		j,
		"resetOnDurationWarningThresholdExceeded",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) ResetOnFailure() {
	_jsii_.InvokeVoid(
		j,
		"resetOnFailure",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) ResetOnStart() {
	_jsii_.InvokeVoid(
		j,
		"resetOnStart",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) ResetOnStreamingBacklogExceeded() {
	_jsii_.InvokeVoid(
		j,
		"resetOnStreamingBacklogExceeded",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) ResetOnSuccess() {
	_jsii_.InvokeVoid(
		j,
		"resetOnSuccess",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobEmailNotificationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobEmailNotificationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

