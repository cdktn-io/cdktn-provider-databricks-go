// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package domain


type DomainIcon struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.132.0/docs/resources/domain#color Domain#color}.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.132.0/docs/resources/domain#name Domain#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

