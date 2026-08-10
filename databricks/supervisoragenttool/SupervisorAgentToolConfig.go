// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package supervisoragenttool

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SupervisorAgentToolConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/supervisor_agent_tool#parent SupervisorAgentTool#parent}.
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/supervisor_agent_tool#tool_id SupervisorAgentTool#tool_id}.
	ToolId *string `field:"required" json:"toolId" yaml:"toolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/supervisor_agent_tool#tool_type SupervisorAgentTool#tool_type}.
	ToolType *string `field:"required" json:"toolType" yaml:"toolType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/supervisor_agent_tool#app SupervisorAgentTool#app}.
	App *SupervisorAgentToolApp `field:"optional" json:"app" yaml:"app"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/supervisor_agent_tool#description SupervisorAgentTool#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/supervisor_agent_tool#genie_space SupervisorAgentTool#genie_space}.
	GenieSpace *SupervisorAgentToolGenieSpace `field:"optional" json:"genieSpace" yaml:"genieSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/supervisor_agent_tool#knowledge_assistant SupervisorAgentTool#knowledge_assistant}.
	KnowledgeAssistant *SupervisorAgentToolKnowledgeAssistant `field:"optional" json:"knowledgeAssistant" yaml:"knowledgeAssistant"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/supervisor_agent_tool#provider_config SupervisorAgentTool#provider_config}.
	ProviderConfig *SupervisorAgentToolProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/supervisor_agent_tool#uc_connection SupervisorAgentTool#uc_connection}.
	UcConnection *SupervisorAgentToolUcConnection `field:"optional" json:"ucConnection" yaml:"ucConnection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/supervisor_agent_tool#uc_function SupervisorAgentTool#uc_function}.
	UcFunction *SupervisorAgentToolUcFunction `field:"optional" json:"ucFunction" yaml:"ucFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/supervisor_agent_tool#volume SupervisorAgentTool#volume}.
	Volume *SupervisorAgentToolVolume `field:"optional" json:"volume" yaml:"volume"`
}

