// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package assumegrouprestrictedsharinggroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AssumeGroupRestrictedSharingGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/assume_group_restricted_sharing_group#restricted_sharing_group_id AssumeGroupRestrictedSharingGroup#restricted_sharing_group_id}.
	RestrictedSharingGroupId *string `field:"required" json:"restrictedSharingGroupId" yaml:"restrictedSharingGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/assume_group_restricted_sharing_group#provider_config AssumeGroupRestrictedSharingGroup#provider_config}.
	ProviderConfig *AssumeGroupRestrictedSharingGroupProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

