// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package externallocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/externallocation/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ExternalLocationEffectiveFileEventQueueOutputReference interface {
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
	InternalValue() *ExternalLocationEffectiveFileEventQueue
	SetInternalValue(val *ExternalLocationEffectiveFileEventQueue)
	ManagedAqs() ExternalLocationEffectiveFileEventQueueManagedAqsOutputReference
	ManagedAqsInput() *ExternalLocationEffectiveFileEventQueueManagedAqs
	ManagedPubsub() ExternalLocationEffectiveFileEventQueueManagedPubsubOutputReference
	ManagedPubsubInput() *ExternalLocationEffectiveFileEventQueueManagedPubsub
	ManagedSqs() ExternalLocationEffectiveFileEventQueueManagedSqsOutputReference
	ManagedSqsInput() *ExternalLocationEffectiveFileEventQueueManagedSqs
	ProvidedAqs() ExternalLocationEffectiveFileEventQueueProvidedAqsOutputReference
	ProvidedAqsInput() *ExternalLocationEffectiveFileEventQueueProvidedAqs
	ProvidedPubsub() ExternalLocationEffectiveFileEventQueueProvidedPubsubOutputReference
	ProvidedPubsubInput() *ExternalLocationEffectiveFileEventQueueProvidedPubsub
	ProvidedSqs() ExternalLocationEffectiveFileEventQueueProvidedSqsOutputReference
	ProvidedSqsInput() *ExternalLocationEffectiveFileEventQueueProvidedSqs
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
	PutManagedAqs(value *ExternalLocationEffectiveFileEventQueueManagedAqs)
	PutManagedPubsub(value *ExternalLocationEffectiveFileEventQueueManagedPubsub)
	PutManagedSqs(value *ExternalLocationEffectiveFileEventQueueManagedSqs)
	PutProvidedAqs(value *ExternalLocationEffectiveFileEventQueueProvidedAqs)
	PutProvidedPubsub(value *ExternalLocationEffectiveFileEventQueueProvidedPubsub)
	PutProvidedSqs(value *ExternalLocationEffectiveFileEventQueueProvidedSqs)
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

// The jsii proxy struct for ExternalLocationEffectiveFileEventQueueOutputReference
type jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) InternalValue() *ExternalLocationEffectiveFileEventQueue {
	var returns *ExternalLocationEffectiveFileEventQueue
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ManagedAqs() ExternalLocationEffectiveFileEventQueueManagedAqsOutputReference {
	var returns ExternalLocationEffectiveFileEventQueueManagedAqsOutputReference
	_jsii_.Get(
		j,
		"managedAqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ManagedAqsInput() *ExternalLocationEffectiveFileEventQueueManagedAqs {
	var returns *ExternalLocationEffectiveFileEventQueueManagedAqs
	_jsii_.Get(
		j,
		"managedAqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ManagedPubsub() ExternalLocationEffectiveFileEventQueueManagedPubsubOutputReference {
	var returns ExternalLocationEffectiveFileEventQueueManagedPubsubOutputReference
	_jsii_.Get(
		j,
		"managedPubsub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ManagedPubsubInput() *ExternalLocationEffectiveFileEventQueueManagedPubsub {
	var returns *ExternalLocationEffectiveFileEventQueueManagedPubsub
	_jsii_.Get(
		j,
		"managedPubsubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ManagedSqs() ExternalLocationEffectiveFileEventQueueManagedSqsOutputReference {
	var returns ExternalLocationEffectiveFileEventQueueManagedSqsOutputReference
	_jsii_.Get(
		j,
		"managedSqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ManagedSqsInput() *ExternalLocationEffectiveFileEventQueueManagedSqs {
	var returns *ExternalLocationEffectiveFileEventQueueManagedSqs
	_jsii_.Get(
		j,
		"managedSqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ProvidedAqs() ExternalLocationEffectiveFileEventQueueProvidedAqsOutputReference {
	var returns ExternalLocationEffectiveFileEventQueueProvidedAqsOutputReference
	_jsii_.Get(
		j,
		"providedAqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ProvidedAqsInput() *ExternalLocationEffectiveFileEventQueueProvidedAqs {
	var returns *ExternalLocationEffectiveFileEventQueueProvidedAqs
	_jsii_.Get(
		j,
		"providedAqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ProvidedPubsub() ExternalLocationEffectiveFileEventQueueProvidedPubsubOutputReference {
	var returns ExternalLocationEffectiveFileEventQueueProvidedPubsubOutputReference
	_jsii_.Get(
		j,
		"providedPubsub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ProvidedPubsubInput() *ExternalLocationEffectiveFileEventQueueProvidedPubsub {
	var returns *ExternalLocationEffectiveFileEventQueueProvidedPubsub
	_jsii_.Get(
		j,
		"providedPubsubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ProvidedSqs() ExternalLocationEffectiveFileEventQueueProvidedSqsOutputReference {
	var returns ExternalLocationEffectiveFileEventQueueProvidedSqsOutputReference
	_jsii_.Get(
		j,
		"providedSqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ProvidedSqsInput() *ExternalLocationEffectiveFileEventQueueProvidedSqs {
	var returns *ExternalLocationEffectiveFileEventQueueProvidedSqs
	_jsii_.Get(
		j,
		"providedSqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewExternalLocationEffectiveFileEventQueueOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ExternalLocationEffectiveFileEventQueueOutputReference {
	_init_.Initialize()

	if err := validateNewExternalLocationEffectiveFileEventQueueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.externalLocation.ExternalLocationEffectiveFileEventQueueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewExternalLocationEffectiveFileEventQueueOutputReference_Override(e ExternalLocationEffectiveFileEventQueueOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.externalLocation.ExternalLocationEffectiveFileEventQueueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference)SetInternalValue(val *ExternalLocationEffectiveFileEventQueue) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) PutManagedAqs(value *ExternalLocationEffectiveFileEventQueueManagedAqs) {
	if err := e.validatePutManagedAqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putManagedAqs",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) PutManagedPubsub(value *ExternalLocationEffectiveFileEventQueueManagedPubsub) {
	if err := e.validatePutManagedPubsubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putManagedPubsub",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) PutManagedSqs(value *ExternalLocationEffectiveFileEventQueueManagedSqs) {
	if err := e.validatePutManagedSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putManagedSqs",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) PutProvidedAqs(value *ExternalLocationEffectiveFileEventQueueProvidedAqs) {
	if err := e.validatePutProvidedAqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putProvidedAqs",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) PutProvidedPubsub(value *ExternalLocationEffectiveFileEventQueueProvidedPubsub) {
	if err := e.validatePutProvidedPubsubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putProvidedPubsub",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) PutProvidedSqs(value *ExternalLocationEffectiveFileEventQueueProvidedSqs) {
	if err := e.validatePutProvidedSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putProvidedSqs",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ResetManagedAqs() {
	_jsii_.InvokeVoid(
		e,
		"resetManagedAqs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ResetManagedPubsub() {
	_jsii_.InvokeVoid(
		e,
		"resetManagedPubsub",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ResetManagedSqs() {
	_jsii_.InvokeVoid(
		e,
		"resetManagedSqs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ResetProvidedAqs() {
	_jsii_.InvokeVoid(
		e,
		"resetProvidedAqs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ResetProvidedPubsub() {
	_jsii_.InvokeVoid(
		e,
		"resetProvidedPubsub",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ResetProvidedSqs() {
	_jsii_.InvokeVoid(
		e,
		"resetProvidedSqs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalLocationEffectiveFileEventQueueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

