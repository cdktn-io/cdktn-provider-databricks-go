// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksaigatewaymodelservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAiGatewayModelServiceConfigAOutputReference interface {
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
	InferenceTable() DataDatabricksAiGatewayModelServiceConfigInferenceTableOutputReference
	InferenceTableInput() interface{}
	InternalValue() *DataDatabricksAiGatewayModelServiceConfigA
	SetInternalValue(val *DataDatabricksAiGatewayModelServiceConfigA)
	RateLimits() DataDatabricksAiGatewayModelServiceConfigRateLimitsList
	RateLimitsInput() interface{}
	Routing() DataDatabricksAiGatewayModelServiceConfigRoutingOutputReference
	RoutingInput() interface{}
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
	PutInferenceTable(value *DataDatabricksAiGatewayModelServiceConfigInferenceTable)
	PutRateLimits(value interface{})
	PutRouting(value *DataDatabricksAiGatewayModelServiceConfigRouting)
	ResetInferenceTable()
	ResetRateLimits()
	ResetRouting()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAiGatewayModelServiceConfigAOutputReference
type jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) InferenceTable() DataDatabricksAiGatewayModelServiceConfigInferenceTableOutputReference {
	var returns DataDatabricksAiGatewayModelServiceConfigInferenceTableOutputReference
	_jsii_.Get(
		j,
		"inferenceTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) InferenceTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) InternalValue() *DataDatabricksAiGatewayModelServiceConfigA {
	var returns *DataDatabricksAiGatewayModelServiceConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) RateLimits() DataDatabricksAiGatewayModelServiceConfigRateLimitsList {
	var returns DataDatabricksAiGatewayModelServiceConfigRateLimitsList
	_jsii_.Get(
		j,
		"rateLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) RateLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rateLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) Routing() DataDatabricksAiGatewayModelServiceConfigRoutingOutputReference {
	var returns DataDatabricksAiGatewayModelServiceConfigRoutingOutputReference
	_jsii_.Get(
		j,
		"routing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) RoutingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAiGatewayModelServiceConfigAOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAiGatewayModelServiceConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAiGatewayModelServiceConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiGatewayModelService.DataDatabricksAiGatewayModelServiceConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAiGatewayModelServiceConfigAOutputReference_Override(d DataDatabricksAiGatewayModelServiceConfigAOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiGatewayModelService.DataDatabricksAiGatewayModelServiceConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference)SetInternalValue(val *DataDatabricksAiGatewayModelServiceConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) PutInferenceTable(value *DataDatabricksAiGatewayModelServiceConfigInferenceTable) {
	if err := d.validatePutInferenceTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInferenceTable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) PutRateLimits(value interface{}) {
	if err := d.validatePutRateLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRateLimits",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) PutRouting(value *DataDatabricksAiGatewayModelServiceConfigRouting) {
	if err := d.validatePutRoutingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRouting",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) ResetInferenceTable() {
	_jsii_.InvokeVoid(
		d,
		"resetInferenceTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) ResetRateLimits() {
	_jsii_.InvokeVoid(
		d,
		"resetRateLimits",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) ResetRouting() {
	_jsii_.InvokeVoid(
		d,
		"resetRouting",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServiceConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

