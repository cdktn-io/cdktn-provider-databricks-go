// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineTestingSourcesGlob struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/pipeline#include Pipeline#include}.
	Include *string `field:"optional" json:"include" yaml:"include"`
}

