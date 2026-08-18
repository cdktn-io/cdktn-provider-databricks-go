// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelservices


type DataDatabricksAiGatewayModelServicesModelServicesConfigRouting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_gateway_model_services#destinations DataDatabricksAiGatewayModelServices#destinations}.
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_gateway_model_services#fallback DataDatabricksAiGatewayModelServices#fallback}.
	Fallback *DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallback `field:"optional" json:"fallback" yaml:"fallback"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_gateway_model_services#first_token_timeout DataDatabricksAiGatewayModelServices#first_token_timeout}.
	FirstTokenTimeout *string `field:"optional" json:"firstTokenTimeout" yaml:"firstTokenTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_gateway_model_services#traffic_splitting DataDatabricksAiGatewayModelServices#traffic_splitting}.
	TrafficSplitting *DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingTrafficSplitting `field:"optional" json:"trafficSplitting" yaml:"trafficSplitting"`
}

