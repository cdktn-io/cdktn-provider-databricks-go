// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakewatchdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/lakewatchdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LakewatchDatasourceAutoloaderOutputReference interface {
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
	EffectiveRawTable() LakewatchDatasourceAutoloaderEffectiveRawTableOutputReference
	Envelope() LakewatchDatasourceAutoloaderEnvelopeOutputReference
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	OverridePartitionColumns() *[]*string
	SetOverridePartitionColumns(val *[]*string)
	OverridePartitionColumnsInput() *[]*string
	QueueUrl() *string
	SetQueueUrl(val *string)
	QueueUrlInput() *string
	RawTable() LakewatchDatasourceAutoloaderRawTableOutputReference
	RawTableInput() interface{}
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
	PutEnvelope(value *LakewatchDatasourceAutoloaderEnvelope)
	PutRawTable(value *LakewatchDatasourceAutoloaderRawTable)
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

// The jsii proxy struct for LakewatchDatasourceAutoloaderOutputReference
type jsiiProxy_LakewatchDatasourceAutoloaderOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) AwsSecretScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSecretScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) AwsSecretScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSecretScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) EffectiveIncludeExistingFiles() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"effectiveIncludeExistingFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) EffectiveRawTable() LakewatchDatasourceAutoloaderEffectiveRawTableOutputReference {
	var returns LakewatchDatasourceAutoloaderEffectiveRawTableOutputReference
	_jsii_.Get(
		j,
		"effectiveRawTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) Envelope() LakewatchDatasourceAutoloaderEnvelopeOutputReference {
	var returns LakewatchDatasourceAutoloaderEnvelopeOutputReference
	_jsii_.Get(
		j,
		"envelope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) EnvelopeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envelopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) FormatOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"formatOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) FormatOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"formatOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) IncludeExistingFiles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeExistingFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) IncludeExistingFilesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeExistingFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) OverridePartitionColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"overridePartitionColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) OverridePartitionColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"overridePartitionColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) QueueUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) QueueUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) RawTable() LakewatchDatasourceAutoloaderRawTableOutputReference {
	var returns LakewatchDatasourceAutoloaderRawTableOutputReference
	_jsii_.Get(
		j,
		"rawTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) RawTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) ServiceCredentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) ServiceCredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLakewatchDatasourceAutoloaderOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LakewatchDatasourceAutoloaderOutputReference {
	_init_.Initialize()

	if err := validateNewLakewatchDatasourceAutoloaderOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LakewatchDatasourceAutoloaderOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.lakewatchDatasource.LakewatchDatasourceAutoloaderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLakewatchDatasourceAutoloaderOutputReference_Override(l LakewatchDatasourceAutoloaderOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.lakewatchDatasource.LakewatchDatasourceAutoloaderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference)SetAwsSecretScope(val *string) {
	if err := j.validateSetAwsSecretScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsSecretScope",
		val,
	)
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference)SetFormatOptions(val *map[string]*string) {
	if err := j.validateSetFormatOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"formatOptions",
		val,
	)
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference)SetIncludeExistingFiles(val interface{}) {
	if err := j.validateSetIncludeExistingFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeExistingFiles",
		val,
	)
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference)SetOverridePartitionColumns(val *[]*string) {
	if err := j.validateSetOverridePartitionColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overridePartitionColumns",
		val,
	)
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference)SetQueueUrl(val *string) {
	if err := j.validateSetQueueUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueUrl",
		val,
	)
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference)SetServiceCredentials(val *string) {
	if err := j.validateSetServiceCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceCredentials",
		val,
	)
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) PutEnvelope(value *LakewatchDatasourceAutoloaderEnvelope) {
	if err := l.validatePutEnvelopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putEnvelope",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) PutRawTable(value *LakewatchDatasourceAutoloaderRawTable) {
	if err := l.validatePutRawTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRawTable",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) ResetAwsSecretScope() {
	_jsii_.InvokeVoid(
		l,
		"resetAwsSecretScope",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) ResetEnvelope() {
	_jsii_.InvokeVoid(
		l,
		"resetEnvelope",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) ResetFormatOptions() {
	_jsii_.InvokeVoid(
		l,
		"resetFormatOptions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) ResetIncludeExistingFiles() {
	_jsii_.InvokeVoid(
		l,
		"resetIncludeExistingFiles",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) ResetOverridePartitionColumns() {
	_jsii_.InvokeVoid(
		l,
		"resetOverridePartitionColumns",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) ResetQueueUrl() {
	_jsii_.InvokeVoid(
		l,
		"resetQueueUrl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) ResetServiceCredentials() {
	_jsii_.InvokeVoid(
		l,
		"resetServiceCredentials",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchDatasourceAutoloaderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

