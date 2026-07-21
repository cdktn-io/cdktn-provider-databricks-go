// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature


type FeatureEngineeringFeatureSourceDeltaTableSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/feature_engineering_feature#full_name FeatureEngineeringFeature#full_name}.
	FullName *string `field:"required" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/feature_engineering_feature#dataframe_schema FeatureEngineeringFeature#dataframe_schema}.
	DataframeSchema *string `field:"optional" json:"dataframeSchema" yaml:"dataframeSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/feature_engineering_feature#entity_columns FeatureEngineeringFeature#entity_columns}.
	EntityColumns *[]*string `field:"optional" json:"entityColumns" yaml:"entityColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/feature_engineering_feature#filter_condition FeatureEngineeringFeature#filter_condition}.
	FilterCondition *string `field:"optional" json:"filterCondition" yaml:"filterCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/feature_engineering_feature#timeseries_column FeatureEngineeringFeature#timeseries_column}.
	TimeseriesColumn *string `field:"optional" json:"timeseriesColumn" yaml:"timeseriesColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/feature_engineering_feature#transformation_sql FeatureEngineeringFeature#transformation_sql}.
	TransformationSql *string `field:"optional" json:"transformationSql" yaml:"transformationSql"`
}

