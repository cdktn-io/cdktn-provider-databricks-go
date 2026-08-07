// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksaigatewaymodelproviderservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference interface {
	cdktn.ComplexObject
	AllowAllTargets() interface{}
	SetAllowAllTargets(val interface{})
	AllowAllTargetsInput() interface{}
	AmazonBedrock() DataDatabricksAiGatewayModelProviderServiceConfigAmazonBedrockOutputReference
	AmazonBedrockInput() interface{}
	Anthropic() DataDatabricksAiGatewayModelProviderServiceConfigAnthropicOutputReference
	AnthropicInput() interface{}
	AzureOpenai() DataDatabricksAiGatewayModelProviderServiceConfigAzureOpenaiOutputReference
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
	Custom() DataDatabricksAiGatewayModelProviderServiceConfigCustomOutputReference
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
	GeminiEnterprise() DataDatabricksAiGatewayModelProviderServiceConfigGeminiEnterpriseOutputReference
	GeminiEnterpriseInput() interface{}
	InferenceTable() DataDatabricksAiGatewayModelProviderServiceConfigInferenceTableOutputReference
	InferenceTableInput() interface{}
	InternalValue() *DataDatabricksAiGatewayModelProviderServiceConfigA
	SetInternalValue(val *DataDatabricksAiGatewayModelProviderServiceConfigA)
	MicrosoftFoundry() DataDatabricksAiGatewayModelProviderServiceConfigMicrosoftFoundryOutputReference
	MicrosoftFoundryInput() interface{}
	Openai() DataDatabricksAiGatewayModelProviderServiceConfigOpenaiOutputReference
	OpenaiInput() interface{}
	ProviderType() *string
	SetProviderType(val *string)
	ProviderTypeInput() *string
	RateLimits() DataDatabricksAiGatewayModelProviderServiceConfigRateLimitsList
	RateLimitsInput() interface{}
	Targets() DataDatabricksAiGatewayModelProviderServiceConfigTargetsList
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
	PutAmazonBedrock(value *DataDatabricksAiGatewayModelProviderServiceConfigAmazonBedrock)
	PutAnthropic(value *DataDatabricksAiGatewayModelProviderServiceConfigAnthropic)
	PutAzureOpenai(value *DataDatabricksAiGatewayModelProviderServiceConfigAzureOpenai)
	PutCustom(value *DataDatabricksAiGatewayModelProviderServiceConfigCustom)
	PutGeminiEnterprise(value *DataDatabricksAiGatewayModelProviderServiceConfigGeminiEnterprise)
	PutInferenceTable(value *DataDatabricksAiGatewayModelProviderServiceConfigInferenceTable)
	PutMicrosoftFoundry(value *DataDatabricksAiGatewayModelProviderServiceConfigMicrosoftFoundry)
	PutOpenai(value *DataDatabricksAiGatewayModelProviderServiceConfigOpenai)
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

