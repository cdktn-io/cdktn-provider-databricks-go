// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdisasterrecoveryfailovergroups


type DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsLocationMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/disaster_recovery_failover_groups#name DataDatabricksDisasterRecoveryFailoverGroups#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/disaster_recovery_failover_groups#uri_by_region DataDatabricksDisasterRecoveryFailoverGroups#uri_by_region}.
	UriByRegion interface{} `field:"required" json:"uriByRegion" yaml:"uriByRegion"`
}

