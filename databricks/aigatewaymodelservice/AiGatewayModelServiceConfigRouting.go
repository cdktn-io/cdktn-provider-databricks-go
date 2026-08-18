// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelservice


type AiGatewayModelServiceConfigRouting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_service#destinations AiGatewayModelService#destinations}.
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_service#fallback AiGatewayModelService#fallback}.
	Fallback *AiGatewayModelServiceConfigRoutingFallback `field:"optional" json:"fallback" yaml:"fallback"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_service#first_token_timeout AiGatewayModelService#first_token_timeout}.
	FirstTokenTimeout *string `field:"optional" json:"firstTokenTimeout" yaml:"firstTokenTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/ai_gateway_model_service#traffic_splitting AiGatewayModelService#traffic_splitting}.
	TrafficSplitting *AiGatewayModelServiceConfigRoutingTrafficSplitting `field:"optional" json:"trafficSplitting" yaml:"trafficSplitting"`
}

