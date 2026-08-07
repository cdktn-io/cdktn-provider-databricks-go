// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservices


type DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAmazonBedrockDirect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_services#aws_access_key_id DataDatabricksAiGatewayModelProviderServices#aws_access_key_id}.
	AwsAccessKeyId *string `field:"optional" json:"awsAccessKeyId" yaml:"awsAccessKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_services#aws_secret_access_key DataDatabricksAiGatewayModelProviderServices#aws_secret_access_key}.
	AwsSecretAccessKey *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAmazonBedrockDirectAwsSecretAccessKey `field:"optional" json:"awsSecretAccessKey" yaml:"awsSecretAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_services#region DataDatabricksAiGatewayModelProviderServices#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_services#service_credential DataDatabricksAiGatewayModelProviderServices#service_credential}.
	ServiceCredential *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAmazonBedrockDirectServiceCredential `field:"optional" json:"serviceCredential" yaml:"serviceCredential"`
}

