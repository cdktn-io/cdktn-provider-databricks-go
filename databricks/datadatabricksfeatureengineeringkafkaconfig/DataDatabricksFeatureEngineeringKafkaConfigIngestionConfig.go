// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfig


type DataDatabricksFeatureEngineeringKafkaConfigIngestionConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/feature_engineering_kafka_config#ingestion_destination DataDatabricksFeatureEngineeringKafkaConfig#ingestion_destination}.
	IngestionDestination *DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigIngestionDestination `field:"required" json:"ingestionDestination" yaml:"ingestionDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/feature_engineering_kafka_config#backfill_source DataDatabricksFeatureEngineeringKafkaConfig#backfill_source}.
	BackfillSource *DataDatabricksFeatureEngineeringKafkaConfigIngestionConfigBackfillSource `field:"optional" json:"backfillSource" yaml:"backfillSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/feature_engineering_kafka_config#deduplication_columns DataDatabricksFeatureEngineeringKafkaConfig#deduplication_columns}.
	DeduplicationColumns *[]*string `field:"optional" json:"deduplicationColumns" yaml:"deduplicationColumns"`
}

