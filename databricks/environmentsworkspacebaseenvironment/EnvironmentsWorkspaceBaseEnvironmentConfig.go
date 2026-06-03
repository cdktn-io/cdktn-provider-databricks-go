// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package environmentsworkspacebaseenvironment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EnvironmentsWorkspaceBaseEnvironmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/environments_workspace_base_environment#display_name EnvironmentsWorkspaceBaseEnvironment#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/environments_workspace_base_environment#base_environment_type EnvironmentsWorkspaceBaseEnvironment#base_environment_type}.
	BaseEnvironmentType *string `field:"optional" json:"baseEnvironmentType" yaml:"baseEnvironmentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/environments_workspace_base_environment#filepath EnvironmentsWorkspaceBaseEnvironment#filepath}.
	Filepath *string `field:"optional" json:"filepath" yaml:"filepath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/environments_workspace_base_environment#provider_config EnvironmentsWorkspaceBaseEnvironment#provider_config}.
	ProviderConfig *EnvironmentsWorkspaceBaseEnvironmentProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/environments_workspace_base_environment#workspace_base_environment_id EnvironmentsWorkspaceBaseEnvironment#workspace_base_environment_id}.
	WorkspaceBaseEnvironmentId *string `field:"optional" json:"workspaceBaseEnvironmentId" yaml:"workspaceBaseEnvironmentId"`
}

