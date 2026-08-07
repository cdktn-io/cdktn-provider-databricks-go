// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksservingendpoints/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference interface {
	cdktn.ComplexObject
	AnnotationsTable() *string
	SetAnnotationsTable(val *string)
	AnnotationsTableInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogsTable() *string
	SetLogsTable(val *string)
	LogsTableInput() *string
	MetricsTable() *string
	SetMetricsTable(val *string)
	MetricsTableInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TracesTable() *string
	SetTracesTable(val *string)
	TracesTableInput() *string
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
	ResetAnnotationsTable()
	ResetLogsTable()
	ResetMetricsTable()
	ResetTracesTable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference
type jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) AnnotationsTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationsTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) AnnotationsTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationsTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) LogsTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logsTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) LogsTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logsTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) MetricsTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) MetricsTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) TracesTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tracesTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) TracesTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tracesTableInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksServingEndpoints.DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference_Override(d DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksServingEndpoints.DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference)SetAnnotationsTable(val *string) {
	if err := j.validateSetAnnotationsTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotationsTable",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference)SetLogsTable(val *string) {
	if err := j.validateSetLogsTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logsTable",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference)SetMetricsTable(val *string) {
	if err := j.validateSetMetricsTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsTable",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference)SetTracesTable(val *string) {
	if err := j.validateSetTracesTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tracesTable",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) ResetAnnotationsTable() {
	_jsii_.InvokeVoid(
		d,
		"resetAnnotationsTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) ResetLogsTable() {
	_jsii_.InvokeVoid(
		d,
		"resetLogsTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) ResetMetricsTable() {
	_jsii_.InvokeVoid(
		d,
		"resetMetricsTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) ResetTracesTable() {
	_jsii_.InvokeVoid(
		d,
		"resetTracesTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNamesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

