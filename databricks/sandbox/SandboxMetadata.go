// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sandbox


type SandboxMetadata struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/sandbox#custom_tags Sandbox#custom_tags}.
	CustomTags *map[string]*string `field:"optional" json:"customTags" yaml:"customTags"`
}

