// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package disasterrecoveryfailovergroup


type DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/disaster_recovery_failover_group#name DisasterRecoveryFailoverGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/disaster_recovery_failover_group#uri_by_region DisasterRecoveryFailoverGroup#uri_by_region}.
	UriByRegion interface{} `field:"required" json:"uriByRegion" yaml:"uriByRegion"`
}

