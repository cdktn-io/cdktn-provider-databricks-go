// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sandbox


type SandboxSpecCompute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.133.0/docs/resources/sandbox#inactivity_timeout Sandbox#inactivity_timeout}.
	InactivityTimeout *string `field:"optional" json:"inactivityTimeout" yaml:"inactivityTimeout"`
}

