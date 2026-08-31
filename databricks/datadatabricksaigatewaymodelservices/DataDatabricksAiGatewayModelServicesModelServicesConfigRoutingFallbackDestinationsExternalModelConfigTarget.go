// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelservices


type DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingFallbackDestinationsExternalModelConfigTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/ai_gateway_model_services#model DataDatabricksAiGatewayModelServices#model}.
	Model *string `field:"required" json:"model" yaml:"model"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/ai_gateway_model_services#native_api_types DataDatabricksAiGatewayModelServices#native_api_types}.
	NativeApiTypes *[]*string `field:"optional" json:"nativeApiTypes" yaml:"nativeApiTypes"`
}

