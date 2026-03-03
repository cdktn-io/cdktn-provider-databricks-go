// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfigs


type DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsBackfillSourceDeltaTableSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/feature_engineering_kafka_configs#entity_columns DataDatabricksFeatureEngineeringKafkaConfigs#entity_columns}.
	EntityColumns *[]*string `field:"required" json:"entityColumns" yaml:"entityColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/feature_engineering_kafka_configs#full_name DataDatabricksFeatureEngineeringKafkaConfigs#full_name}.
	FullName *string `field:"required" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/feature_engineering_kafka_configs#timeseries_column DataDatabricksFeatureEngineeringKafkaConfigs#timeseries_column}.
	TimeseriesColumn *string `field:"required" json:"timeseriesColumn" yaml:"timeseriesColumn"`
}

