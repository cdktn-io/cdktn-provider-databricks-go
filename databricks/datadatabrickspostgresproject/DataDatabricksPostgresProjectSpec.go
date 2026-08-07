// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresproject


type DataDatabricksPostgresProjectSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/postgres_project#budget_policy_id DataDatabricksPostgresProject#budget_policy_id}.
	BudgetPolicyId *string `field:"optional" json:"budgetPolicyId" yaml:"budgetPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/postgres_project#custom_tags DataDatabricksPostgresProject#custom_tags}.
	CustomTags interface{} `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/postgres_project#default_branch DataDatabricksPostgresProject#default_branch}.
	DefaultBranch *string `field:"optional" json:"defaultBranch" yaml:"defaultBranch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/postgres_project#default_endpoint_settings DataDatabricksPostgresProject#default_endpoint_settings}.
	DefaultEndpointSettings *DataDatabricksPostgresProjectSpecDefaultEndpointSettings `field:"optional" json:"defaultEndpointSettings" yaml:"defaultEndpointSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/postgres_project#display_name DataDatabricksPostgresProject#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/postgres_project#enable_pg_native_login DataDatabricksPostgresProject#enable_pg_native_login}.
	EnablePgNativeLogin interface{} `field:"optional" json:"enablePgNativeLogin" yaml:"enablePgNativeLogin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/postgres_project#history_retention_duration DataDatabricksPostgresProject#history_retention_duration}.
	HistoryRetentionDuration *string `field:"optional" json:"historyRetentionDuration" yaml:"historyRetentionDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/postgres_project#pg_version DataDatabricksPostgresProject#pg_version}.
	PgVersion *float64 `field:"optional" json:"pgVersion" yaml:"pgVersion"`
}

