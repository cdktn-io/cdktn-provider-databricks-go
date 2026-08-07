// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservice


type DataDatabricksAiGatewayModelProviderServiceConfigAmazonBedrockDirect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_service#aws_access_key_id DataDatabricksAiGatewayModelProviderService#aws_access_key_id}.
	AwsAccessKeyId *string `field:"optional" json:"awsAccessKeyId" yaml:"awsAccessKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_service#aws_secret_access_key DataDatabricksAiGatewayModelProviderService#aws_secret_access_key}.
	AwsSecretAccessKey *DataDatabricksAiGatewayModelProviderServiceConfigAmazonBedrockDirectAwsSecretAccessKey `field:"optional" json:"awsSecretAccessKey" yaml:"awsSecretAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_service#region DataDatabricksAiGatewayModelProviderService#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_service#service_credential DataDatabricksAiGatewayModelProviderService#service_credential}.
	ServiceCredential *DataDatabricksAiGatewayModelProviderServiceConfigAmazonBedrockDirectServiceCredential `field:"optional" json:"serviceCredential" yaml:"serviceCredential"`
}

