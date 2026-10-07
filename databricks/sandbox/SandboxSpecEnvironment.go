// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sandbox


type SandboxSpecEnvironment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/resources/sandbox#image_uri Sandbox#image_uri}.
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
}

