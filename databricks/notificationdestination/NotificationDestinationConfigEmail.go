// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package notificationdestination


type NotificationDestinationConfigEmail struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/resources/notification_destination#addresses NotificationDestination#addresses}.
	Addresses *[]*string `field:"optional" json:"addresses" yaml:"addresses"`
}

