// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterDriverNodeTypeFlexibility struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/cluster#alternate_node_type_ids Cluster#alternate_node_type_ids}.
	AlternateNodeTypeIds *[]*string `field:"optional" json:"alternateNodeTypeIds" yaml:"alternateNodeTypeIds"`
}

