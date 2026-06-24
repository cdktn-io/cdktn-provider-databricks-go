// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfigs


type DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsBackfillSourceDeltaTableSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/feature_engineering_kafka_configs#full_name DataDatabricksFeatureEngineeringKafkaConfigs#full_name}.
	FullName *string `field:"required" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/feature_engineering_kafka_configs#dataframe_schema DataDatabricksFeatureEngineeringKafkaConfigs#dataframe_schema}.
	DataframeSchema *string `field:"optional" json:"dataframeSchema" yaml:"dataframeSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/feature_engineering_kafka_configs#entity_columns DataDatabricksFeatureEngineeringKafkaConfigs#entity_columns}.
	EntityColumns *[]*string `field:"optional" json:"entityColumns" yaml:"entityColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/feature_engineering_kafka_configs#filter_condition DataDatabricksFeatureEngineeringKafkaConfigs#filter_condition}.
	FilterCondition *string `field:"optional" json:"filterCondition" yaml:"filterCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/feature_engineering_kafka_configs#timeseries_column DataDatabricksFeatureEngineeringKafkaConfigs#timeseries_column}.
	TimeseriesColumn *string `field:"optional" json:"timeseriesColumn" yaml:"timeseriesColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/feature_engineering_kafka_configs#transformation_sql DataDatabricksFeatureEngineeringKafkaConfigs#transformation_sql}.
	TransformationSql *string `field:"optional" json:"transformationSql" yaml:"transformationSql"`
}

