// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package genieagent


type GenieAgentDefinitionInstructions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#example_question_sqls GenieAgent#example_question_sqls}.
	ExampleQuestionSqls interface{} `field:"optional" json:"exampleQuestionSqls" yaml:"exampleQuestionSqls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#join_specs GenieAgent#join_specs}.
	JoinSpecs interface{} `field:"optional" json:"joinSpecs" yaml:"joinSpecs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#sql_functions GenieAgent#sql_functions}.
	SqlFunctions interface{} `field:"optional" json:"sqlFunctions" yaml:"sqlFunctions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#sql_snippets GenieAgent#sql_snippets}.
	SqlSnippets *GenieAgentDefinitionInstructionsSqlSnippets `field:"optional" json:"sqlSnippets" yaml:"sqlSnippets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#text_instructions GenieAgent#text_instructions}.
	TextInstructions interface{} `field:"optional" json:"textInstructions" yaml:"textInstructions"`
}

