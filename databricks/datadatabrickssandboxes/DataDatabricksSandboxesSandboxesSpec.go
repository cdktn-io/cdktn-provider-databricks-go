// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickssandboxes


type DataDatabricksSandboxesSandboxesSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/data-sources/sandboxes#compute DataDatabricksSandboxes#compute}.
	Compute *DataDatabricksSandboxesSandboxesSpecCompute `field:"optional" json:"compute" yaml:"compute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/data-sources/sandboxes#environment DataDatabricksSandboxes#environment}.
	Environment *DataDatabricksSandboxesSandboxesSpecEnvironment `field:"optional" json:"environment" yaml:"environment"`
}

