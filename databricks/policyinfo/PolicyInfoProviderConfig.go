// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package policyinfo


type PolicyInfoProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/policy_info#workspace_id PolicyInfo#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

