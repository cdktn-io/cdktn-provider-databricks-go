// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobJobClusterNewClusterWorkerNodeTypeFlexibility struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/job#alternate_node_type_ids Job#alternate_node_type_ids}.
	AlternateNodeTypeIds *[]*string `field:"optional" json:"alternateNodeTypeIds" yaml:"alternateNodeTypeIds"`
}

