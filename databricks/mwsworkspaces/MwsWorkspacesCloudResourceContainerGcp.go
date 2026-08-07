// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mwsworkspaces


type MwsWorkspacesCloudResourceContainerGcp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/mws_workspaces#project_id MwsWorkspaces#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

