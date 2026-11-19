// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package genieagent


type GenieAgentDefinitionBenchmarksQuestionsAnswer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/genie_agent#content GenieAgent#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/genie_agent#format GenieAgent#format}.
	Format *string `field:"required" json:"format" yaml:"format"`
}

