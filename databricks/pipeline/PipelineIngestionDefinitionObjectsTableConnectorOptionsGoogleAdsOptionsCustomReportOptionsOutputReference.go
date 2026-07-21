// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/pipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptions
	SetInternalValue(val *PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptions)
	Metrics() *[]*string
	SetMetrics(val *[]*string)
	MetricsInput() *[]*string
	Resource() *string
	SetResource(val *string)
	ResourceFields() *[]*string
	SetResourceFields(val *[]*string)
	ResourceFieldsInput() *[]*string
	ResourceInput() *string
	Segments() *[]*string
	SetSegments(val *[]*string)
	SegmentsInput() *[]*string
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
	ResetMetrics()
	ResetResourceFields()
	ResetSegments()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference
type jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) InternalValue() *PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) Metrics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) MetricsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) ResourceFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) ResourceFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) ResourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) Segments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"segments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) SegmentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"segmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference_Override(p PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference)SetInternalValue(val *PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference)SetMetrics(val *[]*string) {
	if err := j.validateSetMetricsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metrics",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference)SetResource(val *string) {
	if err := j.validateSetResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference)SetResourceFields(val *[]*string) {
	if err := j.validateSetResourceFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceFields",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference)SetSegments(val *[]*string) {
	if err := j.validateSetSegmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segments",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) ResetMetrics() {
	_jsii_.InvokeVoid(
		p,
		"resetMetrics",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) ResetResourceFields() {
	_jsii_.InvokeVoid(
		p,
		"resetResourceFields",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) ResetSegments() {
	_jsii_.InvokeVoid(
		p,
		"resetSegments",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

