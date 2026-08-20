// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/job/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JobTriggersModelOutputReference interface {
	cdktn.ComplexObject
	Aliases() *[]*string
	SetAliases(val *[]*string)
	AliasesInput() *[]*string
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
	Condition() *string
	SetCondition(val *string)
	ConditionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *JobTriggersModel
	SetInternalValue(val *JobTriggersModel)
	MinTimeBetweenTriggersSeconds() *float64
	SetMinTimeBetweenTriggersSeconds(val *float64)
	MinTimeBetweenTriggersSecondsInput() *float64
	SecurableName() *string
	SetSecurableName(val *string)
	SecurableNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WaitAfterLastChangeSeconds() *float64
	SetWaitAfterLastChangeSeconds(val *float64)
	WaitAfterLastChangeSecondsInput() *float64
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
	ResetAliases()
	ResetMinTimeBetweenTriggersSeconds()
	ResetSecurableName()
	ResetWaitAfterLastChangeSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTriggersModelOutputReference
type jsiiProxy_JobTriggersModelOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_JobTriggersModelOutputReference) Aliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) AliasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) Condition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) ConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) InternalValue() *JobTriggersModel {
	var returns *JobTriggersModel
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) MinTimeBetweenTriggersSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTimeBetweenTriggersSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) MinTimeBetweenTriggersSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTimeBetweenTriggersSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) SecurableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) SecurableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) WaitAfterLastChangeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitAfterLastChangeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) WaitAfterLastChangeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitAfterLastChangeSecondsInput",
		&returns,
	)
	return returns
}


func NewJobTriggersModelOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) JobTriggersModelOutputReference {
	_init_.Initialize()

	if err := validateNewJobTriggersModelOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTriggersModelOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTriggersModelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTriggersModelOutputReference_Override(j JobTriggersModelOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.job.JobTriggersModelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTriggersModelOutputReference)SetAliases(val *[]*string) {
	if err := j.validateSetAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliases",
		val,
	)
}

func (j *jsiiProxy_JobTriggersModelOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTriggersModelOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTriggersModelOutputReference)SetCondition(val *string) {
	if err := j.validateSetConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"condition",
		val,
	)
}

func (j *jsiiProxy_JobTriggersModelOutputReference)SetInternalValue(val *JobTriggersModel) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTriggersModelOutputReference)SetMinTimeBetweenTriggersSeconds(val *float64) {
	if err := j.validateSetMinTimeBetweenTriggersSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTimeBetweenTriggersSeconds",
		val,
	)
}

func (j *jsiiProxy_JobTriggersModelOutputReference)SetSecurableName(val *string) {
	if err := j.validateSetSecurableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securableName",
		val,
	)
}

func (j *jsiiProxy_JobTriggersModelOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTriggersModelOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTriggersModelOutputReference)SetWaitAfterLastChangeSeconds(val *float64) {
	if err := j.validateSetWaitAfterLastChangeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitAfterLastChangeSeconds",
		val,
	)
}

func (j *jsiiProxy_JobTriggersModelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) ResetAliases() {
	_jsii_.InvokeVoid(
		j,
		"resetAliases",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggersModelOutputReference) ResetMinTimeBetweenTriggersSeconds() {
	_jsii_.InvokeVoid(
		j,
		"resetMinTimeBetweenTriggersSeconds",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggersModelOutputReference) ResetSecurableName() {
	_jsii_.InvokeVoid(
		j,
		"resetSecurableName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggersModelOutputReference) ResetWaitAfterLastChangeSeconds() {
	_jsii_.InvokeVoid(
		j,
		"resetWaitAfterLastChangeSeconds",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTriggersModelOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := j.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTriggersModelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

