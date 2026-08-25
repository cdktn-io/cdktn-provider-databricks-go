// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeature


type DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#approx_count_distinct DataDatabricksFeatureEngineeringFeature#approx_count_distinct}.
	ApproxCountDistinct *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionApproxCountDistinct `field:"optional" json:"approxCountDistinct" yaml:"approxCountDistinct"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#approx_percentile DataDatabricksFeatureEngineeringFeature#approx_percentile}.
	ApproxPercentile *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionApproxPercentile `field:"optional" json:"approxPercentile" yaml:"approxPercentile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#avg DataDatabricksFeatureEngineeringFeature#avg}.
	Avg *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionAvg `field:"optional" json:"avg" yaml:"avg"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#count_function DataDatabricksFeatureEngineeringFeature#count_function}.
	CountFunction *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionCountFunction `field:"optional" json:"countFunction" yaml:"countFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#first DataDatabricksFeatureEngineeringFeature#first}.
	First *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionFirst `field:"optional" json:"first" yaml:"first"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#first_distinct DataDatabricksFeatureEngineeringFeature#first_distinct}.
	FirstDistinct *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionFirstDistinct `field:"optional" json:"firstDistinct" yaml:"firstDistinct"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#first_n DataDatabricksFeatureEngineeringFeature#first_n}.
	FirstN *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionFirstN `field:"optional" json:"firstN" yaml:"firstN"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#last DataDatabricksFeatureEngineeringFeature#last}.
	Last *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionLast `field:"optional" json:"last" yaml:"last"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#last_distinct DataDatabricksFeatureEngineeringFeature#last_distinct}.
	LastDistinct *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionLastDistinct `field:"optional" json:"lastDistinct" yaml:"lastDistinct"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#last_n DataDatabricksFeatureEngineeringFeature#last_n}.
	LastN *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionLastN `field:"optional" json:"lastN" yaml:"lastN"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#max DataDatabricksFeatureEngineeringFeature#max}.
	Max *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionMax `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#min DataDatabricksFeatureEngineeringFeature#min}.
	Min *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionMin `field:"optional" json:"min" yaml:"min"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#stddev_pop DataDatabricksFeatureEngineeringFeature#stddev_pop}.
	StddevPop *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionStddevPop `field:"optional" json:"stddevPop" yaml:"stddevPop"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#stddev_samp DataDatabricksFeatureEngineeringFeature#stddev_samp}.
	StddevSamp *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionStddevSamp `field:"optional" json:"stddevSamp" yaml:"stddevSamp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#sum DataDatabricksFeatureEngineeringFeature#sum}.
	Sum *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionSum `field:"optional" json:"sum" yaml:"sum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#time_window DataDatabricksFeatureEngineeringFeature#time_window}.
	TimeWindow *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindow `field:"optional" json:"timeWindow" yaml:"timeWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#var_pop DataDatabricksFeatureEngineeringFeature#var_pop}.
	VarPop *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionVarPop `field:"optional" json:"varPop" yaml:"varPop"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#var_samp DataDatabricksFeatureEngineeringFeature#var_samp}.
	VarSamp *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionVarSamp `field:"optional" json:"varSamp" yaml:"varSamp"`
}

