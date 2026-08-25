// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdisasterrecoveryfailovergroups


type DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/disaster_recovery_failover_groups#name DataDatabricksDisasterRecoveryFailoverGroups#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/disaster_recovery_failover_groups#workspace_ids DataDatabricksDisasterRecoveryFailoverGroups#workspace_ids}.
	WorkspaceIds *[]*string `field:"required" json:"workspaceIds" yaml:"workspaceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/disaster_recovery_failover_groups#replicate_workspace_assets DataDatabricksDisasterRecoveryFailoverGroups#replicate_workspace_assets}.
	ReplicateWorkspaceAssets interface{} `field:"optional" json:"replicateWorkspaceAssets" yaml:"replicateWorkspaceAssets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/disaster_recovery_failover_groups#stable_url_names DataDatabricksDisasterRecoveryFailoverGroups#stable_url_names}.
	StableUrlNames *[]*string `field:"optional" json:"stableUrlNames" yaml:"stableUrlNames"`
}

