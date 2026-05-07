// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabricksaccountnetworkpolicies/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpRanges() *[]*string
	SetIpRanges(val *[]*string)
	IpRangesInput() *[]*string
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
	ResetIpRanges()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference
type jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) IpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) IpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference_Override(d DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference)SetIpRanges(val *[]*string) {
	if err := j.validateSetIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipRanges",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) ResetIpRanges() {
	_jsii_.InvokeVoid(
		d,
		"resetIpRanges",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

