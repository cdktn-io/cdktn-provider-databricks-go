// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchrules


type DataDatabricksLakewatchRulesRulesSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rules#at_least_every DataDatabricksLakewatchRules#at_least_every}.
	AtLeastEvery *string `field:"optional" json:"atLeastEvery" yaml:"atLeastEvery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rules#cron_expression DataDatabricksLakewatchRules#cron_expression}.
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rules#enabled DataDatabricksLakewatchRules#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

