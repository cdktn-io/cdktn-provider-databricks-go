// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package genieagent


type GenieAgentDefinitionInstructionsSqlSnippetsExpressions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/genie_agent#id GenieAgent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/genie_agent#sql GenieAgent#sql}.
	Sql *string `field:"required" json:"sql" yaml:"sql"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/genie_agent#alias GenieAgent#alias}.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/genie_agent#display_name GenieAgent#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/genie_agent#instruction GenieAgent#instruction}.
	Instruction *string `field:"optional" json:"instruction" yaml:"instruction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/genie_agent#synonyms GenieAgent#synonyms}.
	Synonyms *[]*string `field:"optional" json:"synonyms" yaml:"synonyms"`
}

