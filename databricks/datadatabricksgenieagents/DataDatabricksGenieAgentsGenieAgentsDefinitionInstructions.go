// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksgenieagents


type DataDatabricksGenieAgentsGenieAgentsDefinitionInstructions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/genie_agents#example_question_sqls DataDatabricksGenieAgents#example_question_sqls}.
	ExampleQuestionSqls interface{} `field:"optional" json:"exampleQuestionSqls" yaml:"exampleQuestionSqls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/genie_agents#join_specs DataDatabricksGenieAgents#join_specs}.
	JoinSpecs interface{} `field:"optional" json:"joinSpecs" yaml:"joinSpecs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/genie_agents#sql_functions DataDatabricksGenieAgents#sql_functions}.
	SqlFunctions interface{} `field:"optional" json:"sqlFunctions" yaml:"sqlFunctions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/genie_agents#sql_snippets DataDatabricksGenieAgents#sql_snippets}.
	SqlSnippets *DataDatabricksGenieAgentsGenieAgentsDefinitionInstructionsSqlSnippets `field:"optional" json:"sqlSnippets" yaml:"sqlSnippets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/genie_agents#text_instructions DataDatabricksGenieAgents#text_instructions}.
	TextInstructions interface{} `field:"optional" json:"textInstructions" yaml:"textInstructions"`
}

