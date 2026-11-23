// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchrule


type DataDatabricksLakewatchRuleMitre struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rule#sub_technique DataDatabricksLakewatchRule#sub_technique}.
	SubTechnique *string `field:"optional" json:"subTechnique" yaml:"subTechnique"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rule#sub_technique_id DataDatabricksLakewatchRule#sub_technique_id}.
	SubTechniqueId *string `field:"optional" json:"subTechniqueId" yaml:"subTechniqueId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rule#tactic DataDatabricksLakewatchRule#tactic}.
	Tactic *string `field:"optional" json:"tactic" yaml:"tactic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rule#taxonomy DataDatabricksLakewatchRule#taxonomy}.
	Taxonomy *string `field:"optional" json:"taxonomy" yaml:"taxonomy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rule#technique DataDatabricksLakewatchRule#technique}.
	Technique *string `field:"optional" json:"technique" yaml:"technique"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rule#technique_id DataDatabricksLakewatchRule#technique_id}.
	TechniqueId *string `field:"optional" json:"techniqueId" yaml:"techniqueId"`
}

