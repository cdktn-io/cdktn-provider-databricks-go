// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservices


type DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigTargets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_gateway_model_provider_services#model DataDatabricksAiGatewayModelProviderServices#model}.
	Model *string `field:"required" json:"model" yaml:"model"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_gateway_model_provider_services#native_api_types DataDatabricksAiGatewayModelProviderServices#native_api_types}.
	NativeApiTypes *[]*string `field:"optional" json:"nativeApiTypes" yaml:"nativeApiTypes"`
}

