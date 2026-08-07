// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksfeatureengineeringkafkaconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference interface {
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
	DisableHostnameVerification() interface{}
	SetDisableHostnameVerification(val interface{})
	DisableHostnameVerificationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyPasswordRef() DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeyPasswordRefOutputReference
	KeyPasswordRefInput() *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeyPasswordRef
	KeystoreLocation() *string
	SetKeystoreLocation(val *string)
	KeystoreLocationInput() *string
	KeystorePasswordRef() DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeystorePasswordRefOutputReference
	KeystorePasswordRefInput() *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeystorePasswordRef
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TruststoreLocation() *string
	SetTruststoreLocation(val *string)
	TruststoreLocationInput() *string
	TruststorePasswordRef() DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigTruststorePasswordRefOutputReference
	TruststorePasswordRefInput() *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigTruststorePasswordRef
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
	PutKeyPasswordRef(value *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeyPasswordRef)
	PutKeystorePasswordRef(value *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeystorePasswordRef)
	PutTruststorePasswordRef(value *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigTruststorePasswordRef)
	ResetDisableHostnameVerification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference
type jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) DisableHostnameVerification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableHostnameVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) DisableHostnameVerificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableHostnameVerificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) KeyPasswordRef() DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeyPasswordRefOutputReference {
	var returns DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeyPasswordRefOutputReference
	_jsii_.Get(
		j,
		"keyPasswordRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) KeyPasswordRefInput() *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeyPasswordRef {
	var returns *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeyPasswordRef
	_jsii_.Get(
		j,
		"keyPasswordRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) KeystoreLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystoreLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) KeystoreLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystoreLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) KeystorePasswordRef() DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeystorePasswordRefOutputReference {
	var returns DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeystorePasswordRefOutputReference
	_jsii_.Get(
		j,
		"keystorePasswordRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) KeystorePasswordRefInput() *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeystorePasswordRef {
	var returns *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeystorePasswordRef
	_jsii_.Get(
		j,
		"keystorePasswordRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) TruststoreLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststoreLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) TruststoreLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststoreLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) TruststorePasswordRef() DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigTruststorePasswordRefOutputReference {
	var returns DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigTruststorePasswordRefOutputReference
	_jsii_.Get(
		j,
		"truststorePasswordRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) TruststorePasswordRefInput() *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigTruststorePasswordRef {
	var returns *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigTruststorePasswordRef
	_jsii_.Get(
		j,
		"truststorePasswordRefInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringKafkaConfig.DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference_Override(d DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksFeatureEngineeringKafkaConfig.DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference)SetDisableHostnameVerification(val interface{}) {
	if err := j.validateSetDisableHostnameVerificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableHostnameVerification",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference)SetKeystoreLocation(val *string) {
	if err := j.validateSetKeystoreLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keystoreLocation",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference)SetTruststoreLocation(val *string) {
	if err := j.validateSetTruststoreLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"truststoreLocation",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) PutKeyPasswordRef(value *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeyPasswordRef) {
	if err := d.validatePutKeyPasswordRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKeyPasswordRef",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) PutKeystorePasswordRef(value *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeystorePasswordRef) {
	if err := d.validatePutKeystorePasswordRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKeystorePasswordRef",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) PutTruststorePasswordRef(value *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigTruststorePasswordRef) {
	if err := d.validatePutTruststorePasswordRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTruststorePasswordRef",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) ResetDisableHostnameVerification() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableHostnameVerification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

