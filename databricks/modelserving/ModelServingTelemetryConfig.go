// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingTelemetryConfig struct {
	// inference_table_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/model_serving#inference_table_config ModelServing#inference_table_config}
	InferenceTableConfig *ModelServingTelemetryConfigInferenceTableConfig `field:"optional" json:"inferenceTableConfig" yaml:"inferenceTableConfig"`
}

