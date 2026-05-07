// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksqualitymonitorv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabricksqualitymonitorv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference interface {
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
	InternalValue() *DataDatabricksQualityMonitorV2ValidityCheckConfigurations
	SetInternalValue(val *DataDatabricksQualityMonitorV2ValidityCheckConfigurations)
	Name() *string
	SetName(val *string)
	NameInput() *string
	PercentNullValidityCheck() DataDatabricksQualityMonitorV2ValidityCheckConfigurationsPercentNullValidityCheckOutputReference
	PercentNullValidityCheckInput() interface{}
	RangeValidityCheck() DataDatabricksQualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference
	RangeValidityCheckInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UniquenessValidityCheck() DataDatabricksQualityMonitorV2ValidityCheckConfigurationsUniquenessValidityCheckOutputReference
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
	PutPercentNullValidityCheck(value *DataDatabricksQualityMonitorV2ValidityCheckConfigurationsPercentNullValidityCheck)
	PutRangeValidityCheck(value *DataDatabricksQualityMonitorV2ValidityCheckConfigurationsRangeValidityCheck)
	PutUniquenessValidityCheck(value *DataDatabricksQualityMonitorV2ValidityCheckConfigurationsUniquenessValidityCheck)
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

// The jsii proxy struct for DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference
type jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) InternalValue() *DataDatabricksQualityMonitorV2ValidityCheckConfigurations {
	var returns *DataDatabricksQualityMonitorV2ValidityCheckConfigurations
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) PercentNullValidityCheck() DataDatabricksQualityMonitorV2ValidityCheckConfigurationsPercentNullValidityCheckOutputReference {
	var returns DataDatabricksQualityMonitorV2ValidityCheckConfigurationsPercentNullValidityCheckOutputReference
	_jsii_.Get(
		j,
		"percentNullValidityCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) PercentNullValidityCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"percentNullValidityCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) RangeValidityCheck() DataDatabricksQualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference {
	var returns DataDatabricksQualityMonitorV2ValidityCheckConfigurationsRangeValidityCheckOutputReference
	_jsii_.Get(
		j,
		"rangeValidityCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) RangeValidityCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rangeValidityCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) UniquenessValidityCheck() DataDatabricksQualityMonitorV2ValidityCheckConfigurationsUniquenessValidityCheckOutputReference {
	var returns DataDatabricksQualityMonitorV2ValidityCheckConfigurationsUniquenessValidityCheckOutputReference
	_jsii_.Get(
		j,
		"uniquenessValidityCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) UniquenessValidityCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uniquenessValidityCheckInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksQualityMonitorV2.DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference_Override(d DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksQualityMonitorV2.DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference)SetInternalValue(val *DataDatabricksQualityMonitorV2ValidityCheckConfigurations) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) PutPercentNullValidityCheck(value *DataDatabricksQualityMonitorV2ValidityCheckConfigurationsPercentNullValidityCheck) {
	if err := d.validatePutPercentNullValidityCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPercentNullValidityCheck",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) PutRangeValidityCheck(value *DataDatabricksQualityMonitorV2ValidityCheckConfigurationsRangeValidityCheck) {
	if err := d.validatePutRangeValidityCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRangeValidityCheck",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) PutUniquenessValidityCheck(value *DataDatabricksQualityMonitorV2ValidityCheckConfigurationsUniquenessValidityCheck) {
	if err := d.validatePutUniquenessValidityCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUniquenessValidityCheck",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) ResetPercentNullValidityCheck() {
	_jsii_.InvokeVoid(
		d,
		"resetPercentNullValidityCheck",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) ResetRangeValidityCheck() {
	_jsii_.InvokeVoid(
		d,
		"resetRangeValidityCheck",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) ResetUniquenessValidityCheck() {
	_jsii_.InvokeVoid(
		d,
		"resetUniquenessValidityCheck",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksQualityMonitorV2ValidityCheckConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

