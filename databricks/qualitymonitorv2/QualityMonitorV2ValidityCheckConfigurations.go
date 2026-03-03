// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qualitymonitorv2


type QualityMonitorV2ValidityCheckConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/resources/quality_monitor_v2#name QualityMonitorV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/resources/quality_monitor_v2#percent_null_validity_check QualityMonitorV2#percent_null_validity_check}.
	PercentNullValidityCheck *QualityMonitorV2ValidityCheckConfigurationsPercentNullValidityCheck `field:"optional" json:"percentNullValidityCheck" yaml:"percentNullValidityCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/resources/quality_monitor_v2#range_validity_check QualityMonitorV2#range_validity_check}.
	RangeValidityCheck *QualityMonitorV2ValidityCheckConfigurationsRangeValidityCheck `field:"optional" json:"rangeValidityCheck" yaml:"rangeValidityCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/resources/quality_monitor_v2#uniqueness_validity_check QualityMonitorV2#uniqueness_validity_check}.
	UniquenessValidityCheck *QualityMonitorV2ValidityCheckConfigurationsUniquenessValidityCheck `field:"optional" json:"uniquenessValidityCheck" yaml:"uniquenessValidityCheck"`
}

