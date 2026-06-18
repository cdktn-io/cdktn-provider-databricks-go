// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringkafkaconfig


type FeatureEngineeringKafkaConfigIngestionConfigBackfillSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/feature_engineering_kafka_config#delta_table_name FeatureEngineeringKafkaConfig#delta_table_name}.
	DeltaTableName *string `field:"optional" json:"deltaTableName" yaml:"deltaTableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/feature_engineering_kafka_config#delta_table_source FeatureEngineeringKafkaConfig#delta_table_source}.
	DeltaTableSource *FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSource `field:"optional" json:"deltaTableSource" yaml:"deltaTableSource"`
}

