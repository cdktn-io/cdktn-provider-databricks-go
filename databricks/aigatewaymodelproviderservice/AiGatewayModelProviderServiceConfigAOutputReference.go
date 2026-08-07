// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelproviderservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/aigatewaymodelproviderservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AiGatewayModelProviderServiceConfigAOutputReference interface {
	cdktn.ComplexObject
	AllowAllTargets() interface{}
	SetAllowAllTargets(val interface{})
	AllowAllTargetsInput() interface{}
	AmazonBedrock() AiGatewayModelProviderServiceConfigAmazonBedrockOutputReference
	AmazonBedrockInput() interface{}
	Anthropic() AiGatewayModelProviderServiceConfigAnthropicOutputReference
	AnthropicInput() interface{}
	AzureOpenai() AiGatewayModelProviderServiceConfigAzureOpenaiOutputReference
	AzureOpenaiInput() interface{}
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
	Custom() AiGatewayModelProviderServiceConfigCustomOutputReference
	CustomInput() interface{}
	ForwardHeaders() interface{}
	SetForwardHeaders(val interface{})
	ForwardHeadersInput() interface{}
	ForwardQueryParameters() interface{}
	SetForwardQueryParameters(val interface{})
	ForwardQueryParametersInput() interface{}
	ForwardUnmanagedPaths() interface{}
	SetForwardUnmanagedPaths(val interface{})
	ForwardUnmanagedPathsInput() interface{}
	// Experimental.
	Fqn() *string
	GeminiEnterprise() AiGatewayModelProviderServiceConfigGeminiEnterpriseOutputReference
	GeminiEnterpriseInput() interface{}
	InferenceTable() AiGatewayModelProviderServiceConfigInferenceTableOutputReference
	InferenceTableInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MicrosoftFoundry() AiGatewayModelProviderServiceConfigMicrosoftFoundryOutputReference
	MicrosoftFoundryInput() interface{}
	Openai() AiGatewayModelProviderServiceConfigOpenaiOutputReference
	OpenaiInput() interface{}
	ProviderType() *string
	SetProviderType(val *string)
	ProviderTypeInput() *string
	RateLimits() AiGatewayModelProviderServiceConfigRateLimitsList
	RateLimitsInput() interface{}
	Targets() AiGatewayModelProviderServiceConfigTargetsList
	TargetsInput() interface{}
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
	PutAmazonBedrock(value *AiGatewayModelProviderServiceConfigAmazonBedrock)
	PutAnthropic(value *AiGatewayModelProviderServiceConfigAnthropic)
	PutAzureOpenai(value *AiGatewayModelProviderServiceConfigAzureOpenai)
	PutCustom(value *AiGatewayModelProviderServiceConfigCustom)
	PutGeminiEnterprise(value *AiGatewayModelProviderServiceConfigGeminiEnterprise)
	PutInferenceTable(value *AiGatewayModelProviderServiceConfigInferenceTable)
	PutMicrosoftFoundry(value *AiGatewayModelProviderServiceConfigMicrosoftFoundry)
	PutOpenai(value *AiGatewayModelProviderServiceConfigOpenai)
	PutRateLimits(value interface{})
	PutTargets(value interface{})
	ResetAllowAllTargets()
	ResetAmazonBedrock()
	ResetAnthropic()
	ResetAzureOpenai()
	ResetCustom()
	ResetForwardHeaders()
	ResetForwardQueryParameters()
	ResetForwardUnmanagedPaths()
	ResetGeminiEnterprise()
	ResetInferenceTable()
	ResetMicrosoftFoundry()
	ResetOpenai()
	ResetProviderType()
	ResetRateLimits()
	ResetTargets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AiGatewayModelProviderServiceConfigAOutputReference
type jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) AllowAllTargets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) AllowAllTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) AmazonBedrock() AiGatewayModelProviderServiceConfigAmazonBedrockOutputReference {
	var returns AiGatewayModelProviderServiceConfigAmazonBedrockOutputReference
	_jsii_.Get(
		j,
		"amazonBedrock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) AmazonBedrockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonBedrockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) Anthropic() AiGatewayModelProviderServiceConfigAnthropicOutputReference {
	var returns AiGatewayModelProviderServiceConfigAnthropicOutputReference
	_jsii_.Get(
		j,
		"anthropic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) AnthropicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anthropicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) AzureOpenai() AiGatewayModelProviderServiceConfigAzureOpenaiOutputReference {
	var returns AiGatewayModelProviderServiceConfigAzureOpenaiOutputReference
	_jsii_.Get(
		j,
		"azureOpenai",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) AzureOpenaiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureOpenaiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) Custom() AiGatewayModelProviderServiceConfigCustomOutputReference {
	var returns AiGatewayModelProviderServiceConfigCustomOutputReference
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) CustomInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ForwardHeaders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ForwardHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ForwardQueryParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ForwardQueryParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ForwardUnmanagedPaths() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardUnmanagedPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ForwardUnmanagedPathsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardUnmanagedPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) GeminiEnterprise() AiGatewayModelProviderServiceConfigGeminiEnterpriseOutputReference {
	var returns AiGatewayModelProviderServiceConfigGeminiEnterpriseOutputReference
	_jsii_.Get(
		j,
		"geminiEnterprise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) GeminiEnterpriseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geminiEnterpriseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) InferenceTable() AiGatewayModelProviderServiceConfigInferenceTableOutputReference {
	var returns AiGatewayModelProviderServiceConfigInferenceTableOutputReference
	_jsii_.Get(
		j,
		"inferenceTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) InferenceTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) MicrosoftFoundry() AiGatewayModelProviderServiceConfigMicrosoftFoundryOutputReference {
	var returns AiGatewayModelProviderServiceConfigMicrosoftFoundryOutputReference
	_jsii_.Get(
		j,
		"microsoftFoundry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) MicrosoftFoundryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"microsoftFoundryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) Openai() AiGatewayModelProviderServiceConfigOpenaiOutputReference {
	var returns AiGatewayModelProviderServiceConfigOpenaiOutputReference
	_jsii_.Get(
		j,
		"openai",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) OpenaiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openaiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) RateLimits() AiGatewayModelProviderServiceConfigRateLimitsList {
	var returns AiGatewayModelProviderServiceConfigRateLimitsList
	_jsii_.Get(
		j,
		"rateLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) RateLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rateLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) Targets() AiGatewayModelProviderServiceConfigTargetsList {
	var returns AiGatewayModelProviderServiceConfigTargetsList
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) TargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAiGatewayModelProviderServiceConfigAOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AiGatewayModelProviderServiceConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewAiGatewayModelProviderServiceConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.aiGatewayModelProviderService.AiGatewayModelProviderServiceConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAiGatewayModelProviderServiceConfigAOutputReference_Override(a AiGatewayModelProviderServiceConfigAOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.aiGatewayModelProviderService.AiGatewayModelProviderServiceConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference)SetAllowAllTargets(val interface{}) {
	if err := j.validateSetAllowAllTargetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAllTargets",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference)SetForwardHeaders(val interface{}) {
	if err := j.validateSetForwardHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardHeaders",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference)SetForwardQueryParameters(val interface{}) {
	if err := j.validateSetForwardQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardQueryParameters",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference)SetForwardUnmanagedPaths(val interface{}) {
	if err := j.validateSetForwardUnmanagedPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardUnmanagedPaths",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference)SetProviderType(val *string) {
	if err := j.validateSetProviderTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerType",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) PutAmazonBedrock(value *AiGatewayModelProviderServiceConfigAmazonBedrock) {
	if err := a.validatePutAmazonBedrockParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAmazonBedrock",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) PutAnthropic(value *AiGatewayModelProviderServiceConfigAnthropic) {
	if err := a.validatePutAnthropicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAnthropic",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) PutAzureOpenai(value *AiGatewayModelProviderServiceConfigAzureOpenai) {
	if err := a.validatePutAzureOpenaiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAzureOpenai",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) PutCustom(value *AiGatewayModelProviderServiceConfigCustom) {
	if err := a.validatePutCustomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustom",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) PutGeminiEnterprise(value *AiGatewayModelProviderServiceConfigGeminiEnterprise) {
	if err := a.validatePutGeminiEnterpriseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGeminiEnterprise",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) PutInferenceTable(value *AiGatewayModelProviderServiceConfigInferenceTable) {
	if err := a.validatePutInferenceTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInferenceTable",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) PutMicrosoftFoundry(value *AiGatewayModelProviderServiceConfigMicrosoftFoundry) {
	if err := a.validatePutMicrosoftFoundryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMicrosoftFoundry",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) PutOpenai(value *AiGatewayModelProviderServiceConfigOpenai) {
	if err := a.validatePutOpenaiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpenai",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) PutRateLimits(value interface{}) {
	if err := a.validatePutRateLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRateLimits",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) PutTargets(value interface{}) {
	if err := a.validatePutTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTargets",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ResetAllowAllTargets() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowAllTargets",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ResetAmazonBedrock() {
	_jsii_.InvokeVoid(
		a,
		"resetAmazonBedrock",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ResetAnthropic() {
	_jsii_.InvokeVoid(
		a,
		"resetAnthropic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ResetAzureOpenai() {
	_jsii_.InvokeVoid(
		a,
		"resetAzureOpenai",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ResetCustom() {
	_jsii_.InvokeVoid(
		a,
		"resetCustom",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ResetForwardHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetForwardHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ResetForwardQueryParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetForwardQueryParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ResetForwardUnmanagedPaths() {
	_jsii_.InvokeVoid(
		a,
		"resetForwardUnmanagedPaths",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ResetGeminiEnterprise() {
	_jsii_.InvokeVoid(
		a,
		"resetGeminiEnterprise",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ResetInferenceTable() {
	_jsii_.InvokeVoid(
		a,
		"resetInferenceTable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ResetMicrosoftFoundry() {
	_jsii_.InvokeVoid(
		a,
		"resetMicrosoftFoundry",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ResetOpenai() {
	_jsii_.InvokeVoid(
		a,
		"resetOpenai",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ResetProviderType() {
	_jsii_.InvokeVoid(
		a,
		"resetProviderType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ResetRateLimits() {
	_jsii_.InvokeVoid(
		a,
		"resetRateLimits",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ResetTargets() {
	_jsii_.InvokeVoid(
		a,
		"resetTargets",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AiGatewayModelProviderServiceConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

