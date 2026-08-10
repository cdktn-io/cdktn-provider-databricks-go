// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgressyncedtable


type DataDatabricksPostgresSyncedTableSpecNewPipelineSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/postgres_synced_table#budget_policy_id DataDatabricksPostgresSyncedTable#budget_policy_id}.
	BudgetPolicyId *string `field:"optional" json:"budgetPolicyId" yaml:"budgetPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/postgres_synced_table#storage_catalog DataDatabricksPostgresSyncedTable#storage_catalog}.
	StorageCatalog *string `field:"optional" json:"storageCatalog" yaml:"storageCatalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/postgres_synced_table#storage_schema DataDatabricksPostgresSyncedTable#storage_schema}.
	StorageSchema *string `field:"optional" json:"storageSchema" yaml:"storageSchema"`
}

