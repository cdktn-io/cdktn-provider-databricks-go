// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountiamuserv2


type AccountIamUserV2FullName struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/account_iam_user_v2#family_name AccountIamUserV2#family_name}.
	FamilyName *string `field:"optional" json:"familyName" yaml:"familyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/account_iam_user_v2#given_name AccountIamUserV2#given_name}.
	GivenName *string `field:"optional" json:"givenName" yaml:"givenName"`
}

