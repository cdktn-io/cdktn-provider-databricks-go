// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package externallocation


type ExternalLocationEffectiveFileEventQueueProvidedPubsub struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/external_location#managed_resource_id ExternalLocation#managed_resource_id}.
	ManagedResourceId *string `field:"optional" json:"managedResourceId" yaml:"managedResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/external_location#subscription_name ExternalLocation#subscription_name}.
	SubscriptionName *string `field:"optional" json:"subscriptionName" yaml:"subscriptionName"`
}

