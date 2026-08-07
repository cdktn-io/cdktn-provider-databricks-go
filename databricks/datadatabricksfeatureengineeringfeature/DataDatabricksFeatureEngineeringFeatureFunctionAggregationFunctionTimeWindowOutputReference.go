// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeature

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksfeatureengineeringfeature/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference interface {
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
	Continuous() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowContinuousOutputReference
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
	Rolling() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowRollingOutputReference
	RollingInput() interface{}
	Sawtooth() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSawtoothOutputReference
	SawtoothInput() interface{}
	Sliding() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSlidingOutputReference
	SlidingInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Tumbling() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowTumblingOutputReference
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
	PutContinuous(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowContinuous)
	PutRolling(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowRolling)
	PutSawtooth(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSawtooth)
	PutSliding(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSliding)
	PutTumbling(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowTumbling)
	ResetContinuous()
	ResetRolling()
	ResetSawtooth()
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

// The jsii proxy struct for DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference
type jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) Continuous() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowContinuousOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowContinuousOutputReference
	_jsii_.Get(
		j,
		"continuous",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ContinuousInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continuousInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) Rolling() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowRollingOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowRollingOutputReference
	_jsii_.Get(
		j,
		"rolling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) RollingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rollingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) Sawtooth() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSawtoothOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSawtoothOutputReference
	_jsii_.Get(
		j,
		"sawtooth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) SawtoothInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sawtoothInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) Sliding() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSlidingOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSlidingOutputReference
	_jsii_.Get(
		j,
		"sliding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) SlidingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slidingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) Tumbling() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowTumblingOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowTumblingOutputReference
	_jsii_.Get(
		j,
		"tumbling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) TumblingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tumblingInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringFeature.DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference_Override(d DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringFeature.DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) PutContinuous(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowContinuous) {
	if err := d.validatePutContinuousParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContinuous",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) PutRolling(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowRolling) {
	if err := d.validatePutRollingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRolling",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) PutSawtooth(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSawtooth) {
	if err := d.validatePutSawtoothParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSawtooth",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) PutSliding(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSliding) {
	if err := d.validatePutSlidingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSliding",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) PutTumbling(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowTumbling) {
	if err := d.validatePutTumblingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTumbling",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ResetContinuous() {
	_jsii_.InvokeVoid(
		d,
		"resetContinuous",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ResetRolling() {
	_jsii_.InvokeVoid(
		d,
		"resetRolling",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ResetSawtooth() {
	_jsii_.InvokeVoid(
		d,
		"resetSawtooth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ResetSliding() {
	_jsii_.InvokeVoid(
		d,
		"resetSliding",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ResetTumbling() {
	_jsii_.InvokeVoid(
		d,
		"resetTumbling",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

