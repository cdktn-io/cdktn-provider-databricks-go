// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfigs


type DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsAuthConfigMtlsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/feature_engineering_kafka_configs#key_password_ref DataDatabricksFeatureEngineeringKafkaConfigs#key_password_ref}.
	KeyPasswordRef *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsAuthConfigMtlsConfigKeyPasswordRef `field:"required" json:"keyPasswordRef" yaml:"keyPasswordRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/feature_engineering_kafka_configs#keystore_location DataDatabricksFeatureEngineeringKafkaConfigs#keystore_location}.
	KeystoreLocation *string `field:"required" json:"keystoreLocation" yaml:"keystoreLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/feature_engineering_kafka_configs#keystore_password_ref DataDatabricksFeatureEngineeringKafkaConfigs#keystore_password_ref}.
	KeystorePasswordRef *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsAuthConfigMtlsConfigKeystorePasswordRef `field:"required" json:"keystorePasswordRef" yaml:"keystorePasswordRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/feature_engineering_kafka_configs#truststore_location DataDatabricksFeatureEngineeringKafkaConfigs#truststore_location}.
	TruststoreLocation *string `field:"required" json:"truststoreLocation" yaml:"truststoreLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/feature_engineering_kafka_configs#truststore_password_ref DataDatabricksFeatureEngineeringKafkaConfigs#truststore_password_ref}.
	TruststorePasswordRef *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsAuthConfigMtlsConfigTruststorePasswordRef `field:"required" json:"truststorePasswordRef" yaml:"truststorePasswordRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/feature_engineering_kafka_configs#disable_hostname_verification DataDatabricksFeatureEngineeringKafkaConfigs#disable_hostname_verification}.
	DisableHostnameVerification interface{} `field:"optional" json:"disableHostnameVerification" yaml:"disableHostnameVerification"`
}

