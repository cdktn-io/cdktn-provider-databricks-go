// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringkafkaconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/featureengineeringkafkaconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference interface {
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
	DataframeSchema() *string
	SetDataframeSchema(val *string)
	DataframeSchemaInput() *string
	EntityColumns() *[]*string
	SetEntityColumns(val *[]*string)
	EntityColumnsInput() *[]*string
	FilterCondition() *string
	SetFilterCondition(val *string)
	FilterConditionInput() *string
	// Experimental.
	Fqn() *string
	FullName() *string
	SetFullName(val *string)
	FullNameInput() *string
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
	TimeseriesColumn() *string
	SetTimeseriesColumn(val *string)
	TimeseriesColumnInput() *string
	TransformationSql() *string
	SetTransformationSql(val *string)
	TransformationSqlInput() *string
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
	ResetDataframeSchema()
	ResetEntityColumns()
	ResetFilterCondition()
	ResetTimeseriesColumn()
	ResetTransformationSql()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference
type jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) DataframeSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataframeSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) DataframeSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataframeSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) EntityColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entityColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) EntityColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entityColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) FilterCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) FilterConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) FullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) FullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) TimeseriesColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeseriesColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) TimeseriesColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeseriesColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) TransformationSql() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformationSql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) TransformationSqlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformationSqlInput",
		&returns,
	)
	return returns
}


func NewFeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference {
	_init_.Initialize()

	if err := validateNewFeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.featureEngineeringKafkaConfig.FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference_Override(f FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.featureEngineeringKafkaConfig.FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetDataframeSchema(val *string) {
	if err := j.validateSetDataframeSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataframeSchema",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetEntityColumns(val *[]*string) {
	if err := j.validateSetEntityColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityColumns",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetFilterCondition(val *string) {
	if err := j.validateSetFilterConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterCondition",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetFullName(val *string) {
	if err := j.validateSetFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullName",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetTimeseriesColumn(val *string) {
	if err := j.validateSetTimeseriesColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeseriesColumn",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetTransformationSql(val *string) {
	if err := j.validateSetTransformationSqlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transformationSql",
		val,
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) ResetDataframeSchema() {
	_jsii_.InvokeVoid(
		f,
		"resetDataframeSchema",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) ResetEntityColumns() {
	_jsii_.InvokeVoid(
		f,
		"resetEntityColumns",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) ResetFilterCondition() {
	_jsii_.InvokeVoid(
		f,
		"resetFilterCondition",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) ResetTimeseriesColumn() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeseriesColumn",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) ResetTransformationSql() {
	_jsii_.InvokeVoid(
		f,
		"resetTransformationSql",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

