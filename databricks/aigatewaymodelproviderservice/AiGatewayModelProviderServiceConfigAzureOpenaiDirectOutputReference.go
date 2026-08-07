// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelproviderservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/aigatewaymodelproviderservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference interface {
	cdktn.ComplexObject
	ApiKey() AiGatewayModelProviderServiceConfigAzureOpenaiDirectApiKeyOutputReference
	ApiKeyInput() interface{}
	BaseUrl() *string
	SetBaseUrl(val *string)
	BaseUrlInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() AiGatewayModelProviderServiceConfigAzureOpenaiDirectClientSecretOutputReference
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
	ServiceCredential() AiGatewayModelProviderServiceConfigAzureOpenaiDirectServiceCredentialOutputReference
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
	PutApiKey(value *AiGatewayModelProviderServiceConfigAzureOpenaiDirectApiKey)
	PutClientSecret(value *AiGatewayModelProviderServiceConfigAzureOpenaiDirectClientSecret)
	PutServiceCredential(value *AiGatewayModelProviderServiceConfigAzureOpenaiDirectServiceCredential)
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

// The jsii proxy struct for AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference
type jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) ApiKey() AiGatewayModelProviderServiceConfigAzureOpenaiDirectApiKeyOutputReference {
	var returns AiGatewayModelProviderServiceConfigAzureOpenaiDirectApiKeyOutputReference
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) ApiKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) BaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) BaseUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) ClientSecret() AiGatewayModelProviderServiceConfigAzureOpenaiDirectClientSecretOutputReference {
	var returns AiGatewayModelProviderServiceConfigAzureOpenaiDirectClientSecretOutputReference
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) ClientSecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) ServiceCredential() AiGatewayModelProviderServiceConfigAzureOpenaiDirectServiceCredentialOutputReference {
	var returns AiGatewayModelProviderServiceConfigAzureOpenaiDirectServiceCredentialOutputReference
	_jsii_.Get(
		j,
		"serviceCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) ServiceCredentialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference {
	_init_.Initialize()

	if err := validateNewAiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.aiGatewayModelProviderService.AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference_Override(a AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.aiGatewayModelProviderService.AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference)SetBaseUrl(val *string) {
	if err := j.validateSetBaseUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseUrl",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) PutApiKey(value *AiGatewayModelProviderServiceConfigAzureOpenaiDirectApiKey) {
	if err := a.validatePutApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApiKey",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) PutClientSecret(value *AiGatewayModelProviderServiceConfigAzureOpenaiDirectClientSecret) {
	if err := a.validatePutClientSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putClientSecret",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) PutServiceCredential(value *AiGatewayModelProviderServiceConfigAzureOpenaiDirectServiceCredential) {
	if err := a.validatePutServiceCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServiceCredential",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) ResetApiKey() {
	_jsii_.InvokeVoid(
		a,
		"resetApiKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) ResetBaseUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetBaseUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) ResetServiceCredential() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceCredential",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) ResetTenantId() {
	_jsii_.InvokeVoid(
		a,
		"resetTenantId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAzureOpenaiDirectOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

