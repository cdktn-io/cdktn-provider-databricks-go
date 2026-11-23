// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchrule


type DataDatabricksLakewatchRuleOutput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rule#context DataDatabricksLakewatchRule#context}.
	Context *map[string]*string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rule#default_context DataDatabricksLakewatchRule#default_context}.
	DefaultContext interface{} `field:"optional" json:"defaultContext" yaml:"defaultContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rule#event_time_column DataDatabricksLakewatchRule#event_time_column}.
	EventTimeColumn *string `field:"optional" json:"eventTimeColumn" yaml:"eventTimeColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rule#summary DataDatabricksLakewatchRule#summary}.
	Summary *string `field:"optional" json:"summary" yaml:"summary"`
}

