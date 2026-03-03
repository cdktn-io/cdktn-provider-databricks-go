// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappsspaces


type DataDatabricksAppsSpacesSpacesResources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_spaces#name DataDatabricksAppsSpaces#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_spaces#app DataDatabricksAppsSpaces#app}.
	App *DataDatabricksAppsSpacesSpacesResourcesApp `field:"optional" json:"app" yaml:"app"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_spaces#database DataDatabricksAppsSpaces#database}.
	Database *DataDatabricksAppsSpacesSpacesResourcesDatabase `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_spaces#description DataDatabricksAppsSpaces#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_spaces#experiment DataDatabricksAppsSpaces#experiment}.
	Experiment *DataDatabricksAppsSpacesSpacesResourcesExperiment `field:"optional" json:"experiment" yaml:"experiment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_spaces#genie_space DataDatabricksAppsSpaces#genie_space}.
	GenieSpace *DataDatabricksAppsSpacesSpacesResourcesGenieSpace `field:"optional" json:"genieSpace" yaml:"genieSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_spaces#job DataDatabricksAppsSpaces#job}.
	Job *DataDatabricksAppsSpacesSpacesResourcesJob `field:"optional" json:"job" yaml:"job"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_spaces#secret DataDatabricksAppsSpaces#secret}.
	Secret *DataDatabricksAppsSpacesSpacesResourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_spaces#serving_endpoint DataDatabricksAppsSpaces#serving_endpoint}.
	ServingEndpoint *DataDatabricksAppsSpacesSpacesResourcesServingEndpoint `field:"optional" json:"servingEndpoint" yaml:"servingEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_spaces#sql_warehouse DataDatabricksAppsSpaces#sql_warehouse}.
	SqlWarehouse *DataDatabricksAppsSpacesSpacesResourcesSqlWarehouse `field:"optional" json:"sqlWarehouse" yaml:"sqlWarehouse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_spaces#uc_securable DataDatabricksAppsSpaces#uc_securable}.
	UcSecurable *DataDatabricksAppsSpacesSpacesResourcesUcSecurable `field:"optional" json:"ucSecurable" yaml:"ucSecurable"`
}

