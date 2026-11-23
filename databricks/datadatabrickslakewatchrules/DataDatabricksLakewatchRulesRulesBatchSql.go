// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchrules


type DataDatabricksLakewatchRulesRulesBatchSql struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rules#query DataDatabricksLakewatchRules#query}.
	Query *string `field:"required" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rules#default_catalog_name DataDatabricksLakewatchRules#default_catalog_name}.
	DefaultCatalogName *string `field:"optional" json:"defaultCatalogName" yaml:"defaultCatalogName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rules#default_schema_name DataDatabricksLakewatchRules#default_schema_name}.
	DefaultSchemaName *string `field:"optional" json:"defaultSchemaName" yaml:"defaultSchemaName"`
}

