// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterWorkerNodeTypeFlexibility struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.132.0/docs/resources/cluster#alternate_node_type_ids Cluster#alternate_node_type_ids}.
	AlternateNodeTypeIds *[]*string `field:"optional" json:"alternateNodeTypeIds" yaml:"alternateNodeTypeIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.132.0/docs/resources/cluster#aws_context_id Cluster#aws_context_id}.
	AwsContextId *string `field:"optional" json:"awsContextId" yaml:"awsContextId"`
}

