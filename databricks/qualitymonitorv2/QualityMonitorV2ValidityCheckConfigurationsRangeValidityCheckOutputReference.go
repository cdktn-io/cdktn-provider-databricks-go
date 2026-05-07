// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qualitymonitorv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/qualitymonitorv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference interface {
	cdktn.ComplexObject
	ColumnNames() *[]*string
	SetColumnNames(val *[]*string)
	ColumnNamesInput() *[]*string
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
	LowerBound() *float64
	SetLowerBound(val *float64)
	LowerBoundInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UpperBound() *float64
	SetUpperBound(val *float64)
	UpperBoundInput() *float64
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
	ResetColumnNames()
	ResetLowerBound()
	ResetUpperBound()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference
type jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) ColumnNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"columnNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) ColumnNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"columnNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) LowerBound() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lowerBound",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) LowerBoundInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lowerBoundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) UpperBound() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"upperBound",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) UpperBoundInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"upperBoundInput",
		&returns,
	)
	return returns
}


func NewQualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference {
	_init_.Initialize()

	if err := validateNewQualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.qualityMonitorV2.QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference_Override(q QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.qualityMonitorV2.QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference)SetColumnNames(val *[]*string) {
	if err := j.validateSetColumnNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnNames",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference)SetLowerBound(val *float64) {
	if err := j.validateSetLowerBoundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lowerBound",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference)SetUpperBound(val *float64) {
	if err := j.validateSetUpperBoundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upperBound",
		val,
	)
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) ResetColumnNames() {
	_jsii_.InvokeVoid(
		q,
		"resetColumnNames",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) ResetLowerBound() {
	_jsii_.InvokeVoid(
		q,
		"resetLowerBound",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) ResetUpperBound() {
	_jsii_.InvokeVoid(
		q,
		"resetUpperBound",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := q.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

