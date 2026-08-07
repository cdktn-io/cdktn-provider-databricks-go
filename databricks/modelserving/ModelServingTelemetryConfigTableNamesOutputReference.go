// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package modelserving

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/modelserving/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ModelServingTelemetryConfigTableNamesOutputReference interface {
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
	InternalValue() *ModelServingTelemetryConfigTableNames
	SetInternalValue(val *ModelServingTelemetryConfigTableNames)
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

// The jsii proxy struct for ModelServingTelemetryConfigTableNamesOutputReference
type jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) AnnotationsTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationsTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) AnnotationsTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationsTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) InternalValue() *ModelServingTelemetryConfigTableNames {
	var returns *ModelServingTelemetryConfigTableNames
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) LogsTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logsTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) LogsTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logsTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) MetricsTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) MetricsTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) TracesTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tracesTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) TracesTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tracesTableInput",
		&returns,
	)
	return returns
}


func NewModelServingTelemetryConfigTableNamesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ModelServingTelemetryConfigTableNamesOutputReference {
	_init_.Initialize()

	if err := validateNewModelServingTelemetryConfigTableNamesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.modelServing.ModelServingTelemetryConfigTableNamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewModelServingTelemetryConfigTableNamesOutputReference_Override(m ModelServingTelemetryConfigTableNamesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.modelServing.ModelServingTelemetryConfigTableNamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference)SetAnnotationsTable(val *string) {
	if err := j.validateSetAnnotationsTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotationsTable",
		val,
	)
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference)SetInternalValue(val *ModelServingTelemetryConfigTableNames) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference)SetLogsTable(val *string) {
	if err := j.validateSetLogsTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logsTable",
		val,
	)
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference)SetMetricsTable(val *string) {
	if err := j.validateSetMetricsTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsTable",
		val,
	)
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference)SetTracesTable(val *string) {
	if err := j.validateSetTracesTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tracesTable",
		val,
	)
}

func (m *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) ResetAnnotationsTable() {
	_jsii_.InvokeVoid(
		m,
		"resetAnnotationsTable",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) ResetLogsTable() {
	_jsii_.InvokeVoid(
		m,
		"resetLogsTable",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) ResetMetricsTable() {
	_jsii_.InvokeVoid(
		m,
		"resetMetricsTable",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) ResetTracesTable() {
	_jsii_.InvokeVoid(
		m,
		"resetTracesTable",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingTelemetryConfigTableNamesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

