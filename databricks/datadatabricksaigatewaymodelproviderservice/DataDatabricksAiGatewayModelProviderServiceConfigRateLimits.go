// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservice


type DataDatabricksAiGatewayModelProviderServiceConfigRateLimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/ai_gateway_model_provider_service#key DataDatabricksAiGatewayModelProviderService#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/ai_gateway_model_provider_service#renewal_period DataDatabricksAiGatewayModelProviderService#renewal_period}.
	RenewalPeriod *string `field:"required" json:"renewalPeriod" yaml:"renewalPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/ai_gateway_model_provider_service#principal DataDatabricksAiGatewayModelProviderService#principal}.
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/ai_gateway_model_provider_service#requests DataDatabricksAiGatewayModelProviderService#requests}.
	Requests *float64 `field:"optional" json:"requests" yaml:"requests"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/ai_gateway_model_provider_service#request_tag_key DataDatabricksAiGatewayModelProviderService#request_tag_key}.
	RequestTagKey *string `field:"optional" json:"requestTagKey" yaml:"requestTagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/ai_gateway_model_provider_service#request_tag_value DataDatabricksAiGatewayModelProviderService#request_tag_value}.
	RequestTagValue *string `field:"optional" json:"requestTagValue" yaml:"requestTagValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/ai_gateway_model_provider_service#tokens DataDatabricksAiGatewayModelProviderService#tokens}.
	Tokens *float64 `field:"optional" json:"tokens" yaml:"tokens"`
}

