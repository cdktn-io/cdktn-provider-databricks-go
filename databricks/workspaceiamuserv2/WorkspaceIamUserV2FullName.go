// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspaceiamuserv2


type WorkspaceIamUserV2FullName struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/workspace_iam_user_v2#family_name WorkspaceIamUserV2#family_name}.
	FamilyName *string `field:"optional" json:"familyName" yaml:"familyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/workspace_iam_user_v2#given_name WorkspaceIamUserV2#given_name}.
	GivenName *string `field:"optional" json:"givenName" yaml:"givenName"`
}

