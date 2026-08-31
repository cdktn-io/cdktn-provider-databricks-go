// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelservices


type DataDatabricksAiGatewayModelServicesModelServicesConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/ai_gateway_model_services#inference_table DataDatabricksAiGatewayModelServices#inference_table}.
	InferenceTable *DataDatabricksAiGatewayModelServicesModelServicesConfigInferenceTable `field:"optional" json:"inferenceTable" yaml:"inferenceTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/ai_gateway_model_services#rate_limits DataDatabricksAiGatewayModelServices#rate_limits}.
	RateLimits interface{} `field:"optional" json:"rateLimits" yaml:"rateLimits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/ai_gateway_model_services#routing DataDatabricksAiGatewayModelServices#routing}.
	Routing *DataDatabricksAiGatewayModelServicesModelServicesConfigRouting `field:"optional" json:"routing" yaml:"routing"`
}

