// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterWorkloadTypeClients struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/cluster#jobs Cluster#jobs}.
	Jobs interface{} `field:"optional" json:"jobs" yaml:"jobs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/cluster#notebooks Cluster#notebooks}.
	Notebooks interface{} `field:"optional" json:"notebooks" yaml:"notebooks"`
}

