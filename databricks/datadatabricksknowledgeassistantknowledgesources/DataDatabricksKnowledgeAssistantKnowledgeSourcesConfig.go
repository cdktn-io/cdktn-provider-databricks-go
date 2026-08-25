// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksknowledgeassistantknowledgesources

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksKnowledgeAssistantKnowledgeSourcesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/knowledge_assistant_knowledge_sources#parent DataDatabricksKnowledgeAssistantKnowledgeSources#parent}.
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/knowledge_assistant_knowledge_sources#page_size DataDatabricksKnowledgeAssistantKnowledgeSources#page_size}.
	PageSize *float64 `field:"optional" json:"pageSize" yaml:"pageSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/knowledge_assistant_knowledge_sources#provider_config DataDatabricksKnowledgeAssistantKnowledgeSources#provider_config}.
	ProviderConfig *DataDatabricksKnowledgeAssistantKnowledgeSourcesProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

