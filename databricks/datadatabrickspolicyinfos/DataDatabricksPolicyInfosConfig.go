// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspolicyinfos

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksPolicyInfosConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/policy_infos#on_securable_fullname DataDatabricksPolicyInfos#on_securable_fullname}.
	OnSecurableFullname *string `field:"required" json:"onSecurableFullname" yaml:"onSecurableFullname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/policy_infos#on_securable_type DataDatabricksPolicyInfos#on_securable_type}.
	OnSecurableType *string `field:"required" json:"onSecurableType" yaml:"onSecurableType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/policy_infos#include_inherited DataDatabricksPolicyInfos#include_inherited}.
	IncludeInherited interface{} `field:"optional" json:"includeInherited" yaml:"includeInherited"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/policy_infos#max_results DataDatabricksPolicyInfos#max_results}.
	MaxResults *float64 `field:"optional" json:"maxResults" yaml:"maxResults"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/policy_infos#provider_config DataDatabricksPolicyInfos#provider_config}.
	ProviderConfig *DataDatabricksPolicyInfosProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

