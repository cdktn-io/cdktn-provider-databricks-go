// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sqlquery


type SqlQueryParameterNumber struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/sql_query#value SqlQuery#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

