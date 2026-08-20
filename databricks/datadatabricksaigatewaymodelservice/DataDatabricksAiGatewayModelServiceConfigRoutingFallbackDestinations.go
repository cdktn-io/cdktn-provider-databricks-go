// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelservice


type DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_service#destination_type DataDatabricksAiGatewayModelService#destination_type}.
	DestinationType *string `field:"required" json:"destinationType" yaml:"destinationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_service#name DataDatabricksAiGatewayModelService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_service#external_model_config DataDatabricksAiGatewayModelService#external_model_config}.
	ExternalModelConfig *DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsExternalModelConfig `field:"optional" json:"externalModelConfig" yaml:"externalModelConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_service#pay_per_token_config DataDatabricksAiGatewayModelService#pay_per_token_config}.
	PayPerTokenConfig *DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsPayPerTokenConfig `field:"optional" json:"payPerTokenConfig" yaml:"payPerTokenConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_service#provisioned_throughput_config DataDatabricksAiGatewayModelService#provisioned_throughput_config}.
	ProvisionedThroughputConfig *DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsProvisionedThroughputConfig `field:"optional" json:"provisionedThroughputConfig" yaml:"provisionedThroughputConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_service#traffic_percentage DataDatabricksAiGatewayModelService#traffic_percentage}.
	TrafficPercentage *float64 `field:"optional" json:"trafficPercentage" yaml:"trafficPercentage"`
}

