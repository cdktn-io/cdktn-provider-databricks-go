// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package obotoken


type OboTokenProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/obo_token#workspace_id OboToken#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

