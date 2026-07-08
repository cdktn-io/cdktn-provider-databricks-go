// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeatures

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabricksfeatureengineeringfeatures/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference interface {
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
	Delay() *string
	SetDelay(val *string)
	DelayInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WindowDuration() *string
	SetWindowDuration(val *string)
	WindowDurationInput() *string
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
	ResetDelay()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference
type jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) Delay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) DelayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) WindowDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) WindowDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowDurationInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringFeatures.DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference_Override(d DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringFeatures.DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference)SetDelay(val *string) {
	if err := j.validateSetDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delay",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference)SetWindowDuration(val *string) {
	if err := j.validateSetWindowDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowDuration",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) ResetDelay() {
	_jsii_.InvokeVoid(
		d,
		"resetDelay",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesTimeWindowLongRollingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

