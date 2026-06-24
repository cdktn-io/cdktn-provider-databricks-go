// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfig


type DataDatabricksFeatureEngineeringKafkaConfigBackfillSourceDeltaTableSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/feature_engineering_kafka_config#full_name DataDatabricksFeatureEngineeringKafkaConfig#full_name}.
	FullName *string `field:"required" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/feature_engineering_kafka_config#dataframe_schema DataDatabricksFeatureEngineeringKafkaConfig#dataframe_schema}.
	DataframeSchema *string `field:"optional" json:"dataframeSchema" yaml:"dataframeSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/feature_engineering_kafka_config#entity_columns DataDatabricksFeatureEngineeringKafkaConfig#entity_columns}.
	EntityColumns *[]*string `field:"optional" json:"entityColumns" yaml:"entityColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/feature_engineering_kafka_config#filter_condition DataDatabricksFeatureEngineeringKafkaConfig#filter_condition}.
	FilterCondition *string `field:"optional" json:"filterCondition" yaml:"filterCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/feature_engineering_kafka_config#timeseries_column DataDatabricksFeatureEngineeringKafkaConfig#timeseries_column}.
	TimeseriesColumn *string `field:"optional" json:"timeseriesColumn" yaml:"timeseriesColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/feature_engineering_kafka_config#transformation_sql DataDatabricksFeatureEngineeringKafkaConfig#transformation_sql}.
	TransformationSql *string `field:"optional" json:"transformationSql" yaml:"transformationSql"`
}

