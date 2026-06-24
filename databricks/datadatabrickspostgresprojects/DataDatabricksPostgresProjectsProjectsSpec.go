// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresprojects


type DataDatabricksPostgresProjectsProjectsSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/postgres_projects#budget_policy_id DataDatabricksPostgresProjects#budget_policy_id}.
	BudgetPolicyId *string `field:"optional" json:"budgetPolicyId" yaml:"budgetPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/postgres_projects#custom_tags DataDatabricksPostgresProjects#custom_tags}.
	CustomTags interface{} `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/postgres_projects#default_branch DataDatabricksPostgresProjects#default_branch}.
	DefaultBranch *string `field:"optional" json:"defaultBranch" yaml:"defaultBranch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/postgres_projects#default_endpoint_settings DataDatabricksPostgresProjects#default_endpoint_settings}.
	DefaultEndpointSettings *DataDatabricksPostgresProjectsProjectsSpecDefaultEndpointSettings `field:"optional" json:"defaultEndpointSettings" yaml:"defaultEndpointSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/postgres_projects#display_name DataDatabricksPostgresProjects#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/postgres_projects#enable_pg_native_login DataDatabricksPostgresProjects#enable_pg_native_login}.
	EnablePgNativeLogin interface{} `field:"optional" json:"enablePgNativeLogin" yaml:"enablePgNativeLogin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/postgres_projects#history_retention_duration DataDatabricksPostgresProjects#history_retention_duration}.
	HistoryRetentionDuration *string `field:"optional" json:"historyRetentionDuration" yaml:"historyRetentionDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/postgres_projects#pg_version DataDatabricksPostgresProjects#pg_version}.
	PgVersion *float64 `field:"optional" json:"pgVersion" yaml:"pgVersion"`
}

