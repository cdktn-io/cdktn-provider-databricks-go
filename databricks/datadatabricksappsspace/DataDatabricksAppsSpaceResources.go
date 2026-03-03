// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappsspace


type DataDatabricksAppsSpaceResources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_space#name DataDatabricksAppsSpace#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_space#app DataDatabricksAppsSpace#app}.
	App *DataDatabricksAppsSpaceResourcesApp `field:"optional" json:"app" yaml:"app"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_space#database DataDatabricksAppsSpace#database}.
	Database *DataDatabricksAppsSpaceResourcesDatabase `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_space#description DataDatabricksAppsSpace#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_space#experiment DataDatabricksAppsSpace#experiment}.
	Experiment *DataDatabricksAppsSpaceResourcesExperiment `field:"optional" json:"experiment" yaml:"experiment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_space#genie_space DataDatabricksAppsSpace#genie_space}.
	GenieSpace *DataDatabricksAppsSpaceResourcesGenieSpace `field:"optional" json:"genieSpace" yaml:"genieSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_space#job DataDatabricksAppsSpace#job}.
	Job *DataDatabricksAppsSpaceResourcesJob `field:"optional" json:"job" yaml:"job"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_space#secret DataDatabricksAppsSpace#secret}.
	Secret *DataDatabricksAppsSpaceResourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_space#serving_endpoint DataDatabricksAppsSpace#serving_endpoint}.
	ServingEndpoint *DataDatabricksAppsSpaceResourcesServingEndpoint `field:"optional" json:"servingEndpoint" yaml:"servingEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_space#sql_warehouse DataDatabricksAppsSpace#sql_warehouse}.
	SqlWarehouse *DataDatabricksAppsSpaceResourcesSqlWarehouse `field:"optional" json:"sqlWarehouse" yaml:"sqlWarehouse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_space#uc_securable DataDatabricksAppsSpace#uc_securable}.
	UcSecurable *DataDatabricksAppsSpaceResourcesUcSecurable `field:"optional" json:"ucSecurable" yaml:"ucSecurable"`
}

