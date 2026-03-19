// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeatures

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/datadatabricksfeatureengineeringfeatures/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference interface {
	cdktn.ComplexObject
	ApproxCountDistinct() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionApproxCountDistinctOutputReference
	ApproxCountDistinctInput() interface{}
	ApproxPercentile() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionApproxPercentileOutputReference
	ApproxPercentileInput() interface{}
	Avg() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionAvgOutputReference
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
	CountFunction() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionCountFunctionOutputReference
	CountFunctionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	First() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionFirstOutputReference
	FirstInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Last() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionLastOutputReference
	LastInput() interface{}
	Max() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionMaxOutputReference
	MaxInput() interface{}
	Min() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionMinOutputReference
	MinInput() interface{}
	StddevPop() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionStddevPopOutputReference
	StddevPopInput() interface{}
	StddevSamp() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionStddevSampOutputReference
	StddevSampInput() interface{}
	Sum() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionSumOutputReference
	SumInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeWindow() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference
	TimeWindowInput() interface{}
	VarPop() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionVarPopOutputReference
	VarPopInput() interface{}
	VarSamp() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionVarSampOutputReference
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
	PutApproxCountDistinct(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionApproxCountDistinct)
	PutApproxPercentile(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionApproxPercentile)
	PutAvg(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionAvg)
	PutCountFunction(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionCountFunction)
	PutFirst(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionFirst)
	PutLast(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionLast)
	PutMax(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionMax)
	PutMin(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionMin)
	PutStddevPop(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionStddevPop)
	PutStddevSamp(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionStddevSamp)
	PutSum(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionSum)
	PutTimeWindow(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindow)
	PutVarPop(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionVarPop)
	PutVarSamp(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionVarSamp)
	ResetApproxCountDistinct()
	ResetApproxPercentile()
	ResetAvg()
	ResetCountFunction()
	ResetFirst()
	ResetLast()
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

// The jsii proxy struct for DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference
type jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ApproxCountDistinct() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionApproxCountDistinctOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionApproxCountDistinctOutputReference
	_jsii_.Get(
		j,
		"approxCountDistinct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ApproxCountDistinctInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approxCountDistinctInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ApproxPercentile() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionApproxPercentileOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionApproxPercentileOutputReference
	_jsii_.Get(
		j,
		"approxPercentile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ApproxPercentileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approxPercentileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) Avg() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionAvgOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionAvgOutputReference
	_jsii_.Get(
		j,
		"avg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) AvgInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"avgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) CountFunction() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionCountFunctionOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionCountFunctionOutputReference
	_jsii_.Get(
		j,
		"countFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) CountFunctionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"countFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) First() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionFirstOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionFirstOutputReference
	_jsii_.Get(
		j,
		"first",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) FirstInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) Last() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionLastOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionLastOutputReference
	_jsii_.Get(
		j,
		"last",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) LastInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) Max() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionMaxOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionMaxOutputReference
	_jsii_.Get(
		j,
		"max",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) MaxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) Min() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionMinOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionMinOutputReference
	_jsii_.Get(
		j,
		"min",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) MinInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"minInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) StddevPop() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionStddevPopOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionStddevPopOutputReference
	_jsii_.Get(
		j,
		"stddevPop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) StddevPopInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stddevPopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) StddevSamp() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionStddevSampOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionStddevSampOutputReference
	_jsii_.Get(
		j,
		"stddevSamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) StddevSampInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stddevSampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) Sum() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionSumOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionSumOutputReference
	_jsii_.Get(
		j,
		"sum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) SumInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) TimeWindow() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowOutputReference
	_jsii_.Get(
		j,
		"timeWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) TimeWindowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) VarPop() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionVarPopOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionVarPopOutputReference
	_jsii_.Get(
		j,
		"varPop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) VarPopInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"varPopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) VarSamp() DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionVarSampOutputReference {
	var returns DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionVarSampOutputReference
	_jsii_.Get(
		j,
		"varSamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) VarSampInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"varSampInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringFeatures.DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference_Override(d DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringFeatures.DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) PutApproxCountDistinct(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionApproxCountDistinct) {
	if err := d.validatePutApproxCountDistinctParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApproxCountDistinct",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) PutApproxPercentile(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionApproxPercentile) {
	if err := d.validatePutApproxPercentileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApproxPercentile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) PutAvg(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionAvg) {
	if err := d.validatePutAvgParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAvg",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) PutCountFunction(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionCountFunction) {
	if err := d.validatePutCountFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCountFunction",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) PutFirst(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionFirst) {
	if err := d.validatePutFirstParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFirst",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) PutLast(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionLast) {
	if err := d.validatePutLastParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLast",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) PutMax(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionMax) {
	if err := d.validatePutMaxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMax",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) PutMin(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionMin) {
	if err := d.validatePutMinParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMin",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) PutStddevPop(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionStddevPop) {
	if err := d.validatePutStddevPopParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStddevPop",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) PutStddevSamp(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionStddevSamp) {
	if err := d.validatePutStddevSampParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStddevSamp",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) PutSum(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionSum) {
	if err := d.validatePutSumParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSum",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) PutTimeWindow(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindow) {
	if err := d.validatePutTimeWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeWindow",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) PutVarPop(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionVarPop) {
	if err := d.validatePutVarPopParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVarPop",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) PutVarSamp(value *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionVarSamp) {
	if err := d.validatePutVarSampParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVarSamp",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ResetApproxCountDistinct() {
	_jsii_.InvokeVoid(
		d,
		"resetApproxCountDistinct",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ResetApproxPercentile() {
	_jsii_.InvokeVoid(
		d,
		"resetApproxPercentile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ResetAvg() {
	_jsii_.InvokeVoid(
		d,
		"resetAvg",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ResetCountFunction() {
	_jsii_.InvokeVoid(
		d,
		"resetCountFunction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ResetFirst() {
	_jsii_.InvokeVoid(
		d,
		"resetFirst",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ResetLast() {
	_jsii_.InvokeVoid(
		d,
		"resetLast",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ResetMax() {
	_jsii_.InvokeVoid(
		d,
		"resetMax",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ResetMin() {
	_jsii_.InvokeVoid(
		d,
		"resetMin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ResetStddevPop() {
	_jsii_.InvokeVoid(
		d,
		"resetStddevPop",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ResetStddevSamp() {
	_jsii_.InvokeVoid(
		d,
		"resetStddevSamp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ResetSum() {
	_jsii_.InvokeVoid(
		d,
		"resetSum",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ResetTimeWindow() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeWindow",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ResetVarPop() {
	_jsii_.InvokeVoid(
		d,
		"resetVarPop",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ResetVarSamp() {
	_jsii_.InvokeVoid(
		d,
		"resetVarSamp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

