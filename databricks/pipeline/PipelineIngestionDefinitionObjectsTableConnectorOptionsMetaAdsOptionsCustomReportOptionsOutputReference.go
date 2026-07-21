// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/pipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptions
	SetInternalValue(val *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptions)
	Level() *string
	SetLevel(val *string)
	LevelInput() *string
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
	ResetActionAttributionWindows()
	ResetActionBreakdowns()
	ResetActionReportTime()
	ResetBreakdowns()
	ResetLevel()
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

// The jsii proxy struct for PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference
type jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) ActionAttributionWindows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionAttributionWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) ActionAttributionWindowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionAttributionWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) ActionBreakdowns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionBreakdowns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) ActionBreakdownsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionBreakdownsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) ActionReportTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionReportTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) ActionReportTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionReportTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) Breakdowns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"breakdowns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) BreakdownsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"breakdownsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) InternalValue() *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) TimeIncrement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeIncrement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) TimeIncrementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeIncrementInput",
		&returns,
	)
	return returns
}


func NewPipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference_Override(p PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference)SetActionAttributionWindows(val *[]*string) {
	if err := j.validateSetActionAttributionWindowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionAttributionWindows",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference)SetActionBreakdowns(val *[]*string) {
	if err := j.validateSetActionBreakdownsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionBreakdowns",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference)SetActionReportTime(val *string) {
	if err := j.validateSetActionReportTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionReportTime",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference)SetBreakdowns(val *[]*string) {
	if err := j.validateSetBreakdownsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"breakdowns",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference)SetInternalValue(val *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference)SetTimeIncrement(val *string) {
	if err := j.validateSetTimeIncrementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeIncrement",
		val,
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) ResetActionAttributionWindows() {
	_jsii_.InvokeVoid(
		p,
		"resetActionAttributionWindows",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) ResetActionBreakdowns() {
	_jsii_.InvokeVoid(
		p,
		"resetActionBreakdowns",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) ResetActionReportTime() {
	_jsii_.InvokeVoid(
		p,
		"resetActionReportTime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) ResetBreakdowns() {
	_jsii_.InvokeVoid(
		p,
		"resetBreakdowns",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) ResetLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) ResetTimeIncrement() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeIncrement",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

