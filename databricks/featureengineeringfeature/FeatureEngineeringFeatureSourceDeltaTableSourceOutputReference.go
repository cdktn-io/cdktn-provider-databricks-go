// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/featureengineeringfeature/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference interface {
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
	ResetFilterCondition()
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

// The jsii proxy struct for FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference
type jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) DataframeSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataframeSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) DataframeSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataframeSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) EntityColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entityColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) EntityColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entityColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) FilterCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) FilterConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) FullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) FullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) TimeseriesColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeseriesColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) TimeseriesColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeseriesColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) TransformationSql() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformationSql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) TransformationSqlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformationSqlInput",
		&returns,
	)
	return returns
}


func NewFeatureEngineeringFeatureSourceDeltaTableSourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference {
	_init_.Initialize()

	if err := validateNewFeatureEngineeringFeatureSourceDeltaTableSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.featureEngineeringFeature.FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFeatureEngineeringFeatureSourceDeltaTableSourceOutputReference_Override(f FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.featureEngineeringFeature.FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference)SetDataframeSchema(val *string) {
	if err := j.validateSetDataframeSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataframeSchema",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference)SetEntityColumns(val *[]*string) {
	if err := j.validateSetEntityColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityColumns",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference)SetFilterCondition(val *string) {
	if err := j.validateSetFilterConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterCondition",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference)SetFullName(val *string) {
	if err := j.validateSetFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullName",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference)SetTimeseriesColumn(val *string) {
	if err := j.validateSetTimeseriesColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeseriesColumn",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference)SetTransformationSql(val *string) {
	if err := j.validateSetTransformationSqlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transformationSql",
		val,
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (f *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (f *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) ResetDataframeSchema() {
	_jsii_.InvokeVoid(
		f,
		"resetDataframeSchema",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) ResetFilterCondition() {
	_jsii_.InvokeVoid(
		f,
		"resetFilterCondition",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) ResetTransformationSql() {
	_jsii_.InvokeVoid(
		f,
		"resetTransformationSql",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (f *jsiiProxy_FeatureEngineeringFeatureSourceDeltaTableSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

