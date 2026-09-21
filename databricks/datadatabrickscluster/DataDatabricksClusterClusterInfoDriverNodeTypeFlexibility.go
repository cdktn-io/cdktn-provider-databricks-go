// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster


type DataDatabricksClusterClusterInfoDriverNodeTypeFlexibility struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.133.0/docs/data-sources/cluster#alternate_node_type_ids DataDatabricksCluster#alternate_node_type_ids}.
	AlternateNodeTypeIds *[]*string `field:"optional" json:"alternateNodeTypeIds" yaml:"alternateNodeTypeIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.133.0/docs/data-sources/cluster#aws_context_id DataDatabricksCluster#aws_context_id}.
	AwsContextId *string `field:"optional" json:"awsContextId" yaml:"awsContextId"`
}

