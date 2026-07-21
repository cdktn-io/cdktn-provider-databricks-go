// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FeatureEngineeringFeatureConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/feature_engineering_feature#full_name FeatureEngineeringFeature#full_name}.
	FullName *string `field:"required" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/feature_engineering_feature#function FeatureEngineeringFeature#function}.
	Function *FeatureEngineeringFeatureFunction `field:"required" json:"function" yaml:"function"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/feature_engineering_feature#source FeatureEngineeringFeature#source}.
	Source *FeatureEngineeringFeatureSource `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/feature_engineering_feature#description FeatureEngineeringFeature#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/feature_engineering_feature#entities FeatureEngineeringFeature#entities}.
	Entities interface{} `field:"optional" json:"entities" yaml:"entities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/feature_engineering_feature#filter_condition FeatureEngineeringFeature#filter_condition}.
	FilterCondition *string `field:"optional" json:"filterCondition" yaml:"filterCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/feature_engineering_feature#inputs FeatureEngineeringFeature#inputs}.
	Inputs *[]*string `field:"optional" json:"inputs" yaml:"inputs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/feature_engineering_feature#lineage_context FeatureEngineeringFeature#lineage_context}.
	LineageContext *FeatureEngineeringFeatureLineageContext `field:"optional" json:"lineageContext" yaml:"lineageContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/feature_engineering_feature#provider_config FeatureEngineeringFeature#provider_config}.
	ProviderConfig *FeatureEngineeringFeatureProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/feature_engineering_feature#timeseries_column FeatureEngineeringFeature#timeseries_column}.
	TimeseriesColumn *FeatureEngineeringFeatureTimeseriesColumn `field:"optional" json:"timeseriesColumn" yaml:"timeseriesColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/feature_engineering_feature#time_window FeatureEngineeringFeature#time_window}.
	TimeWindow *FeatureEngineeringFeatureTimeWindow `field:"optional" json:"timeWindow" yaml:"timeWindow"`
}

