// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabrickslakewatchrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksLakewatchRuleMitreOutputReference interface {
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
	InternalValue() *DataDatabricksLakewatchRuleMitre
	SetInternalValue(val *DataDatabricksLakewatchRuleMitre)
	SubTechnique() *string
	SetSubTechnique(val *string)
	SubTechniqueId() *string
	SetSubTechniqueId(val *string)
	SubTechniqueIdInput() *string
	SubTechniqueInput() *string
	Tactic() *string
	SetTactic(val *string)
	TacticInput() *string
	Taxonomy() *string
	SetTaxonomy(val *string)
	TaxonomyInput() *string
	Technique() *string
	SetTechnique(val *string)
	TechniqueId() *string
	SetTechniqueId(val *string)
	TechniqueIdInput() *string
	TechniqueInput() *string
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
	ResetSubTechnique()
	ResetSubTechniqueId()
	ResetTactic()
	ResetTaxonomy()
	ResetTechnique()
	ResetTechniqueId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksLakewatchRuleMitreOutputReference
type jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) InternalValue() *DataDatabricksLakewatchRuleMitre {
	var returns *DataDatabricksLakewatchRuleMitre
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) SubTechnique() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subTechnique",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) SubTechniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subTechniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) SubTechniqueIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subTechniqueIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) SubTechniqueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subTechniqueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) Tactic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tactic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) TacticInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tacticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) Taxonomy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taxonomy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) TaxonomyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taxonomyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) Technique() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technique",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) TechniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"techniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) TechniqueIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"techniqueIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) TechniqueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"techniqueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksLakewatchRuleMitreOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksLakewatchRuleMitreOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksLakewatchRuleMitreOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchRule.DataDatabricksLakewatchRuleMitreOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksLakewatchRuleMitreOutputReference_Override(d DataDatabricksLakewatchRuleMitreOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchRule.DataDatabricksLakewatchRuleMitreOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference)SetInternalValue(val *DataDatabricksLakewatchRuleMitre) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference)SetSubTechnique(val *string) {
	if err := j.validateSetSubTechniqueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subTechnique",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference)SetSubTechniqueId(val *string) {
	if err := j.validateSetSubTechniqueIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subTechniqueId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference)SetTactic(val *string) {
	if err := j.validateSetTacticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tactic",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference)SetTaxonomy(val *string) {
	if err := j.validateSetTaxonomyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taxonomy",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference)SetTechnique(val *string) {
	if err := j.validateSetTechniqueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"technique",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference)SetTechniqueId(val *string) {
	if err := j.validateSetTechniqueIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"techniqueId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) ResetSubTechnique() {
	_jsii_.InvokeVoid(
		d,
		"resetSubTechnique",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) ResetSubTechniqueId() {
	_jsii_.InvokeVoid(
		d,
		"resetSubTechniqueId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) ResetTactic() {
	_jsii_.InvokeVoid(
		d,
		"resetTactic",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) ResetTaxonomy() {
	_jsii_.InvokeVoid(
		d,
		"resetTaxonomy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) ResetTechnique() {
	_jsii_.InvokeVoid(
		d,
		"resetTechnique",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) ResetTechniqueId() {
	_jsii_.InvokeVoid(
		d,
		"resetTechniqueId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksLakewatchRuleMitreOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

