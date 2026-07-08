// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickssupervisoragenttool


type DataDatabricksSupervisorAgentToolKnowledgeAssistant struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/supervisor_agent_tool#knowledge_assistant_id DataDatabricksSupervisorAgentTool#knowledge_assistant_id}.
	KnowledgeAssistantId *string `field:"required" json:"knowledgeAssistantId" yaml:"knowledgeAssistantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/supervisor_agent_tool#serving_endpoint_name DataDatabricksSupervisorAgentTool#serving_endpoint_name}.
	ServingEndpointName *string `field:"optional" json:"servingEndpointName" yaml:"servingEndpointName"`
}

