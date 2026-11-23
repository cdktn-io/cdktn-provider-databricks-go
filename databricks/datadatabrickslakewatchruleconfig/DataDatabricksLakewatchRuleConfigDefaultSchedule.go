// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchruleconfig


type DataDatabricksLakewatchRuleConfigDefaultSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rule_config#at_least_every DataDatabricksLakewatchRuleConfigA#at_least_every}.
	AtLeastEvery *string `field:"optional" json:"atLeastEvery" yaml:"atLeastEvery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rule_config#cron_expression DataDatabricksLakewatchRuleConfigA#cron_expression}.
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rule_config#enabled DataDatabricksLakewatchRuleConfigA#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

