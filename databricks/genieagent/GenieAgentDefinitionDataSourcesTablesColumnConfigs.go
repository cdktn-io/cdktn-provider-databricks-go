// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package genieagent


type GenieAgentDefinitionDataSourcesTablesColumnConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#column GenieAgent#column}.
	Column *string `field:"required" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#description GenieAgent#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#display_name GenieAgent#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#enable_entity_matching GenieAgent#enable_entity_matching}.
	EnableEntityMatching interface{} `field:"optional" json:"enableEntityMatching" yaml:"enableEntityMatching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#enable_format_assistance GenieAgent#enable_format_assistance}.
	EnableFormatAssistance interface{} `field:"optional" json:"enableFormatAssistance" yaml:"enableFormatAssistance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#exclude GenieAgent#exclude}.
	Exclude interface{} `field:"optional" json:"exclude" yaml:"exclude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#synonyms GenieAgent#synonyms}.
	Synonyms *[]*string `field:"optional" json:"synonyms" yaml:"synonyms"`
}

