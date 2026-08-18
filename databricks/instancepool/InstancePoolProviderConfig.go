// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package instancepool


type InstancePoolProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/instance_pool#workspace_id InstancePool#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

