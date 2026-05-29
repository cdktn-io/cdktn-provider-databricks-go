// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package knowledgeassistantknowledgesource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KnowledgeAssistantKnowledgeSourceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/knowledge_assistant_knowledge_source#description KnowledgeAssistantKnowledgeSource#description}.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/knowledge_assistant_knowledge_source#display_name KnowledgeAssistantKnowledgeSource#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/knowledge_assistant_knowledge_source#parent KnowledgeAssistantKnowledgeSource#parent}.
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/knowledge_assistant_knowledge_source#source_type KnowledgeAssistantKnowledgeSource#source_type}.
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/knowledge_assistant_knowledge_source#files KnowledgeAssistantKnowledgeSource#files}.
	Files *KnowledgeAssistantKnowledgeSourceFiles `field:"optional" json:"files" yaml:"files"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/knowledge_assistant_knowledge_source#file_table KnowledgeAssistantKnowledgeSource#file_table}.
	FileTable *KnowledgeAssistantKnowledgeSourceFileTable `field:"optional" json:"fileTable" yaml:"fileTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/knowledge_assistant_knowledge_source#index KnowledgeAssistantKnowledgeSource#index}.
	Index *KnowledgeAssistantKnowledgeSourceIndex `field:"optional" json:"index" yaml:"index"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/knowledge_assistant_knowledge_source#provider_config KnowledgeAssistantKnowledgeSource#provider_config}.
	ProviderConfig *KnowledgeAssistantKnowledgeSourceProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

