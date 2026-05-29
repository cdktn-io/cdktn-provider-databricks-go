// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package onlinestore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OnlineStoreConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/online_store#capacity OnlineStore#capacity}.
	Capacity *string `field:"required" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/online_store#name OnlineStore#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/online_store#provider_config OnlineStore#provider_config}.
	ProviderConfig *OnlineStoreProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/online_store#read_replica_count OnlineStore#read_replica_count}.
	ReadReplicaCount *float64 `field:"optional" json:"readReplicaCount" yaml:"readReplicaCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/online_store#usage_policy_id OnlineStore#usage_policy_id}.
	UsagePolicyId *string `field:"optional" json:"usagePolicyId" yaml:"usagePolicyId"`
}

