// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksawsunitycatalogassumerolepolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAwsUnityCatalogAssumeRolePolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/aws_unity_catalog_assume_role_policy#aws_account_id DataDatabricksAwsUnityCatalogAssumeRolePolicy#aws_account_id}.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/aws_unity_catalog_assume_role_policy#external_id DataDatabricksAwsUnityCatalogAssumeRolePolicy#external_id}.
	ExternalId *string `field:"required" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/aws_unity_catalog_assume_role_policy#role_name DataDatabricksAwsUnityCatalogAssumeRolePolicy#role_name}.
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/aws_unity_catalog_assume_role_policy#aws_partition DataDatabricksAwsUnityCatalogAssumeRolePolicy#aws_partition}.
	AwsPartition *string `field:"optional" json:"awsPartition" yaml:"awsPartition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/aws_unity_catalog_assume_role_policy#unity_catalog_iam_arn DataDatabricksAwsUnityCatalogAssumeRolePolicy#unity_catalog_iam_arn}.
	UnityCatalogIamArn *string `field:"optional" json:"unityCatalogIamArn" yaml:"unityCatalogIamArn"`
}

