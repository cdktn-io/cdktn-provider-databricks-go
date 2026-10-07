// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservice


type DataDatabricksAiGatewayModelProviderServiceConfigAzureOpenaiDirectApiKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/data-sources/ai_gateway_model_provider_service#plaintext DataDatabricksAiGatewayModelProviderService#plaintext}.
	Plaintext *string `field:"optional" json:"plaintext" yaml:"plaintext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/data-sources/ai_gateway_model_provider_service#secret_reference DataDatabricksAiGatewayModelProviderService#secret_reference}.
	SecretReference *DataDatabricksAiGatewayModelProviderServiceConfigAzureOpenaiDirectApiKeySecretReference `field:"optional" json:"secretReference" yaml:"secretReference"`
}

