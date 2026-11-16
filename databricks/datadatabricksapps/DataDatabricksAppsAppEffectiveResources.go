// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapps


type DataDatabricksAppsAppEffectiveResources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/apps#name DataDatabricksApps#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/apps#app DataDatabricksApps#app}.
	App *DataDatabricksAppsAppEffectiveResourcesApp `field:"optional" json:"app" yaml:"app"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/apps#database DataDatabricksApps#database}.
	Database *DataDatabricksAppsAppEffectiveResourcesDatabase `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/apps#description DataDatabricksApps#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/apps#experiment DataDatabricksApps#experiment}.
	Experiment *DataDatabricksAppsAppEffectiveResourcesExperiment `field:"optional" json:"experiment" yaml:"experiment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/apps#genie_space DataDatabricksApps#genie_space}.
	GenieSpace *DataDatabricksAppsAppEffectiveResourcesGenieSpace `field:"optional" json:"genieSpace" yaml:"genieSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/apps#job DataDatabricksApps#job}.
	Job *DataDatabricksAppsAppEffectiveResourcesJob `field:"optional" json:"job" yaml:"job"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/apps#postgres DataDatabricksApps#postgres}.
	Postgres *DataDatabricksAppsAppEffectiveResourcesPostgres `field:"optional" json:"postgres" yaml:"postgres"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/apps#secret DataDatabricksApps#secret}.
	Secret *DataDatabricksAppsAppEffectiveResourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/apps#serving_endpoint DataDatabricksApps#serving_endpoint}.
	ServingEndpoint *DataDatabricksAppsAppEffectiveResourcesServingEndpoint `field:"optional" json:"servingEndpoint" yaml:"servingEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/apps#sql_warehouse DataDatabricksApps#sql_warehouse}.
	SqlWarehouse *DataDatabricksAppsAppEffectiveResourcesSqlWarehouse `field:"optional" json:"sqlWarehouse" yaml:"sqlWarehouse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/apps#uc_securable DataDatabricksApps#uc_securable}.
	UcSecurable *DataDatabricksAppsAppEffectiveResourcesUcSecurable `field:"optional" json:"ucSecurable" yaml:"ucSecurable"`
}

