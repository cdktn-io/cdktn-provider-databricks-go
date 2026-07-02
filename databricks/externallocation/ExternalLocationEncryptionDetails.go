// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package externallocation


type ExternalLocationEncryptionDetails struct {
	// sse_encryption_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/external_location#sse_encryption_details ExternalLocation#sse_encryption_details}
	SseEncryptionDetails *ExternalLocationEncryptionDetailsSseEncryptionDetails `field:"optional" json:"sseEncryptionDetails" yaml:"sseEncryptionDetails"`
}

