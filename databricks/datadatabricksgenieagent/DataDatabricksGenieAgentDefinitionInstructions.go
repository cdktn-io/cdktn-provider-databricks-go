// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksgenieagent


type DataDatabricksGenieAgentDefinitionInstructions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/genie_agent#example_question_sqls DataDatabricksGenieAgent#example_question_sqls}.
	ExampleQuestionSqls interface{} `field:"optional" json:"exampleQuestionSqls" yaml:"exampleQuestionSqls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/genie_agent#join_specs DataDatabricksGenieAgent#join_specs}.
	JoinSpecs interface{} `field:"optional" json:"joinSpecs" yaml:"joinSpecs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/genie_agent#sql_functions DataDatabricksGenieAgent#sql_functions}.
	SqlFunctions interface{} `field:"optional" json:"sqlFunctions" yaml:"sqlFunctions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/genie_agent#sql_snippets DataDatabricksGenieAgent#sql_snippets}.
	SqlSnippets *DataDatabricksGenieAgentDefinitionInstructionsSqlSnippets `field:"optional" json:"sqlSnippets" yaml:"sqlSnippets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/genie_agent#text_instructions DataDatabricksGenieAgent#text_instructions}.
	TextInstructions interface{} `field:"optional" json:"textInstructions" yaml:"textInstructions"`
}

