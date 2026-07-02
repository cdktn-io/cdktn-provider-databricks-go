// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringkafkaconfig


type FeatureEngineeringKafkaConfigIngestionConfigBackfillSourceDeltaTableSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_kafka_config#full_name FeatureEngineeringKafkaConfig#full_name}.
	FullName *string `field:"required" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_kafka_config#dataframe_schema FeatureEngineeringKafkaConfig#dataframe_schema}.
	DataframeSchema *string `field:"optional" json:"dataframeSchema" yaml:"dataframeSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_kafka_config#entity_columns FeatureEngineeringKafkaConfig#entity_columns}.
	EntityColumns *[]*string `field:"optional" json:"entityColumns" yaml:"entityColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_kafka_config#filter_condition FeatureEngineeringKafkaConfig#filter_condition}.
	FilterCondition *string `field:"optional" json:"filterCondition" yaml:"filterCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_kafka_config#timeseries_column FeatureEngineeringKafkaConfig#timeseries_column}.
	TimeseriesColumn *string `field:"optional" json:"timeseriesColumn" yaml:"timeseriesColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_kafka_config#transformation_sql FeatureEngineeringKafkaConfig#transformation_sql}.
	TransformationSql *string `field:"optional" json:"transformationSql" yaml:"transformationSql"`
}

