// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/postgresproject/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PostgresProjectSpecOutputReference interface {
	cdktn.ComplexObject
	BudgetPolicyId() *string
	SetBudgetPolicyId(val *string)
	BudgetPolicyIdInput() *string
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
	CustomTags() PostgresProjectSpecCustomTagsList
	CustomTagsInput() interface{}
	DefaultEndpointSettings() PostgresProjectSpecDefaultEndpointSettingsOutputReference
	DefaultEndpointSettingsInput() interface{}
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	// Experimental.
	Fqn() *string
	HistoryRetentionDuration() *string
	SetHistoryRetentionDuration(val *string)
	HistoryRetentionDurationInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PgVersion() *float64
	SetPgVersion(val *float64)
	PgVersionInput() *float64
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
	PutCustomTags(value interface{})
	PutDefaultEndpointSettings(value *PostgresProjectSpecDefaultEndpointSettings)
	ResetBudgetPolicyId()
	ResetCustomTags()
	ResetDefaultEndpointSettings()
	ResetDisplayName()
	ResetHistoryRetentionDuration()
	ResetPgVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PostgresProjectSpecOutputReference
type jsiiProxy_PostgresProjectSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) BudgetPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) BudgetPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) CustomTags() PostgresProjectSpecCustomTagsList {
	var returns PostgresProjectSpecCustomTagsList
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) CustomTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) DefaultEndpointSettings() PostgresProjectSpecDefaultEndpointSettingsOutputReference {
	var returns PostgresProjectSpecDefaultEndpointSettingsOutputReference
	_jsii_.Get(
		j,
		"defaultEndpointSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) DefaultEndpointSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultEndpointSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) HistoryRetentionDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"historyRetentionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) HistoryRetentionDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"historyRetentionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) PgVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pgVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) PgVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pgVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPostgresProjectSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PostgresProjectSpecOutputReference {
	_init_.Initialize()

	if err := validateNewPostgresProjectSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PostgresProjectSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.postgresProject.PostgresProjectSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPostgresProjectSpecOutputReference_Override(p PostgresProjectSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.postgresProject.PostgresProjectSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference)SetBudgetPolicyId(val *string) {
	if err := j.validateSetBudgetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"budgetPolicyId",
		val,
	)
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference)SetHistoryRetentionDuration(val *string) {
	if err := j.validateSetHistoryRetentionDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"historyRetentionDuration",
		val,
	)
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference)SetPgVersion(val *float64) {
	if err := j.validateSetPgVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pgVersion",
		val,
	)
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PostgresProjectSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) PutCustomTags(value interface{}) {
	if err := p.validatePutCustomTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCustomTags",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) PutDefaultEndpointSettings(value *PostgresProjectSpecDefaultEndpointSettings) {
	if err := p.validatePutDefaultEndpointSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDefaultEndpointSettings",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) ResetBudgetPolicyId() {
	_jsii_.InvokeVoid(
		p,
		"resetBudgetPolicyId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) ResetCustomTags() {
	_jsii_.InvokeVoid(
		p,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) ResetDefaultEndpointSettings() {
	_jsii_.InvokeVoid(
		p,
		"resetDefaultEndpointSettings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		p,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) ResetHistoryRetentionDuration() {
	_jsii_.InvokeVoid(
		p,
		"resetHistoryRetentionDuration",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) ResetPgVersion() {
	_jsii_.InvokeVoid(
		p,
		"resetPgVersion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresProjectSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

