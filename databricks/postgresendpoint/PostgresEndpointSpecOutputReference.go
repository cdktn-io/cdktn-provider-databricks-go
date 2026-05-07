// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/postgresendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PostgresEndpointSpecOutputReference interface {
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
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	EndpointType() *string
	SetEndpointType(val *string)
	EndpointTypeInput() *string
	// Experimental.
	Fqn() *string
	Group() PostgresEndpointSpecGroupOutputReference
	GroupInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NoSuspension() interface{}
	SetNoSuspension(val interface{})
	NoSuspensionInput() interface{}
	Settings() PostgresEndpointSpecSettingsOutputReference
	SettingsInput() interface{}
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
	PutGroup(value *PostgresEndpointSpecGroup)
	PutSettings(value *PostgresEndpointSpecSettings)
	ResetAutoscalingLimitMaxCu()
	ResetAutoscalingLimitMinCu()
	ResetDisabled()
	ResetGroup()
	ResetNoSuspension()
	ResetSettings()
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

// The jsii proxy struct for PostgresEndpointSpecOutputReference
type jsiiProxy_PostgresEndpointSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) AutoscalingLimitMaxCu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscalingLimitMaxCu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) AutoscalingLimitMaxCuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscalingLimitMaxCuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) AutoscalingLimitMinCu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscalingLimitMinCu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) AutoscalingLimitMinCuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscalingLimitMinCuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) Group() PostgresEndpointSpecGroupOutputReference {
	var returns PostgresEndpointSpecGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) GroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) NoSuspension() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSuspension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) NoSuspensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSuspensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) Settings() PostgresEndpointSpecSettingsOutputReference {
	var returns PostgresEndpointSpecSettingsOutputReference
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) SettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"settingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) SuspendTimeoutDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendTimeoutDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) SuspendTimeoutDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendTimeoutDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPostgresEndpointSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PostgresEndpointSpecOutputReference {
	_init_.Initialize()

	if err := validateNewPostgresEndpointSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PostgresEndpointSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.postgresEndpoint.PostgresEndpointSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPostgresEndpointSpecOutputReference_Override(p PostgresEndpointSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.postgresEndpoint.PostgresEndpointSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference)SetAutoscalingLimitMaxCu(val *float64) {
	if err := j.validateSetAutoscalingLimitMaxCuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingLimitMaxCu",
		val,
	)
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference)SetAutoscalingLimitMinCu(val *float64) {
	if err := j.validateSetAutoscalingLimitMinCuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingLimitMinCu",
		val,
	)
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference)SetEndpointType(val *string) {
	if err := j.validateSetEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference)SetNoSuspension(val interface{}) {
	if err := j.validateSetNoSuspensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSuspension",
		val,
	)
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference)SetSuspendTimeoutDuration(val *string) {
	if err := j.validateSetSuspendTimeoutDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendTimeoutDuration",
		val,
	)
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PostgresEndpointSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) PutGroup(value *PostgresEndpointSpecGroup) {
	if err := p.validatePutGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGroup",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) PutSettings(value *PostgresEndpointSpecSettings) {
	if err := p.validatePutSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSettings",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) ResetAutoscalingLimitMaxCu() {
	_jsii_.InvokeVoid(
		p,
		"resetAutoscalingLimitMaxCu",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) ResetAutoscalingLimitMinCu() {
	_jsii_.InvokeVoid(
		p,
		"resetAutoscalingLimitMinCu",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) ResetDisabled() {
	_jsii_.InvokeVoid(
		p,
		"resetDisabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		p,
		"resetGroup",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) ResetNoSuspension() {
	_jsii_.InvokeVoid(
		p,
		"resetNoSuspension",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) ResetSettings() {
	_jsii_.InvokeVoid(
		p,
		"resetSettings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) ResetSuspendTimeoutDuration() {
	_jsii_.InvokeVoid(
		p,
		"resetSuspendTimeoutDuration",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresEndpointSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

