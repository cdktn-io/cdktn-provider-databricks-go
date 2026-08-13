// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package environmentsdefaultworkspacebaseenvironment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EnvironmentsDefaultWorkspaceBaseEnvironmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/environments_default_workspace_base_environment#cpu_workspace_base_environment EnvironmentsDefaultWorkspaceBaseEnvironment#cpu_workspace_base_environment}.
	CpuWorkspaceBaseEnvironment *string `field:"optional" json:"cpuWorkspaceBaseEnvironment" yaml:"cpuWorkspaceBaseEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/environments_default_workspace_base_environment#gpu_workspace_base_environment EnvironmentsDefaultWorkspaceBaseEnvironment#gpu_workspace_base_environment}.
	GpuWorkspaceBaseEnvironment *string `field:"optional" json:"gpuWorkspaceBaseEnvironment" yaml:"gpuWorkspaceBaseEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/environments_default_workspace_base_environment#provider_config EnvironmentsDefaultWorkspaceBaseEnvironment#provider_config}.
	ProviderConfig *EnvironmentsDefaultWorkspaceBaseEnvironmentProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

