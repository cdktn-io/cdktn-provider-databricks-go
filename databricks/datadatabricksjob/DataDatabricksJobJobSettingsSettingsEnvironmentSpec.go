// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsEnvironmentSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/job#base_environment DataDatabricksJob#base_environment}.
	BaseEnvironment *string `field:"optional" json:"baseEnvironment" yaml:"baseEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/job#client DataDatabricksJob#client}.
	Client *string `field:"optional" json:"client" yaml:"client"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/job#dependencies DataDatabricksJob#dependencies}.
	Dependencies *[]*string `field:"optional" json:"dependencies" yaml:"dependencies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/job#environment_version DataDatabricksJob#environment_version}.
	EnvironmentVersion *string `field:"optional" json:"environmentVersion" yaml:"environmentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/job#java_dependencies DataDatabricksJob#java_dependencies}.
	JavaDependencies *[]*string `field:"optional" json:"javaDependencies" yaml:"javaDependencies"`
}

