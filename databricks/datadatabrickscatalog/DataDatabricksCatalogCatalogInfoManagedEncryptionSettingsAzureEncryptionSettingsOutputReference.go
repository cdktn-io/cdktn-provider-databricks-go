// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabrickscatalog/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference interface {
	cdktn.ComplexObject
	AzureCmkAccessConnectorId() *string
	SetAzureCmkAccessConnectorId(val *string)
	AzureCmkAccessConnectorIdInput() *string
	AzureCmkManagedIdentityId() *string
	SetAzureCmkManagedIdentityId(val *string)
	AzureCmkManagedIdentityIdInput() *string
	AzureTenantId() *string
	SetAzureTenantId(val *string)
	AzureTenantIdInput() *string
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
	InternalValue() *DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettings
	SetInternalValue(val *DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettings)
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
	ResetAzureCmkAccessConnectorId()
	ResetAzureCmkManagedIdentityId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference
type jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) AzureCmkAccessConnectorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureCmkAccessConnectorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) AzureCmkAccessConnectorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureCmkAccessConnectorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) AzureCmkManagedIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureCmkManagedIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) AzureCmkManagedIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureCmkManagedIdentityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) AzureTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) AzureTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) InternalValue() *DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettings {
	var returns *DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksCatalog.DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference_Override(d DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksCatalog.DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference)SetAzureCmkAccessConnectorId(val *string) {
	if err := j.validateSetAzureCmkAccessConnectorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureCmkAccessConnectorId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference)SetAzureCmkManagedIdentityId(val *string) {
	if err := j.validateSetAzureCmkManagedIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureCmkManagedIdentityId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference)SetAzureTenantId(val *string) {
	if err := j.validateSetAzureTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureTenantId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference)SetInternalValue(val *DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) ResetAzureCmkAccessConnectorId() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureCmkAccessConnectorId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) ResetAzureCmkManagedIdentityId() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureCmkManagedIdentityId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

