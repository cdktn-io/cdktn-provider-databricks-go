// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymcpservices

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksaigatewaymcpservices/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference interface {
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
	IncludeToolSelectors() *[]*string
	SetIncludeToolSelectors(val *[]*string)
	IncludeToolSelectorsInput() *[]*string
	InternalValue() *DataDatabricksAiGatewayMcpServicesMcpServicesConfig
	SetInternalValue(val *DataDatabricksAiGatewayMcpServicesMcpServicesConfig)
	RateLimits() DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsList
	RateLimitsInput() interface{}
	SourceConnection() DataDatabricksAiGatewayMcpServicesMcpServicesConfigSourceConnectionOutputReference
	SourceConnectionInput() interface{}
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
	PutRateLimits(value interface{})
	PutSourceConnection(value *DataDatabricksAiGatewayMcpServicesMcpServicesConfigSourceConnection)
	ResetIncludeToolSelectors()
	ResetRateLimits()
	ResetSourceConnection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference
type jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) IncludeToolSelectors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeToolSelectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) IncludeToolSelectorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeToolSelectorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) InternalValue() *DataDatabricksAiGatewayMcpServicesMcpServicesConfig {
	var returns *DataDatabricksAiGatewayMcpServicesMcpServicesConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) RateLimits() DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsList {
	var returns DataDatabricksAiGatewayMcpServicesMcpServicesConfigRateLimitsList
	_jsii_.Get(
		j,
		"rateLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) RateLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rateLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) SourceConnection() DataDatabricksAiGatewayMcpServicesMcpServicesConfigSourceConnectionOutputReference {
	var returns DataDatabricksAiGatewayMcpServicesMcpServicesConfigSourceConnectionOutputReference
	_jsii_.Get(
		j,
		"sourceConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) SourceConnectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceConnectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiGatewayMcpServices.DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference_Override(d DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiGatewayMcpServices.DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference)SetIncludeToolSelectors(val *[]*string) {
	if err := j.validateSetIncludeToolSelectorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeToolSelectors",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference)SetInternalValue(val *DataDatabricksAiGatewayMcpServicesMcpServicesConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) PutRateLimits(value interface{}) {
	if err := d.validatePutRateLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRateLimits",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) PutSourceConnection(value *DataDatabricksAiGatewayMcpServicesMcpServicesConfigSourceConnection) {
	if err := d.validatePutSourceConnectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSourceConnection",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) ResetIncludeToolSelectors() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeToolSelectors",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) ResetRateLimits() {
	_jsii_.InvokeVoid(
		d,
		"resetRateLimits",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) ResetSourceConnection() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceConnection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAiGatewayMcpServicesMcpServicesConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

