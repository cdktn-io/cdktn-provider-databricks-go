// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2


type WorkspaceSettingV2ProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/workspace_setting_v2#workspace_id WorkspaceSettingV2#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

