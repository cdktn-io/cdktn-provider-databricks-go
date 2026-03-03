// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qualitymonitorv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/qualitymonitorv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QualityMonitorV2ValidityCheckConfigurationsOutputReference interface {
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	PercentNullValidityCheck() QualityMonitorV2ValidityCheckConfigurationsPercentNullValidityCheckOutputReference
	PercentNullValidityCheckInput() interface{}
	RangeValidityCheck() QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference
	RangeValidityCheckInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UniquenessValidityCheck() QualityMonitorV2ValidityCheckConfigurationsUniquenessValidityCheckOutputReference
	UniquenessValidityCheckInput() interface{}
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
	PutPercentNullValidityCheck(value *QualityMonitorV2ValidityCheckConfigurationsPercentNullValidityCheck)
	PutRangeValidityCheck(value *QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheck)
	PutUniquenessValidityCheck(value *QualityMonitorV2ValidityCheckConfigurationsUniquenessValidityCheck)
	ResetName()
	ResetPercentNullValidityCheck()
	ResetRangeValidityCheck()
	ResetUniquenessValidityCheck()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QualityMonitorV2ValidityCheckConfigurationsOutputReference
type jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) PercentNullValidityCheck() QualityMonitorV2ValidityCheckConfigurationsPercentNullValidityCheckOutputReference {
	var returns QualityMonitorV2ValidityCheckConfigurationsPercentNullValidityCheckOutputReference
	_jsii_.Get(
		j,
		"percentNullValidityCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) PercentNullValidityCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"percentNullValidityCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) RangeValidityCheck() QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference {
	var returns QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference
	_jsii_.Get(
		j,
		"rangeValidityCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) RangeValidityCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rangeValidityCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) UniquenessValidityCheck() QualityMonitorV2ValidityCheckConfigurationsUniquenessValidityCheckOutputReference {
	var returns QualityMonitorV2ValidityCheckConfigurationsUniquenessValidityCheckOutputReference
	_jsii_.Get(
		j,
		"uniquenessValidityCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) UniquenessValidityCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uniquenessValidityCheckInput",
		&returns,
	)
	return returns
}


func NewQualityMonitorV2ValidityCheckConfigurationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) QualityMonitorV2ValidityCheckConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewQualityMonitorV2ValidityCheckConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.qualityMonitorV2.QualityMonitorV2ValidityCheckConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewQualityMonitorV2ValidityCheckConfigurationsOutputReference_Override(q QualityMonitorV2ValidityCheckConfigurationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.qualityMonitorV2.QualityMonitorV2ValidityCheckConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		q,
	)
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) PutPercentNullValidityCheck(value *QualityMonitorV2ValidityCheckConfigurationsPercentNullValidityCheck) {
	if err := q.validatePutPercentNullValidityCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPercentNullValidityCheck",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) PutRangeValidityCheck(value *QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheck) {
	if err := q.validatePutRangeValidityCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putRangeValidityCheck",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) PutUniquenessValidityCheck(value *QualityMonitorV2ValidityCheckConfigurationsUniquenessValidityCheck) {
	if err := q.validatePutUniquenessValidityCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putUniquenessValidityCheck",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		q,
		"resetName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) ResetPercentNullValidityCheck() {
	_jsii_.InvokeVoid(
		q,
		"resetPercentNullValidityCheck",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) ResetRangeValidityCheck() {
	_jsii_.InvokeVoid(
		q,
		"resetRangeValidityCheck",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) ResetUniquenessValidityCheck() {
	_jsii_.InvokeVoid(
		q,
		"resetUniquenessValidityCheck",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QualityMonitorV2ValidityCheckConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

