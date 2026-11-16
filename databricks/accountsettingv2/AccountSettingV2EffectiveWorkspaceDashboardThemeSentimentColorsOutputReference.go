// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountsettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/accountsettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Negative() AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsNegativeOutputReference
	NegativeInput() interface{}
	Neutral() AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsNeutralOutputReference
	NeutralInput() interface{}
	Positive() AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsPositiveOutputReference
	PositiveInput() interface{}
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
	PutNegative(value *AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsNegative)
	PutNeutral(value *AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsNeutral)
	PutPositive(value *AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsPositive)
	ResetNegative()
	ResetNeutral()
	ResetPositive()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference
type jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) Negative() AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsNegativeOutputReference {
	var returns AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsNegativeOutputReference
	_jsii_.Get(
		j,
		"negative",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) NegativeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) Neutral() AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsNeutralOutputReference {
	var returns AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsNeutralOutputReference
	_jsii_.Get(
		j,
		"neutral",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) NeutralInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neutralInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) Positive() AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsPositiveOutputReference {
	var returns AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsPositiveOutputReference
	_jsii_.Get(
		j,
		"positive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) PositiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"positiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference {
	_init_.Initialize()

	if err := validateNewAccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.accountSettingV2.AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference_Override(a AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.accountSettingV2.AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) PutNegative(value *AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsNegative) {
	if err := a.validatePutNegativeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNegative",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) PutNeutral(value *AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsNeutral) {
	if err := a.validatePutNeutralParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNeutral",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) PutPositive(value *AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsPositive) {
	if err := a.validatePutPositiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPositive",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) ResetNegative() {
	_jsii_.InvokeVoid(
		a,
		"resetNegative",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) ResetNeutral() {
	_jsii_.InvokeVoid(
		a,
		"resetNeutral",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) ResetPositive() {
	_jsii_.InvokeVoid(
		a,
		"resetPositive",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccountSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

