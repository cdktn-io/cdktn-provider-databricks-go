// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspaceentitytagassignment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspaceEntityTagAssignmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/workspace_entity_tag_assignment#entity_id WorkspaceEntityTagAssignment#entity_id}.
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/workspace_entity_tag_assignment#entity_type WorkspaceEntityTagAssignment#entity_type}.
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/workspace_entity_tag_assignment#tag_key WorkspaceEntityTagAssignment#tag_key}.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/workspace_entity_tag_assignment#provider_config WorkspaceEntityTagAssignment#provider_config}.
	ProviderConfig *WorkspaceEntityTagAssignmentProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/workspace_entity_tag_assignment#tag_value WorkspaceEntityTagAssignment#tag_value}.
	TagValue *string `field:"optional" json:"tagValue" yaml:"tagValue"`
}

