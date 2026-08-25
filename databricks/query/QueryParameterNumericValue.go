// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package query


type QueryParameterNumericValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/query#value Query#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

