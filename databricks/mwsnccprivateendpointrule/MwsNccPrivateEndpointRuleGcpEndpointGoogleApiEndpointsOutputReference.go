// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mwsnccprivateendpointrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/mwsnccprivateendpointrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference interface {
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
	Endpoints() *[]*string
	SetEndpoints(val *[]*string)
	EndpointsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpoints
	SetInternalValue(val *MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpoints)
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
	ResetEndpoints()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference
type jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) Endpoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) EndpointsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"endpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) InternalValue() *MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpoints {
	var returns *MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpoints
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference {
	_init_.Initialize()

	if err := validateNewMwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.mwsNccPrivateEndpointRule.MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference_Override(m MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.mwsNccPrivateEndpointRule.MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference)SetEndpoints(val *[]*string) {
	if err := j.validateSetEndpointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoints",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference)SetInternalValue(val *MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpoints) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) ResetEndpoints() {
	_jsii_.InvokeVoid(
		m,
		"resetEndpoints",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsNccPrivateEndpointRuleGcpEndpointGoogleApiEndpointsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

