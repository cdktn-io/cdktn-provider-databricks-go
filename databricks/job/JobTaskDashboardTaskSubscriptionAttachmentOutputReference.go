// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/job/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JobTaskDashboardTaskSubscriptionAttachmentOutputReference interface {
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
	DataAttachment() JobTaskDashboardTaskSubscriptionAttachmentDataAttachmentOutputReference
	DataAttachmentInput() *JobTaskDashboardTaskSubscriptionAttachmentDataAttachment
	// Experimental.
	Fqn() *string
	InternalValue() *JobTaskDashboardTaskSubscriptionAttachment
	SetInternalValue(val *JobTaskDashboardTaskSubscriptionAttachment)
	PageAttachment() JobTaskDashboardTaskSubscriptionAttachmentPageAttachmentOutputReference
	PageAttachmentInput() *JobTaskDashboardTaskSubscriptionAttachmentPageAttachment
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
	PutDataAttachment(value *JobTaskDashboardTaskSubscriptionAttachmentDataAttachment)
	PutPageAttachment(value *JobTaskDashboardTaskSubscriptionAttachmentPageAttachment)
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

// The jsii proxy struct for JobTaskDashboardTaskSubscriptionAttachmentOutputReference
type jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) DataAttachment() JobTaskDashboardTaskSubscriptionAttachmentDataAttachmentOutputReference {
	var returns JobTaskDashboardTaskSubscriptionAttachmentDataAttachmentOutputReference
	_jsii_.Get(
		j,
		"dataAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) DataAttachmentInput() *JobTaskDashboardTaskSubscriptionAttachmentDataAttachment {
	var returns *JobTaskDashboardTaskSubscriptionAttachmentDataAttachment
	_jsii_.Get(
		j,
		"dataAttachmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) InternalValue() *JobTaskDashboardTaskSubscriptionAttachment {
	var returns *JobTaskDashboardTaskSubscriptionAttachment
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) PageAttachment() JobTaskDashboardTaskSubscriptionAttachmentPageAttachmentOutputReference {
	var returns JobTaskDashboardTaskSubscriptionAttachmentPageAttachmentOutputReference
	_jsii_.Get(
		j,
		"pageAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) PageAttachmentInput() *JobTaskDashboardTaskSubscriptionAttachmentPageAttachment {
	var returns *JobTaskDashboardTaskSubscriptionAttachmentPageAttachment
	_jsii_.Get(
		j,
		"pageAttachmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobTaskDashboardTaskSubscriptionAttachmentOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) JobTaskDashboardTaskSubscriptionAttachmentOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskDashboardTaskSubscriptionAttachmentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskDashboardTaskSubscriptionAttachmentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskDashboardTaskSubscriptionAttachmentOutputReference_Override(j JobTaskDashboardTaskSubscriptionAttachmentOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTaskDashboardTaskSubscriptionAttachmentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference)SetInternalValue(val *JobTaskDashboardTaskSubscriptionAttachment) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) PutDataAttachment(value *JobTaskDashboardTaskSubscriptionAttachmentDataAttachment) {
	if err := j.validatePutDataAttachmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDataAttachment",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) PutPageAttachment(value *JobTaskDashboardTaskSubscriptionAttachmentPageAttachment) {
	if err := j.validatePutPageAttachmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPageAttachment",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) ResetDataAttachment() {
	_jsii_.InvokeVoid(
		j,
		"resetDataAttachment",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) ResetPageAttachment() {
	_jsii_.InvokeVoid(
		j,
		"resetPageAttachment",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskDashboardTaskSubscriptionAttachmentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

