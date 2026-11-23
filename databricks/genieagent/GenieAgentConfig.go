// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package genieagent

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GenieAgentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#definition GenieAgent#definition}.
	Definition *GenieAgentDefinition `field:"required" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#display_name GenieAgent#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#warehouse_id GenieAgent#warehouse_id}.
	WarehouseId *string `field:"required" json:"warehouseId" yaml:"warehouseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#description GenieAgent#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#genie_agent_id GenieAgent#genie_agent_id}.
	GenieAgentId *string `field:"optional" json:"genieAgentId" yaml:"genieAgentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#parent_path GenieAgent#parent_path}.
	ParentPath *string `field:"optional" json:"parentPath" yaml:"parentPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/genie_agent#provider_config GenieAgent#provider_config}.
	ProviderConfig *GenieAgentProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

