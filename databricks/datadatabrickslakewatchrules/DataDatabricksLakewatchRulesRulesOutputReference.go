// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabrickslakewatchrules/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksLakewatchRulesRulesOutputReference interface {
	cdktn.ComplexObject
	Annotations() cdktn.StringMap
	Batch() DataDatabricksLakewatchRulesRulesBatchOutputReference
	CatalogName() *string
	Category() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ComputeWorkspaceId() *string
	ContentResourceId() *string
	CreatedBy() DataDatabricksLakewatchRulesRulesCreatedByOutputReference
	CreateTime() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Deduplication() *string
	DeduplicationTimeWindow() *string
	Description() *string
	EffectiveCatalogName() *string
	EffectiveRuleName() *string
	EffectiveRunAsPrincipal() *string
	EffectiveSchemaName() *string
	Etag() *string
	Fidelity() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksLakewatchRulesRules
	SetInternalValue(val *DataDatabricksLakewatchRulesRules)
	Mitre() DataDatabricksLakewatchRulesRulesMitreList
	Name() *string
	SetName(val *string)
	NameInput() *string
	Objective() *string
	OriginalContentPack() DataDatabricksLakewatchRulesRulesOriginalContentPackOutputReference
	Output() DataDatabricksLakewatchRulesRulesOutputOutputReference
	Owner() DataDatabricksLakewatchRulesRulesOwnerOutputReference
	ProviderConfig() DataDatabricksLakewatchRulesRulesProviderConfigOutputReference
	ProviderConfigInput() interface{}
	References() *[]*string
	RuleId() *string
	RuleName() *string
	RunAsPrincipal() *string
	Runbook() *string
	Schedule() DataDatabricksLakewatchRulesRulesScheduleOutputReference
	SchemaName() *string
	Severity() *string
	Status() DataDatabricksLakewatchRulesRulesStatusOutputReference
	SuppressUntilTime() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Timerange() DataDatabricksLakewatchRulesRulesTimerangeOutputReference
	UpdatedBy() DataDatabricksLakewatchRulesRulesUpdatedByOutputReference
	UpdateTime() *string
	Version() *float64
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutProviderConfig(value *DataDatabricksLakewatchRulesRulesProviderConfig)
	ResetProviderConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksLakewatchRulesRulesOutputReference
type jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Annotations() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Batch() DataDatabricksLakewatchRulesRulesBatchOutputReference {
	var returns DataDatabricksLakewatchRulesRulesBatchOutputReference
	_jsii_.Get(
		j,
		"batch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) CatalogName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Category() *string {
	var returns *string
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) ComputeWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) ContentResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) CreatedBy() DataDatabricksLakewatchRulesRulesCreatedByOutputReference {
	var returns DataDatabricksLakewatchRulesRulesCreatedByOutputReference
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Deduplication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deduplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) DeduplicationTimeWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deduplicationTimeWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) EffectiveCatalogName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveCatalogName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) EffectiveRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) EffectiveRunAsPrincipal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveRunAsPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) EffectiveSchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveSchemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Fidelity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fidelity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) InternalValue() *DataDatabricksLakewatchRulesRules {
	var returns *DataDatabricksLakewatchRulesRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Mitre() DataDatabricksLakewatchRulesRulesMitreList {
	var returns DataDatabricksLakewatchRulesRulesMitreList
	_jsii_.Get(
		j,
		"mitre",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Objective() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objective",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) OriginalContentPack() DataDatabricksLakewatchRulesRulesOriginalContentPackOutputReference {
	var returns DataDatabricksLakewatchRulesRulesOriginalContentPackOutputReference
	_jsii_.Get(
		j,
		"originalContentPack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Output() DataDatabricksLakewatchRulesRulesOutputOutputReference {
	var returns DataDatabricksLakewatchRulesRulesOutputOutputReference
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Owner() DataDatabricksLakewatchRulesRulesOwnerOutputReference {
	var returns DataDatabricksLakewatchRulesRulesOwnerOutputReference
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) ProviderConfig() DataDatabricksLakewatchRulesRulesProviderConfigOutputReference {
	var returns DataDatabricksLakewatchRulesRulesProviderConfigOutputReference
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) References() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"references",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) RuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) RunAsPrincipal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Runbook() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runbook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Schedule() DataDatabricksLakewatchRulesRulesScheduleOutputReference {
	var returns DataDatabricksLakewatchRulesRulesScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) SchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Severity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Status() DataDatabricksLakewatchRulesRulesStatusOutputReference {
	var returns DataDatabricksLakewatchRulesRulesStatusOutputReference
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) SuppressUntilTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suppressUntilTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Timerange() DataDatabricksLakewatchRulesRulesTimerangeOutputReference {
	var returns DataDatabricksLakewatchRulesRulesTimerangeOutputReference
	_jsii_.Get(
		j,
		"timerange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) UpdatedBy() DataDatabricksLakewatchRulesRulesUpdatedByOutputReference {
	var returns DataDatabricksLakewatchRulesRulesUpdatedByOutputReference
	_jsii_.Get(
		j,
		"updatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewDataDatabricksLakewatchRulesRulesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksLakewatchRulesRulesOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksLakewatchRulesRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchRules.DataDatabricksLakewatchRulesRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksLakewatchRulesRulesOutputReference_Override(d DataDatabricksLakewatchRulesRulesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchRules.DataDatabricksLakewatchRulesRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference)SetInternalValue(val *DataDatabricksLakewatchRulesRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) PutProviderConfig(value *DataDatabricksLakewatchRulesRulesProviderConfig) {
	if err := d.validatePutProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProviderConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

