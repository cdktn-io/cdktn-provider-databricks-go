// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservices

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksaigatewaymodelproviderservices/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference interface {
	cdktn.ComplexObject
	ApiKey() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectApiKeyOutputReference
	ApiKeyInput() interface{}
	BaseUrl() *string
	SetBaseUrl(val *string)
	BaseUrlInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectClientSecretOutputReference
	ClientSecretInput() interface{}
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
	ServiceCredential() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectServiceCredentialOutputReference
	ServiceCredentialInput() interface{}
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
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
	PutApiKey(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectApiKey)
	PutClientSecret(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectClientSecret)
	PutServiceCredential(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectServiceCredential)
	ResetApiKey()
	ResetBaseUrl()
	ResetClientId()
	ResetClientSecret()
	ResetServiceCredential()
	ResetTenantId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference
type jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) ApiKey() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectApiKeyOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectApiKeyOutputReference
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) ApiKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) BaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) BaseUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) ClientSecret() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectClientSecretOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectClientSecretOutputReference
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) ClientSecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) ServiceCredential() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectServiceCredentialOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectServiceCredentialOutputReference
	_jsii_.Get(
		j,
		"serviceCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) ServiceCredentialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiGatewayModelProviderServices.DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference_Override(d DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiGatewayModelProviderServices.DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference)SetBaseUrl(val *string) {
	if err := j.validateSetBaseUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseUrl",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) PutApiKey(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectApiKey) {
	if err := d.validatePutApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApiKey",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) PutClientSecret(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectClientSecret) {
	if err := d.validatePutClientSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putClientSecret",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) PutServiceCredential(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectServiceCredential) {
	if err := d.validatePutServiceCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceCredential",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) ResetApiKey() {
	_jsii_.InvokeVoid(
		d,
		"resetApiKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) ResetBaseUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetBaseUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		d,
		"resetClientId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		d,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) ResetServiceCredential() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceCredential",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) ResetTenantId() {
	_jsii_.InvokeVoid(
		d,
		"resetTenantId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

