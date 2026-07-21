// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/pipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference interface {
	cdktn.ComplexObject
	ActionAttributionWindows() *[]*string
	SetActionAttributionWindows(val *[]*string)
	ActionAttributionWindowsInput() *[]*string
	ActionBreakdowns() *[]*string
	SetActionBreakdowns(val *[]*string)
	ActionBreakdownsInput() *[]*string
	ActionReportTime() *string
	SetActionReportTime(val *string)
	ActionReportTimeInput() *string
	Breakdowns() *[]*string
	SetBreakdowns(val *[]*string)
	BreakdownsInput() *[]*string
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
	CustomInsightsLookbackWindow() *float64
	SetCustomInsightsLookbackWindow(val *float64)
	CustomInsightsLookbackWindowInput() *float64
	CustomReportOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference
	CustomReportOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptions
	// Experimental.
	Fqn() *string
	InternalValue() *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptions
	SetInternalValue(val *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptions)
	Level() *string
	SetLevel(val *string)
	LevelInput() *string
	StartDate() *string
	SetStartDate(val *string)
	StartDateInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeIncrement() *string
	SetTimeIncrement(val *string)
	TimeIncrementInput() *string
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
	PutCustomReportOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptions)
	ResetActionAttributionWindows()
	ResetActionBreakdowns()
	ResetActionReportTime()
	ResetBreakdowns()
	ResetCustomInsightsLookbackWindow()
	ResetCustomReportOptions()
	ResetLevel()
	ResetStartDate()
	ResetTimeIncrement()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference
type jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ActionAttributionWindows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionAttributionWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ActionAttributionWindowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionAttributionWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ActionBreakdowns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionBreakdowns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ActionBreakdownsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionBreakdownsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ActionReportTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionReportTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ActionReportTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionReportTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) Breakdowns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"breakdowns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) BreakdownsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"breakdownsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) CustomInsightsLookbackWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customInsightsLookbackWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) CustomInsightsLookbackWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customInsightsLookbackWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) CustomReportOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference
	_jsii_.Get(
		j,
		"customReportOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) CustomReportOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptions
	_jsii_.Get(
		j,
		"customReportOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) InternalValue() *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) StartDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) StartDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) TimeIncrement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeIncrement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) TimeIncrementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeIncrementInput",
		&returns,
	)
	return returns
}


func NewPipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference_Override(p PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference)SetActionAttributionWindows(val *[]*string) {
	if err := j.validateSetActionAttributionWindowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionAttributionWindows",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference)SetActionBreakdowns(val *[]*string) {
	if err := j.validateSetActionBreakdownsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionBreakdowns",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference)SetActionReportTime(val *string) {
	if err := j.validateSetActionReportTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionReportTime",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference)SetBreakdowns(val *[]*string) {
	if err := j.validateSetBreakdownsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"breakdowns",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference)SetCustomInsightsLookbackWindow(val *float64) {
	if err := j.validateSetCustomInsightsLookbackWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customInsightsLookbackWindow",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference)SetInternalValue(val *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference)SetStartDate(val *string) {
	if err := j.validateSetStartDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startDate",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference)SetTimeIncrement(val *string) {
	if err := j.validateSetTimeIncrementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeIncrement",
		val,
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) PutCustomReportOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptions) {
	if err := p.validatePutCustomReportOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCustomReportOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ResetActionAttributionWindows() {
	_jsii_.InvokeVoid(
		p,
		"resetActionAttributionWindows",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ResetActionBreakdowns() {
	_jsii_.InvokeVoid(
		p,
		"resetActionBreakdowns",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ResetActionReportTime() {
	_jsii_.InvokeVoid(
		p,
		"resetActionReportTime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ResetBreakdowns() {
	_jsii_.InvokeVoid(
		p,
		"resetBreakdowns",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ResetCustomInsightsLookbackWindow() {
	_jsii_.InvokeVoid(
		p,
		"resetCustomInsightsLookbackWindow",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ResetCustomReportOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetCustomReportOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ResetLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ResetStartDate() {
	_jsii_.InvokeVoid(
		p,
		"resetStartDate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ResetTimeIncrement() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeIncrement",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

