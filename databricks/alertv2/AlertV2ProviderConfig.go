// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package alertv2


type AlertV2ProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/alert_v2#workspace_id AlertV2#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

