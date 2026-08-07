// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package disasterrecoveryfailovergroup


type DisasterRecoveryFailoverGroupWorkspaceSets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/disaster_recovery_failover_group#name DisasterRecoveryFailoverGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/disaster_recovery_failover_group#workspace_ids DisasterRecoveryFailoverGroup#workspace_ids}.
	WorkspaceIds *[]*string `field:"required" json:"workspaceIds" yaml:"workspaceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/disaster_recovery_failover_group#replicate_workspace_assets DisasterRecoveryFailoverGroup#replicate_workspace_assets}.
	ReplicateWorkspaceAssets interface{} `field:"optional" json:"replicateWorkspaceAssets" yaml:"replicateWorkspaceAssets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/disaster_recovery_failover_group#stable_url_names DisasterRecoveryFailoverGroup#stable_url_names}.
	StableUrlNames *[]*string `field:"optional" json:"stableUrlNames" yaml:"stableUrlNames"`
}

