// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspacesettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabricksworkspacesettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference interface {
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
	Negative() DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsNegativeOutputReference
	NegativeInput() interface{}
	Neutral() DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsNeutralOutputReference
	NeutralInput() interface{}
	Positive() DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsPositiveOutputReference
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
	PutNegative(value *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsNegative)
	PutNeutral(value *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsNeutral)
	PutPositive(value *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsPositive)
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

// The jsii proxy struct for DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference
type jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) Negative() DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsNegativeOutputReference {
	var returns DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsNegativeOutputReference
	_jsii_.Get(
		j,
		"negative",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) NegativeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) Neutral() DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsNeutralOutputReference {
	var returns DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsNeutralOutputReference
	_jsii_.Get(
		j,
		"neutral",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) NeutralInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neutralInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) Positive() DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsPositiveOutputReference {
	var returns DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsPositiveOutputReference
	_jsii_.Get(
		j,
		"positive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) PositiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"positiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksWorkspaceSettingV2.DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference_Override(d DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksWorkspaceSettingV2.DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) PutNegative(value *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsNegative) {
	if err := d.validatePutNegativeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNegative",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) PutNeutral(value *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsNeutral) {
	if err := d.validatePutNeutralParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNeutral",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) PutPositive(value *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsPositive) {
	if err := d.validatePutPositiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPositive",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) ResetNegative() {
	_jsii_.InvokeVoid(
		d,
		"resetNegative",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) ResetNeutral() {
	_jsii_.InvokeVoid(
		d,
		"resetNeutral",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) ResetPositive() {
	_jsii_.InvokeVoid(
		d,
		"resetPositive",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

