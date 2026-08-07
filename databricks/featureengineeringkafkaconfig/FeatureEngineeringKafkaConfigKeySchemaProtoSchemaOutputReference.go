// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringkafkaconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/featureengineeringkafkaconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MessageName() *string
	SetMessageName(val *string)
	MessageNameInput() *string
	SchemaText() *string
	SetSchemaText(val *string)
	SchemaTextInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference
type jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) MessageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) MessageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) SchemaText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) SchemaTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference {
	_init_.Initialize()

	if err := validateNewFeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.featureEngineeringKafkaConfig.FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference_Override(f FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.featureEngineeringKafkaConfig.FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference)SetMessageName(val *string) {
	if err := j.validateSetMessageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageName",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference)SetSchemaText(val *string) {
	if err := j.validateSetSchemaTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaText",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (f *jsiiProxy_FeatureEngineeringKafkaConfigKeySchemaProtoSchemaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

