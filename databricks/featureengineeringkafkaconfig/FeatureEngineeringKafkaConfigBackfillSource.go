// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringkafkaconfig


type FeatureEngineeringKafkaConfigBackfillSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/feature_engineering_kafka_config#delta_table_source FeatureEngineeringKafkaConfig#delta_table_source}.
	DeltaTableSource *FeatureEngineeringKafkaConfigBackfillSourceDeltaTableSource `field:"optional" json:"deltaTableSource" yaml:"deltaTableSource"`
}

