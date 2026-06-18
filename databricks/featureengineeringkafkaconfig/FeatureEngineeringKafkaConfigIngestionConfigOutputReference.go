// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringkafkaconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/featureengineeringkafkaconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FeatureEngineeringKafkaConfigIngestionConfigOutputReference interface {
	cdktn.ComplexObject
	BackfillJobId() *float64
	BackfillSource() FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceOutputReference
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
	IngestionDestination() FeatureEngineeringKafkaConfigIngestionConfigIngestionDestinationOutputReference
	IngestionDestinationInput() interface{}
	IngestionJobId() *float64
	IngestionPipelineId() *string
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
	PutBackfillSource(value *FeatureEngineeringKafkaConfigIngestionConfigBackfillSource)
	PutIngestionDestination(value *FeatureEngineeringKafkaConfigIngestionConfigIngestionDestination)
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

// The jsii proxy struct for FeatureEngineeringKafkaConfigIngestionConfigOutputReference
type jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) BackfillJobId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backfillJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) BackfillSource() FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceOutputReference {
	var returns FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceOutputReference
	_jsii_.Get(
		j,
		"backfillSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) BackfillSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backfillSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) DeduplicationColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deduplicationColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) DeduplicationColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deduplicationColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) IngestionDestination() FeatureEngineeringKafkaConfigIngestionConfigIngestionDestinationOutputReference {
	var returns FeatureEngineeringKafkaConfigIngestionConfigIngestionDestinationOutputReference
	_jsii_.Get(
		j,
		"ingestionDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) IngestionDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingestionDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) IngestionJobId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingestionJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) IngestionPipelineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingestionPipelineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFeatureEngineeringKafkaConfigIngestionConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) FeatureEngineeringKafkaConfigIngestionConfigOutputReference {
	_init_.Initialize()

	if err := validateNewFeatureEngineeringKafkaConfigIngestionConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.featureEngineeringKafkaConfig.FeatureEngineeringKafkaConfigIngestionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFeatureEngineeringKafkaConfigIngestionConfigOutputReference_Override(f FeatureEngineeringKafkaConfigIngestionConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.featureEngineeringKafkaConfig.FeatureEngineeringKafkaConfigIngestionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference)SetDeduplicationColumns(val *[]*string) {
	if err := j.validateSetDeduplicationColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deduplicationColumns",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) PutBackfillSource(value *FeatureEngineeringKafkaConfigIngestionConfigBackfillSource) {
	if err := f.validatePutBackfillSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putBackfillSource",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) PutIngestionDestination(value *FeatureEngineeringKafkaConfigIngestionConfigIngestionDestination) {
	if err := f.validatePutIngestionDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putIngestionDestination",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) ResetBackfillSource() {
	_jsii_.InvokeVoid(
		f,
		"resetBackfillSource",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) ResetDeduplicationColumns() {
	_jsii_.InvokeVoid(
		f,
		"resetDeduplicationColumns",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigIngestionConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

