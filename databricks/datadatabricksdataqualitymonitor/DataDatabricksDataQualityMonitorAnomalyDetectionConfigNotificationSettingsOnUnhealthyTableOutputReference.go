// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualitymonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabricksdataqualitymonitor/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference interface {
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
	DestinationIds() *[]*string
	SetDestinationIds(val *[]*string)
	DestinationIdsInput() *[]*string
	DestinationUserIds() *[]*float64
	SetDestinationUserIds(val *[]*float64)
	DestinationUserIdsInput() *[]*float64
	EmailAddresses() *[]*string
	SetEmailAddresses(val *[]*string)
	EmailAddressesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetDestinationIds()
	ResetDestinationUserIds()
	ResetEmailAddresses()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference
type jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) DestinationIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) DestinationIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) DestinationUserIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"destinationUserIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) DestinationUserIdsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"destinationUserIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) EmailAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) EmailAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksDataQualityMonitor.DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference_Override(d DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksDataQualityMonitor.DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference)SetDestinationIds(val *[]*string) {
	if err := j.validateSetDestinationIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationIds",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference)SetDestinationUserIds(val *[]*float64) {
	if err := j.validateSetDestinationUserIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationUserIds",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference)SetEmailAddresses(val *[]*string) {
	if err := j.validateSetEmailAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAddresses",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) ResetDestinationIds() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinationIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) ResetDestinationUserIds() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinationUserIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) ResetEmailAddresses() {
	_jsii_.InvokeVoid(
		d,
		"resetEmailAddresses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTableOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

