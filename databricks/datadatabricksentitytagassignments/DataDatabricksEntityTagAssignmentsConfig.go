// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksentitytagassignments

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksEntityTagAssignmentsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/entity_tag_assignments#entity_name DataDatabricksEntityTagAssignments#entity_name}.
	EntityName *string `field:"required" json:"entityName" yaml:"entityName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/entity_tag_assignments#entity_type DataDatabricksEntityTagAssignments#entity_type}.
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/entity_tag_assignments#max_results DataDatabricksEntityTagAssignments#max_results}.
	MaxResults *float64 `field:"optional" json:"maxResults" yaml:"maxResults"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/entity_tag_assignments#provider_config DataDatabricksEntityTagAssignments#provider_config}.
	ProviderConfig *DataDatabricksEntityTagAssignmentsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

