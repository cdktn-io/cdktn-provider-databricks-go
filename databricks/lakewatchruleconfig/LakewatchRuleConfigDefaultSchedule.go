// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakewatchruleconfig


type LakewatchRuleConfigDefaultSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/lakewatch_rule_config#at_least_every LakewatchRuleConfigA#at_least_every}.
	AtLeastEvery *string `field:"optional" json:"atLeastEvery" yaml:"atLeastEvery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/lakewatch_rule_config#cron_expression LakewatchRuleConfigA#cron_expression}.
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/lakewatch_rule_config#enabled LakewatchRuleConfigA#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

