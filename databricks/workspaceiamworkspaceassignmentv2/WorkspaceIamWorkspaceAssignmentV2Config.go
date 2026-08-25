// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspaceiamworkspaceassignmentv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspaceIamWorkspaceAssignmentV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/workspace_iam_workspace_assignment_v2#principal_id WorkspaceIamWorkspaceAssignmentV2#principal_id}.
	PrincipalId *float64 `field:"required" json:"principalId" yaml:"principalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/workspace_iam_workspace_assignment_v2#entitlements WorkspaceIamWorkspaceAssignmentV2#entitlements}.
	Entitlements *[]*string `field:"optional" json:"entitlements" yaml:"entitlements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/workspace_iam_workspace_assignment_v2#provider_config WorkspaceIamWorkspaceAssignmentV2#provider_config}.
	ProviderConfig *WorkspaceIamWorkspaceAssignmentV2ProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/workspace_iam_workspace_assignment_v2#workspace_id WorkspaceIamWorkspaceAssignmentV2#workspace_id}.
	WorkspaceId *float64 `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

