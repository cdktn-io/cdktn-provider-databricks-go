// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package defaultnamespacesetting


type DefaultNamespaceSettingProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/default_namespace_setting#workspace_id DefaultNamespaceSetting#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

