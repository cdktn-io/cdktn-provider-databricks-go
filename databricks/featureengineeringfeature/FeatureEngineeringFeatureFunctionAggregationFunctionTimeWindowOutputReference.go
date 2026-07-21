// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/featureengineeringfeature/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference interface {
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
	Continuous() FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowContinuousOutputReference
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
	Lifetime() FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowLifetimeOutputReference
	LifetimeInput() interface{}
	LongRolling() FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowLongRollingOutputReference
	LongRollingInput() interface{}
	Rolling() FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowRollingOutputReference
	RollingInput() interface{}
	Sliding() FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSlidingOutputReference
	SlidingInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Tumbling() FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowTumblingOutputReference
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
	PutContinuous(value *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowContinuous)
	PutLifetime(value *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowLifetime)
	PutLongRolling(value *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowLongRolling)
	PutRolling(value *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowRolling)
	PutSliding(value *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSliding)
	PutTumbling(value *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowTumbling)
	ResetContinuous()
	ResetLifetime()
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

// The jsii proxy struct for FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference
type jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) Continuous() FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowContinuousOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowContinuousOutputReference
	_jsii_.Get(
		j,
		"continuous",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ContinuousInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continuousInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) Lifetime() FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowLifetimeOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowLifetimeOutputReference
	_jsii_.Get(
		j,
		"lifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) LifetimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) LongRolling() FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowLongRollingOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowLongRollingOutputReference
	_jsii_.Get(
		j,
		"longRolling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) LongRollingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"longRollingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) Rolling() FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowRollingOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowRollingOutputReference
	_jsii_.Get(
		j,
		"rolling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) RollingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rollingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) Sliding() FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSlidingOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSlidingOutputReference
	_jsii_.Get(
		j,
		"sliding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) SlidingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slidingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) Tumbling() FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowTumblingOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowTumblingOutputReference
	_jsii_.Get(
		j,
		"tumbling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) TumblingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tumblingInput",
		&returns,
	)
	return returns
}


func NewFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference {
	_init_.Initialize()

	if err := validateNewFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.featureEngineeringFeature.FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference_Override(f FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.featureEngineeringFeature.FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) PutContinuous(value *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowContinuous) {
	if err := f.validatePutContinuousParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putContinuous",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) PutLifetime(value *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowLifetime) {
	if err := f.validatePutLifetimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putLifetime",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) PutLongRolling(value *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowLongRolling) {
	if err := f.validatePutLongRollingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putLongRolling",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) PutRolling(value *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowRolling) {
	if err := f.validatePutRollingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putRolling",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) PutSliding(value *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSliding) {
	if err := f.validatePutSlidingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putSliding",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) PutTumbling(value *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowTumbling) {
	if err := f.validatePutTumblingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTumbling",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ResetContinuous() {
	_jsii_.InvokeVoid(
		f,
		"resetContinuous",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ResetLifetime() {
	_jsii_.InvokeVoid(
		f,
		"resetLifetime",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ResetLongRolling() {
	_jsii_.InvokeVoid(
		f,
		"resetLongRolling",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ResetRolling() {
	_jsii_.InvokeVoid(
		f,
		"resetRolling",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ResetSliding() {
	_jsii_.InvokeVoid(
		f,
		"resetSliding",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ResetTumbling() {
	_jsii_.InvokeVoid(
		f,
		"resetTumbling",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := f.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

