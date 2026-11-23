// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchrules


type DataDatabricksLakewatchRulesRulesBatchCustom struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rules#notebook DataDatabricksLakewatchRules#notebook}.
	Notebook *string `field:"required" json:"notebook" yaml:"notebook"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rules#options DataDatabricksLakewatchRules#options}.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
}

