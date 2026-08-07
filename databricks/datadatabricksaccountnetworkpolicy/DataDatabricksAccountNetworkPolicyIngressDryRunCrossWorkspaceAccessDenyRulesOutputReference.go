// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksaccountnetworkpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference interface {
	cdktn.ComplexObject
	Authentication() DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference
	AuthenticationInput() interface{}
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
	Destination() DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesDestinationOutputReference
	DestinationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	Origin() DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOriginOutputReference
	OriginInput() interface{}
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
	PutAuthentication(value *DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthentication)
	PutDestination(value *DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesDestination)
	PutOrigin(value *DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOrigin)
	ResetAuthentication()
	ResetDestination()
	ResetLabel()
	ResetOrigin()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference
type jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) Authentication() DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference {
	var returns DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) AuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) Destination() DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesDestinationOutputReference {
	var returns DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesDestinationOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) DestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) Origin() DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOriginOutputReference {
	var returns DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOriginOutputReference
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) OriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicy.DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference_Override(d DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicy.DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) PutAuthentication(value *DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesAuthentication) {
	if err := d.validatePutAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAuthentication",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) PutDestination(value *DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesDestination) {
	if err := d.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDestination",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) PutOrigin(value *DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOrigin) {
	if err := d.validatePutOriginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOrigin",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) ResetAuthentication() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) ResetDestination() {
	_jsii_.InvokeVoid(
		d,
		"resetDestination",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		d,
		"resetLabel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) ResetOrigin() {
	_jsii_.InvokeVoid(
		d,
		"resetOrigin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

