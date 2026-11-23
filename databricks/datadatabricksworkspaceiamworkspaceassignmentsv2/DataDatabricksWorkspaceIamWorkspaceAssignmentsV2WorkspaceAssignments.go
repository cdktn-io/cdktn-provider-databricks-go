// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspaceiamworkspaceassignmentsv2


type DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/workspace_iam_workspace_assignments_v2#name DataDatabricksWorkspaceIamWorkspaceAssignmentsV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/workspace_iam_workspace_assignments_v2#provider_config DataDatabricksWorkspaceIamWorkspaceAssignmentsV2#provider_config}.
	ProviderConfig *DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

