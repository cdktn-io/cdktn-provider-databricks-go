// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountsettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabricksaccountsettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference interface {
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
	Dark() *string
	SetDark(val *string)
	DarkInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Light() *string
	SetLight(val *string)
	LightInput() *string
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
	ResetDark()
	ResetLight()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference
type jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) Dark() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dark",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) DarkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"darkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) Light() *string {
	var returns *string
	_jsii_.Get(
		j,
		"light",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) LightInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountSettingV2.DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference_Override(d DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAccountSettingV2.DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference)SetDark(val *string) {
	if err := j.validateSetDarkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dark",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference)SetLight(val *string) {
	if err := j.validateSetLightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"light",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) ResetDark() {
	_jsii_.InvokeVoid(
		d,
		"resetDark",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) ResetLight() {
	_jsii_.InvokeVoid(
		d,
		"resetLight",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitleFontColorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

