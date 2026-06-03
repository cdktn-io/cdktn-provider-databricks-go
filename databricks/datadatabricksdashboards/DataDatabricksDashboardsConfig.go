// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdashboards

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksDashboardsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/dashboards#dashboard_name_contains DataDatabricksDashboards#dashboard_name_contains}.
	DashboardNameContains *string `field:"optional" json:"dashboardNameContains" yaml:"dashboardNameContains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/dashboards#provider_config DataDatabricksDashboards#provider_config}.
	ProviderConfig *DataDatabricksDashboardsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

