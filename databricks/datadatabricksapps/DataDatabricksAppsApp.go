// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapps


type DataDatabricksAppsApp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/apps#name DataDatabricksApps#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/apps#budget_policy_id DataDatabricksApps#budget_policy_id}.
	BudgetPolicyId *string `field:"optional" json:"budgetPolicyId" yaml:"budgetPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/apps#compute_max_instances DataDatabricksApps#compute_max_instances}.
	ComputeMaxInstances *float64 `field:"optional" json:"computeMaxInstances" yaml:"computeMaxInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/apps#compute_min_instances DataDatabricksApps#compute_min_instances}.
	ComputeMinInstances *float64 `field:"optional" json:"computeMinInstances" yaml:"computeMinInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/apps#compute_size DataDatabricksApps#compute_size}.
	ComputeSize *string `field:"optional" json:"computeSize" yaml:"computeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/apps#description DataDatabricksApps#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/apps#git_repository DataDatabricksApps#git_repository}.
	GitRepository *DataDatabricksAppsAppGitRepository `field:"optional" json:"gitRepository" yaml:"gitRepository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/apps#resources DataDatabricksApps#resources}.
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/apps#space DataDatabricksApps#space}.
	Space *string `field:"optional" json:"space" yaml:"space"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/apps#telemetry_export_destinations DataDatabricksApps#telemetry_export_destinations}.
	TelemetryExportDestinations interface{} `field:"optional" json:"telemetryExportDestinations" yaml:"telemetryExportDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/apps#usage_policy_id DataDatabricksApps#usage_policy_id}.
	UsagePolicyId *string `field:"optional" json:"usagePolicyId" yaml:"usagePolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/apps#user_api_scopes DataDatabricksApps#user_api_scopes}.
	UserApiScopes *[]*string `field:"optional" json:"userApiScopes" yaml:"userApiScopes"`
}

