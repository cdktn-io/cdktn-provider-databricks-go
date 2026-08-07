// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksexternallocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksexternallocation/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference interface {
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
	InternalValue() *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqs
	SetInternalValue(val *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqs)
	ManagedResourceId() *string
	SetManagedResourceId(val *string)
	ManagedResourceIdInput() *string
	QueueUrl() *string
	SetQueueUrl(val *string)
	QueueUrlInput() *string
	ResourceGroup() *string
	SetResourceGroup(val *string)
	ResourceGroupInput() *string
	SubscriptionId() *string
	SetSubscriptionId(val *string)
	SubscriptionIdInput() *string
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
	ResetManagedResourceId()
	ResetQueueUrl()
	ResetResourceGroup()
	ResetSubscriptionId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference
type jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) InternalValue() *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqs {
	var returns *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) ManagedResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) ManagedResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) QueueUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) QueueUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) ResourceGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) ResourceGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) SubscriptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) SubscriptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksExternalLocation.DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference_Override(d DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksExternalLocation.DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference)SetInternalValue(val *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqs) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference)SetManagedResourceId(val *string) {
	if err := j.validateSetManagedResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedResourceId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference)SetQueueUrl(val *string) {
	if err := j.validateSetQueueUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueUrl",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference)SetResourceGroup(val *string) {
	if err := j.validateSetResourceGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroup",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference)SetSubscriptionId(val *string) {
	if err := j.validateSetSubscriptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscriptionId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) ResetManagedResourceId() {
	_jsii_.InvokeVoid(
		d,
		"resetManagedResourceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) ResetQueueUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetQueueUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) ResetResourceGroup() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceGroup",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) ResetSubscriptionId() {
	_jsii_.InvokeVoid(
		d,
		"resetSubscriptionId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

