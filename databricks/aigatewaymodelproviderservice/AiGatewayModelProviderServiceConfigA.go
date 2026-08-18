// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelproviderservice


type AiGatewayModelProviderServiceConfigA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_provider_service#allow_all_targets AiGatewayModelProviderService#allow_all_targets}.
	AllowAllTargets interface{} `field:"optional" json:"allowAllTargets" yaml:"allowAllTargets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_provider_service#amazon_bedrock AiGatewayModelProviderService#amazon_bedrock}.
	AmazonBedrock *AiGatewayModelProviderServiceConfigAmazonBedrock `field:"optional" json:"amazonBedrock" yaml:"amazonBedrock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_provider_service#anthropic AiGatewayModelProviderService#anthropic}.
	Anthropic *AiGatewayModelProviderServiceConfigAnthropic `field:"optional" json:"anthropic" yaml:"anthropic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_provider_service#azure_openai AiGatewayModelProviderService#azure_openai}.
	AzureOpenai *AiGatewayModelProviderServiceConfigAzureOpenai `field:"optional" json:"azureOpenai" yaml:"azureOpenai"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_provider_service#custom AiGatewayModelProviderService#custom}.
	Custom *AiGatewayModelProviderServiceConfigCustom `field:"optional" json:"custom" yaml:"custom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_provider_service#forward_headers AiGatewayModelProviderService#forward_headers}.
	ForwardHeaders interface{} `field:"optional" json:"forwardHeaders" yaml:"forwardHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_provider_service#forward_query_parameters AiGatewayModelProviderService#forward_query_parameters}.
	ForwardQueryParameters interface{} `field:"optional" json:"forwardQueryParameters" yaml:"forwardQueryParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_provider_service#forward_unmanaged_paths AiGatewayModelProviderService#forward_unmanaged_paths}.
	ForwardUnmanagedPaths interface{} `field:"optional" json:"forwardUnmanagedPaths" yaml:"forwardUnmanagedPaths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_provider_service#gemini_enterprise AiGatewayModelProviderService#gemini_enterprise}.
	GeminiEnterprise *AiGatewayModelProviderServiceConfigGeminiEnterprise `field:"optional" json:"geminiEnterprise" yaml:"geminiEnterprise"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_provider_service#inference_table AiGatewayModelProviderService#inference_table}.
	InferenceTable *AiGatewayModelProviderServiceConfigInferenceTable `field:"optional" json:"inferenceTable" yaml:"inferenceTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_provider_service#microsoft_foundry AiGatewayModelProviderService#microsoft_foundry}.
	MicrosoftFoundry *AiGatewayModelProviderServiceConfigMicrosoftFoundry `field:"optional" json:"microsoftFoundry" yaml:"microsoftFoundry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_provider_service#openai AiGatewayModelProviderService#openai}.
	Openai *AiGatewayModelProviderServiceConfigOpenai `field:"optional" json:"openai" yaml:"openai"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_provider_service#provider_type AiGatewayModelProviderService#provider_type}.
	ProviderType *string `field:"optional" json:"providerType" yaml:"providerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_provider_service#rate_limits AiGatewayModelProviderService#rate_limits}.
	RateLimits interface{} `field:"optional" json:"rateLimits" yaml:"rateLimits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_provider_service#targets AiGatewayModelProviderService#targets}.
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}

