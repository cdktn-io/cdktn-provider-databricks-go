// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package recipient


type RecipientPropertiesKvpairs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/recipient#properties Recipient#properties}.
	Properties *map[string]*string `field:"required" json:"properties" yaml:"properties"`
}

