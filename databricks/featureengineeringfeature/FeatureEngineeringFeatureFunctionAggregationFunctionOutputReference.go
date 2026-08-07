// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/featureengineeringfeature/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference interface {
	cdktn.ComplexObject
	ApproxCountDistinct() FeatureEngineeringFeatureFunctionAggregationFunctionApproxCountDistinctOutputReference
	ApproxCountDistinctInput() interface{}
	ApproxPercentile() FeatureEngineeringFeatureFunctionAggregationFunctionApproxPercentileOutputReference
	ApproxPercentileInput() interface{}
	Avg() FeatureEngineeringFeatureFunctionAggregationFunctionAvgOutputReference
	AvgInput() interface{}
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
	CountFunction() FeatureEngineeringFeatureFunctionAggregationFunctionCountFunctionOutputReference
	CountFunctionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	First() FeatureEngineeringFeatureFunctionAggregationFunctionFirstOutputReference
	FirstDistinct() FeatureEngineeringFeatureFunctionAggregationFunctionFirstDistinctOutputReference
	FirstDistinctInput() interface{}
	FirstInput() interface{}
	FirstN() FeatureEngineeringFeatureFunctionAggregationFunctionFirstNOutputReference
	FirstNInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Last() FeatureEngineeringFeatureFunctionAggregationFunctionLastOutputReference
	LastDistinct() FeatureEngineeringFeatureFunctionAggregationFunctionLastDistinctOutputReference
	LastDistinctInput() interface{}
	LastInput() interface{}
	LastN() FeatureEngineeringFeatureFunctionAggregationFunctionLastNOutputReference
	LastNInput() interface{}
	Max() FeatureEngineeringFeatureFunctionAggregationFunctionMaxOutputReference
	MaxInput() interface{}
	Min() FeatureEngineeringFeatureFunctionAggregationFunctionMinOutputReference
	MinInput() interface{}
	StddevPop() FeatureEngineeringFeatureFunctionAggregationFunctionStddevPopOutputReference
	StddevPopInput() interface{}
	StddevSamp() FeatureEngineeringFeatureFunctionAggregationFunctionStddevSampOutputReference
	StddevSampInput() interface{}
	Sum() FeatureEngineeringFeatureFunctionAggregationFunctionSumOutputReference
	SumInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeWindow() FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference
	TimeWindowInput() interface{}
	VarPop() FeatureEngineeringFeatureFunctionAggregationFunctionVarPopOutputReference
	VarPopInput() interface{}
	VarSamp() FeatureEngineeringFeatureFunctionAggregationFunctionVarSampOutputReference
	VarSampInput() interface{}
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
	PutApproxCountDistinct(value *FeatureEngineeringFeatureFunctionAggregationFunctionApproxCountDistinct)
	PutApproxPercentile(value *FeatureEngineeringFeatureFunctionAggregationFunctionApproxPercentile)
	PutAvg(value *FeatureEngineeringFeatureFunctionAggregationFunctionAvg)
	PutCountFunction(value *FeatureEngineeringFeatureFunctionAggregationFunctionCountFunction)
	PutFirst(value *FeatureEngineeringFeatureFunctionAggregationFunctionFirst)
	PutFirstDistinct(value *FeatureEngineeringFeatureFunctionAggregationFunctionFirstDistinct)
	PutFirstN(value *FeatureEngineeringFeatureFunctionAggregationFunctionFirstN)
	PutLast(value *FeatureEngineeringFeatureFunctionAggregationFunctionLast)
	PutLastDistinct(value *FeatureEngineeringFeatureFunctionAggregationFunctionLastDistinct)
	PutLastN(value *FeatureEngineeringFeatureFunctionAggregationFunctionLastN)
	PutMax(value *FeatureEngineeringFeatureFunctionAggregationFunctionMax)
	PutMin(value *FeatureEngineeringFeatureFunctionAggregationFunctionMin)
	PutStddevPop(value *FeatureEngineeringFeatureFunctionAggregationFunctionStddevPop)
	PutStddevSamp(value *FeatureEngineeringFeatureFunctionAggregationFunctionStddevSamp)
	PutSum(value *FeatureEngineeringFeatureFunctionAggregationFunctionSum)
	PutTimeWindow(value *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindow)
	PutVarPop(value *FeatureEngineeringFeatureFunctionAggregationFunctionVarPop)
	PutVarSamp(value *FeatureEngineeringFeatureFunctionAggregationFunctionVarSamp)
	ResetApproxCountDistinct()
	ResetApproxPercentile()
	ResetAvg()
	ResetCountFunction()
	ResetFirst()
	ResetFirstDistinct()
	ResetFirstN()
	ResetLast()
	ResetLastDistinct()
	ResetLastN()
	ResetMax()
	ResetMin()
	ResetStddevPop()
	ResetStddevSamp()
	ResetSum()
	ResetTimeWindow()
	ResetVarPop()
	ResetVarSamp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference
type jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ApproxCountDistinct() FeatureEngineeringFeatureFunctionAggregationFunctionApproxCountDistinctOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionApproxCountDistinctOutputReference
	_jsii_.Get(
		j,
		"approxCountDistinct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ApproxCountDistinctInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approxCountDistinctInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ApproxPercentile() FeatureEngineeringFeatureFunctionAggregationFunctionApproxPercentileOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionApproxPercentileOutputReference
	_jsii_.Get(
		j,
		"approxPercentile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ApproxPercentileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approxPercentileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) Avg() FeatureEngineeringFeatureFunctionAggregationFunctionAvgOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionAvgOutputReference
	_jsii_.Get(
		j,
		"avg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) AvgInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"avgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) CountFunction() FeatureEngineeringFeatureFunctionAggregationFunctionCountFunctionOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionCountFunctionOutputReference
	_jsii_.Get(
		j,
		"countFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) CountFunctionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"countFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) First() FeatureEngineeringFeatureFunctionAggregationFunctionFirstOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionFirstOutputReference
	_jsii_.Get(
		j,
		"first",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) FirstDistinct() FeatureEngineeringFeatureFunctionAggregationFunctionFirstDistinctOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionFirstDistinctOutputReference
	_jsii_.Get(
		j,
		"firstDistinct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) FirstDistinctInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstDistinctInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) FirstInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) FirstN() FeatureEngineeringFeatureFunctionAggregationFunctionFirstNOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionFirstNOutputReference
	_jsii_.Get(
		j,
		"firstN",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) FirstNInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstNInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) Last() FeatureEngineeringFeatureFunctionAggregationFunctionLastOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionLastOutputReference
	_jsii_.Get(
		j,
		"last",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) LastDistinct() FeatureEngineeringFeatureFunctionAggregationFunctionLastDistinctOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionLastDistinctOutputReference
	_jsii_.Get(
		j,
		"lastDistinct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) LastDistinctInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastDistinctInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) LastInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) LastN() FeatureEngineeringFeatureFunctionAggregationFunctionLastNOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionLastNOutputReference
	_jsii_.Get(
		j,
		"lastN",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) LastNInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastNInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) Max() FeatureEngineeringFeatureFunctionAggregationFunctionMaxOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionMaxOutputReference
	_jsii_.Get(
		j,
		"max",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) MaxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) Min() FeatureEngineeringFeatureFunctionAggregationFunctionMinOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionMinOutputReference
	_jsii_.Get(
		j,
		"min",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) MinInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"minInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) StddevPop() FeatureEngineeringFeatureFunctionAggregationFunctionStddevPopOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionStddevPopOutputReference
	_jsii_.Get(
		j,
		"stddevPop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) StddevPopInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stddevPopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) StddevSamp() FeatureEngineeringFeatureFunctionAggregationFunctionStddevSampOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionStddevSampOutputReference
	_jsii_.Get(
		j,
		"stddevSamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) StddevSampInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stddevSampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) Sum() FeatureEngineeringFeatureFunctionAggregationFunctionSumOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionSumOutputReference
	_jsii_.Get(
		j,
		"sum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) SumInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) TimeWindow() FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference
	_jsii_.Get(
		j,
		"timeWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) TimeWindowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) VarPop() FeatureEngineeringFeatureFunctionAggregationFunctionVarPopOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionVarPopOutputReference
	_jsii_.Get(
		j,
		"varPop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) VarPopInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"varPopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) VarSamp() FeatureEngineeringFeatureFunctionAggregationFunctionVarSampOutputReference {
	var returns FeatureEngineeringFeatureFunctionAggregationFunctionVarSampOutputReference
	_jsii_.Get(
		j,
		"varSamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) VarSampInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"varSampInput",
		&returns,
	)
	return returns
}


func NewFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference {
	_init_.Initialize()

	if err := validateNewFeatureEngineeringFeatureFunctionAggregationFunctionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.featureEngineeringFeature.FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference_Override(f FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.featureEngineeringFeature.FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutApproxCountDistinct(value *FeatureEngineeringFeatureFunctionAggregationFunctionApproxCountDistinct) {
	if err := f.validatePutApproxCountDistinctParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putApproxCountDistinct",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutApproxPercentile(value *FeatureEngineeringFeatureFunctionAggregationFunctionApproxPercentile) {
	if err := f.validatePutApproxPercentileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putApproxPercentile",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutAvg(value *FeatureEngineeringFeatureFunctionAggregationFunctionAvg) {
	if err := f.validatePutAvgParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putAvg",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutCountFunction(value *FeatureEngineeringFeatureFunctionAggregationFunctionCountFunction) {
	if err := f.validatePutCountFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putCountFunction",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutFirst(value *FeatureEngineeringFeatureFunctionAggregationFunctionFirst) {
	if err := f.validatePutFirstParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putFirst",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutFirstDistinct(value *FeatureEngineeringFeatureFunctionAggregationFunctionFirstDistinct) {
	if err := f.validatePutFirstDistinctParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putFirstDistinct",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutFirstN(value *FeatureEngineeringFeatureFunctionAggregationFunctionFirstN) {
	if err := f.validatePutFirstNParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putFirstN",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutLast(value *FeatureEngineeringFeatureFunctionAggregationFunctionLast) {
	if err := f.validatePutLastParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putLast",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutLastDistinct(value *FeatureEngineeringFeatureFunctionAggregationFunctionLastDistinct) {
	if err := f.validatePutLastDistinctParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putLastDistinct",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutLastN(value *FeatureEngineeringFeatureFunctionAggregationFunctionLastN) {
	if err := f.validatePutLastNParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putLastN",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutMax(value *FeatureEngineeringFeatureFunctionAggregationFunctionMax) {
	if err := f.validatePutMaxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putMax",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutMin(value *FeatureEngineeringFeatureFunctionAggregationFunctionMin) {
	if err := f.validatePutMinParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putMin",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutStddevPop(value *FeatureEngineeringFeatureFunctionAggregationFunctionStddevPop) {
	if err := f.validatePutStddevPopParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putStddevPop",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutStddevSamp(value *FeatureEngineeringFeatureFunctionAggregationFunctionStddevSamp) {
	if err := f.validatePutStddevSampParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putStddevSamp",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutSum(value *FeatureEngineeringFeatureFunctionAggregationFunctionSum) {
	if err := f.validatePutSumParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putSum",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutTimeWindow(value *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindow) {
	if err := f.validatePutTimeWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTimeWindow",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutVarPop(value *FeatureEngineeringFeatureFunctionAggregationFunctionVarPop) {
	if err := f.validatePutVarPopParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putVarPop",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutVarSamp(value *FeatureEngineeringFeatureFunctionAggregationFunctionVarSamp) {
	if err := f.validatePutVarSampParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putVarSamp",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetApproxCountDistinct() {
	_jsii_.InvokeVoid(
		f,
		"resetApproxCountDistinct",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetApproxPercentile() {
	_jsii_.InvokeVoid(
		f,
		"resetApproxPercentile",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetAvg() {
	_jsii_.InvokeVoid(
		f,
		"resetAvg",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetCountFunction() {
	_jsii_.InvokeVoid(
		f,
		"resetCountFunction",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetFirst() {
	_jsii_.InvokeVoid(
		f,
		"resetFirst",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetFirstDistinct() {
	_jsii_.InvokeVoid(
		f,
		"resetFirstDistinct",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetFirstN() {
	_jsii_.InvokeVoid(
		f,
		"resetFirstN",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetLast() {
	_jsii_.InvokeVoid(
		f,
		"resetLast",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetLastDistinct() {
	_jsii_.InvokeVoid(
		f,
		"resetLastDistinct",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetLastN() {
	_jsii_.InvokeVoid(
		f,
		"resetLastN",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetMax() {
	_jsii_.InvokeVoid(
		f,
		"resetMax",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetMin() {
	_jsii_.InvokeVoid(
		f,
		"resetMin",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetStddevPop() {
	_jsii_.InvokeVoid(
		f,
		"resetStddevPop",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetStddevSamp() {
	_jsii_.InvokeVoid(
		f,
		"resetStddevSamp",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetSum() {
	_jsii_.InvokeVoid(
		f,
		"resetSum",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetTimeWindow() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeWindow",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetVarPop() {
	_jsii_.InvokeVoid(
		f,
		"resetVarPop",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetVarSamp() {
	_jsii_.InvokeVoid(
		f,
		"resetVarSamp",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (f *jsiiProxy_FeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

