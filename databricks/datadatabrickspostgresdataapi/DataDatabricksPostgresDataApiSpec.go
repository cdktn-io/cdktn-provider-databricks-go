// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresdataapi


type DataDatabricksPostgresDataApiSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/postgres_data_api#db_aggregates_enabled DataDatabricksPostgresDataApi#db_aggregates_enabled}.
	DbAggregatesEnabled interface{} `field:"optional" json:"dbAggregatesEnabled" yaml:"dbAggregatesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/postgres_data_api#db_extra_search_path DataDatabricksPostgresDataApi#db_extra_search_path}.
	DbExtraSearchPath *[]*string `field:"optional" json:"dbExtraSearchPath" yaml:"dbExtraSearchPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/postgres_data_api#db_max_rows DataDatabricksPostgresDataApi#db_max_rows}.
	DbMaxRows *float64 `field:"optional" json:"dbMaxRows" yaml:"dbMaxRows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/postgres_data_api#db_schemas DataDatabricksPostgresDataApi#db_schemas}.
	DbSchemas *[]*string `field:"optional" json:"dbSchemas" yaml:"dbSchemas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/postgres_data_api#jwt_cache_max_lifetime DataDatabricksPostgresDataApi#jwt_cache_max_lifetime}.
	JwtCacheMaxLifetime *string `field:"optional" json:"jwtCacheMaxLifetime" yaml:"jwtCacheMaxLifetime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/postgres_data_api#jwt_role_claim_key DataDatabricksPostgresDataApi#jwt_role_claim_key}.
	JwtRoleClaimKey *string `field:"optional" json:"jwtRoleClaimKey" yaml:"jwtRoleClaimKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/postgres_data_api#openapi_mode DataDatabricksPostgresDataApi#openapi_mode}.
	OpenapiMode *string `field:"optional" json:"openapiMode" yaml:"openapiMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/postgres_data_api#server_cors_allowed_origins DataDatabricksPostgresDataApi#server_cors_allowed_origins}.
	ServerCorsAllowedOrigins *[]*string `field:"optional" json:"serverCorsAllowedOrigins" yaml:"serverCorsAllowedOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/postgres_data_api#server_timing_enabled DataDatabricksPostgresDataApi#server_timing_enabled}.
	ServerTimingEnabled interface{} `field:"optional" json:"serverTimingEnabled" yaml:"serverTimingEnabled"`
}

