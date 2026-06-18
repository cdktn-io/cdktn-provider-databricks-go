// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresprojects


type DataDatabricksPostgresProjectsProjectsStatusDefaultEndpointSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/postgres_projects#autoscaling_limit_max_cu DataDatabricksPostgresProjects#autoscaling_limit_max_cu}.
	AutoscalingLimitMaxCu *float64 `field:"optional" json:"autoscalingLimitMaxCu" yaml:"autoscalingLimitMaxCu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/postgres_projects#autoscaling_limit_min_cu DataDatabricksPostgresProjects#autoscaling_limit_min_cu}.
	AutoscalingLimitMinCu *float64 `field:"optional" json:"autoscalingLimitMinCu" yaml:"autoscalingLimitMinCu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/postgres_projects#no_suspension DataDatabricksPostgresProjects#no_suspension}.
	NoSuspension interface{} `field:"optional" json:"noSuspension" yaml:"noSuspension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/postgres_projects#pg_settings DataDatabricksPostgresProjects#pg_settings}.
	PgSettings *map[string]*string `field:"optional" json:"pgSettings" yaml:"pgSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/postgres_projects#suspend_timeout_duration DataDatabricksPostgresProjects#suspend_timeout_duration}.
	SuspendTimeoutDuration *string `field:"optional" json:"suspendTimeoutDuration" yaml:"suspendTimeoutDuration"`
}

