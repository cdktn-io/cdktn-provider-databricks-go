// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresendpoints


type DataDatabricksPostgresEndpointsEndpointsSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/postgres_endpoints#endpoint_type DataDatabricksPostgresEndpoints#endpoint_type}.
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/postgres_endpoints#autoscaling_limit_max_cu DataDatabricksPostgresEndpoints#autoscaling_limit_max_cu}.
	AutoscalingLimitMaxCu *float64 `field:"optional" json:"autoscalingLimitMaxCu" yaml:"autoscalingLimitMaxCu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/postgres_endpoints#autoscaling_limit_min_cu DataDatabricksPostgresEndpoints#autoscaling_limit_min_cu}.
	AutoscalingLimitMinCu *float64 `field:"optional" json:"autoscalingLimitMinCu" yaml:"autoscalingLimitMinCu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/postgres_endpoints#disabled DataDatabricksPostgresEndpoints#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/postgres_endpoints#group DataDatabricksPostgresEndpoints#group}.
	Group *DataDatabricksPostgresEndpointsEndpointsSpecGroup `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/postgres_endpoints#no_suspension DataDatabricksPostgresEndpoints#no_suspension}.
	NoSuspension interface{} `field:"optional" json:"noSuspension" yaml:"noSuspension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/postgres_endpoints#settings DataDatabricksPostgresEndpoints#settings}.
	Settings *DataDatabricksPostgresEndpointsEndpointsSpecSettings `field:"optional" json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/postgres_endpoints#suspend_timeout_duration DataDatabricksPostgresEndpoints#suspend_timeout_duration}.
	SuspendTimeoutDuration *string `field:"optional" json:"suspendTimeoutDuration" yaml:"suspendTimeoutDuration"`
}

