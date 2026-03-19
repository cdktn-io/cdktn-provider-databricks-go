// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappspace


type DataDatabricksAppSpaceResources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/app_space#name DataDatabricksAppSpace#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/app_space#app DataDatabricksAppSpace#app}.
	App *DataDatabricksAppSpaceResourcesApp `field:"optional" json:"app" yaml:"app"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/app_space#database DataDatabricksAppSpace#database}.
	Database *DataDatabricksAppSpaceResourcesDatabase `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/app_space#description DataDatabricksAppSpace#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/app_space#experiment DataDatabricksAppSpace#experiment}.
	Experiment *DataDatabricksAppSpaceResourcesExperiment `field:"optional" json:"experiment" yaml:"experiment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/app_space#genie_space DataDatabricksAppSpace#genie_space}.
	GenieSpace *DataDatabricksAppSpaceResourcesGenieSpace `field:"optional" json:"genieSpace" yaml:"genieSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/app_space#job DataDatabricksAppSpace#job}.
	Job *DataDatabricksAppSpaceResourcesJob `field:"optional" json:"job" yaml:"job"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/app_space#postgres DataDatabricksAppSpace#postgres}.
	Postgres *DataDatabricksAppSpaceResourcesPostgres `field:"optional" json:"postgres" yaml:"postgres"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/app_space#secret DataDatabricksAppSpace#secret}.
	Secret *DataDatabricksAppSpaceResourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/app_space#serving_endpoint DataDatabricksAppSpace#serving_endpoint}.
	ServingEndpoint *DataDatabricksAppSpaceResourcesServingEndpoint `field:"optional" json:"servingEndpoint" yaml:"servingEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/app_space#sql_warehouse DataDatabricksAppSpace#sql_warehouse}.
	SqlWarehouse *DataDatabricksAppSpaceResourcesSqlWarehouse `field:"optional" json:"sqlWarehouse" yaml:"sqlWarehouse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/app_space#uc_securable DataDatabricksAppSpace#uc_securable}.
	UcSecurable *DataDatabricksAppSpaceResourcesUcSecurable `field:"optional" json:"ucSecurable" yaml:"ucSecurable"`
}

