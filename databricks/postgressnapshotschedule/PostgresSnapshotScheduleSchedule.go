// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgressnapshotschedule


type PostgresSnapshotScheduleSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.133.0/docs/resources/postgres_snapshot_schedule#retention PostgresSnapshotSchedule#retention}.
	Retention *string `field:"required" json:"retention" yaml:"retention"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.133.0/docs/resources/postgres_snapshot_schedule#daily_schedule PostgresSnapshotSchedule#daily_schedule}.
	DailySchedule *PostgresSnapshotScheduleScheduleDailySchedule `field:"optional" json:"dailySchedule" yaml:"dailySchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.133.0/docs/resources/postgres_snapshot_schedule#monthly_schedule PostgresSnapshotSchedule#monthly_schedule}.
	MonthlySchedule *PostgresSnapshotScheduleScheduleMonthlySchedule `field:"optional" json:"monthlySchedule" yaml:"monthlySchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.133.0/docs/resources/postgres_snapshot_schedule#weekly_schedule PostgresSnapshotSchedule#weekly_schedule}.
	WeeklySchedule *PostgresSnapshotScheduleScheduleWeeklySchedule `field:"optional" json:"weeklySchedule" yaml:"weeklySchedule"`
}

