// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints


type DataDatabricksServingEndpointsEndpointsTelemetryConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/serving_endpoints#enabled_telemetry_features DataDatabricksServingEndpoints#enabled_telemetry_features}.
	EnabledTelemetryFeatures *[]*string `field:"optional" json:"enabledTelemetryFeatures" yaml:"enabledTelemetryFeatures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/serving_endpoints#inference_table_config DataDatabricksServingEndpoints#inference_table_config}.
	InferenceTableConfig interface{} `field:"optional" json:"inferenceTableConfig" yaml:"inferenceTableConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/serving_endpoints#table_names DataDatabricksServingEndpoints#table_names}.
	TableNames interface{} `field:"optional" json:"tableNames" yaml:"tableNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/serving_endpoints#telemetry_profile_id DataDatabricksServingEndpoints#telemetry_profile_id}.
	TelemetryProfileId *string `field:"optional" json:"telemetryProfileId" yaml:"telemetryProfileId"`
}

