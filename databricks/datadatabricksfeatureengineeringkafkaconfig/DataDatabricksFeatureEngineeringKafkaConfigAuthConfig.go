// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfig


type DataDatabricksFeatureEngineeringKafkaConfigAuthConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/feature_engineering_kafka_config#mtls_config DataDatabricksFeatureEngineeringKafkaConfig#mtls_config}.
	MtlsConfig *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfig `field:"optional" json:"mtlsConfig" yaml:"mtlsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/feature_engineering_kafka_config#uc_service_credential_name DataDatabricksFeatureEngineeringKafkaConfig#uc_service_credential_name}.
	UcServiceCredentialName *string `field:"optional" json:"ucServiceCredentialName" yaml:"ucServiceCredentialName"`
}

