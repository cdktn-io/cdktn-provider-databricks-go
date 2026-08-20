// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksexternallocation


type DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueue struct {
	// managed_aqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/external_location#managed_aqs DataDatabricksExternalLocation#managed_aqs}
	ManagedAqs *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedAqs `field:"optional" json:"managedAqs" yaml:"managedAqs"`
	// managed_pubsub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/external_location#managed_pubsub DataDatabricksExternalLocation#managed_pubsub}
	ManagedPubsub *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedPubsub `field:"optional" json:"managedPubsub" yaml:"managedPubsub"`
	// managed_sqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/external_location#managed_sqs DataDatabricksExternalLocation#managed_sqs}
	ManagedSqs *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueManagedSqs `field:"optional" json:"managedSqs" yaml:"managedSqs"`
	// provided_aqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/external_location#provided_aqs DataDatabricksExternalLocation#provided_aqs}
	ProvidedAqs *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedAqs `field:"optional" json:"providedAqs" yaml:"providedAqs"`
	// provided_pubsub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/external_location#provided_pubsub DataDatabricksExternalLocation#provided_pubsub}
	ProvidedPubsub *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedPubsub `field:"optional" json:"providedPubsub" yaml:"providedPubsub"`
	// provided_sqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/external_location#provided_sqs DataDatabricksExternalLocation#provided_sqs}
	ProvidedSqs *DataDatabricksExternalLocationExternalLocationInfoEffectiveFileEventQueueProvidedSqs `field:"optional" json:"providedSqs" yaml:"providedSqs"`
}

