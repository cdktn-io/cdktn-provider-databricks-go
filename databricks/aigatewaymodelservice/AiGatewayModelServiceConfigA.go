// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelservice


type AiGatewayModelServiceConfigA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/ai_gateway_model_service#inference_table AiGatewayModelService#inference_table}.
	InferenceTable *AiGatewayModelServiceConfigInferenceTable `field:"optional" json:"inferenceTable" yaml:"inferenceTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/ai_gateway_model_service#rate_limits AiGatewayModelService#rate_limits}.
	RateLimits interface{} `field:"optional" json:"rateLimits" yaml:"rateLimits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/ai_gateway_model_service#routing AiGatewayModelService#routing}.
	Routing *AiGatewayModelServiceConfigRouting `field:"optional" json:"routing" yaml:"routing"`
}

