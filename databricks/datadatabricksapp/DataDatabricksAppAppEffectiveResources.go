// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapp


type DataDatabricksAppAppEffectiveResources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/app#name DataDatabricksApp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/app#app DataDatabricksApp#app}.
	App *DataDatabricksAppAppEffectiveResourcesApp `field:"optional" json:"app" yaml:"app"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/app#database DataDatabricksApp#database}.
	Database *DataDatabricksAppAppEffectiveResourcesDatabase `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/app#description DataDatabricksApp#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/app#experiment DataDatabricksApp#experiment}.
	Experiment *DataDatabricksAppAppEffectiveResourcesExperiment `field:"optional" json:"experiment" yaml:"experiment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/app#genie_space DataDatabricksApp#genie_space}.
	GenieSpace *DataDatabricksAppAppEffectiveResourcesGenieSpace `field:"optional" json:"genieSpace" yaml:"genieSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/app#job DataDatabricksApp#job}.
	Job *DataDatabricksAppAppEffectiveResourcesJob `field:"optional" json:"job" yaml:"job"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/app#postgres DataDatabricksApp#postgres}.
	Postgres *DataDatabricksAppAppEffectiveResourcesPostgres `field:"optional" json:"postgres" yaml:"postgres"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/app#secret DataDatabricksApp#secret}.
	Secret *DataDatabricksAppAppEffectiveResourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/app#serving_endpoint DataDatabricksApp#serving_endpoint}.
	ServingEndpoint *DataDatabricksAppAppEffectiveResourcesServingEndpoint `field:"optional" json:"servingEndpoint" yaml:"servingEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/app#sql_warehouse DataDatabricksApp#sql_warehouse}.
	SqlWarehouse *DataDatabricksAppAppEffectiveResourcesSqlWarehouse `field:"optional" json:"sqlWarehouse" yaml:"sqlWarehouse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/app#uc_securable DataDatabricksApp#uc_securable}.
	UcSecurable *DataDatabricksAppAppEffectiveResourcesUcSecurable `field:"optional" json:"ucSecurable" yaml:"ucSecurable"`
}

