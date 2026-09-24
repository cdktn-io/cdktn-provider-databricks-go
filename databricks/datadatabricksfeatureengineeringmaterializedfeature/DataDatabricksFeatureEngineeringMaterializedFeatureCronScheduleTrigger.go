// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringmaterializedfeature


type DataDatabricksFeatureEngineeringMaterializedFeatureCronScheduleTrigger struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/data-sources/feature_engineering_materialized_feature#cron_expression DataDatabricksFeatureEngineeringMaterializedFeature#cron_expression}.
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/data-sources/feature_engineering_materialized_feature#mode DataDatabricksFeatureEngineeringMaterializedFeature#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/data-sources/feature_engineering_materialized_feature#timezone_id DataDatabricksFeatureEngineeringMaterializedFeature#timezone_id}.
	TimezoneId *string `field:"optional" json:"timezoneId" yaml:"timezoneId"`
}

