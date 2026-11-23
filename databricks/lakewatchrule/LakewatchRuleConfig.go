// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakewatchrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LakewatchRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#catalog_name LakewatchRule#catalog_name}.
	CatalogName *string `field:"required" json:"catalogName" yaml:"catalogName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#fidelity LakewatchRule#fidelity}.
	Fidelity *string `field:"required" json:"fidelity" yaml:"fidelity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#output LakewatchRule#output}.
	Output *LakewatchRuleOutput `field:"required" json:"output" yaml:"output"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#rule_name LakewatchRule#rule_name}.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#schedule LakewatchRule#schedule}.
	Schedule *LakewatchRuleSchedule `field:"required" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#schema_name LakewatchRule#schema_name}.
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#severity LakewatchRule#severity}.
	Severity *string `field:"required" json:"severity" yaml:"severity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#annotations LakewatchRule#annotations}.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#batch LakewatchRule#batch}.
	Batch *LakewatchRuleBatch `field:"optional" json:"batch" yaml:"batch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#category LakewatchRule#category}.
	Category *string `field:"optional" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#content_resource_id LakewatchRule#content_resource_id}.
	ContentResourceId *string `field:"optional" json:"contentResourceId" yaml:"contentResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#deduplication LakewatchRule#deduplication}.
	Deduplication *string `field:"optional" json:"deduplication" yaml:"deduplication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#deduplication_time_window LakewatchRule#deduplication_time_window}.
	DeduplicationTimeWindow *string `field:"optional" json:"deduplicationTimeWindow" yaml:"deduplicationTimeWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#description LakewatchRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#mitre LakewatchRule#mitre}.
	Mitre interface{} `field:"optional" json:"mitre" yaml:"mitre"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#objective LakewatchRule#objective}.
	Objective *string `field:"optional" json:"objective" yaml:"objective"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#provider_config LakewatchRule#provider_config}.
	ProviderConfig *LakewatchRuleProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#references LakewatchRule#references}.
	References *[]*string `field:"optional" json:"references" yaml:"references"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#run_as_principal LakewatchRule#run_as_principal}.
	RunAsPrincipal *string `field:"optional" json:"runAsPrincipal" yaml:"runAsPrincipal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#runbook LakewatchRule#runbook}.
	Runbook *string `field:"optional" json:"runbook" yaml:"runbook"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#suppress_until_time LakewatchRule#suppress_until_time}.
	SuppressUntilTime *string `field:"optional" json:"suppressUntilTime" yaml:"suppressUntilTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_rule#timerange LakewatchRule#timerange}.
	Timerange *LakewatchRuleTimerange `field:"optional" json:"timerange" yaml:"timerange"`
}

