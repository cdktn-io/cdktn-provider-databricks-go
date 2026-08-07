// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservices

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksaigatewaymodelproviderservices/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference interface {
	cdktn.ComplexObject
	AllowAllTargets() interface{}
	SetAllowAllTargets(val interface{})
	AllowAllTargetsInput() interface{}
	AmazonBedrock() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAmazonBedrockOutputReference
	AmazonBedrockInput() interface{}
	Anthropic() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAnthropicOutputReference
	AnthropicInput() interface{}
	AzureOpenai() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAzureOpenaiOutputReference
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
	Custom() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigCustomOutputReference
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
	GeminiEnterprise() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigGeminiEnterpriseOutputReference
	GeminiEnterpriseInput() interface{}
	InferenceTable() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigInferenceTableOutputReference
	InferenceTableInput() interface{}
	InternalValue() *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfig
	SetInternalValue(val *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfig)
	MicrosoftFoundry() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryOutputReference
	MicrosoftFoundryInput() interface{}
	Openai() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOpenaiOutputReference
	OpenaiInput() interface{}
	ProviderType() *string
	SetProviderType(val *string)
	ProviderTypeInput() *string
	RateLimits() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigRateLimitsList
	RateLimitsInput() interface{}
	Targets() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigTargetsList
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
	PutAmazonBedrock(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAmazonBedrock)
	PutAnthropic(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAnthropic)
	PutAzureOpenai(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAzureOpenai)
	PutCustom(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigCustom)
	PutGeminiEnterprise(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigGeminiEnterprise)
	PutInferenceTable(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigInferenceTable)
	PutMicrosoftFoundry(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundry)
	PutOpenai(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOpenai)
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

// The jsii proxy struct for DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference
type jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) AllowAllTargets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) AllowAllTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) AmazonBedrock() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAmazonBedrockOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAmazonBedrockOutputReference
	_jsii_.Get(
		j,
		"amazonBedrock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) AmazonBedrockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonBedrockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) Anthropic() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAnthropicOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAnthropicOutputReference
	_jsii_.Get(
		j,
		"anthropic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) AnthropicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anthropicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) AzureOpenai() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAzureOpenaiOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAzureOpenaiOutputReference
	_jsii_.Get(
		j,
		"azureOpenai",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) AzureOpenaiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureOpenaiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) Custom() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigCustomOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigCustomOutputReference
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) CustomInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ForwardHeaders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ForwardHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ForwardQueryParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ForwardQueryParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ForwardUnmanagedPaths() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardUnmanagedPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ForwardUnmanagedPathsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardUnmanagedPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) GeminiEnterprise() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigGeminiEnterpriseOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigGeminiEnterpriseOutputReference
	_jsii_.Get(
		j,
		"geminiEnterprise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) GeminiEnterpriseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geminiEnterpriseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) InferenceTable() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigInferenceTableOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigInferenceTableOutputReference
	_jsii_.Get(
		j,
		"inferenceTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) InferenceTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) InternalValue() *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfig {
	var returns *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) MicrosoftFoundry() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryOutputReference
	_jsii_.Get(
		j,
		"microsoftFoundry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) MicrosoftFoundryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"microsoftFoundryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) Openai() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOpenaiOutputReference {
	var returns DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOpenaiOutputReference
	_jsii_.Get(
		j,
		"openai",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) OpenaiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openaiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) RateLimits() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigRateLimitsList {
	var returns DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigRateLimitsList
	_jsii_.Get(
		j,
		"rateLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) RateLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rateLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) Targets() DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigTargetsList {
	var returns DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigTargetsList
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) TargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiGatewayModelProviderServices.DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference_Override(d DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksAiGatewayModelProviderServices.DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference)SetAllowAllTargets(val interface{}) {
	if err := j.validateSetAllowAllTargetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAllTargets",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference)SetForwardHeaders(val interface{}) {
	if err := j.validateSetForwardHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardHeaders",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference)SetForwardQueryParameters(val interface{}) {
	if err := j.validateSetForwardQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardQueryParameters",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference)SetForwardUnmanagedPaths(val interface{}) {
	if err := j.validateSetForwardUnmanagedPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardUnmanagedPaths",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference)SetInternalValue(val *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference)SetProviderType(val *string) {
	if err := j.validateSetProviderTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) PutAmazonBedrock(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAmazonBedrock) {
	if err := d.validatePutAmazonBedrockParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAmazonBedrock",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) PutAnthropic(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAnthropic) {
	if err := d.validatePutAnthropicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAnthropic",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) PutAzureOpenai(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAzureOpenai) {
	if err := d.validatePutAzureOpenaiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureOpenai",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) PutCustom(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigCustom) {
	if err := d.validatePutCustomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustom",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) PutGeminiEnterprise(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigGeminiEnterprise) {
	if err := d.validatePutGeminiEnterpriseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGeminiEnterprise",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) PutInferenceTable(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigInferenceTable) {
	if err := d.validatePutInferenceTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInferenceTable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) PutMicrosoftFoundry(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundry) {
	if err := d.validatePutMicrosoftFoundryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMicrosoftFoundry",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) PutOpenai(value *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOpenai) {
	if err := d.validatePutOpenaiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOpenai",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) PutRateLimits(value interface{}) {
	if err := d.validatePutRateLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRateLimits",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) PutTargets(value interface{}) {
	if err := d.validatePutTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTargets",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ResetAllowAllTargets() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowAllTargets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ResetAmazonBedrock() {
	_jsii_.InvokeVoid(
		d,
		"resetAmazonBedrock",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ResetAnthropic() {
	_jsii_.InvokeVoid(
		d,
		"resetAnthropic",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ResetAzureOpenai() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureOpenai",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ResetCustom() {
	_jsii_.InvokeVoid(
		d,
		"resetCustom",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ResetForwardHeaders() {
	_jsii_.InvokeVoid(
		d,
		"resetForwardHeaders",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ResetForwardQueryParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetForwardQueryParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ResetForwardUnmanagedPaths() {
	_jsii_.InvokeVoid(
		d,
		"resetForwardUnmanagedPaths",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ResetGeminiEnterprise() {
	_jsii_.InvokeVoid(
		d,
		"resetGeminiEnterprise",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ResetInferenceTable() {
	_jsii_.InvokeVoid(
		d,
		"resetInferenceTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ResetMicrosoftFoundry() {
	_jsii_.InvokeVoid(
		d,
		"resetMicrosoftFoundry",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ResetOpenai() {
	_jsii_.InvokeVoid(
		d,
		"resetOpenai",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ResetProviderType() {
	_jsii_.InvokeVoid(
		d,
		"resetProviderType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ResetRateLimits() {
	_jsii_.InvokeVoid(
		d,
		"resetRateLimits",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ResetTargets() {
	_jsii_.InvokeVoid(
		d,
		"resetTargets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

