// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package entitlements


type EntitlementsProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/entitlements#workspace_id Entitlements#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

