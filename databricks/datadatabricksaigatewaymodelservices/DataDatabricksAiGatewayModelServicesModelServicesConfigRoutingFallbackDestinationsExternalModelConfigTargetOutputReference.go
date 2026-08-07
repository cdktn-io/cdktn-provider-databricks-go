// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelservices

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksaigatewaymodelservices/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference interface {
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
	InternalValue() *DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTarget
	SetInternalValue(val *DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTarget)
	Model() *string
	SetModel(val *string)
	ModelInput() *string
	NativeApiTypes() *[]*string
	SetNativeApiTypes(val *[]*string)
	NativeApiTypesInput() *[]*string
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
	ResetNativeApiTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference
type jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) InternalValue() *DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTarget {
	var returns *DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTarget
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) Model() *string {
	var returns *string
	_jsii_.Get(
		j,
		"model",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) ModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) NativeApiTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nativeApiTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) NativeApiTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nativeApiTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiGatewayModelServices.DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference_Override(d DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiGatewayModelServices.DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference)SetInternalValue(val *DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTarget) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference)SetModel(val *string) {
	if err := j.validateSetModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"model",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference)SetNativeApiTypes(val *[]*string) {
	if err := j.validateSetNativeApiTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nativeApiTypes",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) ResetNativeApiTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetNativeApiTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTargetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

