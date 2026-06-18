// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfigs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabricksfeatureengineeringkafkaconfigs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference interface {
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
	DeltaTableName() *string
	SetDeltaTableName(val *string)
	DeltaTableNameInput() *string
	DeltaTableSource() DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference
	DeltaTableSourceInput() interface{}
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
	PutDeltaTableSource(value *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSource)
	ResetDeltaTableName()
	ResetDeltaTableSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference
type jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) DeltaTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) DeltaTableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaTableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) DeltaTableSource() DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference {
	var returns DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSourceOutputReference
	_jsii_.Get(
		j,
		"deltaTableSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) DeltaTableSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deltaTableSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringKafkaConfigs.DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference_Override(d DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringKafkaConfigs.DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference)SetDeltaTableName(val *string) {
	if err := j.validateSetDeltaTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deltaTableName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) PutDeltaTableSource(value *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSource) {
	if err := d.validatePutDeltaTableSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDeltaTableSource",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) ResetDeltaTableName() {
	_jsii_.InvokeVoid(
		d,
		"resetDeltaTableName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) ResetDeltaTableSource() {
	_jsii_.InvokeVoid(
		d,
		"resetDeltaTableSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

