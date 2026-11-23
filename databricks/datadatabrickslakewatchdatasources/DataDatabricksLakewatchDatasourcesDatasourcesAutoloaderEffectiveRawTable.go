// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchdatasources


type DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderEffectiveRawTable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_datasources#table_name DataDatabricksLakewatchDatasources#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_datasources#catalog_name DataDatabricksLakewatchDatasources#catalog_name}.
	CatalogName *string `field:"optional" json:"catalogName" yaml:"catalogName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_datasources#schema_name DataDatabricksLakewatchDatasources#schema_name}.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
}

