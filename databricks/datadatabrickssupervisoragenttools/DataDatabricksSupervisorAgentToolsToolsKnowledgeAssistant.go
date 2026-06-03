// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickssupervisoragenttools


type DataDatabricksSupervisorAgentToolsToolsKnowledgeAssistant struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/supervisor_agent_tools#knowledge_assistant_id DataDatabricksSupervisorAgentTools#knowledge_assistant_id}.
	KnowledgeAssistantId *string `field:"required" json:"knowledgeAssistantId" yaml:"knowledgeAssistantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/supervisor_agent_tools#serving_endpoint_name DataDatabricksSupervisorAgentTools#serving_endpoint_name}.
	ServingEndpointName *string `field:"optional" json:"servingEndpointName" yaml:"servingEndpointName"`
}

