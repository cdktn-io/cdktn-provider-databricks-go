// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountsettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/accountsettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference interface {
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
	DarkMode() AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference
	DarkModeInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LightMode() AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference
	LightModeInput() interface{}
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
	PutDarkMode(value *AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkMode)
	PutLightMode(value *AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightMode)
	ResetDarkMode()
	ResetLightMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference
type jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) DarkMode() AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference {
	var returns AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeOutputReference
	_jsii_.Get(
		j,
		"darkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) DarkModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"darkModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) LightMode() AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference {
	var returns AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeOutputReference
	_jsii_.Get(
		j,
		"lightMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) LightModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lightModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference {
	_init_.Initialize()

	if err := validateNewAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.accountSettingV2.AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference_Override(a AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.accountSettingV2.AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) PutDarkMode(value *AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkMode) {
	if err := a.validatePutDarkModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDarkMode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) PutLightMode(value *AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightMode) {
	if err := a.validatePutLightModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLightMode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) ResetDarkMode() {
	_jsii_.InvokeVoid(
		a,
		"resetDarkMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) ResetLightMode() {
	_jsii_.InvokeVoid(
		a,
		"resetLightMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

