// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgressnapshotschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/postgressnapshotschedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PostgresSnapshotScheduleScheduleOutputReference interface {
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
	DailySchedule() PostgresSnapshotScheduleScheduleDailyScheduleOutputReference
	DailyScheduleInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MonthlySchedule() PostgresSnapshotScheduleScheduleMonthlyScheduleOutputReference
	MonthlyScheduleInput() interface{}
	Retention() *string
	SetRetention(val *string)
	RetentionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WeeklySchedule() PostgresSnapshotScheduleScheduleWeeklyScheduleOutputReference
	WeeklyScheduleInput() interface{}
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
	PutDailySchedule(value *PostgresSnapshotScheduleScheduleDailySchedule)
	PutMonthlySchedule(value *PostgresSnapshotScheduleScheduleMonthlySchedule)
	PutWeeklySchedule(value *PostgresSnapshotScheduleScheduleWeeklySchedule)
	ResetDailySchedule()
	ResetMonthlySchedule()
	ResetWeeklySchedule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PostgresSnapshotScheduleScheduleOutputReference
type jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) DailySchedule() PostgresSnapshotScheduleScheduleDailyScheduleOutputReference {
	var returns PostgresSnapshotScheduleScheduleDailyScheduleOutputReference
	_jsii_.Get(
		j,
		"dailySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) DailyScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dailyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) MonthlySchedule() PostgresSnapshotScheduleScheduleMonthlyScheduleOutputReference {
	var returns PostgresSnapshotScheduleScheduleMonthlyScheduleOutputReference
	_jsii_.Get(
		j,
		"monthlySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) MonthlyScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monthlyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) Retention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) RetentionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) WeeklySchedule() PostgresSnapshotScheduleScheduleWeeklyScheduleOutputReference {
	var returns PostgresSnapshotScheduleScheduleWeeklyScheduleOutputReference
	_jsii_.Get(
		j,
		"weeklySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) WeeklyScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"weeklyScheduleInput",
		&returns,
	)
	return returns
}


func NewPostgresSnapshotScheduleScheduleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PostgresSnapshotScheduleScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewPostgresSnapshotScheduleScheduleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.postgresSnapshotSchedule.PostgresSnapshotScheduleScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPostgresSnapshotScheduleScheduleOutputReference_Override(p PostgresSnapshotScheduleScheduleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.postgresSnapshotSchedule.PostgresSnapshotScheduleScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference)SetRetention(val *string) {
	if err := j.validateSetRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) PutDailySchedule(value *PostgresSnapshotScheduleScheduleDailySchedule) {
	if err := p.validatePutDailyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDailySchedule",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) PutMonthlySchedule(value *PostgresSnapshotScheduleScheduleMonthlySchedule) {
	if err := p.validatePutMonthlyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMonthlySchedule",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) PutWeeklySchedule(value *PostgresSnapshotScheduleScheduleWeeklySchedule) {
	if err := p.validatePutWeeklyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putWeeklySchedule",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) ResetDailySchedule() {
	_jsii_.InvokeVoid(
		p,
		"resetDailySchedule",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) ResetMonthlySchedule() {
	_jsii_.InvokeVoid(
		p,
		"resetMonthlySchedule",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) ResetWeeklySchedule() {
	_jsii_.InvokeVoid(
		p,
		"resetWeeklySchedule",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PostgresSnapshotScheduleScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

