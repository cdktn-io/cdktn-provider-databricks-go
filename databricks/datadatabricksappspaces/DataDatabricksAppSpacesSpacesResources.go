// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappspaces


type DataDatabricksAppSpacesSpacesResources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/app_spaces#name DataDatabricksAppSpaces#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/app_spaces#app DataDatabricksAppSpaces#app}.
	App *DataDatabricksAppSpacesSpacesResourcesApp `field:"optional" json:"app" yaml:"app"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/app_spaces#database DataDatabricksAppSpaces#database}.
	Database *DataDatabricksAppSpacesSpacesResourcesDatabase `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/app_spaces#description DataDatabricksAppSpaces#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/app_spaces#experiment DataDatabricksAppSpaces#experiment}.
	Experiment *DataDatabricksAppSpacesSpacesResourcesExperiment `field:"optional" json:"experiment" yaml:"experiment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/app_spaces#genie_space DataDatabricksAppSpaces#genie_space}.
	GenieSpace *DataDatabricksAppSpacesSpacesResourcesGenieSpace `field:"optional" json:"genieSpace" yaml:"genieSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/app_spaces#job DataDatabricksAppSpaces#job}.
	Job *DataDatabricksAppSpacesSpacesResourcesJob `field:"optional" json:"job" yaml:"job"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/app_spaces#postgres DataDatabricksAppSpaces#postgres}.
	Postgres *DataDatabricksAppSpacesSpacesResourcesPostgres `field:"optional" json:"postgres" yaml:"postgres"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/app_spaces#secret DataDatabricksAppSpaces#secret}.
	Secret *DataDatabricksAppSpacesSpacesResourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/app_spaces#serving_endpoint DataDatabricksAppSpaces#serving_endpoint}.
	ServingEndpoint *DataDatabricksAppSpacesSpacesResourcesServingEndpoint `field:"optional" json:"servingEndpoint" yaml:"servingEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/app_spaces#sql_warehouse DataDatabricksAppSpaces#sql_warehouse}.
	SqlWarehouse *DataDatabricksAppSpacesSpacesResourcesSqlWarehouse `field:"optional" json:"sqlWarehouse" yaml:"sqlWarehouse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/app_spaces#uc_securable DataDatabricksAppSpaces#uc_securable}.
	UcSecurable *DataDatabricksAppSpacesSpacesResourcesUcSecurable `field:"optional" json:"ucSecurable" yaml:"ucSecurable"`
}

