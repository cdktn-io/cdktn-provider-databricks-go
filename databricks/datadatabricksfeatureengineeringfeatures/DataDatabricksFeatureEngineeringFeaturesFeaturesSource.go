// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeatures


type DataDatabricksFeatureEngineeringFeaturesFeaturesSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_features#delta_table_source DataDatabricksFeatureEngineeringFeatures#delta_table_source}.
	DeltaTableSource *DataDatabricksFeatureEngineeringFeaturesFeaturesSourceDeltaTableSource `field:"optional" json:"deltaTableSource" yaml:"deltaTableSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_features#kafka_source DataDatabricksFeatureEngineeringFeatures#kafka_source}.
	KafkaSource *DataDatabricksFeatureEngineeringFeaturesFeaturesSourceKafkaSource `field:"optional" json:"kafkaSource" yaml:"kafkaSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_features#request_source DataDatabricksFeatureEngineeringFeatures#request_source}.
	RequestSource *DataDatabricksFeatureEngineeringFeaturesFeaturesSourceRequestSource `field:"optional" json:"requestSource" yaml:"requestSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_features#stream_source DataDatabricksFeatureEngineeringFeatures#stream_source}.
	StreamSource *DataDatabricksFeatureEngineeringFeaturesFeaturesSourceStreamSource `field:"optional" json:"streamSource" yaml:"streamSource"`
}

