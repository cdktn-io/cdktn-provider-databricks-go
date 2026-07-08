// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeature

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabricksfeatureengineeringfeature/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference interface {
	cdktn.ComplexObject
	ApproxCountDistinct() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionApproxCountDistinctOutputReference
	ApproxCountDistinctInput() interface{}
	ApproxPercentile() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionApproxPercentileOutputReference
	ApproxPercentileInput() interface{}
	Avg() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionAvgOutputReference
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
	CountFunction() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionCountFunctionOutputReference
	CountFunctionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	First() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionFirstOutputReference
	FirstDistinct() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionFirstDistinctOutputReference
	FirstDistinctInput() interface{}
	FirstInput() interface{}
	FirstN() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionFirstNOutputReference
	FirstNInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Last() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionLastOutputReference
	LastDistinct() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionLastDistinctOutputReference
	LastDistinctInput() interface{}
	LastInput() interface{}
	LastN() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionLastNOutputReference
	LastNInput() interface{}
	Max() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionMaxOutputReference
	MaxInput() interface{}
	Min() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionMinOutputReference
	MinInput() interface{}
	StddevPop() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionStddevPopOutputReference
	StddevPopInput() interface{}
	StddevSamp() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionStddevSampOutputReference
	StddevSampInput() interface{}
	Sum() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionSumOutputReference
	SumInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeWindow() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference
	TimeWindowInput() interface{}
	VarPop() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionVarPopOutputReference
	VarPopInput() interface{}
	VarSamp() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionVarSampOutputReference
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
	PutApproxCountDistinct(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionApproxCountDistinct)
	PutApproxPercentile(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionApproxPercentile)
	PutAvg(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionAvg)
	PutCountFunction(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionCountFunction)
	PutFirst(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionFirst)
	PutFirstDistinct(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionFirstDistinct)
	PutFirstN(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionFirstN)
	PutLast(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionLast)
	PutLastDistinct(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionLastDistinct)
	PutLastN(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionLastN)
	PutMax(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionMax)
	PutMin(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionMin)
	PutStddevPop(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionStddevPop)
	PutStddevSamp(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionStddevSamp)
	PutSum(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionSum)
	PutTimeWindow(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindow)
	PutVarPop(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionVarPop)
	PutVarSamp(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionVarSamp)
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

// The jsii proxy struct for DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference
type jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ApproxCountDistinct() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionApproxCountDistinctOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionApproxCountDistinctOutputReference
	_jsii_.Get(
		j,
		"approxCountDistinct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ApproxCountDistinctInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approxCountDistinctInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ApproxPercentile() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionApproxPercentileOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionApproxPercentileOutputReference
	_jsii_.Get(
		j,
		"approxPercentile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ApproxPercentileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approxPercentileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) Avg() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionAvgOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionAvgOutputReference
	_jsii_.Get(
		j,
		"avg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) AvgInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"avgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) CountFunction() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionCountFunctionOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionCountFunctionOutputReference
	_jsii_.Get(
		j,
		"countFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) CountFunctionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"countFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) First() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionFirstOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionFirstOutputReference
	_jsii_.Get(
		j,
		"first",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) FirstDistinct() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionFirstDistinctOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionFirstDistinctOutputReference
	_jsii_.Get(
		j,
		"firstDistinct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) FirstDistinctInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstDistinctInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) FirstInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) FirstN() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionFirstNOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionFirstNOutputReference
	_jsii_.Get(
		j,
		"firstN",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) FirstNInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstNInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) Last() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionLastOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionLastOutputReference
	_jsii_.Get(
		j,
		"last",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) LastDistinct() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionLastDistinctOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionLastDistinctOutputReference
	_jsii_.Get(
		j,
		"lastDistinct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) LastDistinctInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastDistinctInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) LastInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) LastN() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionLastNOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionLastNOutputReference
	_jsii_.Get(
		j,
		"lastN",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) LastNInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastNInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) Max() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionMaxOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionMaxOutputReference
	_jsii_.Get(
		j,
		"max",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) MaxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) Min() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionMinOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionMinOutputReference
	_jsii_.Get(
		j,
		"min",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) MinInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"minInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) StddevPop() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionStddevPopOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionStddevPopOutputReference
	_jsii_.Get(
		j,
		"stddevPop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) StddevPopInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stddevPopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) StddevSamp() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionStddevSampOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionStddevSampOutputReference
	_jsii_.Get(
		j,
		"stddevSamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) StddevSampInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stddevSampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) Sum() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionSumOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionSumOutputReference
	_jsii_.Get(
		j,
		"sum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) SumInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) TimeWindow() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowOutputReference
	_jsii_.Get(
		j,
		"timeWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) TimeWindowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) VarPop() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionVarPopOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionVarPopOutputReference
	_jsii_.Get(
		j,
		"varPop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) VarPopInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"varPopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) VarSamp() DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionVarSampOutputReference {
	var returns DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionVarSampOutputReference
	_jsii_.Get(
		j,
		"varSamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) VarSampInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"varSampInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringFeature.DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference_Override(d DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringFeature.DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutApproxCountDistinct(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionApproxCountDistinct) {
	if err := d.validatePutApproxCountDistinctParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApproxCountDistinct",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutApproxPercentile(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionApproxPercentile) {
	if err := d.validatePutApproxPercentileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApproxPercentile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutAvg(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionAvg) {
	if err := d.validatePutAvgParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAvg",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutCountFunction(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionCountFunction) {
	if err := d.validatePutCountFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCountFunction",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutFirst(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionFirst) {
	if err := d.validatePutFirstParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFirst",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutFirstDistinct(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionFirstDistinct) {
	if err := d.validatePutFirstDistinctParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFirstDistinct",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutFirstN(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionFirstN) {
	if err := d.validatePutFirstNParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFirstN",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutLast(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionLast) {
	if err := d.validatePutLastParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLast",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutLastDistinct(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionLastDistinct) {
	if err := d.validatePutLastDistinctParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLastDistinct",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutLastN(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionLastN) {
	if err := d.validatePutLastNParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLastN",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutMax(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionMax) {
	if err := d.validatePutMaxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMax",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutMin(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionMin) {
	if err := d.validatePutMinParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMin",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutStddevPop(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionStddevPop) {
	if err := d.validatePutStddevPopParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStddevPop",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutStddevSamp(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionStddevSamp) {
	if err := d.validatePutStddevSampParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStddevSamp",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutSum(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionSum) {
	if err := d.validatePutSumParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSum",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutTimeWindow(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindow) {
	if err := d.validatePutTimeWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeWindow",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutVarPop(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionVarPop) {
	if err := d.validatePutVarPopParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVarPop",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) PutVarSamp(value *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionVarSamp) {
	if err := d.validatePutVarSampParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVarSamp",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetApproxCountDistinct() {
	_jsii_.InvokeVoid(
		d,
		"resetApproxCountDistinct",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetApproxPercentile() {
	_jsii_.InvokeVoid(
		d,
		"resetApproxPercentile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetAvg() {
	_jsii_.InvokeVoid(
		d,
		"resetAvg",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetCountFunction() {
	_jsii_.InvokeVoid(
		d,
		"resetCountFunction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetFirst() {
	_jsii_.InvokeVoid(
		d,
		"resetFirst",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetFirstDistinct() {
	_jsii_.InvokeVoid(
		d,
		"resetFirstDistinct",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetFirstN() {
	_jsii_.InvokeVoid(
		d,
		"resetFirstN",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetLast() {
	_jsii_.InvokeVoid(
		d,
		"resetLast",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetLastDistinct() {
	_jsii_.InvokeVoid(
		d,
		"resetLastDistinct",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetLastN() {
	_jsii_.InvokeVoid(
		d,
		"resetLastN",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetMax() {
	_jsii_.InvokeVoid(
		d,
		"resetMax",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetMin() {
	_jsii_.InvokeVoid(
		d,
		"resetMin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetStddevPop() {
	_jsii_.InvokeVoid(
		d,
		"resetStddevPop",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetStddevSamp() {
	_jsii_.InvokeVoid(
		d,
		"resetStddevSamp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetSum() {
	_jsii_.InvokeVoid(
		d,
		"resetSum",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetTimeWindow() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeWindow",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetVarPop() {
	_jsii_.InvokeVoid(
		d,
		"resetVarPop",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ResetVarSamp() {
	_jsii_.InvokeVoid(
		d,
		"resetVarSamp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

