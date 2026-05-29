// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package onlinetable


type OnlineTableProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/online_table#workspace_id OnlineTable#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

