// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package app


type AppTelemetryExportDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/app#unity_catalog App#unity_catalog}.
	UnityCatalog *AppTelemetryExportDestinationsUnityCatalog `field:"optional" json:"unityCatalog" yaml:"unityCatalog"`
}

