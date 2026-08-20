// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/pipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference interface {
	cdktn.ComplexObject
	ApiSourceConnectorOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsApiSourceConnectorOptionsOutputReference
	ApiSourceConnectorOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsApiSourceConnectorOptions
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
	ConfluenceOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsConfluenceOptionsOutputReference
	ConfluenceOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsConfluenceOptions
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GdriveOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsGdriveOptionsOutputReference
	GdriveOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsGdriveOptions
	GoogleAdsOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsOutputReference
	GoogleAdsOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptions
	InternalValue() *PipelineIngestionDefinitionObjectsTableConnectorOptions
	SetInternalValue(val *PipelineIngestionDefinitionObjectsTableConnectorOptions)
	JiraOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsJiraOptionsOutputReference
	JiraOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsJiraOptions
	KafkaOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsOutputReference
	KafkaOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptions
	LinkedinAdsOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsLinkedinAdsOptionsOutputReference
	LinkedinAdsOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsLinkedinAdsOptions
	MarketoOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsMarketoOptionsOutputReference
	MarketoOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsMarketoOptions
	MetaAdsOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference
	MetaAdsOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptions
	OutlookOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsOutlookOptionsOutputReference
	OutlookOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsOutlookOptions
	RedditAdsOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsRedditAdsOptionsOutputReference
	RedditAdsOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsRedditAdsOptions
	SharepointOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsSharepointOptionsOutputReference
	SharepointOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsSharepointOptions
	SmartsheetOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsSmartsheetOptionsOutputReference
	SmartsheetOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsSmartsheetOptions
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TiktokAdsOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference
	TiktokAdsOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptions
	ZendeskSupportOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsZendeskSupportOptionsOutputReference
	ZendeskSupportOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsZendeskSupportOptions
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
	PutApiSourceConnectorOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsApiSourceConnectorOptions)
	PutConfluenceOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsConfluenceOptions)
	PutGdriveOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsGdriveOptions)
	PutGoogleAdsOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptions)
	PutJiraOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsJiraOptions)
	PutKafkaOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptions)
	PutLinkedinAdsOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsLinkedinAdsOptions)
	PutMarketoOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsMarketoOptions)
	PutMetaAdsOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptions)
	PutOutlookOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsOutlookOptions)
	PutRedditAdsOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsRedditAdsOptions)
	PutSharepointOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsSharepointOptions)
	PutSmartsheetOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsSmartsheetOptions)
	PutTiktokAdsOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptions)
	PutZendeskSupportOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsZendeskSupportOptions)
	ResetApiSourceConnectorOptions()
	ResetConfluenceOptions()
	ResetGdriveOptions()
	ResetGoogleAdsOptions()
	ResetJiraOptions()
	ResetKafkaOptions()
	ResetLinkedinAdsOptions()
	ResetMarketoOptions()
	ResetMetaAdsOptions()
	ResetOutlookOptions()
	ResetRedditAdsOptions()
	ResetSharepointOptions()
	ResetSmartsheetOptions()
	ResetTiktokAdsOptions()
	ResetZendeskSupportOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference
type jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ApiSourceConnectorOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsApiSourceConnectorOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableConnectorOptionsApiSourceConnectorOptionsOutputReference
	_jsii_.Get(
		j,
		"apiSourceConnectorOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ApiSourceConnectorOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsApiSourceConnectorOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsApiSourceConnectorOptions
	_jsii_.Get(
		j,
		"apiSourceConnectorOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ConfluenceOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsConfluenceOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableConnectorOptionsConfluenceOptionsOutputReference
	_jsii_.Get(
		j,
		"confluenceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ConfluenceOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsConfluenceOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsConfluenceOptions
	_jsii_.Get(
		j,
		"confluenceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) GdriveOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsGdriveOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableConnectorOptionsGdriveOptionsOutputReference
	_jsii_.Get(
		j,
		"gdriveOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) GdriveOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsGdriveOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsGdriveOptions
	_jsii_.Get(
		j,
		"gdriveOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) GoogleAdsOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsOutputReference
	_jsii_.Get(
		j,
		"googleAdsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) GoogleAdsOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptions
	_jsii_.Get(
		j,
		"googleAdsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) InternalValue() *PipelineIngestionDefinitionObjectsTableConnectorOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) JiraOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsJiraOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableConnectorOptionsJiraOptionsOutputReference
	_jsii_.Get(
		j,
		"jiraOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) JiraOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsJiraOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsJiraOptions
	_jsii_.Get(
		j,
		"jiraOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) KafkaOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsOutputReference
	_jsii_.Get(
		j,
		"kafkaOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) KafkaOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptions
	_jsii_.Get(
		j,
		"kafkaOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) LinkedinAdsOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsLinkedinAdsOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableConnectorOptionsLinkedinAdsOptionsOutputReference
	_jsii_.Get(
		j,
		"linkedinAdsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) LinkedinAdsOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsLinkedinAdsOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsLinkedinAdsOptions
	_jsii_.Get(
		j,
		"linkedinAdsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) MarketoOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsMarketoOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableConnectorOptionsMarketoOptionsOutputReference
	_jsii_.Get(
		j,
		"marketoOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) MarketoOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsMarketoOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsMarketoOptions
	_jsii_.Get(
		j,
		"marketoOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) MetaAdsOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsOutputReference
	_jsii_.Get(
		j,
		"metaAdsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) MetaAdsOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptions
	_jsii_.Get(
		j,
		"metaAdsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) OutlookOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsOutlookOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableConnectorOptionsOutlookOptionsOutputReference
	_jsii_.Get(
		j,
		"outlookOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) OutlookOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsOutlookOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsOutlookOptions
	_jsii_.Get(
		j,
		"outlookOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) RedditAdsOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsRedditAdsOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableConnectorOptionsRedditAdsOptionsOutputReference
	_jsii_.Get(
		j,
		"redditAdsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) RedditAdsOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsRedditAdsOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsRedditAdsOptions
	_jsii_.Get(
		j,
		"redditAdsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) SharepointOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsSharepointOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableConnectorOptionsSharepointOptionsOutputReference
	_jsii_.Get(
		j,
		"sharepointOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) SharepointOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsSharepointOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsSharepointOptions
	_jsii_.Get(
		j,
		"sharepointOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) SmartsheetOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsSmartsheetOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableConnectorOptionsSmartsheetOptionsOutputReference
	_jsii_.Get(
		j,
		"smartsheetOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) SmartsheetOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsSmartsheetOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsSmartsheetOptions
	_jsii_.Get(
		j,
		"smartsheetOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) TiktokAdsOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsOutputReference
	_jsii_.Get(
		j,
		"tiktokAdsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) TiktokAdsOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptions
	_jsii_.Get(
		j,
		"tiktokAdsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ZendeskSupportOptions() PipelineIngestionDefinitionObjectsTableConnectorOptionsZendeskSupportOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsTableConnectorOptionsZendeskSupportOptionsOutputReference
	_jsii_.Get(
		j,
		"zendeskSupportOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ZendeskSupportOptionsInput() *PipelineIngestionDefinitionObjectsTableConnectorOptionsZendeskSupportOptions {
	var returns *PipelineIngestionDefinitionObjectsTableConnectorOptionsZendeskSupportOptions
	_jsii_.Get(
		j,
		"zendeskSupportOptionsInput",
		&returns,
	)
	return returns
}


func NewPipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference_Override(p PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference)SetInternalValue(val *PipelineIngestionDefinitionObjectsTableConnectorOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) PutApiSourceConnectorOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsApiSourceConnectorOptions) {
	if err := p.validatePutApiSourceConnectorOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApiSourceConnectorOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) PutConfluenceOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsConfluenceOptions) {
	if err := p.validatePutConfluenceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putConfluenceOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) PutGdriveOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsGdriveOptions) {
	if err := p.validatePutGdriveOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGdriveOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) PutGoogleAdsOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptions) {
	if err := p.validatePutGoogleAdsOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGoogleAdsOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) PutJiraOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsJiraOptions) {
	if err := p.validatePutJiraOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putJiraOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) PutKafkaOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptions) {
	if err := p.validatePutKafkaOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putKafkaOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) PutLinkedinAdsOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsLinkedinAdsOptions) {
	if err := p.validatePutLinkedinAdsOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLinkedinAdsOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) PutMarketoOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsMarketoOptions) {
	if err := p.validatePutMarketoOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMarketoOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) PutMetaAdsOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptions) {
	if err := p.validatePutMetaAdsOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMetaAdsOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) PutOutlookOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsOutlookOptions) {
	if err := p.validatePutOutlookOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putOutlookOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) PutRedditAdsOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsRedditAdsOptions) {
	if err := p.validatePutRedditAdsOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRedditAdsOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) PutSharepointOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsSharepointOptions) {
	if err := p.validatePutSharepointOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSharepointOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) PutSmartsheetOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsSmartsheetOptions) {
	if err := p.validatePutSmartsheetOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSmartsheetOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) PutTiktokAdsOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptions) {
	if err := p.validatePutTiktokAdsOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTiktokAdsOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) PutZendeskSupportOptions(value *PipelineIngestionDefinitionObjectsTableConnectorOptionsZendeskSupportOptions) {
	if err := p.validatePutZendeskSupportOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putZendeskSupportOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ResetApiSourceConnectorOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetApiSourceConnectorOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ResetConfluenceOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetConfluenceOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ResetGdriveOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetGdriveOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ResetGoogleAdsOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetGoogleAdsOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ResetJiraOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetJiraOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ResetKafkaOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetKafkaOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ResetLinkedinAdsOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetLinkedinAdsOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ResetMarketoOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetMarketoOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ResetMetaAdsOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetMetaAdsOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ResetOutlookOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetOutlookOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ResetRedditAdsOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetRedditAdsOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ResetSharepointOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetSharepointOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ResetSmartsheetOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetSmartsheetOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ResetTiktokAdsOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetTiktokAdsOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ResetZendeskSupportOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetZendeskSupportOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsTableConnectorOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

