// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchdatasource


type DataDatabricksLakewatchDatasourceSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/lakewatch_datasource#at_least_every DataDatabricksLakewatchDatasource#at_least_every}.
	AtLeastEvery *string `field:"optional" json:"atLeastEvery" yaml:"atLeastEvery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/lakewatch_datasource#compute_group DataDatabricksLakewatchDatasource#compute_group}.
	ComputeGroup *string `field:"optional" json:"computeGroup" yaml:"computeGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/lakewatch_datasource#continuous DataDatabricksLakewatchDatasource#continuous}.
	Continuous interface{} `field:"optional" json:"continuous" yaml:"continuous"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/lakewatch_datasource#cron_expression DataDatabricksLakewatchDatasource#cron_expression}.
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/lakewatch_datasource#enabled DataDatabricksLakewatchDatasource#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

