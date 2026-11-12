// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchdatasources

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabrickslakewatchdatasources/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference interface {
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
	EffectiveRawTable() DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderEffectiveRawTableOutputReference
	Envelope() DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderEnvelopeOutputReference
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
	InternalValue() *DataDatabricksLakewatchDatasourcesDatasourcesAutoloader
	SetInternalValue(val *DataDatabricksLakewatchDatasourcesDatasourcesAutoloader)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	OverridePartitionColumns() *[]*string
	SetOverridePartitionColumns(val *[]*string)
	OverridePartitionColumnsInput() *[]*string
	QueueUrl() *string
	SetQueueUrl(val *string)
	QueueUrlInput() *string
	RawTable() DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderRawTableOutputReference
	RawTableInput() *DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderRawTable
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
	PutEnvelope(value *DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderEnvelope)
	PutRawTable(value *DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderRawTable)
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

// The jsii proxy struct for DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference
type jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) AwsSecretScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSecretScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) AwsSecretScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSecretScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) EffectiveIncludeExistingFiles() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"effectiveIncludeExistingFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) EffectiveRawTable() DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderEffectiveRawTableOutputReference {
	var returns DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderEffectiveRawTableOutputReference
	_jsii_.Get(
		j,
		"effectiveRawTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) Envelope() DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderEnvelopeOutputReference {
	var returns DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderEnvelopeOutputReference
	_jsii_.Get(
		j,
		"envelope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) EnvelopeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envelopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) FormatOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"formatOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) FormatOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"formatOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) IncludeExistingFiles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeExistingFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) IncludeExistingFilesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeExistingFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) InternalValue() *DataDatabricksLakewatchDatasourcesDatasourcesAutoloader {
	var returns *DataDatabricksLakewatchDatasourcesDatasourcesAutoloader
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) OverridePartitionColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"overridePartitionColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) OverridePartitionColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"overridePartitionColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) QueueUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) QueueUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) RawTable() DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderRawTableOutputReference {
	var returns DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderRawTableOutputReference
	_jsii_.Get(
		j,
		"rawTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) RawTableInput() *DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderRawTable {
	var returns *DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderRawTable
	_jsii_.Get(
		j,
		"rawTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) ServiceCredentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) ServiceCredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchDatasources.DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference_Override(d DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchDatasources.DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference)SetAwsSecretScope(val *string) {
	if err := j.validateSetAwsSecretScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsSecretScope",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference)SetFormatOptions(val *map[string]*string) {
	if err := j.validateSetFormatOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"formatOptions",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference)SetIncludeExistingFiles(val interface{}) {
	if err := j.validateSetIncludeExistingFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeExistingFiles",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference)SetInternalValue(val *DataDatabricksLakewatchDatasourcesDatasourcesAutoloader) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference)SetOverridePartitionColumns(val *[]*string) {
	if err := j.validateSetOverridePartitionColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overridePartitionColumns",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference)SetQueueUrl(val *string) {
	if err := j.validateSetQueueUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueUrl",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference)SetServiceCredentials(val *string) {
	if err := j.validateSetServiceCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceCredentials",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) PutEnvelope(value *DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderEnvelope) {
	if err := d.validatePutEnvelopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEnvelope",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) PutRawTable(value *DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderRawTable) {
	if err := d.validatePutRawTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRawTable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) ResetAwsSecretScope() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsSecretScope",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) ResetEnvelope() {
	_jsii_.InvokeVoid(
		d,
		"resetEnvelope",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) ResetFormatOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetFormatOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) ResetIncludeExistingFiles() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeExistingFiles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) ResetOverridePartitionColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetOverridePartitionColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) ResetQueueUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetQueueUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) ResetServiceCredentials() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceCredentials",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

