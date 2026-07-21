// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/pipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference interface {
	cdktn.ComplexObject
	AutoFullRefreshPolicy() PipelineIngestionDefinitionObjectsTableTableConfigurationAutoFullRefreshPolicyOutputReference
	AutoFullRefreshPolicyInput() *PipelineIngestionDefinitionObjectsTableTableConfigurationAutoFullRefreshPolicy
	ClusteringColumns() *[]*string
	SetClusteringColumns(val *[]*string)
	ClusteringColumnsInput() *[]*string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableAutoClustering() interface{}
	SetEnableAutoClustering(val interface{})
	EnableAutoClusteringInput() interface{}
	ExcludeColumns() *[]*string
	SetExcludeColumns(val *[]*string)
	ExcludeColumnsInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludeColumns() *[]*string
	SetIncludeColumns(val *[]*string)
	IncludeColumnsInput() *[]*string
	InternalValue() *PipelineIngestionDefinitionObjectsTableTableConfiguration
	SetInternalValue(val *PipelineIngestionDefinitionObjectsTableTableConfiguration)
	PrimaryKeys() *[]*string
	SetPrimaryKeys(val *[]*string)
	PrimaryKeysInput() *[]*string
	QueryBasedConnectorConfig() PipelineIngestionDefinitionObjectsTableTableConfigurationQueryBasedConnectorConfigOutputReference
	QueryBasedConnectorConfigInput() *PipelineIngestionDefinitionObjectsTableTableConfigurationQueryBasedConnectorConfig
	RowFilter() *string
	SetRowFilter(val *string)
	RowFilterInput() *string
	SalesforceIncludeFormulaFields() interface{}
	SetSalesforceIncludeFormulaFields(val interface{})
	SalesforceIncludeFormulaFieldsInput() interface{}
	ScdType() *string
	SetScdType(val *string)
	ScdTypeInput() *string
	SequenceBy() *[]*string
	SetSequenceBy(val *[]*string)
	SequenceByInput() *[]*string
	SourceMetadataColumn() *string
	SetSourceMetadataColumn(val *string)
	SourceMetadataColumnInput() *string
	TableProperties() *map[string]*string
	SetTableProperties(val *map[string]*string)
	TablePropertiesInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WorkdayReportParameters() PipelineIngestionDefinitionObjectsTableTableConfigurationWorkdayReportParametersOutputReference
	WorkdayReportParametersInput() *PipelineIngestionDefinitionObjectsTableTableConfigurationWorkdayReportParameters
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
	PutAutoFullRefreshPolicy(value *PipelineIngestionDefinitionObjectsTableTableConfigurationAutoFullRefreshPolicy)
	PutQueryBasedConnectorConfig(value *PipelineIngestionDefinitionObjectsTableTableConfigurationQueryBasedConnectorConfig)
	PutWorkdayReportParameters(value *PipelineIngestionDefinitionObjectsTableTableConfigurationWorkdayReportParameters)
	ResetAutoFullRefreshPolicy()
	ResetClusteringColumns()
	ResetEnableAutoClustering()
	ResetExcludeColumns()
	ResetIncludeColumns()
	ResetPrimaryKeys()
	ResetQueryBasedConnectorConfig()
	ResetRowFilter()
	ResetSalesforceIncludeFormulaFields()
	ResetScdType()
	ResetSequenceBy()
	ResetSourceMetadataColumn()
	ResetTableProperties()
	ResetWorkdayReportParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference
type jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) AutoFullRefreshPolicy() PipelineIngestionDefinitionObjectsTableTableConfigurationAutoFullRefreshPolicyOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableTableConfigurationAutoFullRefreshPolicyOutputReference
	_jsii_.Get(
		j,
		"autoFullRefreshPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) AutoFullRefreshPolicyInput() *PipelineIngestionDefinitionObjectsTableTableConfigurationAutoFullRefreshPolicy {
	var returns *PipelineIngestionDefinitionObjectsTableTableConfigurationAutoFullRefreshPolicy
	_jsii_.Get(
		j,
		"autoFullRefreshPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ClusteringColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusteringColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ClusteringColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusteringColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) EnableAutoClustering() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoClustering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) EnableAutoClusteringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoClusteringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ExcludeColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ExcludeColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) IncludeColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) IncludeColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) InternalValue() *PipelineIngestionDefinitionObjectsTableTableConfiguration {
	var returns *PipelineIngestionDefinitionObjectsTableTableConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) PrimaryKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"primaryKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) PrimaryKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"primaryKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) QueryBasedConnectorConfig() PipelineIngestionDefinitionObjectsTableTableConfigurationQueryBasedConnectorConfigOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableTableConfigurationQueryBasedConnectorConfigOutputReference
	_jsii_.Get(
		j,
		"queryBasedConnectorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) QueryBasedConnectorConfigInput() *PipelineIngestionDefinitionObjectsTableTableConfigurationQueryBasedConnectorConfig {
	var returns *PipelineIngestionDefinitionObjectsTableTableConfigurationQueryBasedConnectorConfig
	_jsii_.Get(
		j,
		"queryBasedConnectorConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) RowFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rowFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) RowFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rowFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) SalesforceIncludeFormulaFields() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"salesforceIncludeFormulaFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) SalesforceIncludeFormulaFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"salesforceIncludeFormulaFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ScdType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scdType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ScdTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scdTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) SequenceBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sequenceBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) SequenceByInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sequenceByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) SourceMetadataColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceMetadataColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) SourceMetadataColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceMetadataColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) TableProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tableProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) TablePropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tablePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) WorkdayReportParameters() PipelineIngestionDefinitionObjectsTableTableConfigurationWorkdayReportParametersOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableTableConfigurationWorkdayReportParametersOutputReference
	_jsii_.Get(
		j,
		"workdayReportParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) WorkdayReportParametersInput() *PipelineIngestionDefinitionObjectsTableTableConfigurationWorkdayReportParameters {
	var returns *PipelineIngestionDefinitionObjectsTableTableConfigurationWorkdayReportParameters
	_jsii_.Get(
		j,
		"workdayReportParametersInput",
		&returns,
	)
	return returns
}


func NewPipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineIngestionDefinitionObjectsTableTableConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference_Override(p PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetClusteringColumns(val *[]*string) {
	if err := j.validateSetClusteringColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusteringColumns",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetEnableAutoClustering(val interface{}) {
	if err := j.validateSetEnableAutoClusteringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutoClustering",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetExcludeColumns(val *[]*string) {
	if err := j.validateSetExcludeColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeColumns",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetIncludeColumns(val *[]*string) {
	if err := j.validateSetIncludeColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeColumns",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetInternalValue(val *PipelineIngestionDefinitionObjectsTableTableConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetPrimaryKeys(val *[]*string) {
	if err := j.validateSetPrimaryKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryKeys",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetRowFilter(val *string) {
	if err := j.validateSetRowFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowFilter",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetSalesforceIncludeFormulaFields(val interface{}) {
	if err := j.validateSetSalesforceIncludeFormulaFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"salesforceIncludeFormulaFields",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetScdType(val *string) {
	if err := j.validateSetScdTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scdType",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetSequenceBy(val *[]*string) {
	if err := j.validateSetSequenceByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sequenceBy",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetSourceMetadataColumn(val *string) {
	if err := j.validateSetSourceMetadataColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceMetadataColumn",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetTableProperties(val *map[string]*string) {
	if err := j.validateSetTablePropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableProperties",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) PutAutoFullRefreshPolicy(value *PipelineIngestionDefinitionObjectsTableTableConfigurationAutoFullRefreshPolicy) {
	if err := p.validatePutAutoFullRefreshPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAutoFullRefreshPolicy",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) PutQueryBasedConnectorConfig(value *PipelineIngestionDefinitionObjectsTableTableConfigurationQueryBasedConnectorConfig) {
	if err := p.validatePutQueryBasedConnectorConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQueryBasedConnectorConfig",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) PutWorkdayReportParameters(value *PipelineIngestionDefinitionObjectsTableTableConfigurationWorkdayReportParameters) {
	if err := p.validatePutWorkdayReportParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putWorkdayReportParameters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetAutoFullRefreshPolicy() {
	_jsii_.InvokeVoid(
		p,
		"resetAutoFullRefreshPolicy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetClusteringColumns() {
	_jsii_.InvokeVoid(
		p,
		"resetClusteringColumns",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetEnableAutoClustering() {
	_jsii_.InvokeVoid(
		p,
		"resetEnableAutoClustering",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetExcludeColumns() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludeColumns",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetIncludeColumns() {
	_jsii_.InvokeVoid(
		p,
		"resetIncludeColumns",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetPrimaryKeys() {
	_jsii_.InvokeVoid(
		p,
		"resetPrimaryKeys",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetQueryBasedConnectorConfig() {
	_jsii_.InvokeVoid(
		p,
		"resetQueryBasedConnectorConfig",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetRowFilter() {
	_jsii_.InvokeVoid(
		p,
		"resetRowFilter",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetSalesforceIncludeFormulaFields() {
	_jsii_.InvokeVoid(
		p,
		"resetSalesforceIncludeFormulaFields",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetScdType() {
	_jsii_.InvokeVoid(
		p,
		"resetScdType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetSequenceBy() {
	_jsii_.InvokeVoid(
		p,
		"resetSequenceBy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetSourceMetadataColumn() {
	_jsii_.InvokeVoid(
		p,
		"resetSourceMetadataColumn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetTableProperties() {
	_jsii_.InvokeVoid(
		p,
		"resetTableProperties",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ResetWorkdayReportParameters() {
	_jsii_.InvokeVoid(
		p,
		"resetWorkdayReportParameters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableTableConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

