// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package domain

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/domain#tag_key Domain#tag_key}.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/domain#business_owner_ids Domain#business_owner_ids}.
	BusinessOwnerIds *[]*float64 `field:"optional" json:"businessOwnerIds" yaml:"businessOwnerIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/domain#description Domain#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/domain#domain_id Domain#domain_id}.
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/domain#draft Domain#draft}.
	Draft interface{} `field:"optional" json:"draft" yaml:"draft"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/domain#icon Domain#icon}.
	Icon *DomainIcon `field:"optional" json:"icon" yaml:"icon"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/domain#parent_domain_id Domain#parent_domain_id}.
	ParentDomainId *string `field:"optional" json:"parentDomainId" yaml:"parentDomainId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/domain#provider_config Domain#provider_config}.
	ProviderConfig *DomainProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/domain#subtitle Domain#subtitle}.
	Subtitle *string `field:"optional" json:"subtitle" yaml:"subtitle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/domain#technical_owner_ids Domain#technical_owner_ids}.
	TechnicalOwnerIds *[]*float64 `field:"optional" json:"technicalOwnerIds" yaml:"technicalOwnerIds"`
}

