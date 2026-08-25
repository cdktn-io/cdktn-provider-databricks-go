// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package environmentsworkspacebaseenvironment


type EnvironmentsWorkspaceBaseEnvironmentSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/environments_workspace_base_environment#dependencies EnvironmentsWorkspaceBaseEnvironment#dependencies}.
	Dependencies *[]*string `field:"optional" json:"dependencies" yaml:"dependencies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/environments_workspace_base_environment#environment_version EnvironmentsWorkspaceBaseEnvironment#environment_version}.
	EnvironmentVersion *string `field:"optional" json:"environmentVersion" yaml:"environmentVersion"`
}

