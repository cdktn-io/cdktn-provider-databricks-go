// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfigs


type DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/feature_engineering_kafka_configs#ingestion_destination DataDatabricksFeatureEngineeringKafkaConfigs#ingestion_destination}.
	IngestionDestination *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigIngestionDestination `field:"required" json:"ingestionDestination" yaml:"ingestionDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/feature_engineering_kafka_configs#backfill_source DataDatabricksFeatureEngineeringKafkaConfigs#backfill_source}.
	BackfillSource *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsIngestionConfigBackfillSource `field:"optional" json:"backfillSource" yaml:"backfillSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/feature_engineering_kafka_configs#deduplication_columns DataDatabricksFeatureEngineeringKafkaConfigs#deduplication_columns}.
	DeduplicationColumns *[]*string `field:"optional" json:"deduplicationColumns" yaml:"deduplicationColumns"`
}

