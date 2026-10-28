// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/job/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference interface {
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
	DataAttachment() JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentDataAttachmentOutputReference
	DataAttachmentInput() *JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentDataAttachment
	// Experimental.
	Fqn() *string
	InternalValue() *JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachment
	SetInternalValue(val *JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachment)
	PageAttachment() JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentPageAttachmentOutputReference
	PageAttachmentInput() *JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentPageAttachment
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
	PutDataAttachment(value *JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentDataAttachment)
	PutPageAttachment(value *JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentPageAttachment)
	ResetDataAttachment()
	ResetPageAttachment()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference
type jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) DataAttachment() JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentDataAttachmentOutputReference {
	var returns JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentDataAttachmentOutputReference
	_jsii_.Get(
		j,
		"dataAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) DataAttachmentInput() *JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentDataAttachment {
	var returns *JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentDataAttachment
	_jsii_.Get(
		j,
		"dataAttachmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) InternalValue() *JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachment {
	var returns *JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachment
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) PageAttachment() JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentPageAttachmentOutputReference {
	var returns JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentPageAttachmentOutputReference
	_jsii_.Get(
		j,
		"pageAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) PageAttachmentInput() *JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentPageAttachment {
	var returns *JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentPageAttachment
	_jsii_.Get(
		j,
		"pageAttachmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference_Override(j JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference)SetInternalValue(val *JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachment) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) PutDataAttachment(value *JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentDataAttachment) {
	if err := j.validatePutDataAttachmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDataAttachment",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) PutPageAttachment(value *JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentPageAttachment) {
	if err := j.validatePutPageAttachmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPageAttachment",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) ResetDataAttachment() {
	_jsii_.InvokeVoid(
		j,
		"resetDataAttachment",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) ResetPageAttachment() {
	_jsii_.InvokeVoid(
		j,
		"resetPageAttachment",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

