// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobEnvironmentSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/job#base_environment Job#base_environment}.
	BaseEnvironment *string `field:"optional" json:"baseEnvironment" yaml:"baseEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/job#client Job#client}.
	Client *string `field:"optional" json:"client" yaml:"client"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/job#dependencies Job#dependencies}.
	Dependencies *[]*string `field:"optional" json:"dependencies" yaml:"dependencies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/job#environment_version Job#environment_version}.
	EnvironmentVersion *string `field:"optional" json:"environmentVersion" yaml:"environmentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/job#java_dependencies Job#java_dependencies}.
	JavaDependencies *[]*string `field:"optional" json:"javaDependencies" yaml:"javaDependencies"`
}

