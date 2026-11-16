// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakewatchrule


type LakewatchRuleBatchSql struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/lakewatch_rule#query LakewatchRule#query}.
	Query *string `field:"required" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/lakewatch_rule#default_catalog_name LakewatchRule#default_catalog_name}.
	DefaultCatalogName *string `field:"optional" json:"defaultCatalogName" yaml:"defaultCatalogName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/lakewatch_rule#default_schema_name LakewatchRule#default_schema_name}.
	DefaultSchemaName *string `field:"optional" json:"defaultSchemaName" yaml:"defaultSchemaName"`
}

