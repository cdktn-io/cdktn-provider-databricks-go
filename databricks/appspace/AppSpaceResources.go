// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appspace


type AppSpaceResources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/app_space#name AppSpace#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/app_space#app AppSpace#app}.
	App *AppSpaceResourcesApp `field:"optional" json:"app" yaml:"app"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/app_space#database AppSpace#database}.
	Database *AppSpaceResourcesDatabase `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/app_space#description AppSpace#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/app_space#experiment AppSpace#experiment}.
	Experiment *AppSpaceResourcesExperiment `field:"optional" json:"experiment" yaml:"experiment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/app_space#genie_space AppSpace#genie_space}.
	GenieSpace *AppSpaceResourcesGenieSpace `field:"optional" json:"genieSpace" yaml:"genieSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/app_space#job AppSpace#job}.
	Job *AppSpaceResourcesJob `field:"optional" json:"job" yaml:"job"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/app_space#postgres AppSpace#postgres}.
	Postgres *AppSpaceResourcesPostgres `field:"optional" json:"postgres" yaml:"postgres"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/app_space#secret AppSpace#secret}.
	Secret *AppSpaceResourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/app_space#serving_endpoint AppSpace#serving_endpoint}.
	ServingEndpoint *AppSpaceResourcesServingEndpoint `field:"optional" json:"servingEndpoint" yaml:"servingEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/app_space#sql_warehouse AppSpace#sql_warehouse}.
	SqlWarehouse *AppSpaceResourcesSqlWarehouse `field:"optional" json:"sqlWarehouse" yaml:"sqlWarehouse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/app_space#uc_securable AppSpace#uc_securable}.
	UcSecurable *AppSpaceResourcesUcSecurable `field:"optional" json:"ucSecurable" yaml:"ucSecurable"`
}

