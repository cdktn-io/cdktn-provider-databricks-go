// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakewatchrule


type LakewatchRuleOutput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/lakewatch_rule#context LakewatchRule#context}.
	Context *map[string]*string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/lakewatch_rule#default_context LakewatchRule#default_context}.
	DefaultContext interface{} `field:"optional" json:"defaultContext" yaml:"defaultContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/lakewatch_rule#event_time_column LakewatchRule#event_time_column}.
	EventTimeColumn *string `field:"optional" json:"eventTimeColumn" yaml:"eventTimeColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/lakewatch_rule#summary LakewatchRule#summary}.
	Summary *string `field:"optional" json:"summary" yaml:"summary"`
}

