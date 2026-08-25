// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountiamworkspaceassignmentsv2


type DataDatabricksAccountIamWorkspaceAssignmentsV2WorkspaceAssignments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/account_iam_workspace_assignments_v2#principal_id DataDatabricksAccountIamWorkspaceAssignmentsV2#principal_id}.
	PrincipalId *float64 `field:"required" json:"principalId" yaml:"principalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/account_iam_workspace_assignments_v2#workspace_id DataDatabricksAccountIamWorkspaceAssignmentsV2#workspace_id}.
	WorkspaceId *float64 `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

