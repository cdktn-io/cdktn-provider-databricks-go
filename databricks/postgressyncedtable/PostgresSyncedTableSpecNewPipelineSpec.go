// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgressyncedtable


type PostgresSyncedTableSpecNewPipelineSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/postgres_synced_table#budget_policy_id PostgresSyncedTable#budget_policy_id}.
	BudgetPolicyId *string `field:"optional" json:"budgetPolicyId" yaml:"budgetPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/postgres_synced_table#storage_catalog PostgresSyncedTable#storage_catalog}.
	StorageCatalog *string `field:"optional" json:"storageCatalog" yaml:"storageCatalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/postgres_synced_table#storage_schema PostgresSyncedTable#storage_schema}.
	StorageSchema *string `field:"optional" json:"storageSchema" yaml:"storageSchema"`
}

