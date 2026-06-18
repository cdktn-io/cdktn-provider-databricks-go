// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfigs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabricksfeatureengineeringkafkaconfigs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference interface {
	cdktn.ComplexObject
	BackfillJobId() *float64
	BackfillSource() DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference
	BackfillSourceInput() interface{}
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
	DeduplicationColumns() *[]*string
	SetDeduplicationColumns(val *[]*string)
	DeduplicationColumnsInput() *[]*string
	// Experimental.
	Fqn() *string
	IngestionDestination() DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigIngestionDestinationOutputReference
	IngestionDestinationInput() *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigIngestionDestination
	IngestionJobId() *float64
	IngestionPipelineId() *string
	InternalValue() *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfig
	SetInternalValue(val *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfig)
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
	PutBackfillSource(value *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSource)
	PutIngestionDestination(value *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigIngestionDestination)
	ResetBackfillSource()
	ResetDeduplicationColumns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference
type jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) BackfillJobId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backfillJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) BackfillSource() DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference {
	var returns DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceOutputReference
	_jsii_.Get(
		j,
		"backfillSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) BackfillSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backfillSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) DeduplicationColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deduplicationColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) DeduplicationColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deduplicationColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) IngestionDestination() DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigIngestionDestinationOutputReference {
	var returns DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigIngestionDestinationOutputReference
	_jsii_.Get(
		j,
		"ingestionDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) IngestionDestinationInput() *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigIngestionDestination {
	var returns *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigIngestionDestination
	_jsii_.Get(
		j,
		"ingestionDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) IngestionJobId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingestionJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) IngestionPipelineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingestionPipelineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) InternalValue() *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfig {
	var returns *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringKafkaConfigs.DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference_Override(d DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringKafkaConfigs.DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference)SetDeduplicationColumns(val *[]*string) {
	if err := j.validateSetDeduplicationColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deduplicationColumns",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference)SetInternalValue(val *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) PutBackfillSource(value *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSource) {
	if err := d.validatePutBackfillSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBackfillSource",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) PutIngestionDestination(value *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigIngestionDestination) {
	if err := d.validatePutIngestionDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIngestionDestination",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) ResetBackfillSource() {
	_jsii_.InvokeVoid(
		d,
		"resetBackfillSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) ResetDeduplicationColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetDeduplicationColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

