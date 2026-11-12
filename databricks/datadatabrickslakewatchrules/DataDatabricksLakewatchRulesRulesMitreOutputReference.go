// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabrickslakewatchrules/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksLakewatchRulesRulesMitreOutputReference interface {
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
	InternalValue() *DataDatabricksLakewatchRulesRulesMitre
	SetInternalValue(val *DataDatabricksLakewatchRulesRulesMitre)
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

// The jsii proxy struct for DataDatabricksLakewatchRulesRulesMitreOutputReference
type jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) InternalValue() *DataDatabricksLakewatchRulesRulesMitre {
	var returns *DataDatabricksLakewatchRulesRulesMitre
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) SubTechnique() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subTechnique",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) SubTechniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subTechniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) SubTechniqueIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subTechniqueIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) SubTechniqueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subTechniqueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) Tactic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tactic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) TacticInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tacticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) Taxonomy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taxonomy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) TaxonomyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taxonomyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) Technique() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technique",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) TechniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"techniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) TechniqueIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"techniqueIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) TechniqueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"techniqueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksLakewatchRulesRulesMitreOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksLakewatchRulesRulesMitreOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksLakewatchRulesRulesMitreOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchRules.DataDatabricksLakewatchRulesRulesMitreOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksLakewatchRulesRulesMitreOutputReference_Override(d DataDatabricksLakewatchRulesRulesMitreOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchRules.DataDatabricksLakewatchRulesRulesMitreOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference)SetInternalValue(val *DataDatabricksLakewatchRulesRulesMitre) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference)SetSubTechnique(val *string) {
	if err := j.validateSetSubTechniqueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subTechnique",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference)SetSubTechniqueId(val *string) {
	if err := j.validateSetSubTechniqueIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subTechniqueId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference)SetTactic(val *string) {
	if err := j.validateSetTacticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tactic",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference)SetTaxonomy(val *string) {
	if err := j.validateSetTaxonomyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taxonomy",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference)SetTechnique(val *string) {
	if err := j.validateSetTechniqueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"technique",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference)SetTechniqueId(val *string) {
	if err := j.validateSetTechniqueIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"techniqueId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) ResetSubTechnique() {
	_jsii_.InvokeVoid(
		d,
		"resetSubTechnique",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) ResetSubTechniqueId() {
	_jsii_.InvokeVoid(
		d,
		"resetSubTechniqueId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) ResetTactic() {
	_jsii_.InvokeVoid(
		d,
		"resetTactic",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) ResetTaxonomy() {
	_jsii_.InvokeVoid(
		d,
		"resetTaxonomy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) ResetTechnique() {
	_jsii_.InvokeVoid(
		d,
		"resetTechnique",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) ResetTechniqueId() {
	_jsii_.InvokeVoid(
		d,
		"resetTechniqueId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksLakewatchRulesRulesMitreOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

