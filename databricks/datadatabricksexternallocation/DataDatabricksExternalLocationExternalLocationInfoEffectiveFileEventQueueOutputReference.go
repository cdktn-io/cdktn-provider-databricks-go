// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksexternallocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksexternallocation/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference interface {
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
	InternalValue() *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueue
	SetInternalValue(val *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueue)
	ManagedAqs() DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference
	ManagedAqsInput() *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqs
	ManagedPubsub() DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedPubsubOutputReference
	ManagedPubsubInput() *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedPubsub
	ManagedSqs() DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedSqsOutputReference
	ManagedSqsInput() *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedSqs
	ProvidedAqs() DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedAqsOutputReference
	ProvidedAqsInput() *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedAqs
	ProvidedPubsub() DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedPubsubOutputReference
	ProvidedPubsubInput() *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedPubsub
	ProvidedSqs() DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedSqsOutputReference
	ProvidedSqsInput() *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedSqs
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
	PutManagedAqs(value *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqs)
	PutManagedPubsub(value *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedPubsub)
	PutManagedSqs(value *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedSqs)
	PutProvidedAqs(value *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedAqs)
	PutProvidedPubsub(value *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedPubsub)
	PutProvidedSqs(value *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedSqs)
	ResetManagedAqs()
	ResetManagedPubsub()
	ResetManagedSqs()
	ResetProvidedAqs()
	ResetProvidedPubsub()
	ResetProvidedSqs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference
type jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) InternalValue() *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueue {
	var returns *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueue
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ManagedAqs() DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference {
	var returns DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqsOutputReference
	_jsii_.Get(
		j,
		"managedAqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ManagedAqsInput() *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqs {
	var returns *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqs
	_jsii_.Get(
		j,
		"managedAqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ManagedPubsub() DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedPubsubOutputReference {
	var returns DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedPubsubOutputReference
	_jsii_.Get(
		j,
		"managedPubsub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ManagedPubsubInput() *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedPubsub {
	var returns *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedPubsub
	_jsii_.Get(
		j,
		"managedPubsubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ManagedSqs() DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedSqsOutputReference {
	var returns DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedSqsOutputReference
	_jsii_.Get(
		j,
		"managedSqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ManagedSqsInput() *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedSqs {
	var returns *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedSqs
	_jsii_.Get(
		j,
		"managedSqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ProvidedAqs() DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedAqsOutputReference {
	var returns DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedAqsOutputReference
	_jsii_.Get(
		j,
		"providedAqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ProvidedAqsInput() *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedAqs {
	var returns *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedAqs
	_jsii_.Get(
		j,
		"providedAqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ProvidedPubsub() DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedPubsubOutputReference {
	var returns DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedPubsubOutputReference
	_jsii_.Get(
		j,
		"providedPubsub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ProvidedPubsubInput() *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedPubsub {
	var returns *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedPubsub
	_jsii_.Get(
		j,
		"providedPubsubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ProvidedSqs() DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedSqsOutputReference {
	var returns DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedSqsOutputReference
	_jsii_.Get(
		j,
		"providedSqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ProvidedSqsInput() *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedSqs {
	var returns *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedSqs
	_jsii_.Get(
		j,
		"providedSqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksExternalLocation.DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference_Override(d DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksExternalLocation.DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference)SetInternalValue(val *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueue) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) PutManagedAqs(value *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqs) {
	if err := d.validatePutManagedAqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putManagedAqs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) PutManagedPubsub(value *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedPubsub) {
	if err := d.validatePutManagedPubsubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putManagedPubsub",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) PutManagedSqs(value *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedSqs) {
	if err := d.validatePutManagedSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putManagedSqs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) PutProvidedAqs(value *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedAqs) {
	if err := d.validatePutProvidedAqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProvidedAqs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) PutProvidedPubsub(value *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedPubsub) {
	if err := d.validatePutProvidedPubsubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProvidedPubsub",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) PutProvidedSqs(value *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedSqs) {
	if err := d.validatePutProvidedSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProvidedSqs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ResetManagedAqs() {
	_jsii_.InvokeVoid(
		d,
		"resetManagedAqs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ResetManagedPubsub() {
	_jsii_.InvokeVoid(
		d,
		"resetManagedPubsub",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ResetManagedSqs() {
	_jsii_.InvokeVoid(
		d,
		"resetManagedSqs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ResetProvidedAqs() {
	_jsii_.InvokeVoid(
		d,
		"resetProvidedAqs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ResetProvidedPubsub() {
	_jsii_.InvokeVoid(
		d,
		"resetProvidedPubsub",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ResetProvidedSqs() {
	_jsii_.InvokeVoid(
		d,
		"resetProvidedSqs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

