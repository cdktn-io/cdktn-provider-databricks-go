// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelproviderservice


type AiGatewayModelProviderServiceConfigAmazonBedrockDirectAwsAccessKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/ai_gateway_model_provider_service#access_key_id AiGatewayModelProviderService#access_key_id}.
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/ai_gateway_model_provider_service#secret_access_key AiGatewayModelProviderService#secret_access_key}.
	SecretAccessKey *AiGatewayModelProviderServiceConfigAmazonBedrockDirectAwsAccessKeySecretAccessKey `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
}

