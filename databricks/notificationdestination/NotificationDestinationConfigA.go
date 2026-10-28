// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package notificationdestination


type NotificationDestinationConfigA struct {
	// databricks_app_slack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/notification_destination#databricks_app_slack NotificationDestination#databricks_app_slack}
	DatabricksAppSlack *NotificationDestinationConfigDatabricksAppSlack `field:"optional" json:"databricksAppSlack" yaml:"databricksAppSlack"`
	// databricks_app_teams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/notification_destination#databricks_app_teams NotificationDestination#databricks_app_teams}
	DatabricksAppTeams *NotificationDestinationConfigDatabricksAppTeams `field:"optional" json:"databricksAppTeams" yaml:"databricksAppTeams"`
	// email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/notification_destination#email NotificationDestination#email}
	Email *NotificationDestinationConfigEmail `field:"optional" json:"email" yaml:"email"`
	// generic_webhook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/notification_destination#generic_webhook NotificationDestination#generic_webhook}
	GenericWebhook *NotificationDestinationConfigGenericWebhook `field:"optional" json:"genericWebhook" yaml:"genericWebhook"`
	// microsoft_teams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/notification_destination#microsoft_teams NotificationDestination#microsoft_teams}
	MicrosoftTeams *NotificationDestinationConfigMicrosoftTeams `field:"optional" json:"microsoftTeams" yaml:"microsoftTeams"`
	// pagerduty block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/notification_destination#pagerduty NotificationDestination#pagerduty}
	Pagerduty *NotificationDestinationConfigPagerduty `field:"optional" json:"pagerduty" yaml:"pagerduty"`
	// slack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/notification_destination#slack NotificationDestination#slack}
	Slack *NotificationDestinationConfigSlack `field:"optional" json:"slack" yaml:"slack"`
}

