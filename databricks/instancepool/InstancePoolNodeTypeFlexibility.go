// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package instancepool


type InstancePoolNodeTypeFlexibility struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/instance_pool#alternate_node_type_ids InstancePool#alternate_node_type_ids}.
	AlternateNodeTypeIds *[]*string `field:"required" json:"alternateNodeTypeIds" yaml:"alternateNodeTypeIds"`
}

