// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakewatchdatasource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LakewatchDatasourceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_datasource#catalog_name LakewatchDatasource#catalog_name}.
	CatalogName *string `field:"required" json:"catalogName" yaml:"catalogName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_datasource#datasource_name LakewatchDatasource#datasource_name}.
	DatasourceName *string `field:"required" json:"datasourceName" yaml:"datasourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_datasource#schedule LakewatchDatasource#schedule}.
	Schedule *LakewatchDatasourceSchedule `field:"required" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_datasource#schema_name LakewatchDatasource#schema_name}.
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_datasource#annotations LakewatchDatasource#annotations}.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_datasource#autoloader LakewatchDatasource#autoloader}.
	Autoloader *LakewatchDatasourceAutoloader `field:"optional" json:"autoloader" yaml:"autoloader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_datasource#compute_mode LakewatchDatasource#compute_mode}.
	ComputeMode *string `field:"optional" json:"computeMode" yaml:"computeMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_datasource#description LakewatchDatasource#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_datasource#existing_table LakewatchDatasource#existing_table}.
	ExistingTable *LakewatchDatasourceExistingTable `field:"optional" json:"existingTable" yaml:"existingTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_datasource#parsers LakewatchDatasource#parsers}.
	Parsers interface{} `field:"optional" json:"parsers" yaml:"parsers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_datasource#provider_config LakewatchDatasource#provider_config}.
	ProviderConfig *LakewatchDatasourceProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_datasource#run_as_principal LakewatchDatasource#run_as_principal}.
	RunAsPrincipal *string `field:"optional" json:"runAsPrincipal" yaml:"runAsPrincipal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_datasource#unity_catalog_connection LakewatchDatasource#unity_catalog_connection}.
	UnityCatalogConnection *LakewatchDatasourceUnityCatalogConnection `field:"optional" json:"unityCatalogConnection" yaml:"unityCatalogConnection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_datasource#warehouse_id LakewatchDatasource#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

