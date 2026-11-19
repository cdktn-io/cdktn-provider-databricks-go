// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package genieagent


type GenieAgentDefinitionInstructionsExampleQuestionSqlsParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/genie_agent#name GenieAgent#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/genie_agent#default_value GenieAgent#default_value}.
	DefaultValue *GenieAgentDefinitionInstructionsExampleQuestionSqlsParametersDefaultValue `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/genie_agent#description GenieAgent#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/genie_agent#type_hint GenieAgent#type_hint}.
	TypeHint *string `field:"optional" json:"typeHint" yaml:"typeHint"`
}

