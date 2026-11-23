// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchdatasources


type DataDatabricksLakewatchDatasourcesDatasourcesUnityCatalogConnection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_datasources#connection_name DataDatabricksLakewatchDatasources#connection_name}.
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_datasources#configuration DataDatabricksLakewatchDatasources#configuration}.
	Configuration *map[string]*string `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_datasources#start_time DataDatabricksLakewatchDatasources#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

