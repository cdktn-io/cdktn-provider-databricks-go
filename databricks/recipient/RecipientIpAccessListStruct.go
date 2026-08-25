// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package recipient


type RecipientIpAccessListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/recipient#allowed_ip_addresses Recipient#allowed_ip_addresses}.
	AllowedIpAddresses *[]*string `field:"optional" json:"allowedIpAddresses" yaml:"allowedIpAddresses"`
}

