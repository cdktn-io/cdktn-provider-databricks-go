// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfigs


type DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/feature_engineering_kafka_configs#delta_table_name DataDatabricksFeatureEngineeringKafkaConfigs#delta_table_name}.
	DeltaTableName *string `field:"optional" json:"deltaTableName" yaml:"deltaTableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/feature_engineering_kafka_configs#delta_table_source DataDatabricksFeatureEngineeringKafkaConfigs#delta_table_source}.
	DeltaTableSource *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSourceDeltaTableSource `field:"optional" json:"deltaTableSource" yaml:"deltaTableSource"`
}

