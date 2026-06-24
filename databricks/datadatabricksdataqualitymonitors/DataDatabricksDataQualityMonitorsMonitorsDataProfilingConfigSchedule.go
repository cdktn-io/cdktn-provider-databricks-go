// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualitymonitors


type DataDatabricksDataQualityMonitorsMonitorsDataProfilingConfigSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/data_quality_monitors#quartz_cron_expression DataDatabricksDataQualityMonitors#quartz_cron_expression}.
	QuartzCronExpression *string `field:"required" json:"quartzCronExpression" yaml:"quartzCronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/data_quality_monitors#timezone_id DataDatabricksDataQualityMonitors#timezone_id}.
	TimezoneId *string `field:"required" json:"timezoneId" yaml:"timezoneId"`
}

