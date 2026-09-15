// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgressnapshotschedule


type DataDatabricksPostgresSnapshotScheduleSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.132.0/docs/data-sources/postgres_snapshot_schedule#retention DataDatabricksPostgresSnapshotSchedule#retention}.
	Retention *string `field:"required" json:"retention" yaml:"retention"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.132.0/docs/data-sources/postgres_snapshot_schedule#daily_schedule DataDatabricksPostgresSnapshotSchedule#daily_schedule}.
	DailySchedule *DataDatabricksPostgresSnapshotScheduleScheduleDailySchedule `field:"optional" json:"dailySchedule" yaml:"dailySchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.132.0/docs/data-sources/postgres_snapshot_schedule#monthly_schedule DataDatabricksPostgresSnapshotSchedule#monthly_schedule}.
	MonthlySchedule *DataDatabricksPostgresSnapshotScheduleScheduleMonthlySchedule `field:"optional" json:"monthlySchedule" yaml:"monthlySchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.132.0/docs/data-sources/postgres_snapshot_schedule#weekly_schedule DataDatabricksPostgresSnapshotSchedule#weekly_schedule}.
	WeeklySchedule *DataDatabricksPostgresSnapshotScheduleScheduleWeeklySchedule `field:"optional" json:"weeklySchedule" yaml:"weeklySchedule"`
}

