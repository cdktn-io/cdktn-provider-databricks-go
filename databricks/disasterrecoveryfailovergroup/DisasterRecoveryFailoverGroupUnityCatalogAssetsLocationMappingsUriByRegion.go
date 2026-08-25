// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package disasterrecoveryfailovergroup


type DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsUriByRegion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/disaster_recovery_failover_group#region DisasterRecoveryFailoverGroup#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/disaster_recovery_failover_group#uri DisasterRecoveryFailoverGroup#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

