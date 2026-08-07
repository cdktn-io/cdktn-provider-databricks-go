// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelservice


type DataDatabricksAiGatewayModelServiceConfigRouting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_service#destinations DataDatabricksAiGatewayModelService#destinations}.
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_service#fallback DataDatabricksAiGatewayModelService#fallback}.
	Fallback *DataDatabricksAiGatewayModelServiceConfigRoutingFallback `field:"optional" json:"fallback" yaml:"fallback"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_service#first_token_timeout DataDatabricksAiGatewayModelService#first_token_timeout}.
	FirstTokenTimeout *string `field:"optional" json:"firstTokenTimeout" yaml:"firstTokenTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_service#traffic_splitting DataDatabricksAiGatewayModelService#traffic_splitting}.
	TrafficSplitting *DataDatabricksAiGatewayModelServiceConfigRoutingTrafficSplitting `field:"optional" json:"trafficSplitting" yaml:"trafficSplitting"`
}

