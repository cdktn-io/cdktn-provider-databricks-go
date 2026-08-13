// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelservice


type DataDatabricksAiGatewayModelServiceConfigA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/ai_gateway_model_service#inference_table DataDatabricksAiGatewayModelService#inference_table}.
	InferenceTable *DataDatabricksAiGatewayModelServiceConfigInferenceTable `field:"optional" json:"inferenceTable" yaml:"inferenceTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/ai_gateway_model_service#rate_limits DataDatabricksAiGatewayModelService#rate_limits}.
	RateLimits interface{} `field:"optional" json:"rateLimits" yaml:"rateLimits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/ai_gateway_model_service#routing DataDatabricksAiGatewayModelService#routing}.
	Routing *DataDatabricksAiGatewayModelServiceConfigRouting `field:"optional" json:"routing" yaml:"routing"`
}

