// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterInitScriptsGcs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/cluster#destination Cluster#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
}

