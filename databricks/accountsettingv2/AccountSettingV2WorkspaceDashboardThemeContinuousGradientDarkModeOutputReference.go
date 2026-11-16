// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountsettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/accountsettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference interface {
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
	CustomDiverging() AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference
	CustomDivergingInput() interface{}
	CustomSequential() AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference
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
	PutCustomDiverging(value *AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDiverging)
	PutCustomSequential(value *AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequential)
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

// The jsii proxy struct for AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference
type jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) Color() *string {
	var returns *string
	_jsii_.Get(
		j,
		"color",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ColorScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ColorSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) CustomDiverging() AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference {
	var returns AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDivergingOutputReference
	_jsii_.Get(
		j,
		"customDiverging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) CustomDivergingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customDivergingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) CustomSequential() AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference {
	var returns AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequentialOutputReference
	_jsii_.Get(
		j,
		"customSequential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) CustomSequentialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customSequentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) FromCategoricalIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromCategoricalIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) FromCategoricalIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromCategoricalIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference {
	_init_.Initialize()

	if err := validateNewAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.accountSettingV2.AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference_Override(a AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.accountSettingV2.AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference)SetColor(val *string) {
	if err := j.validateSetColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"color",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference)SetColorScheme(val *string) {
	if err := j.validateSetColorSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"colorScheme",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference)SetFromCategoricalIndex(val *float64) {
	if err := j.validateSetFromCategoricalIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fromCategoricalIndex",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) PutCustomDiverging(value *AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomDiverging) {
	if err := a.validatePutCustomDivergingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomDiverging",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) PutCustomSequential(value *AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeCustomSequential) {
	if err := a.validatePutCustomSequentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomSequential",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ResetColor() {
	_jsii_.InvokeVoid(
		a,
		"resetColor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ResetColorScheme() {
	_jsii_.InvokeVoid(
		a,
		"resetColorScheme",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ResetCustomDiverging() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomDiverging",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ResetCustomSequential() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomSequential",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ResetFromCategoricalIndex() {
	_jsii_.InvokeVoid(
		a,
		"resetFromCategoricalIndex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientDarkModeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

