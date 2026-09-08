// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package policyinfo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/policyinfo/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference interface {
	cdktn.ComplexObject
	ColumnTagValue() PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference
	ColumnTagValueInput() interface{}
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
	TagValue() PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionTagValueOutputReference
	TagValueInput() interface{}
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
	PutColumnTagValue(value *PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValue)
	PutTagValue(value *PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionTagValue)
	ResetColumnTagValue()
	ResetTagValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference
type jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) ColumnTagValue() PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference {
	var returns PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference
	_jsii_.Get(
		j,
		"columnTagValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) ColumnTagValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnTagValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) TagValue() PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionTagValueOutputReference {
	var returns PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionTagValueOutputReference
	_jsii_.Get(
		j,
		"tagValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) TagValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference {
	_init_.Initialize()

	if err := validateNewPolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.policyInfo.PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference_Override(p PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.policyInfo.PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) PutColumnTagValue(value *PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValue) {
	if err := p.validatePutColumnTagValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putColumnTagValue",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) PutTagValue(value *PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionTagValue) {
	if err := p.validatePutTagValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTagValue",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) ResetColumnTagValue() {
	_jsii_.InvokeVoid(
		p,
		"resetColumnTagValue",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) ResetTagValue() {
	_jsii_.InvokeVoid(
		p,
		"resetTagValue",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PolicyInfoColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

