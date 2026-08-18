// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgrescdfconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PostgresCdfConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/postgres_cdf_config#catalog PostgresCdfConfig#catalog}.
	Catalog *string `field:"required" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/postgres_cdf_config#parent PostgresCdfConfig#parent}.
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/postgres_cdf_config#postgres_schema PostgresCdfConfig#postgres_schema}.
	PostgresSchema *string `field:"required" json:"postgresSchema" yaml:"postgresSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/postgres_cdf_config#schema PostgresCdfConfig#schema}.
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/postgres_cdf_config#cdf_config_id PostgresCdfConfig#cdf_config_id}.
	CdfConfigId *string `field:"optional" json:"cdfConfigId" yaml:"cdfConfigId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/postgres_cdf_config#provider_config PostgresCdfConfig#provider_config}.
	ProviderConfig *PostgresCdfConfigProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

