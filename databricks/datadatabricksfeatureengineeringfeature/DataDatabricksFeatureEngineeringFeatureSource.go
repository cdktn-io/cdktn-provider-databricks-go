// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeature


type DataDatabricksFeatureEngineeringFeatureSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#delta_table_source DataDatabricksFeatureEngineeringFeature#delta_table_source}.
	DeltaTableSource *DataDatabricksFeatureEngineeringFeatureSourceDeltaTableSource `field:"optional" json:"deltaTableSource" yaml:"deltaTableSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#kafka_source DataDatabricksFeatureEngineeringFeature#kafka_source}.
	KafkaSource *DataDatabricksFeatureEngineeringFeatureSourceKafkaSource `field:"optional" json:"kafkaSource" yaml:"kafkaSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#request_source DataDatabricksFeatureEngineeringFeature#request_source}.
	RequestSource *DataDatabricksFeatureEngineeringFeatureSourceRequestSource `field:"optional" json:"requestSource" yaml:"requestSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#stream_source DataDatabricksFeatureEngineeringFeature#stream_source}.
	StreamSource *DataDatabricksFeatureEngineeringFeatureSourceStreamSource `field:"optional" json:"streamSource" yaml:"streamSource"`
}

