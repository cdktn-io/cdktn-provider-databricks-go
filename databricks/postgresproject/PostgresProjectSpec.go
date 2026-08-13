// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresproject


type PostgresProjectSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/postgres_project#budget_policy_id PostgresProject#budget_policy_id}.
	BudgetPolicyId *string `field:"optional" json:"budgetPolicyId" yaml:"budgetPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/postgres_project#custom_tags PostgresProject#custom_tags}.
	CustomTags interface{} `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/postgres_project#default_branch PostgresProject#default_branch}.
	DefaultBranch *string `field:"optional" json:"defaultBranch" yaml:"defaultBranch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/postgres_project#default_endpoint_settings PostgresProject#default_endpoint_settings}.
	DefaultEndpointSettings *PostgresProjectSpecDefaultEndpointSettings `field:"optional" json:"defaultEndpointSettings" yaml:"defaultEndpointSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/postgres_project#display_name PostgresProject#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/postgres_project#enable_pg_native_login PostgresProject#enable_pg_native_login}.
	EnablePgNativeLogin interface{} `field:"optional" json:"enablePgNativeLogin" yaml:"enablePgNativeLogin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/postgres_project#history_retention_duration PostgresProject#history_retention_duration}.
	HistoryRetentionDuration *string `field:"optional" json:"historyRetentionDuration" yaml:"historyRetentionDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/postgres_project#pg_version PostgresProject#pg_version}.
	PgVersion *float64 `field:"optional" json:"pgVersion" yaml:"pgVersion"`
}

