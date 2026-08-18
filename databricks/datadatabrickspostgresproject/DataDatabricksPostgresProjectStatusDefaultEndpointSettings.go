// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresproject


type DataDatabricksPostgresProjectStatusDefaultEndpointSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/postgres_project#autoscaling_limit_max_cu DataDatabricksPostgresProject#autoscaling_limit_max_cu}.
	AutoscalingLimitMaxCu *float64 `field:"optional" json:"autoscalingLimitMaxCu" yaml:"autoscalingLimitMaxCu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/postgres_project#autoscaling_limit_min_cu DataDatabricksPostgresProject#autoscaling_limit_min_cu}.
	AutoscalingLimitMinCu *float64 `field:"optional" json:"autoscalingLimitMinCu" yaml:"autoscalingLimitMinCu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/postgres_project#no_suspension DataDatabricksPostgresProject#no_suspension}.
	NoSuspension interface{} `field:"optional" json:"noSuspension" yaml:"noSuspension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/postgres_project#pg_settings DataDatabricksPostgresProject#pg_settings}.
	PgSettings *map[string]*string `field:"optional" json:"pgSettings" yaml:"pgSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/postgres_project#suspend_timeout_duration DataDatabricksPostgresProject#suspend_timeout_duration}.
	SuspendTimeoutDuration *string `field:"optional" json:"suspendTimeoutDuration" yaml:"suspendTimeoutDuration"`
}

