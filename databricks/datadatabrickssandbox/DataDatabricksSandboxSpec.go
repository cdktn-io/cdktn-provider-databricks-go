// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickssandbox


type DataDatabricksSandboxSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.133.0/docs/data-sources/sandbox#compute DataDatabricksSandbox#compute}.
	Compute *DataDatabricksSandboxSpecCompute `field:"optional" json:"compute" yaml:"compute"`
}

