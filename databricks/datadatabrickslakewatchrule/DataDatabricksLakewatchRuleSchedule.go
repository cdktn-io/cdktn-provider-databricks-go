// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchrule


type DataDatabricksLakewatchRuleSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/lakewatch_rule#at_least_every DataDatabricksLakewatchRule#at_least_every}.
	AtLeastEvery *string `field:"optional" json:"atLeastEvery" yaml:"atLeastEvery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/lakewatch_rule#cron_expression DataDatabricksLakewatchRule#cron_expression}.
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/lakewatch_rule#enabled DataDatabricksLakewatchRule#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

