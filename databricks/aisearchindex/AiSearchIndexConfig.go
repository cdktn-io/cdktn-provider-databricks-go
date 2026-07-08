// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchindex

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AiSearchIndexConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/ai_search_index#index_type AiSearchIndex#index_type}.
	IndexType *string `field:"required" json:"indexType" yaml:"indexType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/ai_search_index#parent AiSearchIndex#parent}.
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/ai_search_index#primary_key AiSearchIndex#primary_key}.
	PrimaryKey *string `field:"required" json:"primaryKey" yaml:"primaryKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/ai_search_index#delta_sync_index_spec AiSearchIndex#delta_sync_index_spec}.
	DeltaSyncIndexSpec *AiSearchIndexDeltaSyncIndexSpec `field:"optional" json:"deltaSyncIndexSpec" yaml:"deltaSyncIndexSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/ai_search_index#direct_access_index_spec AiSearchIndex#direct_access_index_spec}.
	DirectAccessIndexSpec *AiSearchIndexDirectAccessIndexSpec `field:"optional" json:"directAccessIndexSpec" yaml:"directAccessIndexSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/ai_search_index#index_id AiSearchIndex#index_id}.
	IndexId *string `field:"optional" json:"indexId" yaml:"indexId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/ai_search_index#index_subtype AiSearchIndex#index_subtype}.
	IndexSubtype *string `field:"optional" json:"indexSubtype" yaml:"indexSubtype"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/ai_search_index#provider_config AiSearchIndex#provider_config}.
	ProviderConfig *AiSearchIndexProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

