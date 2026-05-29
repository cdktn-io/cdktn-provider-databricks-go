// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qualitymonitor


type QualityMonitorNotificationsOnNewClassificationTagDetected struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/quality_monitor#email_addresses QualityMonitor#email_addresses}.
	EmailAddresses *[]*string `field:"optional" json:"emailAddresses" yaml:"emailAddresses"`
}

