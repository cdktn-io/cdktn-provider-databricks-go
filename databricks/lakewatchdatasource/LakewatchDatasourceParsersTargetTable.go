// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakewatchdatasource


type LakewatchDatasourceParsersTargetTable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_datasource#table_name LakewatchDatasource#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_datasource#catalog_name LakewatchDatasource#catalog_name}.
	CatalogName *string `field:"optional" json:"catalogName" yaml:"catalogName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_datasource#schema_name LakewatchDatasource#schema_name}.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
}

