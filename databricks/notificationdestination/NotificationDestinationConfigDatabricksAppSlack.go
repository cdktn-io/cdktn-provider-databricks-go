// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package notificationdestination


type NotificationDestinationConfigDatabricksAppSlack struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/notification_destination#channel_id NotificationDestination#channel_id}.
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/notification_destination#channel_id_set NotificationDestination#channel_id_set}.
	ChannelIdSet interface{} `field:"optional" json:"channelIdSet" yaml:"channelIdSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/notification_destination#team_id NotificationDestination#team_id}.
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/notification_destination#team_id_set NotificationDestination#team_id_set}.
	TeamIdSet interface{} `field:"optional" json:"teamIdSet" yaml:"teamIdSet"`
}

