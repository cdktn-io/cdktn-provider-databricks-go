// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qualitymonitorv2


type QualityMonitorV2AnomalyDetectionConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/quality_monitor_v2#excluded_table_full_names QualityMonitorV2#excluded_table_full_names}.
	ExcludedTableFullNames *[]*string `field:"optional" json:"excludedTableFullNames" yaml:"excludedTableFullNames"`
}

