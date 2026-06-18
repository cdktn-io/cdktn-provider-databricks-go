// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringkafkaconfig


type FeatureEngineeringKafkaConfigIngestionConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/feature_engineering_kafka_config#ingestion_destination FeatureEngineeringKafkaConfig#ingestion_destination}.
	IngestionDestination *FeatureEngineeringKafkaConfigIngestionConfigIngestionDestination `field:"required" json:"ingestionDestination" yaml:"ingestionDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/feature_engineering_kafka_config#backfill_source FeatureEngineeringKafkaConfig#backfill_source}.
	BackfillSource *FeatureEngineeringKafkaConfigIngestionConfigBackfillSource `field:"optional" json:"backfillSource" yaml:"backfillSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/feature_engineering_kafka_config#deduplication_columns FeatureEngineeringKafkaConfig#deduplication_columns}.
	DeduplicationColumns *[]*string `field:"optional" json:"deduplicationColumns" yaml:"deduplicationColumns"`
}

