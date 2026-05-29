// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringkafkaconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/featureengineeringkafkaconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference interface {
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
	KeyPasswordRef() FeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeyPasswordRefOutputReference
	KeyPasswordRefInput() interface{}
	KeystoreLocation() *string
	SetKeystoreLocation(val *string)
	KeystoreLocationInput() *string
	KeystorePasswordRef() FeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeystorePasswordRefOutputReference
	KeystorePasswordRefInput() interface{}
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
	TruststorePasswordRef() FeatureEngineeringKafkaConfigAuthConfigMtlsConfigTruststorePasswordRefOutputReference
	TruststorePasswordRefInput() interface{}
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
	PutKeyPasswordRef(value *FeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeyPasswordRef)
	PutKeystorePasswordRef(value *FeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeystorePasswordRef)
	PutTruststorePasswordRef(value *FeatureEngineeringKafkaConfigAuthConfigMtlsConfigTruststorePasswordRef)
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

// The jsii proxy struct for FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference
type jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) DisableHostnameVerification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableHostnameVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) DisableHostnameVerificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableHostnameVerificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) KeyPasswordRef() FeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeyPasswordRefOutputReference {
	var returns FeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeyPasswordRefOutputReference
	_jsii_.Get(
		j,
		"keyPasswordRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) KeyPasswordRefInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyPasswordRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) KeystoreLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystoreLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) KeystoreLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystoreLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) KeystorePasswordRef() FeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeystorePasswordRefOutputReference {
	var returns FeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeystorePasswordRefOutputReference
	_jsii_.Get(
		j,
		"keystorePasswordRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) KeystorePasswordRefInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keystorePasswordRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) TruststoreLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststoreLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) TruststoreLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststoreLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) TruststorePasswordRef() FeatureEngineeringKafkaConfigAuthConfigMtlsConfigTruststorePasswordRefOutputReference {
	var returns FeatureEngineeringKafkaConfigAuthConfigMtlsConfigTruststorePasswordRefOutputReference
	_jsii_.Get(
		j,
		"truststorePasswordRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) TruststorePasswordRefInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"truststorePasswordRefInput",
		&returns,
	)
	return returns
}


func NewFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.featureEngineeringKafkaConfig.FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference_Override(f FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.featureEngineeringKafkaConfig.FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference)SetDisableHostnameVerification(val interface{}) {
	if err := j.validateSetDisableHostnameVerificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableHostnameVerification",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference)SetKeystoreLocation(val *string) {
	if err := j.validateSetKeystoreLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keystoreLocation",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference)SetTruststoreLocation(val *string) {
	if err := j.validateSetTruststoreLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"truststoreLocation",
		val,
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) PutKeyPasswordRef(value *FeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeyPasswordRef) {
	if err := f.validatePutKeyPasswordRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putKeyPasswordRef",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) PutKeystorePasswordRef(value *FeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeystorePasswordRef) {
	if err := f.validatePutKeystorePasswordRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putKeystorePasswordRef",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) PutTruststorePasswordRef(value *FeatureEngineeringKafkaConfigAuthConfigMtlsConfigTruststorePasswordRef) {
	if err := f.validatePutTruststorePasswordRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTruststorePasswordRef",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) ResetDisableHostnameVerification() {
	_jsii_.InvokeVoid(
		f,
		"resetDisableHostnameVerification",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := f.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureEngineeringKafkaConfigAuthConfigMtlsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

