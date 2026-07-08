// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingTelemetryConfigInferenceTableConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/model_serving#name ModelServing#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/model_serving#sampling_fraction ModelServing#sampling_fraction}.
	SamplingFraction *float64 `field:"optional" json:"samplingFraction" yaml:"samplingFraction"`
}

