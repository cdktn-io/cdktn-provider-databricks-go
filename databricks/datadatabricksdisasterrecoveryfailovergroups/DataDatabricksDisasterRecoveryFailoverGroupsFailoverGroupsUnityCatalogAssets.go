// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdisasterrecoveryfailovergroups


type DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/disaster_recovery_failover_groups#catalogs DataDatabricksDisasterRecoveryFailoverGroups#catalogs}.
	Catalogs interface{} `field:"required" json:"catalogs" yaml:"catalogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/disaster_recovery_failover_groups#data_replication_workspace_set DataDatabricksDisasterRecoveryFailoverGroups#data_replication_workspace_set}.
	DataReplicationWorkspaceSet *string `field:"required" json:"dataReplicationWorkspaceSet" yaml:"dataReplicationWorkspaceSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/disaster_recovery_failover_groups#location_mappings DataDatabricksDisasterRecoveryFailoverGroups#location_mappings}.
	LocationMappings interface{} `field:"optional" json:"locationMappings" yaml:"locationMappings"`
}

