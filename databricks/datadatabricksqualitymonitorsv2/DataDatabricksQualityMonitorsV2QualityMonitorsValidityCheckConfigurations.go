// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksqualitymonitorsv2


type DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/quality_monitors_v2#name DataDatabricksQualityMonitorsV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/quality_monitors_v2#percent_null_validity_check DataDatabricksQualityMonitorsV2#percent_null_validity_check}.
	PercentNullValidityCheck *DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsPercentNullValidityCheck `field:"optional" json:"percentNullValidityCheck" yaml:"percentNullValidityCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/quality_monitors_v2#range_validity_check DataDatabricksQualityMonitorsV2#range_validity_check}.
	RangeValidityCheck *DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheck `field:"optional" json:"rangeValidityCheck" yaml:"rangeValidityCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/quality_monitors_v2#uniqueness_validity_check DataDatabricksQualityMonitorsV2#uniqueness_validity_check}.
	UniquenessValidityCheck *DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsUniquenessValidityCheck `field:"optional" json:"uniquenessValidityCheck" yaml:"uniquenessValidityCheck"`
}

