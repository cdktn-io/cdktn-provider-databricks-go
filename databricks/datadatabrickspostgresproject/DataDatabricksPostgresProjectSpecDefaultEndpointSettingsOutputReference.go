// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabrickspostgresproject/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference interface {
	cdktn.ComplexObject
	AutoscalingLimitMaxCu() *float64
	SetAutoscalingLimitMaxCu(val *float64)
	AutoscalingLimitMaxCuInput() *float64
	AutoscalingLimitMinCu() *float64
	SetAutoscalingLimitMinCu(val *float64)
	AutoscalingLimitMinCuInput() *float64
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
	NoSuspension() interface{}
	SetNoSuspension(val interface{})
	NoSuspensionInput() interface{}
	PgSettings() *map[string]*string
	SetPgSettings(val *map[string]*string)
	PgSettingsInput() *map[string]*string
	SuspendTimeoutDuration() *string
	SetSuspendTimeoutDuration(val *string)
	SuspendTimeoutDurationInput() *string
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
	ResetAutoscalingLimitMaxCu()
	ResetAutoscalingLimitMinCu()
	ResetNoSuspension()
	ResetPgSettings()
	ResetSuspendTimeoutDuration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference
type jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) AutoscalingLimitMaxCu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscalingLimitMaxCu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) AutoscalingLimitMaxCuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscalingLimitMaxCuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) AutoscalingLimitMinCu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscalingLimitMinCu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) AutoscalingLimitMinCuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscalingLimitMinCuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) NoSuspension() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSuspension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) NoSuspensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSuspensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) PgSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"pgSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) PgSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"pgSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) SuspendTimeoutDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendTimeoutDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) SuspendTimeoutDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendTimeoutDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresProject.DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference_Override(d DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresProject.DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference)SetAutoscalingLimitMaxCu(val *float64) {
	if err := j.validateSetAutoscalingLimitMaxCuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingLimitMaxCu",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference)SetAutoscalingLimitMinCu(val *float64) {
	if err := j.validateSetAutoscalingLimitMinCuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingLimitMinCu",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference)SetNoSuspension(val interface{}) {
	if err := j.validateSetNoSuspensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSuspension",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference)SetPgSettings(val *map[string]*string) {
	if err := j.validateSetPgSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pgSettings",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference)SetSuspendTimeoutDuration(val *string) {
	if err := j.validateSetSuspendTimeoutDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendTimeoutDuration",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) ResetAutoscalingLimitMaxCu() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoscalingLimitMaxCu",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) ResetAutoscalingLimitMinCu() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoscalingLimitMinCu",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) ResetNoSuspension() {
	_jsii_.InvokeVoid(
		d,
		"resetNoSuspension",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) ResetPgSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetPgSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) ResetSuspendTimeoutDuration() {
	_jsii_.InvokeVoid(
		d,
		"resetSuspendTimeoutDuration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectSpecDefaultEndpointSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

