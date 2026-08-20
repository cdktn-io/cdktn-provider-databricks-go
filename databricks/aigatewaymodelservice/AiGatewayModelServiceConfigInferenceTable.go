// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelservice


type AiGatewayModelServiceConfigInferenceTable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/ai_gateway_model_service#parent AiGatewayModelService#parent}.
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/ai_gateway_model_service#disabled AiGatewayModelService#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/ai_gateway_model_service#table_name_prefix AiGatewayModelService#table_name_prefix}.
	TableNamePrefix *string `field:"optional" json:"tableNamePrefix" yaml:"tableNamePrefix"`
}

