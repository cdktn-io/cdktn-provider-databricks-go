// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelproviderservice


type AiGatewayModelProviderServiceConfigAmazonBedrockDirect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/ai_gateway_model_provider_service#aws_access_key_id AiGatewayModelProviderService#aws_access_key_id}.
	AwsAccessKeyId *string `field:"optional" json:"awsAccessKeyId" yaml:"awsAccessKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/ai_gateway_model_provider_service#aws_secret_access_key AiGatewayModelProviderService#aws_secret_access_key}.
	AwsSecretAccessKey *AiGatewayModelProviderServiceConfigAmazonBedrockDirectAwsSecretAccessKey `field:"optional" json:"awsSecretAccessKey" yaml:"awsSecretAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/ai_gateway_model_provider_service#region AiGatewayModelProviderService#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/ai_gateway_model_provider_service#service_credential AiGatewayModelProviderService#service_credential}.
	ServiceCredential *AiGatewayModelProviderServiceConfigAmazonBedrockDirectServiceCredential `field:"optional" json:"serviceCredential" yaml:"serviceCredential"`
}

