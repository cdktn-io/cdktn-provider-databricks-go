// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingAiGatewayFallbackConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/model_serving#enabled ModelServing#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

