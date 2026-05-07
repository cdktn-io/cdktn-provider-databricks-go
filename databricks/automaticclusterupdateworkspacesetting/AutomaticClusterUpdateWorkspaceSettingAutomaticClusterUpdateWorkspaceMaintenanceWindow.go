// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package automaticclusterupdateworkspacesetting


type AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceMaintenanceWindow struct {
	// week_day_based_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/automatic_cluster_update_workspace_setting#week_day_based_schedule AutomaticClusterUpdateWorkspaceSetting#week_day_based_schedule}
	WeekDayBasedSchedule *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceMaintenanceWindowWeekDayBasedSchedule `field:"optional" json:"weekDayBasedSchedule" yaml:"weekDayBasedSchedule"`
}

