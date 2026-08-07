// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfigs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksfeatureengineeringkafkaconfigs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference interface {
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

// The jsii proxy struct for DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference
type jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) DataframeSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataframeSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) DataframeSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataframeSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) EntityColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entityColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) EntityColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entityColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) FilterCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) FilterConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) FullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) FullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) TimeseriesColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeseriesColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) TimeseriesColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeseriesColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) TransformationSql() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformationSql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) TransformationSqlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformationSqlInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringKafkaConfigs.DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference_Override(d DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringKafkaConfigs.DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetDataframeSchema(val *string) {
	if err := j.validateSetDataframeSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataframeSchema",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetEntityColumns(val *[]*string) {
	if err := j.validateSetEntityColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityColumns",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetFilterCondition(val *string) {
	if err := j.validateSetFilterConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterCondition",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetFullName(val *string) {
	if err := j.validateSetFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetTimeseriesColumn(val *string) {
	if err := j.validateSetTimeseriesColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeseriesColumn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference)SetTransformationSql(val *string) {
	if err := j.validateSetTransformationSqlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transformationSql",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) ResetDataframeSchema() {
	_jsii_.InvokeVoid(
		d,
		"resetDataframeSchema",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) ResetEntityColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetEntityColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) ResetFilterCondition() {
	_jsii_.InvokeVoid(
		d,
		"resetFilterCondition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) ResetTimeseriesColumn() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeseriesColumn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) ResetTransformationSql() {
	_jsii_.InvokeVoid(
		d,
		"resetTransformationSql",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

