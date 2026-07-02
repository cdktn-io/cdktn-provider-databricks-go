// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature


type FeatureEngineeringFeatureFunctionAggregationFunction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_feature#approx_count_distinct FeatureEngineeringFeature#approx_count_distinct}.
	ApproxCountDistinct *FeatureEngineeringFeatureFunctionAggregationFunctionApproxCountDistinct `field:"optional" json:"approxCountDistinct" yaml:"approxCountDistinct"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_feature#approx_percentile FeatureEngineeringFeature#approx_percentile}.
	ApproxPercentile *FeatureEngineeringFeatureFunctionAggregationFunctionApproxPercentile `field:"optional" json:"approxPercentile" yaml:"approxPercentile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_feature#avg FeatureEngineeringFeature#avg}.
	Avg *FeatureEngineeringFeatureFunctionAggregationFunctionAvg `field:"optional" json:"avg" yaml:"avg"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_feature#count_function FeatureEngineeringFeature#count_function}.
	CountFunction *FeatureEngineeringFeatureFunctionAggregationFunctionCountFunction `field:"optional" json:"countFunction" yaml:"countFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_feature#first FeatureEngineeringFeature#first}.
	First *FeatureEngineeringFeatureFunctionAggregationFunctionFirst `field:"optional" json:"first" yaml:"first"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_feature#last FeatureEngineeringFeature#last}.
	Last *FeatureEngineeringFeatureFunctionAggregationFunctionLast `field:"optional" json:"last" yaml:"last"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_feature#max FeatureEngineeringFeature#max}.
	Max *FeatureEngineeringFeatureFunctionAggregationFunctionMax `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_feature#min FeatureEngineeringFeature#min}.
	Min *FeatureEngineeringFeatureFunctionAggregationFunctionMin `field:"optional" json:"min" yaml:"min"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_feature#stddev_pop FeatureEngineeringFeature#stddev_pop}.
	StddevPop *FeatureEngineeringFeatureFunctionAggregationFunctionStddevPop `field:"optional" json:"stddevPop" yaml:"stddevPop"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_feature#stddev_samp FeatureEngineeringFeature#stddev_samp}.
	StddevSamp *FeatureEngineeringFeatureFunctionAggregationFunctionStddevSamp `field:"optional" json:"stddevSamp" yaml:"stddevSamp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_feature#sum FeatureEngineeringFeature#sum}.
	Sum *FeatureEngineeringFeatureFunctionAggregationFunctionSum `field:"optional" json:"sum" yaml:"sum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_feature#time_window FeatureEngineeringFeature#time_window}.
	TimeWindow *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindow `field:"optional" json:"timeWindow" yaml:"timeWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_feature#var_pop FeatureEngineeringFeature#var_pop}.
	VarPop *FeatureEngineeringFeatureFunctionAggregationFunctionVarPop `field:"optional" json:"varPop" yaml:"varPop"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_feature#var_samp FeatureEngineeringFeature#var_samp}.
	VarSamp *FeatureEngineeringFeatureFunctionAggregationFunctionVarSamp `field:"optional" json:"varSamp" yaml:"varSamp"`
}

