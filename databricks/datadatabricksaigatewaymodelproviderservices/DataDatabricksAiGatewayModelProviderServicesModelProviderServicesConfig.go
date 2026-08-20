// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservices


type DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#allow_all_targets DataDatabricksAiGatewayModelProviderServices#allow_all_targets}.
	AllowAllTargets interface{} `field:"optional" json:"allowAllTargets" yaml:"allowAllTargets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#amazon_bedrock DataDatabricksAiGatewayModelProviderServices#amazon_bedrock}.
	AmazonBedrock *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAmazonBedrock `field:"optional" json:"amazonBedrock" yaml:"amazonBedrock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#anthropic DataDatabricksAiGatewayModelProviderServices#anthropic}.
	Anthropic *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAnthropic `field:"optional" json:"anthropic" yaml:"anthropic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#azure_openai DataDatabricksAiGatewayModelProviderServices#azure_openai}.
	AzureOpenai *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAzureOpenai `field:"optional" json:"azureOpenai" yaml:"azureOpenai"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#custom DataDatabricksAiGatewayModelProviderServices#custom}.
	Custom *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigCustom `field:"optional" json:"custom" yaml:"custom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#forward_headers DataDatabricksAiGatewayModelProviderServices#forward_headers}.
	ForwardHeaders interface{} `field:"optional" json:"forwardHeaders" yaml:"forwardHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#forward_query_parameters DataDatabricksAiGatewayModelProviderServices#forward_query_parameters}.
	ForwardQueryParameters interface{} `field:"optional" json:"forwardQueryParameters" yaml:"forwardQueryParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#forward_unmanaged_paths DataDatabricksAiGatewayModelProviderServices#forward_unmanaged_paths}.
	ForwardUnmanagedPaths interface{} `field:"optional" json:"forwardUnmanagedPaths" yaml:"forwardUnmanagedPaths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#gemini_enterprise DataDatabricksAiGatewayModelProviderServices#gemini_enterprise}.
	GeminiEnterprise *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigGeminiEnterprise `field:"optional" json:"geminiEnterprise" yaml:"geminiEnterprise"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#inference_table DataDatabricksAiGatewayModelProviderServices#inference_table}.
	InferenceTable *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigInferenceTable `field:"optional" json:"inferenceTable" yaml:"inferenceTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#microsoft_foundry DataDatabricksAiGatewayModelProviderServices#microsoft_foundry}.
	MicrosoftFoundry *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundry `field:"optional" json:"microsoftFoundry" yaml:"microsoftFoundry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#openai DataDatabricksAiGatewayModelProviderServices#openai}.
	Openai *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigOpenai `field:"optional" json:"openai" yaml:"openai"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#provider_type DataDatabricksAiGatewayModelProviderServices#provider_type}.
	ProviderType *string `field:"optional" json:"providerType" yaml:"providerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#rate_limits DataDatabricksAiGatewayModelProviderServices#rate_limits}.
	RateLimits interface{} `field:"optional" json:"rateLimits" yaml:"rateLimits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#targets DataDatabricksAiGatewayModelProviderServices#targets}.
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}

