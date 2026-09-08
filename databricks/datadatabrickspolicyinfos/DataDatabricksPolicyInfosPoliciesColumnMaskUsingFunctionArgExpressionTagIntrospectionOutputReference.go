// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspolicyinfos

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabrickspolicyinfos/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference interface {
	cdktn.ComplexObject
	ColumnTagValue() DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference
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
	TagValue() DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionTagValueOutputReference
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
	PutColumnTagValue(value *DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValue)
	PutTagValue(value *DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionTagValue)
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

// The jsii proxy struct for DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference
type jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) ColumnTagValue() DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference {
	var returns DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValueOutputReference
	_jsii_.Get(
		j,
		"columnTagValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) ColumnTagValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnTagValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) TagValue() DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionTagValueOutputReference {
	var returns DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionTagValueOutputReference
	_jsii_.Get(
		j,
		"tagValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) TagValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPolicyInfos.DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference_Override(d DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPolicyInfos.DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) PutColumnTagValue(value *DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionColumnTagValue) {
	if err := d.validatePutColumnTagValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putColumnTagValue",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) PutTagValue(value *DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionTagValue) {
	if err := d.validatePutTagValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTagValue",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) ResetColumnTagValue() {
	_jsii_.InvokeVoid(
		d,
		"resetColumnTagValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) ResetTagValue() {
	_jsii_.InvokeVoid(
		d,
		"resetTagValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksPolicyInfosPoliciesColumnMaskUsingFunctionArgExpressionTagIntrospectionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

