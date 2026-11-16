// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountsettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/accountsettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference interface {
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
	CustomDiverging() AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference
	CustomDivergingInput() interface{}
	CustomSequential() AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomSequentialOutputReference
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
	PutCustomDiverging(value *AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDiverging)
	PutCustomSequential(value *AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomSequential)
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

// The jsii proxy struct for AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference
type jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) Color() *string {
	var returns *string
	_jsii_.Get(
		j,
		"color",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ColorScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ColorSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) CustomDiverging() AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference {
	var returns AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDivergingOutputReference
	_jsii_.Get(
		j,
		"customDiverging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) CustomDivergingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customDivergingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) CustomSequential() AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomSequentialOutputReference {
	var returns AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomSequentialOutputReference
	_jsii_.Get(
		j,
		"customSequential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) CustomSequentialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customSequentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) FromCategoricalIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromCategoricalIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) FromCategoricalIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromCategoricalIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference {
	_init_.Initialize()

	if err := validateNewAccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.accountSettingV2.AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference_Override(a AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.accountSettingV2.AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference)SetColor(val *string) {
	if err := j.validateSetColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"color",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference)SetColorScheme(val *string) {
	if err := j.validateSetColorSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"colorScheme",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference)SetFromCategoricalIndex(val *float64) {
	if err := j.validateSetFromCategoricalIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fromCategoricalIndex",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) PutCustomDiverging(value *AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDiverging) {
	if err := a.validatePutCustomDivergingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomDiverging",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) PutCustomSequential(value *AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomSequential) {
	if err := a.validatePutCustomSequentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomSequential",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ResetColor() {
	_jsii_.InvokeVoid(
		a,
		"resetColor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ResetColorScheme() {
	_jsii_.InvokeVoid(
		a,
		"resetColorScheme",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ResetCustomDiverging() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomDiverging",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ResetCustomSequential() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomSequential",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ResetFromCategoricalIndex() {
	_jsii_.InvokeVoid(
		a,
		"resetFromCategoricalIndex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeContinuousGradientLightModeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

