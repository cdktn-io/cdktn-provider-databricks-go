// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountsettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabricksaccountsettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference interface {
	cdktn.ComplexObject
	Color() *string
	SetColor(val *string)
	ColorInput() *string
	ColorScheme() *string
	SetColorScheme(val *string)
	ColorSchemeInput() *string
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
	CustomDiverging() DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference
	CustomDivergingInput() interface{}
	CustomSequential() DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference
	CustomSequentialInput() interface{}
	// Experimental.
	Fqn() *string
	FromCategoricalIndex() *float64
	SetFromCategoricalIndex(val *float64)
	FromCategoricalIndexInput() *float64
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
	PutCustomDiverging(value *DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDiverging)
	PutCustomSequential(value *DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequential)
	ResetColor()
	ResetColorScheme()
	ResetCustomDiverging()
	ResetCustomSequential()
	ResetFromCategoricalIndex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference
type jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) Color() *string {
	var returns *string
	_jsii_.Get(
		j,
		"color",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ColorScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ColorSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) CustomDiverging() DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference {
	var returns DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference
	_jsii_.Get(
		j,
		"customDiverging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) CustomDivergingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customDivergingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) CustomSequential() DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference {
	var returns DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference
	_jsii_.Get(
		j,
		"customSequential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) CustomSequentialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customSequentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) FromCategoricalIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromCategoricalIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) FromCategoricalIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromCategoricalIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountSettingV2.DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference_Override(d DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountSettingV2.DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference)SetColor(val *string) {
	if err := j.validateSetColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"color",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference)SetColorScheme(val *string) {
	if err := j.validateSetColorSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"colorScheme",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference)SetFromCategoricalIndex(val *float64) {
	if err := j.validateSetFromCategoricalIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fromCategoricalIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) PutCustomDiverging(value *DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDiverging) {
	if err := d.validatePutCustomDivergingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomDiverging",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) PutCustomSequential(value *DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequential) {
	if err := d.validatePutCustomSequentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomSequential",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ResetColor() {
	_jsii_.InvokeVoid(
		d,
		"resetColor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ResetColorScheme() {
	_jsii_.InvokeVoid(
		d,
		"resetColorScheme",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ResetCustomDiverging() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomDiverging",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ResetCustomSequential() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomSequential",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ResetFromCategoricalIndex() {
	_jsii_.InvokeVoid(
		d,
		"resetFromCategoricalIndex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

