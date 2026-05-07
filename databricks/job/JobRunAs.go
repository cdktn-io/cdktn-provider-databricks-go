// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobRunAs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/job#group_name Job#group_name}.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/job#service_principal_name Job#service_principal_name}.
	ServicePrincipalName *string `field:"optional" json:"servicePrincipalName" yaml:"servicePrincipalName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/job#user_name Job#user_name}.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

