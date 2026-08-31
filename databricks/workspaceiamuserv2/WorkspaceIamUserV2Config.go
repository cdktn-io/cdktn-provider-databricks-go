// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspaceiamuserv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspaceIamUserV2Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/workspace_iam_user_v2#account_user_status WorkspaceIamUserV2#account_user_status}.
	AccountUserStatus *string `field:"required" json:"accountUserStatus" yaml:"accountUserStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/workspace_iam_user_v2#full_name WorkspaceIamUserV2#full_name}.
	FullName *WorkspaceIamUserV2FullName `field:"required" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/workspace_iam_user_v2#username WorkspaceIamUserV2#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/workspace_iam_user_v2#external_id WorkspaceIamUserV2#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/workspace_iam_user_v2#provider_config WorkspaceIamUserV2#provider_config}.
	ProviderConfig *WorkspaceIamUserV2ProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

