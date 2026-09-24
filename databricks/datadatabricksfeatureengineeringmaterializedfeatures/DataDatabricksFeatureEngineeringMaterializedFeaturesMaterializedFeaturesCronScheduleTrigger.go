// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringmaterializedfeatures


type DataDatabricksFeatureEngineeringMaterializedFeaturesMaterializedFeaturesCronScheduleTrigger struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/data-sources/feature_engineering_materialized_features#cron_expression DataDatabricksFeatureEngineeringMaterializedFeatures#cron_expression}.
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/data-sources/feature_engineering_materialized_features#mode DataDatabricksFeatureEngineeringMaterializedFeatures#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/data-sources/feature_engineering_materialized_features#timezone_id DataDatabricksFeatureEngineeringMaterializedFeatures#timezone_id}.
	TimezoneId *string `field:"optional" json:"timezoneId" yaml:"timezoneId"`
}

