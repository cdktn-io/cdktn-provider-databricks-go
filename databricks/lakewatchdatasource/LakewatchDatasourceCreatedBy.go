// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakewatchdatasource


type LakewatchDatasourceCreatedBy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/lakewatch_datasource#display_name LakewatchDatasource#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/lakewatch_datasource#email LakewatchDatasource#email}.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/lakewatch_datasource#id LakewatchDatasource#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/lakewatch_datasource#principal_type LakewatchDatasource#principal_type}.
	PrincipalType *string `field:"optional" json:"principalType" yaml:"principalType"`
}

