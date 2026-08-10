// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksenvironmentsworkspacebaseenvironment


type DataDatabricksEnvironmentsWorkspaceBaseEnvironmentSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/environments_workspace_base_environment#dependencies DataDatabricksEnvironmentsWorkspaceBaseEnvironment#dependencies}.
	Dependencies *[]*string `field:"optional" json:"dependencies" yaml:"dependencies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/environments_workspace_base_environment#environment_version DataDatabricksEnvironmentsWorkspaceBaseEnvironment#environment_version}.
	EnvironmentVersion *string `field:"optional" json:"environmentVersion" yaml:"environmentVersion"`
}

