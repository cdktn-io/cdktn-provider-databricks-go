// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mount


type MountTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/mount#default Mount#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
}

