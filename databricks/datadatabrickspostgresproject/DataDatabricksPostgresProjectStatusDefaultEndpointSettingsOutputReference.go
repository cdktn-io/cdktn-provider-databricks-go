// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/datadatabrickspostgresproject/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference interface {
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
	InternalValue() *DataDatabricksPostgresProjectStatusDefaultEndpointSettings
	SetInternalValue(val *DataDatabricksPostgresProjectStatusDefaultEndpointSettings)
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

// The jsii proxy struct for DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference
type jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) AutoscalingLimitMaxCu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscalingLimitMaxCu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) AutoscalingLimitMaxCuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscalingLimitMaxCuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) AutoscalingLimitMinCu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscalingLimitMinCu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) AutoscalingLimitMinCuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscalingLimitMinCuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) InternalValue() *DataDatabricksPostgresProjectStatusDefaultEndpointSettings {
	var returns *DataDatabricksPostgresProjectStatusDefaultEndpointSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) NoSuspension() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSuspension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) NoSuspensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSuspensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) PgSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"pgSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) PgSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"pgSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) SuspendTimeoutDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendTimeoutDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) SuspendTimeoutDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendTimeoutDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresProject.DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference_Override(d DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresProject.DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference)SetAutoscalingLimitMaxCu(val *float64) {
	if err := j.validateSetAutoscalingLimitMaxCuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingLimitMaxCu",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference)SetAutoscalingLimitMinCu(val *float64) {
	if err := j.validateSetAutoscalingLimitMinCuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingLimitMinCu",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference)SetInternalValue(val *DataDatabricksPostgresProjectStatusDefaultEndpointSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference)SetNoSuspension(val interface{}) {
	if err := j.validateSetNoSuspensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSuspension",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference)SetPgSettings(val *map[string]*string) {
	if err := j.validateSetPgSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pgSettings",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference)SetSuspendTimeoutDuration(val *string) {
	if err := j.validateSetSuspendTimeoutDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendTimeoutDuration",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) ResetAutoscalingLimitMaxCu() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoscalingLimitMaxCu",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) ResetAutoscalingLimitMinCu() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoscalingLimitMinCu",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) ResetNoSuspension() {
	_jsii_.InvokeVoid(
		d,
		"resetNoSuspension",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) ResetPgSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetPgSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) ResetSuspendTimeoutDuration() {
	_jsii_.InvokeVoid(
		d,
		"resetSuspendTimeoutDuration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectStatusDefaultEndpointSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

