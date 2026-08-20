// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresproject


type PostgresProjectInitialEndpointSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/postgres_project#autoscaling_limit_max_cu PostgresProject#autoscaling_limit_max_cu}.
	AutoscalingLimitMaxCu *float64 `field:"optional" json:"autoscalingLimitMaxCu" yaml:"autoscalingLimitMaxCu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/postgres_project#autoscaling_limit_min_cu PostgresProject#autoscaling_limit_min_cu}.
	AutoscalingLimitMinCu *float64 `field:"optional" json:"autoscalingLimitMinCu" yaml:"autoscalingLimitMinCu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/postgres_project#group PostgresProject#group}.
	Group *PostgresProjectInitialEndpointSpecGroup `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/postgres_project#no_suspension PostgresProject#no_suspension}.
	NoSuspension interface{} `field:"optional" json:"noSuspension" yaml:"noSuspension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/postgres_project#suspend_timeout_duration PostgresProject#suspend_timeout_duration}.
	SuspendTimeoutDuration *string `field:"optional" json:"suspendTimeoutDuration" yaml:"suspendTimeoutDuration"`
}

