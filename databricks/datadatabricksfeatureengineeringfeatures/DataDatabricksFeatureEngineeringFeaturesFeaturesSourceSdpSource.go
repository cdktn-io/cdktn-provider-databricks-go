// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeatures


type DataDatabricksFeatureEngineeringFeaturesFeaturesSourceSdpSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/feature_engineering_features#dataframe_schema DataDatabricksFeatureEngineeringFeatures#dataframe_schema}.
	DataframeSchema *string `field:"optional" json:"dataframeSchema" yaml:"dataframeSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/feature_engineering_features#filter_condition DataDatabricksFeatureEngineeringFeatures#filter_condition}.
	FilterCondition *string `field:"optional" json:"filterCondition" yaml:"filterCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/feature_engineering_features#input_tables DataDatabricksFeatureEngineeringFeatures#input_tables}.
	InputTables *[]*string `field:"optional" json:"inputTables" yaml:"inputTables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/feature_engineering_features#output_table DataDatabricksFeatureEngineeringFeatures#output_table}.
	OutputTable *string `field:"optional" json:"outputTable" yaml:"outputTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/feature_engineering_features#pipeline_id DataDatabricksFeatureEngineeringFeatures#pipeline_id}.
	PipelineId *string `field:"optional" json:"pipelineId" yaml:"pipelineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/feature_engineering_features#transformation_sql DataDatabricksFeatureEngineeringFeatures#transformation_sql}.
	TransformationSql *string `field:"optional" json:"transformationSql" yaml:"transformationSql"`
}

