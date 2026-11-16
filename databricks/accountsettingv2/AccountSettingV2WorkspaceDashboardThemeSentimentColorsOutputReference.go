// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountsettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/accountsettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference interface {
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
	Negative() AccountSettingV2WorkspaceDashboardThemeSentimentColorsNegativeOutputReference
	NegativeInput() interface{}
	Neutral() AccountSettingV2WorkspaceDashboardThemeSentimentColorsNeutralOutputReference
	NeutralInput() interface{}
	Positive() AccountSettingV2WorkspaceDashboardThemeSentimentColorsPositiveOutputReference
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
	PutNegative(value *AccountSettingV2WorkspaceDashboardThemeSentimentColorsNegative)
	PutNeutral(value *AccountSettingV2WorkspaceDashboardThemeSentimentColorsNeutral)
	PutPositive(value *AccountSettingV2WorkspaceDashboardThemeSentimentColorsPositive)
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

// The jsii proxy struct for AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference
type jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) Negative() AccountSettingV2WorkspaceDashboardThemeSentimentColorsNegativeOutputReference {
	var returns AccountSettingV2WorkspaceDashboardThemeSentimentColorsNegativeOutputReference
	_jsii_.Get(
		j,
		"negative",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) NegativeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) Neutral() AccountSettingV2WorkspaceDashboardThemeSentimentColorsNeutralOutputReference {
	var returns AccountSettingV2WorkspaceDashboardThemeSentimentColorsNeutralOutputReference
	_jsii_.Get(
		j,
		"neutral",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) NeutralInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neutralInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) Positive() AccountSettingV2WorkspaceDashboardThemeSentimentColorsPositiveOutputReference {
	var returns AccountSettingV2WorkspaceDashboardThemeSentimentColorsPositiveOutputReference
	_jsii_.Get(
		j,
		"positive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) PositiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"positiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference {
	_init_.Initialize()

	if err := validateNewAccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.accountSettingV2.AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference_Override(a AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.accountSettingV2.AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) PutNegative(value *AccountSettingV2WorkspaceDashboardThemeSentimentColorsNegative) {
	if err := a.validatePutNegativeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNegative",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) PutNeutral(value *AccountSettingV2WorkspaceDashboardThemeSentimentColorsNeutral) {
	if err := a.validatePutNeutralParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNeutral",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) PutPositive(value *AccountSettingV2WorkspaceDashboardThemeSentimentColorsPositive) {
	if err := a.validatePutPositiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPositive",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) ResetNegative() {
	_jsii_.InvokeVoid(
		a,
		"resetNegative",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) ResetNeutral() {
	_jsii_.InvokeVoid(
		a,
		"resetNeutral",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) ResetPositive() {
	_jsii_.InvokeVoid(
		a,
		"resetPositive",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccountSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

