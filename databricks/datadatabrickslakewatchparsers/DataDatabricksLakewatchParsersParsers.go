// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchparsers


type DataDatabricksLakewatchParsersParsers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_parsers#name DataDatabricksLakewatchParsers#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_parsers#provider_config DataDatabricksLakewatchParsers#provider_config}.
	ProviderConfig *DataDatabricksLakewatchParsersParsersProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

