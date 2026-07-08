// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeatures

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabricksfeatureengineeringfeatures/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference interface {
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
	Continuous() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowContinuousOutputReference
	ContinuousInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LongRolling() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowLongRollingOutputReference
	LongRollingInput() interface{}
	Rolling() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowRollingOutputReference
	RollingInput() interface{}
	Sliding() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowSlidingOutputReference
	SlidingInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Tumbling() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowTumblingOutputReference
	TumblingInput() interface{}
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
	PutContinuous(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowContinuous)
	PutLongRolling(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowLongRolling)
	PutRolling(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowRolling)
	PutSliding(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowSliding)
	PutTumbling(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowTumbling)
	ResetContinuous()
	ResetLongRolling()
	ResetRolling()
	ResetSliding()
	ResetTumbling()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference
type jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) Continuous() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowContinuousOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowContinuousOutputReference
	_jsii_.Get(
		j,
		"continuous",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) ContinuousInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continuousInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) LongRolling() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowLongRollingOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowLongRollingOutputReference
	_jsii_.Get(
		j,
		"longRolling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) LongRollingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"longRollingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) Rolling() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowRollingOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowRollingOutputReference
	_jsii_.Get(
		j,
		"rolling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) RollingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rollingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) Sliding() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowSlidingOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowSlidingOutputReference
	_jsii_.Get(
		j,
		"sliding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) SlidingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slidingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) Tumbling() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowTumblingOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowTumblingOutputReference
	_jsii_.Get(
		j,
		"tumbling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) TumblingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tumblingInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringFeatures.DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference_Override(d DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringFeatures.DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) PutContinuous(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowContinuous) {
	if err := d.validatePutContinuousParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContinuous",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) PutLongRolling(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowLongRolling) {
	if err := d.validatePutLongRollingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLongRolling",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) PutRolling(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowRolling) {
	if err := d.validatePutRollingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRolling",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) PutSliding(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowSliding) {
	if err := d.validatePutSlidingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSliding",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) PutTumbling(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowTumbling) {
	if err := d.validatePutTumblingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTumbling",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) ResetContinuous() {
	_jsii_.InvokeVoid(
		d,
		"resetContinuous",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) ResetLongRolling() {
	_jsii_.InvokeVoid(
		d,
		"resetLongRolling",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) ResetRolling() {
	_jsii_.InvokeVoid(
		d,
		"resetRolling",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) ResetSliding() {
	_jsii_.InvokeVoid(
		d,
		"resetSliding",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) ResetTumbling() {
	_jsii_.InvokeVoid(
		d,
		"resetTumbling",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

