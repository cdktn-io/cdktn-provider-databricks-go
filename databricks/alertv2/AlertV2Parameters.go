// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package alertv2


type AlertV2Parameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/alert_v2#name AlertV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/alert_v2#type AlertV2#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/alert_v2#value AlertV2#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

