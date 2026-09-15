// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgressnapshotschedule


type PostgresSnapshotScheduleScheduleWeeklySchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.132.0/docs/resources/postgres_snapshot_schedule#day_of_week PostgresSnapshotSchedule#day_of_week}.
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.132.0/docs/resources/postgres_snapshot_schedule#hour PostgresSnapshotSchedule#hour}.
	Hour *float64 `field:"optional" json:"hour" yaml:"hour"`
}

