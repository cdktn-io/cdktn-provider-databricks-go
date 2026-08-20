// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdisasterrecoveryfailovergroup


type DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/disaster_recovery_failover_group#name DataDatabricksDisasterRecoveryFailoverGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/disaster_recovery_failover_group#uri_by_region DataDatabricksDisasterRecoveryFailoverGroup#uri_by_region}.
	UriByRegion interface{} `field:"required" json:"uriByRegion" yaml:"uriByRegion"`
}

