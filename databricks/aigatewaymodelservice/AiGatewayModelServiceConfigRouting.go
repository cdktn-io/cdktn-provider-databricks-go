// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelservice


type AiGatewayModelServiceConfigRouting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.133.0/docs/resources/ai_gateway_model_service#destinations AiGatewayModelService#destinations}.
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.133.0/docs/resources/ai_gateway_model_service#fallback AiGatewayModelService#fallback}.
	Fallback *AiGatewayModelServiceConfigRoutingFallback `field:"optional" json:"fallback" yaml:"fallback"`
}

