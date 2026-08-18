// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelservice


type AiGatewayModelServiceConfigRoutingFallbackDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_service#destination_type AiGatewayModelService#destination_type}.
	DestinationType *string `field:"required" json:"destinationType" yaml:"destinationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_service#name AiGatewayModelService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_service#external_model_config AiGatewayModelService#external_model_config}.
	ExternalModelConfig *AiGatewayModelServiceConfigRoutingFallbackDestinationsExternalModelConfig `field:"optional" json:"externalModelConfig" yaml:"externalModelConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_service#pay_per_token_config AiGatewayModelService#pay_per_token_config}.
	PayPerTokenConfig *AiGatewayModelServiceConfigRoutingFallbackDestinationsPayPerTokenConfig `field:"optional" json:"payPerTokenConfig" yaml:"payPerTokenConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_service#provisioned_throughput_config AiGatewayModelService#provisioned_throughput_config}.
	ProvisionedThroughputConfig *AiGatewayModelServiceConfigRoutingFallbackDestinationsProvisionedThroughputConfig `field:"optional" json:"provisionedThroughputConfig" yaml:"provisionedThroughputConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_service#traffic_percentage AiGatewayModelService#traffic_percentage}.
	TrafficPercentage *float64 `field:"optional" json:"trafficPercentage" yaml:"trafficPercentage"`
}

