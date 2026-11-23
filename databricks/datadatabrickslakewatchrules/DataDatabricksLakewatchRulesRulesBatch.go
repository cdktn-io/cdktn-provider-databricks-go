// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchrules


type DataDatabricksLakewatchRulesRulesBatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rules#custom DataDatabricksLakewatchRules#custom}.
	Custom *DataDatabricksLakewatchRulesRulesBatchCustom `field:"optional" json:"custom" yaml:"custom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rules#sql DataDatabricksLakewatchRules#sql}.
	Sql *DataDatabricksLakewatchRulesRulesBatchSql `field:"optional" json:"sql" yaml:"sql"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_rules#warehouse_id DataDatabricksLakewatchRules#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

