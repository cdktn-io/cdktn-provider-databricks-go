// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservices


type DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigBedrockMantleDirect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/ai_gateway_model_provider_services#aws_access_key DataDatabricksAiGatewayModelProviderServices#aws_access_key}.
	AwsAccessKey *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigBedrockMantleDirectAwsAccessKey `field:"optional" json:"awsAccessKey" yaml:"awsAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/ai_gateway_model_provider_services#region DataDatabricksAiGatewayModelProviderServices#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/ai_gateway_model_provider_services#service_credential DataDatabricksAiGatewayModelProviderServices#service_credential}.
	ServiceCredential *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigBedrockMantleDirectServiceCredential `field:"optional" json:"serviceCredential" yaml:"serviceCredential"`
}

