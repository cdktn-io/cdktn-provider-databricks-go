// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservice


type DataDatabricksAiGatewayModelProviderServiceConfigAmazonBedrockDirectAwsAccessKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_service#access_key_id DataDatabricksAiGatewayModelProviderService#access_key_id}.
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_service#secret_access_key DataDatabricksAiGatewayModelProviderService#secret_access_key}.
	SecretAccessKey *DataDatabricksAiGatewayModelProviderServiceConfigAmazonBedrockDirectAwsAccessKeySecretAccessKey `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
}

