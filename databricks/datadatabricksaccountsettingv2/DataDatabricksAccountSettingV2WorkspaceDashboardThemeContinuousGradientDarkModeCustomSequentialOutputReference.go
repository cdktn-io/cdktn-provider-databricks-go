// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountsettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabricksaccountsettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference interface {
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
	End() DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialEndOutputReference
	EndInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Start() DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialStartOutputReference
	StartInput() interface{}
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
	PutEnd(value *DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialEnd)
	PutStart(value *DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialStart)
	ResetEnd()
	ResetStart()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference
type jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) End() DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialEndOutputReference {
	var returns DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialEndOutputReference
	_jsii_.Get(
		j,
		"end",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) EndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) Start() DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialStartOutputReference {
	var returns DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialStartOutputReference
	_jsii_.Get(
		j,
		"start",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) StartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountSettingV2.DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference_Override(d DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountSettingV2.DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) PutEnd(value *DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialEnd) {
	if err := d.validatePutEndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEnd",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) PutStart(value *DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialStart) {
	if err := d.validatePutStartParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStart",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) ResetEnd() {
	_jsii_.InvokeVoid(
		d,
		"resetEnd",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) ResetStart() {
	_jsii_.InvokeVoid(
		d,
		"resetStart",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

