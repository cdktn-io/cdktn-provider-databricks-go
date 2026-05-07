// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/pipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference interface {
	cdktn.ComplexObject
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
	DataLevel() *string
	SetDataLevel(val *string)
	DataLevelInput() *string
	Dimensions() *[]*string
	SetDimensions(val *[]*string)
	DimensionsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptions
	SetInternalValue(val *PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptions)
	LookbackWindowDays() *float64
	SetLookbackWindowDays(val *float64)
	LookbackWindowDaysInput() *float64
	Metrics() *[]*string
	SetMetrics(val *[]*string)
	MetricsInput() *[]*string
	QueryLifetime() interface{}
	SetQueryLifetime(val interface{})
	QueryLifetimeInput() interface{}
	ReportType() *string
	SetReportType(val *string)
	ReportTypeInput() *string
	SyncStartDate() *string
	SetSyncStartDate(val *string)
	SyncStartDateInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	ResetDataLevel()
	ResetDimensions()
	ResetLookbackWindowDays()
	ResetMetrics()
	ResetQueryLifetime()
	ResetReportType()
	ResetSyncStartDate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference
type jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) DataLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) DataLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) Dimensions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) DimensionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) InternalValue() *PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) LookbackWindowDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lookbackWindowDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) LookbackWindowDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lookbackWindowDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) Metrics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) MetricsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) QueryLifetime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) QueryLifetimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) ReportType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) ReportTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) SyncStartDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncStartDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) SyncStartDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncStartDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference_Override(p PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference)SetDataLevel(val *string) {
	if err := j.validateSetDataLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataLevel",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference)SetDimensions(val *[]*string) {
	if err := j.validateSetDimensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimensions",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference)SetInternalValue(val *PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference)SetLookbackWindowDays(val *float64) {
	if err := j.validateSetLookbackWindowDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lookbackWindowDays",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference)SetMetrics(val *[]*string) {
	if err := j.validateSetMetricsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metrics",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference)SetQueryLifetime(val interface{}) {
	if err := j.validateSetQueryLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryLifetime",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference)SetReportType(val *string) {
	if err := j.validateSetReportTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportType",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference)SetSyncStartDate(val *string) {
	if err := j.validateSetSyncStartDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncStartDate",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) ResetDataLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetDataLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) ResetDimensions() {
	_jsii_.InvokeVoid(
		p,
		"resetDimensions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) ResetLookbackWindowDays() {
	_jsii_.InvokeVoid(
		p,
		"resetLookbackWindowDays",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) ResetMetrics() {
	_jsii_.InvokeVoid(
		p,
		"resetMetrics",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) ResetQueryLifetime() {
	_jsii_.InvokeVoid(
		p,
		"resetQueryLifetime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) ResetReportType() {
	_jsii_.InvokeVoid(
		p,
		"resetReportType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) ResetSyncStartDate() {
	_jsii_.InvokeVoid(
		p,
		"resetSyncStartDate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

