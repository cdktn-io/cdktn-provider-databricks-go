// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresendpoint


type PostgresEndpointSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/postgres_endpoint#endpoint_type PostgresEndpoint#endpoint_type}.
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/postgres_endpoint#autoscaling_limit_max_cu PostgresEndpoint#autoscaling_limit_max_cu}.
	AutoscalingLimitMaxCu *float64 `field:"optional" json:"autoscalingLimitMaxCu" yaml:"autoscalingLimitMaxCu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/postgres_endpoint#autoscaling_limit_min_cu PostgresEndpoint#autoscaling_limit_min_cu}.
	AutoscalingLimitMinCu *float64 `field:"optional" json:"autoscalingLimitMinCu" yaml:"autoscalingLimitMinCu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/postgres_endpoint#disabled PostgresEndpoint#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/postgres_endpoint#group PostgresEndpoint#group}.
	Group *PostgresEndpointSpecGroup `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/postgres_endpoint#no_suspension PostgresEndpoint#no_suspension}.
	NoSuspension interface{} `field:"optional" json:"noSuspension" yaml:"noSuspension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/postgres_endpoint#settings PostgresEndpoint#settings}.
	Settings *PostgresEndpointSpecSettings `field:"optional" json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/postgres_endpoint#suspend_timeout_duration PostgresEndpoint#suspend_timeout_duration}.
	SuspendTimeoutDuration *string `field:"optional" json:"suspendTimeoutDuration" yaml:"suspendTimeoutDuration"`
}

