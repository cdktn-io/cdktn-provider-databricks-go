// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package supervisoragent

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SupervisorAgentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/supervisor_agent#display_name SupervisorAgent#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/supervisor_agent#description SupervisorAgent#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/supervisor_agent#instructions SupervisorAgent#instructions}.
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/supervisor_agent#provider_config SupervisorAgent#provider_config}.
	ProviderConfig *SupervisorAgentProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

