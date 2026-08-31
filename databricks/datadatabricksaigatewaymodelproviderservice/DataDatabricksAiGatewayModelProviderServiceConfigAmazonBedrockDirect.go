// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservice


type DataDatabricksAiGatewayModelProviderServiceConfigAmazonBedrockDirect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/ai_gateway_model_provider_service#aws_access_key DataDatabricksAiGatewayModelProviderService#aws_access_key}.
	AwsAccessKey *DataDatabricksAiGatewayModelProviderServiceConfigAmazonBedrockDirectAwsAccessKey `field:"optional" json:"awsAccessKey" yaml:"awsAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/ai_gateway_model_provider_service#region DataDatabricksAiGatewayModelProviderService#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/ai_gateway_model_provider_service#service_credential DataDatabricksAiGatewayModelProviderService#service_credential}.
	ServiceCredential *DataDatabricksAiGatewayModelProviderServiceConfigAmazonBedrockDirectServiceCredential `field:"optional" json:"serviceCredential" yaml:"serviceCredential"`
}

