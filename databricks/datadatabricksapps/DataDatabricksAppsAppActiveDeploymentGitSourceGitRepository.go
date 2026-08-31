// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapps


type DataDatabricksAppsAppActiveDeploymentGitSourceGitRepository struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/apps#provider DataDatabricksApps#provider}.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/apps#url DataDatabricksApps#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/apps#auto_deploy DataDatabricksApps#auto_deploy}.
	AutoDeploy interface{} `field:"optional" json:"autoDeploy" yaml:"autoDeploy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/apps#caller_credential_id DataDatabricksApps#caller_credential_id}.
	CallerCredentialId *float64 `field:"optional" json:"callerCredentialId" yaml:"callerCredentialId"`
}