// The jsii proxy struct for DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference
type jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) AllowAllTargets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) AllowAllTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) AmazonBedrock() DataDatabricksAiGatewayModelProviderServiceConfigAmazonBedrockOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServiceConfigAmazonBedrockOutputReference
	_jsii_.Get(
		j,
		"amazonBedrock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) AmazonBedrockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonBedrockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) Anthropic() DataDatabricksAiGatewayModelProviderServiceConfigAnthropicOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServiceConfigAnthropicOutputReference
	_jsii_.Get(
		j,
		"anthropic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) AnthropicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anthropicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) AzureOpenai() DataDatabricksAiGatewayModelProviderServiceConfigAzureOpenaiOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServiceConfigAzureOpenaiOutputReference
	_jsii_.Get(
		j,
		"azureOpenai",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) AzureOpenaiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureOpenaiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) Custom() DataDatabricksAiGatewayModelProviderServiceConfigCustomOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServiceConfigCustomOutputReference
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) CustomInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ForwardHeaders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ForwardHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ForwardQueryParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ForwardQueryParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ForwardUnmanagedPaths() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardUnmanagedPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ForwardUnmanagedPathsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardUnmanagedPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) GeminiEnterprise() DataDatabricksAiGatewayModelProviderServiceConfigGeminiEnterpriseOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServiceConfigGeminiEnterpriseOutputReference
	_jsii_.Get(
		j,
		"geminiEnterprise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) GeminiEnterpriseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geminiEnterpriseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) InferenceTable() DataDatabricksAiGatewayModelProviderServiceConfigInferenceTableOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServiceConfigInferenceTableOutputReference
	_jsii_.Get(
		j,
		"inferenceTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) InferenceTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) InternalValue() *DataDatabricksAiGatewayModelProviderServiceConfigA {
	var returns *DataDatabricksAiGatewayModelProviderServiceConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) MicrosoftFoundry() DataDatabricksAiGatewayModelProviderServiceConfigMicrosoftFoundryOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServiceConfigMicrosoftFoundryOutputReference
	_jsii_.Get(
		j,
		"microsoftFoundry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) MicrosoftFoundryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"microsoftFoundryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) Openai() DataDatabricksAiGatewayModelProviderServiceConfigOpenaiOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServiceConfigOpenaiOutputReference
	_jsii_.Get(
		j,
		"openai",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) OpenaiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openaiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) RateLimits() DataDatabricksAiGatewayModelProviderServiceConfigRateLimitsList {
	var returns DataDatabricksAiGatewayModelProviderServiceConfigRateLimitsList
	_jsii_.Get(
		j,
		"rateLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) RateLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rateLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) Targets() DataDatabricksAiGatewayModelProviderServiceConfigTargetsList {
	var returns DataDatabricksAiGatewayModelProviderServiceConfigTargetsList
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) TargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAiGatewayModelProviderServiceConfigAOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAiGatewayModelProviderServiceConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiGatewayModelProviderService.DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAiGatewayModelProviderServiceConfigAOutputReference_Override(d DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiGatewayModelProviderService.DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference)SetAllowAllTargets(val interface{}) {
	if err := j.validateSetAllowAllTargetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAllTargets",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference)SetForwardHeaders(val interface{}) {
	if err := j.validateSetForwardHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardHeaders",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference)SetForwardQueryParameters(val interface{}) {
	if err := j.validateSetForwardQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardQueryParameters",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference)SetForwardUnmanagedPaths(val interface{}) {
	if err := j.validateSetForwardUnmanagedPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardUnmanagedPaths",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference)SetInternalValue(val *DataDatabricksAiGatewayModelProviderServiceConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference)SetProviderType(val *string) {
	if err := j.validateSetProviderTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) PutAmazonBedrock(value *DataDatabricksAiGatewayModelProviderServiceConfigAmazonBedrock) {
	if err := d.validatePutAmazonBedrockParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAmazonBedrock",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) PutAnthropic(value *DataDatabricksAiGatewayModelProviderServiceConfigAnthropic) {
	if err := d.validatePutAnthropicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAnthropic",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) PutAzureOpenai(value *DataDatabricksAiGatewayModelProviderServiceConfigAzureOpenai) {
	if err := d.validatePutAzureOpenaiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureOpenai",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) PutCustom(value *DataDatabricksAiGatewayModelProviderServiceConfigCustom) {
	if err := d.validatePutCustomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustom",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) PutGeminiEnterprise(value *DataDatabricksAiGatewayModelProviderServiceConfigGeminiEnterprise) {
	if err := d.validatePutGeminiEnterpriseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGeminiEnterprise",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) PutInferenceTable(value *DataDatabricksAiGatewayModelProviderServiceConfigInferenceTable) {
	if err := d.validatePutInferenceTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInferenceTable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) PutMicrosoftFoundry(value *DataDatabricksAiGatewayModelProviderServiceConfigMicrosoftFoundry) {
	if err := d.validatePutMicrosoftFoundryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMicrosoftFoundry",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) PutOpenai(value *DataDatabricksAiGatewayModelProviderServiceConfigOpenai) {
	if err := d.validatePutOpenaiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOpenai",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) PutRateLimits(value interface{}) {
	if err := d.validatePutRateLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRateLimits",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) PutTargets(value interface{}) {
	if err := d.validatePutTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTargets",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ResetAllowAllTargets() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowAllTargets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ResetAmazonBedrock() {
	_jsii_.InvokeVoid(
		d,
		"resetAmazonBedrock",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ResetAnthropic() {
	_jsii_.InvokeVoid(
		d,
		"resetAnthropic",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ResetAzureOpenai() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureOpenai",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ResetCustom() {
	_jsii_.InvokeVoid(
		d,
		"resetCustom",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ResetForwardHeaders() {
	_jsii_.InvokeVoid(
		d,
		"resetForwardHeaders",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ResetForwardQueryParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetForwardQueryParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ResetForwardUnmanagedPaths() {
	_jsii_.InvokeVoid(
		d,
		"resetForwardUnmanagedPaths",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ResetGeminiEnterprise() {
	_jsii_.InvokeVoid(
		d,
		"resetGeminiEnterprise",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ResetInferenceTable() {
	_jsii_.InvokeVoid(
		d,
		"resetInferenceTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ResetMicrosoftFoundry() {
	_jsii_.InvokeVoid(
		d,
		"resetMicrosoftFoundry",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ResetOpenai() {
	_jsii_.InvokeVoid(
		d,
		"resetOpenai",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ResetProviderType() {
	_jsii_.InvokeVoid(
		d,
		"resetProviderType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ResetRateLimits() {
	_jsii_.InvokeVoid(
		d,
		"resetRateLimits",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ResetTargets() {
	_jsii_.InvokeVoid(
		d,
		"resetTargets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServiceConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

