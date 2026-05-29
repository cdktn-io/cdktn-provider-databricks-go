// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringkafkaconfig


type FeatureEngineeringKafkaConfigAuthConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/feature_engineering_kafka_config#mtls_config FeatureEngineeringKafkaConfig#mtls_config}.
	MtlsConfig *FeatureEngineeringKafkaConfigAuthConfigMtlsConfig `field:"optional" json:"mtlsConfig" yaml:"mtlsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/feature_engineering_kafka_config#uc_service_credential_name FeatureEngineeringKafkaConfig#uc_service_credential_name}.
	UcServiceCredentialName *string `field:"optional" json:"ucServiceCredentialName" yaml:"ucServiceCredentialName"`
}

