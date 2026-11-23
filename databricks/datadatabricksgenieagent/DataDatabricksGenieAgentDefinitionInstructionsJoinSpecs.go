// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksgenieagent


type DataDatabricksGenieAgentDefinitionInstructionsJoinSpecs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/genie_agent#id DataDatabricksGenieAgent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/genie_agent#left DataDatabricksGenieAgent#left}.
	Left *DataDatabricksGenieAgentDefinitionInstructionsJoinSpecsLeft `field:"required" json:"left" yaml:"left"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/genie_agent#right DataDatabricksGenieAgent#right}.
	Right *DataDatabricksGenieAgentDefinitionInstructionsJoinSpecsRight `field:"required" json:"right" yaml:"right"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/genie_agent#sql DataDatabricksGenieAgent#sql}.
	Sql *string `field:"required" json:"sql" yaml:"sql"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/genie_agent#comment DataDatabricksGenieAgent#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/genie_agent#instruction DataDatabricksGenieAgent#instruction}.
	Instruction *string `field:"optional" json:"instruction" yaml:"instruction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/genie_agent#relationship_type DataDatabricksGenieAgent#relationship_type}.
	RelationshipType *string `field:"optional" json:"relationshipType" yaml:"relationshipType"`
}

