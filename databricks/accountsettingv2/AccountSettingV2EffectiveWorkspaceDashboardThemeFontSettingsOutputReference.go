// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountsettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/accountsettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference interface {
	cdktn.ComplexObject
	Base() AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference
	BaseInput() interface{}
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
	FieldTitle() AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldTitleOutputReference
	FieldTitleInput() interface{}
	FieldValue() AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldValueOutputReference
	FieldValueInput() interface{}
	// Experimental.
	Fqn() *string
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
	WidgetDescription() AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetDescriptionOutputReference
	WidgetDescriptionInput() interface{}
	WidgetTitle() AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference
	WidgetTitleInput() interface{}
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
	PutBase(value *AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBase)
	PutFieldTitle(value *AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldTitle)
	PutFieldValue(value *AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldValue)
	PutWidgetDescription(value *AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetDescription)
	PutWidgetTitle(value *AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitle)
	ResetBase()
	ResetFieldTitle()
	ResetFieldValue()
	ResetWidgetDescription()
	ResetWidgetTitle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference
type jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) Base() AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference {
	var returns AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBaseOutputReference
	_jsii_.Get(
		j,
		"base",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) BaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"baseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) FieldTitle() AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldTitleOutputReference {
	var returns AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldTitleOutputReference
	_jsii_.Get(
		j,
		"fieldTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) FieldTitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldTitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) FieldValue() AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldValueOutputReference {
	var returns AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldValueOutputReference
	_jsii_.Get(
		j,
		"fieldValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) FieldValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) WidgetDescription() AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetDescriptionOutputReference {
	var returns AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetDescriptionOutputReference
	_jsii_.Get(
		j,
		"widgetDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) WidgetDescriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) WidgetTitle() AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference {
	var returns AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitleOutputReference
	_jsii_.Get(
		j,
		"widgetTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) WidgetTitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetTitleInput",
		&returns,
	)
	return returns
}


func NewAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.accountSettingV2.AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference_Override(a AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.accountSettingV2.AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) PutBase(value *AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsBase) {
	if err := a.validatePutBaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBase",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) PutFieldTitle(value *AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldTitle) {
	if err := a.validatePutFieldTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFieldTitle",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) PutFieldValue(value *AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsFieldValue) {
	if err := a.validatePutFieldValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFieldValue",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) PutWidgetDescription(value *AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetDescription) {
	if err := a.validatePutWidgetDescriptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWidgetDescription",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) PutWidgetTitle(value *AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsWidgetTitle) {
	if err := a.validatePutWidgetTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWidgetTitle",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ResetBase() {
	_jsii_.InvokeVoid(
		a,
		"resetBase",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ResetFieldTitle() {
	_jsii_.InvokeVoid(
		a,
		"resetFieldTitle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ResetFieldValue() {
	_jsii_.InvokeVoid(
		a,
		"resetFieldValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ResetWidgetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetWidgetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ResetWidgetTitle() {
	_jsii_.InvokeVoid(
		a,
		"resetWidgetTitle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

