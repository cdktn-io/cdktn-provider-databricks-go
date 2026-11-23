// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package genieagent


type GenieAgentDefinitionConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#sample_questions GenieAgent#sample_questions}.
	SampleQuestions interface{} `field:"optional" json:"sampleQuestions" yaml:"sampleQuestions"`
}

