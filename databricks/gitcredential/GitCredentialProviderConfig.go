// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gitcredential


type GitCredentialProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/git_credential#workspace_id GitCredential#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

