// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksassumegrouprestrictedsharinggroups

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAssumeGroupRestrictedSharingGroupsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/assume_group_restricted_sharing_groups#page_size DataDatabricksAssumeGroupRestrictedSharingGroups#page_size}.
	PageSize *float64 `field:"optional" json:"pageSize" yaml:"pageSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/assume_group_restricted_sharing_groups#provider_config DataDatabricksAssumeGroupRestrictedSharingGroups#provider_config}.
	ProviderConfig *DataDatabricksAssumeGroupRestrictedSharingGroupsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

