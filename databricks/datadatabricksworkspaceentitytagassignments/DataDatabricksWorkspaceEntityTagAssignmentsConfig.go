// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspaceentitytagassignments

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksWorkspaceEntityTagAssignmentsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/workspace_entity_tag_assignments#entity_id DataDatabricksWorkspaceEntityTagAssignments#entity_id}.
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/workspace_entity_tag_assignments#entity_type DataDatabricksWorkspaceEntityTagAssignments#entity_type}.
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/workspace_entity_tag_assignments#page_size DataDatabricksWorkspaceEntityTagAssignments#page_size}.
	PageSize *float64 `field:"optional" json:"pageSize" yaml:"pageSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/workspace_entity_tag_assignments#provider_config DataDatabricksWorkspaceEntityTagAssignments#provider_config}.
	ProviderConfig *DataDatabricksWorkspaceEntityTagAssignmentsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

