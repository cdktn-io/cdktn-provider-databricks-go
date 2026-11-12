// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabrickslakewatchdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksLakewatchDatasourceAutoloaderOutputReference interface {
	cdktn.ComplexObject
	AwsSecretScope() *string
	SetAwsSecretScope(val *string)
	AwsSecretScopeInput() *string
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
	EffectiveIncludeExistingFiles() cdktn.IResolvable
	EffectiveRawTable() DataDatabricksLakewatchDatasourceAutoloaderEffectiveRawTableOutputReference
	Envelope() DataDatabricksLakewatchDatasourceAutoloaderEnvelopeOutputReference
	EnvelopeInput() interface{}
	Format() *string
	SetFormat(val *string)
	FormatInput() *string
	FormatOptions() *map[string]*string
	SetFormatOptions(val *map[string]*string)
	FormatOptionsInput() *map[string]*string
	// Experimental.
	Fqn() *string
	IncludeExistingFiles() interface{}
	SetIncludeExistingFiles(val interface{})
	IncludeExistingFilesInput() interface{}
	InternalValue() *DataDatabricksLakewatchDatasourceAutoloader
	SetInternalValue(val *DataDatabricksLakewatchDatasourceAutoloader)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	OverridePartitionColumns() *[]*string
	SetOverridePartitionColumns(val *[]*string)
	OverridePartitionColumnsInput() *[]*string
	QueueUrl() *string
	SetQueueUrl(val *string)
	QueueUrlInput() *string
	RawTable() DataDatabricksLakewatchDatasourceAutoloaderRawTableOutputReference
	RawTableInput() *DataDatabricksLakewatchDatasourceAutoloaderRawTable
	ServiceCredentials() *string
	SetServiceCredentials(val *string)
	ServiceCredentialsInput() *string
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
	PutEnvelope(value *DataDatabricksLakewatchDatasourceAutoloaderEnvelope)
	PutRawTable(value *DataDatabricksLakewatchDatasourceAutoloaderRawTable)
	ResetAwsSecretScope()
	ResetEnvelope()
	ResetFormatOptions()
	ResetIncludeExistingFiles()
	ResetOverridePartitionColumns()
	ResetQueueUrl()
	ResetServiceCredentials()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksLakewatchDatasourceAutoloaderOutputReference
type jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) AwsSecretScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSecretScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) AwsSecretScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSecretScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) EffectiveIncludeExistingFiles() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"effectiveIncludeExistingFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) EffectiveRawTable() DataDatabricksLakewatchDatasourceAutoloaderEffectiveRawTableOutputReference {
	var returns DataDatabricksLakewatchDatasourceAutoloaderEffectiveRawTableOutputReference
	_jsii_.Get(
		j,
		"effectiveRawTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) Envelope() DataDatabricksLakewatchDatasourceAutoloaderEnvelopeOutputReference {
	var returns DataDatabricksLakewatchDatasourceAutoloaderEnvelopeOutputReference
	_jsii_.Get(
		j,
		"envelope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) EnvelopeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envelopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) FormatOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"formatOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) FormatOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"formatOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) IncludeExistingFiles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeExistingFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) IncludeExistingFilesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeExistingFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) InternalValue() *DataDatabricksLakewatchDatasourceAutoloader {
	var returns *DataDatabricksLakewatchDatasourceAutoloader
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) OverridePartitionColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"overridePartitionColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) OverridePartitionColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"overridePartitionColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) QueueUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) QueueUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) RawTable() DataDatabricksLakewatchDatasourceAutoloaderRawTableOutputReference {
	var returns DataDatabricksLakewatchDatasourceAutoloaderRawTableOutputReference
	_jsii_.Get(
		j,
		"rawTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) RawTableInput() *DataDatabricksLakewatchDatasourceAutoloaderRawTable {
	var returns *DataDatabricksLakewatchDatasourceAutoloaderRawTable
	_jsii_.Get(
		j,
		"rawTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) ServiceCredentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) ServiceCredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksLakewatchDatasourceAutoloaderOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksLakewatchDatasourceAutoloaderOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksLakewatchDatasourceAutoloaderOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchDatasource.DataDatabricksLakewatchDatasourceAutoloaderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksLakewatchDatasourceAutoloaderOutputReference_Override(d DataDatabricksLakewatchDatasourceAutoloaderOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchDatasource.DataDatabricksLakewatchDatasourceAutoloaderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference)SetAwsSecretScope(val *string) {
	if err := j.validateSetAwsSecretScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsSecretScope",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference)SetFormatOptions(val *map[string]*string) {
	if err := j.validateSetFormatOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"formatOptions",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference)SetIncludeExistingFiles(val interface{}) {
	if err := j.validateSetIncludeExistingFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeExistingFiles",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference)SetInternalValue(val *DataDatabricksLakewatchDatasourceAutoloader) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference)SetOverridePartitionColumns(val *[]*string) {
	if err := j.validateSetOverridePartitionColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overridePartitionColumns",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference)SetQueueUrl(val *string) {
	if err := j.validateSetQueueUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueUrl",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference)SetServiceCredentials(val *string) {
	if err := j.validateSetServiceCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceCredentials",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) PutEnvelope(value *DataDatabricksLakewatchDatasourceAutoloaderEnvelope) {
	if err := d.validatePutEnvelopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEnvelope",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) PutRawTable(value *DataDatabricksLakewatchDatasourceAutoloaderRawTable) {
	if err := d.validatePutRawTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRawTable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) ResetAwsSecretScope() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsSecretScope",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) ResetEnvelope() {
	_jsii_.InvokeVoid(
		d,
		"resetEnvelope",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) ResetFormatOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetFormatOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) ResetIncludeExistingFiles() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeExistingFiles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) ResetOverridePartitionColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetOverridePartitionColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) ResetQueueUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetQueueUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) ResetServiceCredentials() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceCredentials",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourceAutoloaderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

