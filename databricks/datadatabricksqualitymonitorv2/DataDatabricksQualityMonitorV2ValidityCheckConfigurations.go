// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksqualitymonitorv2


type DataDatabricksQualityMonitorV2ValidityCheckConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.111.0/docs/data-sources/quality_monitor_v2#name DataDatabricksQualityMonitorV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.111.0/docs/data-sources/quality_monitor_v2#percent_null_validity_check DataDatabricksQualityMonitorV2#percent_null_validity_check}.
	PercentNullValidityCheck *DataDatabricksQualityMonitorV2ValidityCheckConfigurationsPercentNullValidityCheck `field:"optional" json:"percentNullValidityCheck" yaml:"percentNullValidityCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.111.0/docs/data-sources/quality_monitor_v2#range_validity_check DataDatabricksQualityMonitorV2#range_validity_check}.
	RangeValidityCheck *DataDatabricksQualityMonitorV2ValidityCheckConfigurationsRangeValidityCheck `field:"optional" json:"rangeValidityCheck" yaml:"rangeValidityCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.111.0/docs/data-sources/quality_monitor_v2#uniqueness_validity_check DataDatabricksQualityMonitorV2#uniqueness_validity_check}.
	UniquenessValidityCheck *DataDatabricksQualityMonitorV2ValidityCheckConfigurationsUniquenessValidityCheck `field:"optional" json:"uniquenessValidityCheck" yaml:"uniquenessValidityCheck"`
}

