// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabricksaccountnetworkpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference interface {
	cdktn.ComplexObject
	AllIpRanges() interface{}
	SetAllIpRanges(val interface{})
	AllIpRangesInput() interface{}
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
	ExcludedIpRanges() DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginExcludedIpRangesOutputReference
	ExcludedIpRangesInput() interface{}
	// Experimental.
	Fqn() *string
	IncludedIpRanges() DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference
	IncludedIpRangesInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutExcludedIpRanges(value *DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginExcludedIpRanges)
	PutIncludedIpRanges(value *DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginIncludedIpRanges)
	ResetAllIpRanges()
	ResetExcludedIpRanges()
	ResetIncludedIpRanges()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference
type jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) AllIpRanges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) AllIpRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) ExcludedIpRanges() DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginExcludedIpRangesOutputReference {
	var returns DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginExcludedIpRangesOutputReference
	_jsii_.Get(
		j,
		"excludedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) ExcludedIpRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) IncludedIpRanges() DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference {
	var returns DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference
	_jsii_.Get(
		j,
		"includedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) IncludedIpRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicy.DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference_Override(d DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicy.DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference)SetAllIpRanges(val interface{}) {
	if err := j.validateSetAllIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allIpRanges",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) PutExcludedIpRanges(value *DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginExcludedIpRanges) {
	if err := d.validatePutExcludedIpRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExcludedIpRanges",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) PutIncludedIpRanges(value *DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginIncludedIpRanges) {
	if err := d.validatePutIncludedIpRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIncludedIpRanges",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) ResetAllIpRanges() {
	_jsii_.InvokeVoid(
		d,
		"resetAllIpRanges",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) ResetExcludedIpRanges() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludedIpRanges",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) ResetIncludedIpRanges() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludedIpRanges",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessDenyRulesOriginOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

