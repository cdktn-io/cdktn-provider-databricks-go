// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabricksfeatureengineeringkafkaconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference interface {
	cdktn.ComplexObject
	BackfillJobId() *float64
	BackfillSource() DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigBackfillSourceOutputReference
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
	IngestionDestination() DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigIngestionDestinationOutputReference
	IngestionDestinationInput() *DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigIngestionDestination
	IngestionJobId() *float64
	IngestionPipelineId() *string
	InternalValue() *DataDatabricksFeatureEngineeringKafkaConfigIngestionConfig
	SetInternalValue(val *DataDatabricksFeatureEngineeringKafkaConfigIngestionConfig)
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
	PutBackfillSource(value *DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigBackfillSource)
	PutIngestionDestination(value *DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigIngestionDestination)
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

// The jsii proxy struct for DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference
type jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) BackfillJobId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backfillJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) BackfillSource() DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigBackfillSourceOutputReference {
	var returns DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigBackfillSourceOutputReference
	_jsii_.Get(
		j,
		"backfillSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) BackfillSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backfillSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) DeduplicationColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deduplicationColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) DeduplicationColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deduplicationColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) IngestionDestination() DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigIngestionDestinationOutputReference {
	var returns DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigIngestionDestinationOutputReference
	_jsii_.Get(
		j,
		"ingestionDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) IngestionDestinationInput() *DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigIngestionDestination {
	var returns *DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigIngestionDestination
	_jsii_.Get(
		j,
		"ingestionDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) IngestionJobId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingestionJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) IngestionPipelineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingestionPipelineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) InternalValue() *DataDatabricksFeatureEngineeringKafkaConfigIngestionConfig {
	var returns *DataDatabricksFeatureEngineeringKafkaConfigIngestionConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringKafkaConfig.DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference_Override(d DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringKafkaConfig.DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference)SetDeduplicationColumns(val *[]*string) {
	if err := j.validateSetDeduplicationColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deduplicationColumns",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference)SetInternalValue(val *DataDatabricksFeatureEngineeringKafkaConfigIngestionConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) PutBackfillSource(value *DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigBackfillSource) {
	if err := d.validatePutBackfillSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBackfillSource",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) PutIngestionDestination(value *DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigIngestionDestination) {
	if err := d.validatePutIngestionDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIngestionDestination",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) ResetBackfillSource() {
	_jsii_.InvokeVoid(
		d,
		"resetBackfillSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) ResetDeduplicationColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetDeduplicationColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

