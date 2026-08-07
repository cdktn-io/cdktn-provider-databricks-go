// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/pipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference interface {
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
	CustomReportOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference
	CustomReportOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsCustomReportOptions
	// Experimental.
	Fqn() *string
	InternalValue() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptions
	SetInternalValue(val *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptions)
	LookbackWindowDays() *float64
	SetLookbackWindowDays(val *float64)
	LookbackWindowDaysInput() *float64
	ManagerAccountId() *string
	SetManagerAccountId(val *string)
	ManagerAccountIdInput() *string
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
	PutCustomReportOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsCustomReportOptions)
	ResetCustomReportOptions()
	ResetLookbackWindowDays()
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

// The jsii proxy struct for PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference
type jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) CustomReportOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsCustomReportOptionsOutputReference
	_jsii_.Get(
		j,
		"customReportOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) CustomReportOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsCustomReportOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsCustomReportOptions
	_jsii_.Get(
		j,
		"customReportOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) InternalValue() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) LookbackWindowDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lookbackWindowDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) LookbackWindowDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lookbackWindowDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) ManagerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) ManagerAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managerAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) SyncStartDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncStartDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) SyncStartDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncStartDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference_Override(p PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference)SetInternalValue(val *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference)SetLookbackWindowDays(val *float64) {
	if err := j.validateSetLookbackWindowDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lookbackWindowDays",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference)SetManagerAccountId(val *string) {
	if err := j.validateSetManagerAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managerAccountId",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference)SetSyncStartDate(val *string) {
	if err := j.validateSetSyncStartDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncStartDate",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) PutCustomReportOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsCustomReportOptions) {
	if err := p.validatePutCustomReportOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCustomReportOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) ResetCustomReportOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetCustomReportOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) ResetLookbackWindowDays() {
	_jsii_.InvokeVoid(
		p,
		"resetLookbackWindowDays",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) ResetSyncStartDate() {
	_jsii_.InvokeVoid(
		p,
		"resetSyncStartDate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

