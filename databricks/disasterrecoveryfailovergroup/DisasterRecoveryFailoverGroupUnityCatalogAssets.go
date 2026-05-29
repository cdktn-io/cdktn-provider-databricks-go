// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package disasterrecoveryfailovergroup


type DisasterRecoveryFailoverGroupUnityCatalogAssets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/disaster_recovery_failover_group#catalogs DisasterRecoveryFailoverGroup#catalogs}.
	Catalogs interface{} `field:"required" json:"catalogs" yaml:"catalogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/disaster_recovery_failover_group#data_replication_workspace_set DisasterRecoveryFailoverGroup#data_replication_workspace_set}.
	DataReplicationWorkspaceSet *string `field:"required" json:"dataReplicationWorkspaceSet" yaml:"dataReplicationWorkspaceSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/disaster_recovery_failover_group#location_mappings DisasterRecoveryFailoverGroup#location_mappings}.
	LocationMappings interface{} `field:"optional" json:"locationMappings" yaml:"locationMappings"`
}

