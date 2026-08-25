// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature


type FeatureEngineeringFeatureSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/feature_engineering_feature#delta_table_source FeatureEngineeringFeature#delta_table_source}.
	DeltaTableSource *FeatureEngineeringFeatureSourceDeltaTableSource `field:"optional" json:"deltaTableSource" yaml:"deltaTableSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/feature_engineering_feature#kafka_source FeatureEngineeringFeature#kafka_source}.
	KafkaSource *FeatureEngineeringFeatureSourceKafkaSource `field:"optional" json:"kafkaSource" yaml:"kafkaSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/feature_engineering_feature#request_source FeatureEngineeringFeature#request_source}.
	RequestSource *FeatureEngineeringFeatureSourceRequestSource `field:"optional" json:"requestSource" yaml:"requestSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/feature_engineering_feature#stream_source FeatureEngineeringFeature#stream_source}.
	StreamSource *FeatureEngineeringFeatureSourceStreamSource `field:"optional" json:"streamSource" yaml:"streamSource"`
}

