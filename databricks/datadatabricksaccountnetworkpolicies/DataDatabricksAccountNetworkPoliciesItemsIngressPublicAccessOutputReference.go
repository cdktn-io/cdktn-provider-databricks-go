// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabricksaccountnetworkpolicies/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference interface {
	cdktn.ComplexObject
	AllowRules() DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesList
	AllowRulesInput() interface{}
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
	DenyRules() DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesList
	DenyRulesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RestrictionMode() *string
	SetRestrictionMode(val *string)
	RestrictionModeInput() *string
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
	PutAllowRules(value interface{})
	PutDenyRules(value interface{})
	ResetAllowRules()
	ResetDenyRules()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference
type jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) AllowRules() DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesList {
	var returns DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesList
	_jsii_.Get(
		j,
		"allowRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) AllowRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) DenyRules() DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesList {
	var returns DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesList
	_jsii_.Get(
		j,
		"denyRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) DenyRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"denyRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) RestrictionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restrictionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) RestrictionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restrictionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference_Override(d DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference)SetRestrictionMode(val *string) {
	if err := j.validateSetRestrictionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictionMode",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) PutAllowRules(value interface{}) {
	if err := d.validatePutAllowRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAllowRules",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) PutDenyRules(value interface{}) {
	if err := d.validatePutDenyRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDenyRules",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) ResetAllowRules() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowRules",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) ResetDenyRules() {
	_jsii_.InvokeVoid(
		d,
		"resetDenyRules",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

