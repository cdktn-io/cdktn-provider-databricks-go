// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspacesettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabricksworkspacesettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference interface {
	cdktn.ComplexObject
	AxisLineColor() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeAxisLineColorOutputReference
	AxisLineColorInput() interface{}
	CanvasBackgroundColor() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeCanvasBackgroundColorOutputReference
	CanvasBackgroundColorInput() interface{}
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
	ContinuousGradient() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference
	ContinuousGradientInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FontColor() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontColorOutputReference
	FontColorInput() interface{}
	FontFamily() *string
	SetFontFamily(val *string)
	FontFamilyInput() *string
	FontSettings() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference
	FontSettingsInput() interface{}
	// Experimental.
	Fqn() *string
	GridLineColor() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeGridLineColorOutputReference
	GridLineColorInput() interface{}
	InternalValue() *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardTheme
	SetInternalValue(val *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardTheme)
	MarkCornerRadius() *float64
	SetMarkCornerRadius(val *float64)
	MarkCornerRadiusInput() *float64
	SelectionColor() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeSelectionColorOutputReference
	SelectionColorInput() interface{}
	SentimentColors() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference
	SentimentColorsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TextBoxVerticalAlign() *string
	SetTextBoxVerticalAlign(val *string)
	TextBoxVerticalAlignInput() *string
	VisualizationColor01() *string
	SetVisualizationColor01(val *string)
	VisualizationColor01Input() *string
	VisualizationColor02() *string
	SetVisualizationColor02(val *string)
	VisualizationColor02Input() *string
	VisualizationColor03() *string
	SetVisualizationColor03(val *string)
	VisualizationColor03Input() *string
	VisualizationColor04() *string
	SetVisualizationColor04(val *string)
	VisualizationColor04Input() *string
	VisualizationColor05() *string
	SetVisualizationColor05(val *string)
	VisualizationColor05Input() *string
	VisualizationColor06() *string
	SetVisualizationColor06(val *string)
	VisualizationColor06Input() *string
	VisualizationColor07() *string
	SetVisualizationColor07(val *string)
	VisualizationColor07Input() *string
	VisualizationColor08() *string
	SetVisualizationColor08(val *string)
	VisualizationColor08Input() *string
	VisualizationColor09() *string
	SetVisualizationColor09(val *string)
	VisualizationColor09Input() *string
	VisualizationColor10() *string
	SetVisualizationColor10(val *string)
	VisualizationColor10Input() *string
	VisualizationColor11() *string
	SetVisualizationColor11(val *string)
	VisualizationColor11Input() *string
	VisualizationColor12() *string
	SetVisualizationColor12(val *string)
	VisualizationColor12Input() *string
	VisualizationColor13() *string
	SetVisualizationColor13(val *string)
	VisualizationColor13Input() *string
	VisualizationColor14() *string
	SetVisualizationColor14(val *string)
	VisualizationColor14Input() *string
	VisualizationColor15() *string
	SetVisualizationColor15(val *string)
	VisualizationColor15Input() *string
	VisualizationColor16() *string
	SetVisualizationColor16(val *string)
	VisualizationColor16Input() *string
	VisualizationColor17() *string
	SetVisualizationColor17(val *string)
	VisualizationColor17Input() *string
	VisualizationColor18() *string
	SetVisualizationColor18(val *string)
	VisualizationColor18Input() *string
	VisualizationColor19() *string
	SetVisualizationColor19(val *string)
	VisualizationColor19Input() *string
	VisualizationColor20() *string
	SetVisualizationColor20(val *string)
	VisualizationColor20Input() *string
	WidgetBackgroundColor() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBackgroundColorOutputReference
	WidgetBackgroundColorInput() interface{}
	WidgetBorderColor() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBorderColorOutputReference
	WidgetBorderColorInput() interface{}
	WidgetCornerRadius() *float64
	SetWidgetCornerRadius(val *float64)
	WidgetCornerRadiusInput() *float64
	WidgetHeaderAlignment() *string
	SetWidgetHeaderAlignment(val *string)
	WidgetHeaderAlignmentInput() *string
	WidgetMargin() *float64
	SetWidgetMargin(val *float64)
	WidgetMarginInput() *float64
	WidgetPadding() *float64
	SetWidgetPadding(val *float64)
	WidgetPaddingInput() *float64
	WidgetShadow() *float64
	SetWidgetShadow(val *float64)
	WidgetShadowInput() *float64
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
	PutAxisLineColor(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeAxisLineColor)
	PutCanvasBackgroundColor(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeCanvasBackgroundColor)
	PutContinuousGradient(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradient)
	PutFontColor(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontColor)
	PutFontSettings(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettings)
	PutGridLineColor(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeGridLineColor)
	PutSelectionColor(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeSelectionColor)
	PutSentimentColors(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeSentimentColors)
	PutWidgetBackgroundColor(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBackgroundColor)
	PutWidgetBorderColor(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBorderColor)
	ResetAxisLineColor()
	ResetCanvasBackgroundColor()
	ResetContinuousGradient()
	ResetFontColor()
	ResetFontFamily()
	ResetFontSettings()
	ResetGridLineColor()
	ResetMarkCornerRadius()
	ResetSelectionColor()
	ResetSentimentColors()
	ResetTextBoxVerticalAlign()
	ResetVisualizationColor01()
	ResetVisualizationColor02()
	ResetVisualizationColor03()
	ResetVisualizationColor04()
	ResetVisualizationColor05()
	ResetVisualizationColor06()
	ResetVisualizationColor07()
	ResetVisualizationColor08()
	ResetVisualizationColor09()
	ResetVisualizationColor10()
	ResetVisualizationColor11()
	ResetVisualizationColor12()
	ResetVisualizationColor13()
	ResetVisualizationColor14()
	ResetVisualizationColor15()
	ResetVisualizationColor16()
	ResetVisualizationColor17()
	ResetVisualizationColor18()
	ResetVisualizationColor19()
	ResetVisualizationColor20()
	ResetWidgetBackgroundColor()
	ResetWidgetBorderColor()
	ResetWidgetCornerRadius()
	ResetWidgetHeaderAlignment()
	ResetWidgetMargin()
	ResetWidgetPadding()
	ResetWidgetShadow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference
type jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) AxisLineColor() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeAxisLineColorOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeAxisLineColorOutputReference
	_jsii_.Get(
		j,
		"axisLineColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) AxisLineColorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"axisLineColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) CanvasBackgroundColor() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeCanvasBackgroundColorOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeCanvasBackgroundColorOutputReference
	_jsii_.Get(
		j,
		"canvasBackgroundColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) CanvasBackgroundColorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canvasBackgroundColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ContinuousGradient() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference
	_jsii_.Get(
		j,
		"continuousGradient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ContinuousGradientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continuousGradientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) FontColor() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontColorOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontColorOutputReference
	_jsii_.Get(
		j,
		"fontColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) FontColorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fontColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) FontFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) FontFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) FontSettings() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference
	_jsii_.Get(
		j,
		"fontSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) FontSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fontSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GridLineColor() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeGridLineColorOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeGridLineColorOutputReference
	_jsii_.Get(
		j,
		"gridLineColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GridLineColorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gridLineColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) InternalValue() *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardTheme {
	var returns *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardTheme
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) MarkCornerRadius() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"markCornerRadius",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) MarkCornerRadiusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"markCornerRadiusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) SelectionColor() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeSelectionColorOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeSelectionColorOutputReference
	_jsii_.Get(
		j,
		"selectionColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) SelectionColorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selectionColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) SentimentColors() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference
	_jsii_.Get(
		j,
		"sentimentColors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) SentimentColorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sentimentColorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) TextBoxVerticalAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textBoxVerticalAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) TextBoxVerticalAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textBoxVerticalAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor01() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor01",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor01Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor01Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor02() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor02",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor02Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor02Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor03() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor03",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor03Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor03Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor04() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor04",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor04Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor04Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor05() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor05",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor05Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor05Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor06() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor06",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor06Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor06Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor07() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor07",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor07Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor07Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor08() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor08",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor08Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor08Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor09() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor09",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor09Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor09Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor10() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor10",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor10Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor10Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor11() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor11",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor11Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor11Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor12() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor12",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor12Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor12Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor13() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor13",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor13Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor13Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor14() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor14",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor14Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor14Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor15() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor15",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor15Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor15Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor16() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor16",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor16Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor16Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor17() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor17",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor17Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor17Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor18() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor18",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor18Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor18Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor19() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor19",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor19Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor19Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor20() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor20",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor20Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor20Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetBackgroundColor() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBackgroundColorOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBackgroundColorOutputReference
	_jsii_.Get(
		j,
		"widgetBackgroundColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetBackgroundColorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetBackgroundColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetBorderColor() DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBorderColorOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBorderColorOutputReference
	_jsii_.Get(
		j,
		"widgetBorderColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetBorderColorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetBorderColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetCornerRadius() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widgetCornerRadius",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetCornerRadiusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widgetCornerRadiusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetHeaderAlignment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widgetHeaderAlignment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetHeaderAlignmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widgetHeaderAlignmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetMargin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widgetMargin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetMarginInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widgetMarginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetPadding() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widgetPadding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetPaddingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widgetPaddingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetShadow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widgetShadow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetShadowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widgetShadowInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksWorkspaceSettingV2.DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference_Override(d DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksWorkspaceSettingV2.DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetFontFamily(val *string) {
	if err := j.validateSetFontFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontFamily",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetInternalValue(val *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardTheme) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetMarkCornerRadius(val *float64) {
	if err := j.validateSetMarkCornerRadiusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"markCornerRadius",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetTextBoxVerticalAlign(val *string) {
	if err := j.validateSetTextBoxVerticalAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textBoxVerticalAlign",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor01(val *string) {
	if err := j.validateSetVisualizationColor01Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor01",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor02(val *string) {
	if err := j.validateSetVisualizationColor02Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor02",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor03(val *string) {
	if err := j.validateSetVisualizationColor03Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor03",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor04(val *string) {
	if err := j.validateSetVisualizationColor04Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor04",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor05(val *string) {
	if err := j.validateSetVisualizationColor05Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor05",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor06(val *string) {
	if err := j.validateSetVisualizationColor06Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor06",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor07(val *string) {
	if err := j.validateSetVisualizationColor07Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor07",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor08(val *string) {
	if err := j.validateSetVisualizationColor08Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor08",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor09(val *string) {
	if err := j.validateSetVisualizationColor09Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor09",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor10(val *string) {
	if err := j.validateSetVisualizationColor10Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor10",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor11(val *string) {
	if err := j.validateSetVisualizationColor11Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor11",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor12(val *string) {
	if err := j.validateSetVisualizationColor12Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor12",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor13(val *string) {
	if err := j.validateSetVisualizationColor13Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor13",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor14(val *string) {
	if err := j.validateSetVisualizationColor14Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor14",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor15(val *string) {
	if err := j.validateSetVisualizationColor15Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor15",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor16(val *string) {
	if err := j.validateSetVisualizationColor16Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor16",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor17(val *string) {
	if err := j.validateSetVisualizationColor17Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor17",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor18(val *string) {
	if err := j.validateSetVisualizationColor18Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor18",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor19(val *string) {
	if err := j.validateSetVisualizationColor19Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor19",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor20(val *string) {
	if err := j.validateSetVisualizationColor20Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor20",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetWidgetCornerRadius(val *float64) {
	if err := j.validateSetWidgetCornerRadiusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"widgetCornerRadius",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetWidgetHeaderAlignment(val *string) {
	if err := j.validateSetWidgetHeaderAlignmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"widgetHeaderAlignment",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetWidgetMargin(val *float64) {
	if err := j.validateSetWidgetMarginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"widgetMargin",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetWidgetPadding(val *float64) {
	if err := j.validateSetWidgetPaddingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"widgetPadding",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetWidgetShadow(val *float64) {
	if err := j.validateSetWidgetShadowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"widgetShadow",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutAxisLineColor(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeAxisLineColor) {
	if err := d.validatePutAxisLineColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAxisLineColor",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutCanvasBackgroundColor(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeCanvasBackgroundColor) {
	if err := d.validatePutCanvasBackgroundColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCanvasBackgroundColor",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutContinuousGradient(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradient) {
	if err := d.validatePutContinuousGradientParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContinuousGradient",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutFontColor(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontColor) {
	if err := d.validatePutFontColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFontColor",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutFontSettings(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettings) {
	if err := d.validatePutFontSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFontSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutGridLineColor(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeGridLineColor) {
	if err := d.validatePutGridLineColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGridLineColor",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutSelectionColor(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeSelectionColor) {
	if err := d.validatePutSelectionColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSelectionColor",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutSentimentColors(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeSentimentColors) {
	if err := d.validatePutSentimentColorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSentimentColors",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutWidgetBackgroundColor(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBackgroundColor) {
	if err := d.validatePutWidgetBackgroundColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWidgetBackgroundColor",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutWidgetBorderColor(value *DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBorderColor) {
	if err := d.validatePutWidgetBorderColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWidgetBorderColor",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetAxisLineColor() {
	_jsii_.InvokeVoid(
		d,
		"resetAxisLineColor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetCanvasBackgroundColor() {
	_jsii_.InvokeVoid(
		d,
		"resetCanvasBackgroundColor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetContinuousGradient() {
	_jsii_.InvokeVoid(
		d,
		"resetContinuousGradient",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetFontColor() {
	_jsii_.InvokeVoid(
		d,
		"resetFontColor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetFontFamily() {
	_jsii_.InvokeVoid(
		d,
		"resetFontFamily",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetFontSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetFontSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetGridLineColor() {
	_jsii_.InvokeVoid(
		d,
		"resetGridLineColor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetMarkCornerRadius() {
	_jsii_.InvokeVoid(
		d,
		"resetMarkCornerRadius",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetSelectionColor() {
	_jsii_.InvokeVoid(
		d,
		"resetSelectionColor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetSentimentColors() {
	_jsii_.InvokeVoid(
		d,
		"resetSentimentColors",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetTextBoxVerticalAlign() {
	_jsii_.InvokeVoid(
		d,
		"resetTextBoxVerticalAlign",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor01() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor01",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor02() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor02",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor03() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor03",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor04() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor04",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor05() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor05",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor06() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor06",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor07() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor07",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor08() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor08",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor09() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor09",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor10() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor10",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor11() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor11",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor12() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor12",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor13() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor13",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor14() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor14",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor15() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor15",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor16() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor16",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor17() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor17",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor18() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor18",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor19() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor19",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor20() {
	_jsii_.InvokeVoid(
		d,
		"resetVisualizationColor20",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetWidgetBackgroundColor() {
	_jsii_.InvokeVoid(
		d,
		"resetWidgetBackgroundColor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetWidgetBorderColor() {
	_jsii_.InvokeVoid(
		d,
		"resetWidgetBorderColor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetWidgetCornerRadius() {
	_jsii_.InvokeVoid(
		d,
		"resetWidgetCornerRadius",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetWidgetHeaderAlignment() {
	_jsii_.InvokeVoid(
		d,
		"resetWidgetHeaderAlignment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetWidgetMargin() {
	_jsii_.InvokeVoid(
		d,
		"resetWidgetMargin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetWidgetPadding() {
	_jsii_.InvokeVoid(
		d,
		"resetWidgetPadding",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetWidgetShadow() {
	_jsii_.InvokeVoid(
		d,
		"resetWidgetShadow",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

