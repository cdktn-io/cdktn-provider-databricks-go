// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringmaterializedfeature


type FeatureEngineeringMaterializedFeatureStreamingMode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/feature_engineering_materialized_feature#freshness_target FeatureEngineeringMaterializedFeature#freshness_target}.
	FreshnessTarget *string `field:"optional" json:"freshnessTarget" yaml:"freshnessTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/feature_engineering_materialized_feature#mode FeatureEngineeringMaterializedFeature#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

