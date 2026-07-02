// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresprojects

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabrickspostgresprojects/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference interface {
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
	Group() DataDatabricksPostgresProjectsProjectsInitialEndpointSpecGroupOutputReference
	GroupInput() interface{}
	InternalValue() *DataDatabricksPostgresProjectsProjectsInitialEndpointSpec
	SetInternalValue(val *DataDatabricksPostgresProjectsProjectsInitialEndpointSpec)
	NoSuspension() interface{}
	SetNoSuspension(val interface{})
	NoSuspensionInput() interface{}
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
	PutGroup(value *DataDatabricksPostgresProjectsProjectsInitialEndpointSpecGroup)
	ResetAutoscalingLimitMaxCu()
	ResetAutoscalingLimitMinCu()
	ResetGroup()
	ResetNoSuspension()
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

// The jsii proxy struct for DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference
type jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) AutoscalingLimitMaxCu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscalingLimitMaxCu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) AutoscalingLimitMaxCuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscalingLimitMaxCuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) AutoscalingLimitMinCu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscalingLimitMinCu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) AutoscalingLimitMinCuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscalingLimitMinCuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) Group() DataDatabricksPostgresProjectsProjectsInitialEndpointSpecGroupOutputReference {
	var returns DataDatabricksPostgresProjectsProjectsInitialEndpointSpecGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) GroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) InternalValue() *DataDatabricksPostgresProjectsProjectsInitialEndpointSpec {
	var returns *DataDatabricksPostgresProjectsProjectsInitialEndpointSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) NoSuspension() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSuspension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) NoSuspensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSuspensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) SuspendTimeoutDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendTimeoutDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) SuspendTimeoutDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendTimeoutDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresProjects.DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference_Override(d DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresProjects.DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference)SetAutoscalingLimitMaxCu(val *float64) {
	if err := j.validateSetAutoscalingLimitMaxCuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingLimitMaxCu",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference)SetAutoscalingLimitMinCu(val *float64) {
	if err := j.validateSetAutoscalingLimitMinCuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingLimitMinCu",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference)SetInternalValue(val *DataDatabricksPostgresProjectsProjectsInitialEndpointSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference)SetNoSuspension(val interface{}) {
	if err := j.validateSetNoSuspensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSuspension",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference)SetSuspendTimeoutDuration(val *string) {
	if err := j.validateSetSuspendTimeoutDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendTimeoutDuration",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) PutGroup(value *DataDatabricksPostgresProjectsProjectsInitialEndpointSpecGroup) {
	if err := d.validatePutGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGroup",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) ResetAutoscalingLimitMaxCu() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoscalingLimitMaxCu",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) ResetAutoscalingLimitMinCu() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoscalingLimitMinCu",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		d,
		"resetGroup",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) ResetNoSuspension() {
	_jsii_.InvokeVoid(
		d,
		"resetNoSuspension",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) ResetSuspendTimeoutDuration() {
	_jsii_.InvokeVoid(
		d,
		"resetSuspendTimeoutDuration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresProjectsProjectsInitialEndpointSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

