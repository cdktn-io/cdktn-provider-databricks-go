// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package genieagent


type GenieAgentDefinitionBenchmarksQuestions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#id GenieAgent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#question GenieAgent#question}.
	Question *string `field:"required" json:"question" yaml:"question"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#answer GenieAgent#answer}.
	Answer interface{} `field:"optional" json:"answer" yaml:"answer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#evaluation_note GenieAgent#evaluation_note}.
	EvaluationNote *string `field:"optional" json:"evaluationNote" yaml:"evaluationNote"`
}

