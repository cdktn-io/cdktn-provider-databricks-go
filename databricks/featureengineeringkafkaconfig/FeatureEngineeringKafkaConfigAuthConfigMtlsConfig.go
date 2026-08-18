// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringkafkaconfig


type FeatureEngineeringKafkaConfigAuthConfigMtlsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/feature_engineering_kafka_config#key_password_ref FeatureEngineeringKafkaConfig#key_password_ref}.
	KeyPasswordRef *FeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeyPasswordRef `field:"required" json:"keyPasswordRef" yaml:"keyPasswordRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/feature_engineering_kafka_config#keystore_location FeatureEngineeringKafkaConfig#keystore_location}.
	KeystoreLocation *string `field:"required" json:"keystoreLocation" yaml:"keystoreLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/feature_engineering_kafka_config#keystore_password_ref FeatureEngineeringKafkaConfig#keystore_password_ref}.
	KeystorePasswordRef *FeatureEngineeringKafkaConfigAuthConfigMtlsConfigKeystorePasswordRef `field:"required" json:"keystorePasswordRef" yaml:"keystorePasswordRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/feature_engineering_kafka_config#truststore_location FeatureEngineeringKafkaConfig#truststore_location}.
	TruststoreLocation *string `field:"required" json:"truststoreLocation" yaml:"truststoreLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/feature_engineering_kafka_config#truststore_password_ref FeatureEngineeringKafkaConfig#truststore_password_ref}.
	TruststorePasswordRef *FeatureEngineeringKafkaConfigAuthConfigMtlsConfigTruststorePasswordRef `field:"required" json:"truststorePasswordRef" yaml:"truststorePasswordRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/feature_engineering_kafka_config#disable_hostname_verification FeatureEngineeringKafkaConfig#disable_hostname_verification}.
	DisableHostnameVerification interface{} `field:"optional" json:"disableHostnameVerification" yaml:"disableHostnameVerification"`
}

