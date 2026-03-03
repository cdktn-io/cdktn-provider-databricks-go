// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksentitytagassignments


type DataDatabricksEntityTagAssignmentsTagAssignments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/entity_tag_assignments#entity_name DataDatabricksEntityTagAssignments#entity_name}.
	EntityName *string `field:"required" json:"entityName" yaml:"entityName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/entity_tag_assignments#entity_type DataDatabricksEntityTagAssignments#entity_type}.
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/entity_tag_assignments#tag_key DataDatabricksEntityTagAssignments#tag_key}.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/entity_tag_assignments#provider_config DataDatabricksEntityTagAssignments#provider_config}.
	ProviderConfig *DataDatabricksEntityTagAssignmentsTagAssignmentsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

