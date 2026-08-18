// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelservices


type DataDatabricksAiGatewayModelServicesModelServicesConfigRateLimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_gateway_model_services#key DataDatabricksAiGatewayModelServices#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_gateway_model_services#renewal_period DataDatabricksAiGatewayModelServices#renewal_period}.
	RenewalPeriod *string `field:"required" json:"renewalPeriod" yaml:"renewalPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_gateway_model_services#principal DataDatabricksAiGatewayModelServices#principal}.
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_gateway_model_services#requests DataDatabricksAiGatewayModelServices#requests}.
	Requests *float64 `field:"optional" json:"requests" yaml:"requests"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_gateway_model_services#request_tag_key DataDatabricksAiGatewayModelServices#request_tag_key}.
	RequestTagKey *string `field:"optional" json:"requestTagKey" yaml:"requestTagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_gateway_model_services#request_tag_value DataDatabricksAiGatewayModelServices#request_tag_value}.
	RequestTagValue *string `field:"optional" json:"requestTagValue" yaml:"requestTagValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_gateway_model_services#tokens DataDatabricksAiGatewayModelServices#tokens}.
	Tokens *float64 `field:"optional" json:"tokens" yaml:"tokens"`
}

