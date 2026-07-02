// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfigs


type DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsAuthConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/feature_engineering_kafka_configs#mtls_config DataDatabricksFeatureEngineeringKafkaConfigs#mtls_config}.
	MtlsConfig *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsAuthConfigMtlsConfig `field:"optional" json:"mtlsConfig" yaml:"mtlsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/feature_engineering_kafka_configs#uc_service_credential_name DataDatabricksFeatureEngineeringKafkaConfigs#uc_service_credential_name}.
	UcServiceCredentialName *string `field:"optional" json:"ucServiceCredentialName" yaml:"ucServiceCredentialName"`
}

