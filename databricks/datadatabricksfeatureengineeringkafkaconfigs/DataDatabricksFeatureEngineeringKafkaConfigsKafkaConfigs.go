// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfigs


type DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/feature_engineering_kafka_configs#name DataDatabricksFeatureEngineeringKafkaConfigs#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/feature_engineering_kafka_configs#provider_config DataDatabricksFeatureEngineeringKafkaConfigs#provider_config}.
	ProviderConfig *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

