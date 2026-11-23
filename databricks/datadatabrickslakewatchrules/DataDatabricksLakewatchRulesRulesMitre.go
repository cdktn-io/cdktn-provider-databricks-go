// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchrules


type DataDatabricksLakewatchRulesRulesMitre struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rules#sub_technique DataDatabricksLakewatchRules#sub_technique}.
	SubTechnique *string `field:"optional" json:"subTechnique" yaml:"subTechnique"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rules#sub_technique_id DataDatabricksLakewatchRules#sub_technique_id}.
	SubTechniqueId *string `field:"optional" json:"subTechniqueId" yaml:"subTechniqueId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rules#tactic DataDatabricksLakewatchRules#tactic}.
	Tactic *string `field:"optional" json:"tactic" yaml:"tactic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rules#taxonomy DataDatabricksLakewatchRules#taxonomy}.
	Taxonomy *string `field:"optional" json:"taxonomy" yaml:"taxonomy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rules#technique DataDatabricksLakewatchRules#technique}.
	Technique *string `field:"optional" json:"technique" yaml:"technique"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rules#technique_id DataDatabricksLakewatchRules#technique_id}.
	TechniqueId *string `field:"optional" json:"techniqueId" yaml:"techniqueId"`
}

