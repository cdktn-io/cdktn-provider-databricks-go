// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ipaccesslist


type IpAccessListProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/ip_access_list#workspace_id IpAccessList#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

