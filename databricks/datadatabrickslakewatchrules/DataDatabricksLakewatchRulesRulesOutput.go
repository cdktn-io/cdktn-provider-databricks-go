// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchrules


type DataDatabricksLakewatchRulesRulesOutput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/lakewatch_rules#context DataDatabricksLakewatchRules#context}.
	Context *map[string]*string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/lakewatch_rules#default_context DataDatabricksLakewatchRules#default_context}.
	DefaultContext interface{} `field:"optional" json:"defaultContext" yaml:"defaultContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/lakewatch_rules#event_time_column DataDatabricksLakewatchRules#event_time_column}.
	EventTimeColumn *string `field:"optional" json:"eventTimeColumn" yaml:"eventTimeColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/lakewatch_rules#summary DataDatabricksLakewatchRules#summary}.
	Summary *string `field:"optional" json:"summary" yaml:"summary"`
}

