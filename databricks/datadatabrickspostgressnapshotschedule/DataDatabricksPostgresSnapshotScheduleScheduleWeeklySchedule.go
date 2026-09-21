// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgressnapshotschedule


type DataDatabricksPostgresSnapshotScheduleScheduleWeeklySchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.133.0/docs/data-sources/postgres_snapshot_schedule#day_of_week DataDatabricksPostgresSnapshotSchedule#day_of_week}.
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.133.0/docs/data-sources/postgres_snapshot_schedule#hour DataDatabricksPostgresSnapshotSchedule#hour}.
	Hour *float64 `field:"optional" json:"hour" yaml:"hour"`
}

