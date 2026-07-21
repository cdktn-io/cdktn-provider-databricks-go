// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdisasterrecoveryfailovergroup


type DataDatabricksDisasterRecoveryFailoverGroupWorkspaceSets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/disaster_recovery_failover_group#name DataDatabricksDisasterRecoveryFailoverGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/disaster_recovery_failover_group#workspace_ids DataDatabricksDisasterRecoveryFailoverGroup#workspace_ids}.
	WorkspaceIds *[]*string `field:"required" json:"workspaceIds" yaml:"workspaceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/disaster_recovery_failover_group#replicate_workspace_assets DataDatabricksDisasterRecoveryFailoverGroup#replicate_workspace_assets}.
	ReplicateWorkspaceAssets interface{} `field:"optional" json:"replicateWorkspaceAssets" yaml:"replicateWorkspaceAssets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/disaster_recovery_failover_group#stable_url_names DataDatabricksDisasterRecoveryFailoverGroup#stable_url_names}.
	StableUrlNames *[]*string `field:"optional" json:"stableUrlNames" yaml:"stableUrlNames"`
}

