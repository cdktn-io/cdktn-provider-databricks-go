// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringmaterializedfeatures


type DataDatabricksFeatureEngineeringMaterializedFeaturesMaterializedFeatures struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/feature_engineering_materialized_features#materialized_feature_id DataDatabricksFeatureEngineeringMaterializedFeatures#materialized_feature_id}.
	MaterializedFeatureId *string `field:"required" json:"materializedFeatureId" yaml:"materializedFeatureId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/feature_engineering_materialized_features#provider_config DataDatabricksFeatureEngineeringMaterializedFeatures#provider_config}.
	ProviderConfig *DataDatabricksFeatureEngineeringMaterializedFeaturesMaterializedFeaturesProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

