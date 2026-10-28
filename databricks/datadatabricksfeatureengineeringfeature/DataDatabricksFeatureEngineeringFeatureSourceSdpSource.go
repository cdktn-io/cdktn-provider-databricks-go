// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeature


type DataDatabricksFeatureEngineeringFeatureSourceSdpSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/feature_engineering_feature#dataframe_schema DataDatabricksFeatureEngineeringFeature#dataframe_schema}.
	DataframeSchema *string `field:"optional" json:"dataframeSchema" yaml:"dataframeSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/feature_engineering_feature#filter_condition DataDatabricksFeatureEngineeringFeature#filter_condition}.
	FilterCondition *string `field:"optional" json:"filterCondition" yaml:"filterCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/feature_engineering_feature#input_tables DataDatabricksFeatureEngineeringFeature#input_tables}.
	InputTables *[]*string `field:"optional" json:"inputTables" yaml:"inputTables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/feature_engineering_feature#output_table DataDatabricksFeatureEngineeringFeature#output_table}.
	OutputTable *string `field:"optional" json:"outputTable" yaml:"outputTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/feature_engineering_feature#pipeline_id DataDatabricksFeatureEngineeringFeature#pipeline_id}.
	PipelineId *string `field:"optional" json:"pipelineId" yaml:"pipelineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/feature_engineering_feature#transformation_sql DataDatabricksFeatureEngineeringFeature#transformation_sql}.
	TransformationSql *string `field:"optional" json:"transformationSql" yaml:"transformationSql"`
}

