// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksrfaaccessrequestdestinations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/datadatabricksrfaaccessrequestdestinations/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference interface {
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
	FullName() *string
	SetFullName(val *string)
	FullNameInput() *string
	InternalValue() *DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurable
	SetInternalValue(val *DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurable)
	ProviderShare() *string
	SetProviderShare(val *string)
	ProviderShareInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetFullName()
	ResetProviderShare()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference
type jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) FullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) FullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) InternalValue() *DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurable {
	var returns *DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurable
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) ProviderShare() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerShare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) ProviderShareInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerShareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksRfaAccessRequestDestinations.DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference_Override(d DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksRfaAccessRequestDestinations.DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference)SetFullName(val *string) {
	if err := j.validateSetFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference)SetInternalValue(val *DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurable) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference)SetProviderShare(val *string) {
	if err := j.validateSetProviderShareParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerShare",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) ResetFullName() {
	_jsii_.InvokeVoid(
		d,
		"resetFullName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) ResetProviderShare() {
	_jsii_.InvokeVoid(
		d,
		"resetProviderShare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		d,
		"resetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksRfaAccessRequestDestinationsDestinationSourceSecurableOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

