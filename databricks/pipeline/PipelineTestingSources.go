// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineTestingSources struct {
	// glob block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/pipeline#glob Pipeline#glob}
	Glob *PipelineTestingSourcesGlob `field:"optional" json:"glob" yaml:"glob"`
}

