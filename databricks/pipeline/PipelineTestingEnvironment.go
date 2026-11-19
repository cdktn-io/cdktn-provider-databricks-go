// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineTestingEnvironment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/pipeline#dependencies Pipeline#dependencies}.
	Dependencies *[]*string `field:"optional" json:"dependencies" yaml:"dependencies"`
}

