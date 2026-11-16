// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchdatasources


type DataDatabricksLakewatchDatasourcesDatasources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/lakewatch_datasources#name DataDatabricksLakewatchDatasources#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/lakewatch_datasources#provider_config DataDatabricksLakewatchDatasources#provider_config}.
	ProviderConfig *DataDatabricksLakewatchDatasourcesDatasourcesProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

