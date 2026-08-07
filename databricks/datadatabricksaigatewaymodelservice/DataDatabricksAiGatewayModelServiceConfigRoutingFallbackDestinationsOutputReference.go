// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksaigatewaymodelservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference interface {
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
	DestinationType() *string
	SetDestinationType(val *string)
	DestinationTypeInput() *string
	ExternalModelConfig() DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsExternalModelConfigOutputReference
	ExternalModelConfigInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsDeleted() cdktn.IResolvable
	Name() *string
	SetName(val *string)
	NameInput() *string
	PayPerTokenConfig() DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsPayPerTokenConfigOutputReference
	PayPerTokenConfigInput() interface{}
	ProvisionedThroughputConfig() DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsProvisionedThroughputConfigOutputReference
	ProvisionedThroughputConfigInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrafficPercentage() *float64
	SetTrafficPercentage(val *float64)
	TrafficPercentageInput() *float64
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
	PutExternalModelConfig(value *DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsExternalModelConfig)
	PutPayPerTokenConfig(value *DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsPayPerTokenConfig)
	PutProvisionedThroughputConfig(value *DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsProvisionedThroughputConfig)
	ResetExternalModelConfig()
	ResetPayPerTokenConfig()
	ResetProvisionedThroughputConfig()
	ResetTrafficPercentage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference
type jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) DestinationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) DestinationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ExternalModelConfig() DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsExternalModelConfigOutputReference {
	var returns DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsExternalModelConfigOutputReference
	_jsii_.Get(
		j,
		"externalModelConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ExternalModelConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalModelConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) IsDeleted() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isDeleted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) PayPerTokenConfig() DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsPayPerTokenConfigOutputReference {
	var returns DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsPayPerTokenConfigOutputReference
	_jsii_.Get(
		j,
		"payPerTokenConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) PayPerTokenConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"payPerTokenConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ProvisionedThroughputConfig() DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsProvisionedThroughputConfigOutputReference {
	var returns DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsProvisionedThroughputConfigOutputReference
	_jsii_.Get(
		j,
		"provisionedThroughputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ProvisionedThroughputConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedThroughputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) TrafficPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trafficPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) TrafficPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trafficPercentageInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiGatewayModelService.DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference_Override(d DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiGatewayModelService.DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference)SetDestinationType(val *string) {
	if err := j.validateSetDestinationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference)SetTrafficPercentage(val *float64) {
	if err := j.validateSetTrafficPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficPercentage",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) PutExternalModelConfig(value *DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsExternalModelConfig) {
	if err := d.validatePutExternalModelConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExternalModelConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) PutPayPerTokenConfig(value *DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsPayPerTokenConfig) {
	if err := d.validatePutPayPerTokenConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPayPerTokenConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) PutProvisionedThroughputConfig(value *DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsProvisionedThroughputConfig) {
	if err := d.validatePutProvisionedThroughputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProvisionedThroughputConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ResetExternalModelConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetExternalModelConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ResetPayPerTokenConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetPayPerTokenConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ResetProvisionedThroughputConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetProvisionedThroughputConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ResetTrafficPercentage() {
	_jsii_.InvokeVoid(
		d,
		"resetTrafficPercentage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

