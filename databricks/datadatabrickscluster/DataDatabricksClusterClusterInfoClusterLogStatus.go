// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster


type DataDatabricksClusterClusterInfoClusterLogStatus struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/cluster#last_attempted DataDatabricksCluster#last_attempted}.
	LastAttempted *float64 `field:"optional" json:"lastAttempted" yaml:"lastAttempted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/cluster#last_exception DataDatabricksCluster#last_exception}.
	LastException *string `field:"optional" json:"lastException" yaml:"lastException"`
}

