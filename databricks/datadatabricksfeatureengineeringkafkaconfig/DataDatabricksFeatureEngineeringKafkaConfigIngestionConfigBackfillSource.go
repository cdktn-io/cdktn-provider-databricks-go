// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfig


type DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigBackfillSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/feature_engineering_kafka_config#delta_table_name DataDatabricksFeatureEngineeringKafkaConfig#delta_table_name}.
	DeltaTableName *string `field:"optional" json:"deltaTableName" yaml:"deltaTableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/feature_engineering_kafka_config#delta_table_source DataDatabricksFeatureEngineeringKafkaConfig#delta_table_source}.
	DeltaTableSource *DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSource `field:"optional" json:"deltaTableSource" yaml:"deltaTableSource"`
}

