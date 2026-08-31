// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservices


type DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAmazonBedrockDirectAwsAccessKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/ai_gateway_model_provider_services#access_key_id DataDatabricksAiGatewayModelProviderServices#access_key_id}.
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/ai_gateway_model_provider_services#secret_access_key DataDatabricksAiGatewayModelProviderServices#secret_access_key}.
	SecretAccessKey *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAmazonBedrockDirectAwsAccessKeySecretAccessKey `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
}

