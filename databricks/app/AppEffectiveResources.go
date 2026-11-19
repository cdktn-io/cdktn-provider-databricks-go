// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package app


type AppEffectiveResources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/app#name App#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/app#app App#app}.
	App *AppEffectiveResourcesApp `field:"optional" json:"app" yaml:"app"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/app#database App#database}.
	Database *AppEffectiveResourcesDatabase `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/app#description App#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/app#experiment App#experiment}.
	Experiment *AppEffectiveResourcesExperiment `field:"optional" json:"experiment" yaml:"experiment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/app#genie_space App#genie_space}.
	GenieSpace *AppEffectiveResourcesGenieSpace `field:"optional" json:"genieSpace" yaml:"genieSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/app#job App#job}.
	Job *AppEffectiveResourcesJob `field:"optional" json:"job" yaml:"job"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/app#postgres App#postgres}.
	Postgres *AppEffectiveResourcesPostgres `field:"optional" json:"postgres" yaml:"postgres"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/app#secret App#secret}.
	Secret *AppEffectiveResourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/app#serving_endpoint App#serving_endpoint}.
	ServingEndpoint *AppEffectiveResourcesServingEndpoint `field:"optional" json:"servingEndpoint" yaml:"servingEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/app#sql_warehouse App#sql_warehouse}.
	SqlWarehouse *AppEffectiveResourcesSqlWarehouse `field:"optional" json:"sqlWarehouse" yaml:"sqlWarehouse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/app#uc_securable App#uc_securable}.
	UcSecurable *AppEffectiveResourcesUcSecurable `field:"optional" json:"ucSecurable" yaml:"ucSecurable"`
}

