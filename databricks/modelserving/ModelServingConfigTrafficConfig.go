// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingConfigTrafficConfig struct {
	// routes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/model_serving#routes ModelServing#routes}
	Routes interface{} `field:"optional" json:"routes" yaml:"routes"`
}

