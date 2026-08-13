// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfig


type DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/feature_engineering_kafka_config#key_password_ref DataDatabricksFeatureEngineeringKafkaConfig#key_password_ref}.
	KeyPasswordRef *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeyPasswordRef `field:"required" json:"keyPasswordRef" yaml:"keyPasswordRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/feature_engineering_kafka_config#keystore_location DataDatabricksFeatureEngineeringKafkaConfig#keystore_location}.
	KeystoreLocation *string `field:"required" json:"keystoreLocation" yaml:"keystoreLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/feature_engineering_kafka_config#keystore_password_ref DataDatabricksFeatureEngineeringKafkaConfig#keystore_password_ref}.
	KeystorePasswordRef *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeystorePasswordRef `field:"required" json:"keystorePasswordRef" yaml:"keystorePasswordRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/feature_engineering_kafka_config#truststore_location DataDatabricksFeatureEngineeringKafkaConfig#truststore_location}.
	TruststoreLocation *string `field:"required" json:"truststoreLocation" yaml:"truststoreLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/feature_engineering_kafka_config#truststore_password_ref DataDatabricksFeatureEngineeringKafkaConfig#truststore_password_ref}.
	TruststorePasswordRef *DataDatabricksFeatureEngineeringKafkaConfigAuthConfigMtlsConfigTruststorePasswordRef `field:"required" json:"truststorePasswordRef" yaml:"truststorePasswordRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/feature_engineering_kafka_config#disable_hostname_verification DataDatabricksFeatureEngineeringKafkaConfig#disable_hostname_verification}.
	DisableHostnameVerification interface{} `field:"optional" json:"disableHostnameVerification" yaml:"disableHostnameVerification"`
}

