// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package budget


type BudgetAlertConfigurationsPrincipalOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/budget#override_threshold Budget#override_threshold}.
	OverrideThreshold *string `field:"optional" json:"overrideThreshold" yaml:"overrideThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/budget#principal_id Budget#principal_id}.
	PrincipalId *float64 `field:"optional" json:"principalId" yaml:"principalId"`
}

