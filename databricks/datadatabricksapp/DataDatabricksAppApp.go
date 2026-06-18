// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapp


type DataDatabricksAppApp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/app#name DataDatabricksApp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/app#budget_policy_id DataDatabricksApp#budget_policy_id}.
	BudgetPolicyId *string `field:"optional" json:"budgetPolicyId" yaml:"budgetPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/app#compute_max_instances DataDatabricksApp#compute_max_instances}.
	ComputeMaxInstances *float64 `field:"optional" json:"computeMaxInstances" yaml:"computeMaxInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/app#compute_min_instances DataDatabricksApp#compute_min_instances}.
	ComputeMinInstances *float64 `field:"optional" json:"computeMinInstances" yaml:"computeMinInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/app#compute_size DataDatabricksApp#compute_size}.
	ComputeSize *string `field:"optional" json:"computeSize" yaml:"computeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/app#description DataDatabricksApp#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/app#git_repository DataDatabricksApp#git_repository}.
	GitRepository *DataDatabricksAppAppGitRepository `field:"optional" json:"gitRepository" yaml:"gitRepository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/app#resources DataDatabricksApp#resources}.
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/app#space DataDatabricksApp#space}.
	Space *string `field:"optional" json:"space" yaml:"space"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/app#telemetry_export_destinations DataDatabricksApp#telemetry_export_destinations}.
	TelemetryExportDestinations interface{} `field:"optional" json:"telemetryExportDestinations" yaml:"telemetryExportDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/app#usage_policy_id DataDatabricksApp#usage_policy_id}.
	UsagePolicyId *string `field:"optional" json:"usagePolicyId" yaml:"usagePolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/app#user_api_scopes DataDatabricksApp#user_api_scopes}.
	UserApiScopes *[]*string `field:"optional" json:"userApiScopes" yaml:"userApiScopes"`
}

