// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/job/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JobWebhookNotificationsOutputReference interface {
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
	InternalValue() *JobWebhookNotifications
	SetInternalValue(val *JobWebhookNotifications)
	OnDurationWarningThresholdExceeded() JobWebhookNotificationsOnDurationWarningThresholdExceededList
	OnDurationWarningThresholdExceededInput() interface{}
	OnFailure() JobWebhookNotificationsOnFailureList
	OnFailureInput() interface{}
	OnStart() JobWebhookNotificationsOnStartList
	OnStartInput() interface{}
	OnStreamingBacklogExceeded() JobWebhookNotificationsOnStreamingBacklogExceededList
	OnStreamingBacklogExceededInput() interface{}
	OnSuccess() JobWebhookNotificationsOnSuccessList
	OnSuccessInput() interface{}
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
	PutOnDurationWarningThresholdExceeded(value interface{})
	PutOnFailure(value interface{})
	PutOnStart(value interface{})
	PutOnStreamingBacklogExceeded(value interface{})
	PutOnSuccess(value interface{})
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

// The jsii proxy struct for JobWebhookNotificationsOutputReference
type jsiiProxy_JobWebhookNotificationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) InternalValue() *JobWebhookNotifications {
	var returns *JobWebhookNotifications
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) OnDurationWarningThresholdExceeded() JobWebhookNotificationsOnDurationWarningThresholdExceededList {
	var returns JobWebhookNotificationsOnDurationWarningThresholdExceededList
	_jsii_.Get(
		j,
		"onDurationWarningThresholdExceeded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) OnDurationWarningThresholdExceededInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onDurationWarningThresholdExceededInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) OnFailure() JobWebhookNotificationsOnFailureList {
	var returns JobWebhookNotificationsOnFailureList
	_jsii_.Get(
		j,
		"onFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) OnFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) OnStart() JobWebhookNotificationsOnStartList {
	var returns JobWebhookNotificationsOnStartList
	_jsii_.Get(
		j,
		"onStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) OnStartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) OnStreamingBacklogExceeded() JobWebhookNotificationsOnStreamingBacklogExceededList {
	var returns JobWebhookNotificationsOnStreamingBacklogExceededList
	_jsii_.Get(
		j,
		"onStreamingBacklogExceeded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) OnStreamingBacklogExceededInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onStreamingBacklogExceededInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) OnSuccess() JobWebhookNotificationsOnSuccessList {
	var returns JobWebhookNotificationsOnSuccessList
	_jsii_.Get(
		j,
		"onSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) OnSuccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onSuccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobWebhookNotificationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) JobWebhookNotificationsOutputReference {
	_init_.Initialize()

	if err := validateNewJobWebhookNotificationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobWebhookNotificationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobWebhookNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobWebhookNotificationsOutputReference_Override(j JobWebhookNotificationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobWebhookNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference)SetInternalValue(val *JobWebhookNotifications) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) PutOnDurationWarningThresholdExceeded(value interface{}) {
	if err := j.validatePutOnDurationWarningThresholdExceededParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putOnDurationWarningThresholdExceeded",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) PutOnFailure(value interface{}) {
	if err := j.validatePutOnFailureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putOnFailure",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) PutOnStart(value interface{}) {
	if err := j.validatePutOnStartParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putOnStart",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) PutOnStreamingBacklogExceeded(value interface{}) {
	if err := j.validatePutOnStreamingBacklogExceededParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putOnStreamingBacklogExceeded",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) PutOnSuccess(value interface{}) {
	if err := j.validatePutOnSuccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putOnSuccess",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) ResetOnDurationWarningThresholdExceeded() {
	_jsii_.InvokeVoid(
		j,
		"resetOnDurationWarningThresholdExceeded",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) ResetOnFailure() {
	_jsii_.InvokeVoid(
		j,
		"resetOnFailure",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) ResetOnStart() {
	_jsii_.InvokeVoid(
		j,
		"resetOnStart",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) ResetOnStreamingBacklogExceeded() {
	_jsii_.InvokeVoid(
		j,
		"resetOnStreamingBacklogExceeded",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) ResetOnSuccess() {
	_jsii_.InvokeVoid(
		j,
		"resetOnSuccess",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobWebhookNotificationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

