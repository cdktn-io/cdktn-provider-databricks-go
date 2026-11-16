// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2


type WorkspaceSettingV2WorkspaceDashboardTheme struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#axis_line_color WorkspaceSettingV2#axis_line_color}.
	AxisLineColor *WorkspaceSettingV2WorkspaceDashboardThemeAxisLineColor `field:"optional" json:"axisLineColor" yaml:"axisLineColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#canvas_background_color WorkspaceSettingV2#canvas_background_color}.
	CanvasBackgroundColor *WorkspaceSettingV2WorkspaceDashboardThemeCanvasBackgroundColor `field:"optional" json:"canvasBackgroundColor" yaml:"canvasBackgroundColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#continuous_gradient WorkspaceSettingV2#continuous_gradient}.
	ContinuousGradient *WorkspaceSettingV2WorkspaceDashboardThemeContinuousGradient `field:"optional" json:"continuousGradient" yaml:"continuousGradient"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#font_color WorkspaceSettingV2#font_color}.
	FontColor *WorkspaceSettingV2WorkspaceDashboardThemeFontColor `field:"optional" json:"fontColor" yaml:"fontColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#font_family WorkspaceSettingV2#font_family}.
	FontFamily *string `field:"optional" json:"fontFamily" yaml:"fontFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#font_settings WorkspaceSettingV2#font_settings}.
	FontSettings *WorkspaceSettingV2WorkspaceDashboardThemeFontSettings `field:"optional" json:"fontSettings" yaml:"fontSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#grid_line_color WorkspaceSettingV2#grid_line_color}.
	GridLineColor *WorkspaceSettingV2WorkspaceDashboardThemeGridLineColor `field:"optional" json:"gridLineColor" yaml:"gridLineColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#mark_corner_radius WorkspaceSettingV2#mark_corner_radius}.
	MarkCornerRadius *float64 `field:"optional" json:"markCornerRadius" yaml:"markCornerRadius"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#selection_color WorkspaceSettingV2#selection_color}.
	SelectionColor *WorkspaceSettingV2WorkspaceDashboardThemeSelectionColor `field:"optional" json:"selectionColor" yaml:"selectionColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#sentiment_colors WorkspaceSettingV2#sentiment_colors}.
	SentimentColors *WorkspaceSettingV2WorkspaceDashboardThemeSentimentColors `field:"optional" json:"sentimentColors" yaml:"sentimentColors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#text_box_vertical_align WorkspaceSettingV2#text_box_vertical_align}.
	TextBoxVerticalAlign *string `field:"optional" json:"textBoxVerticalAlign" yaml:"textBoxVerticalAlign"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_01 WorkspaceSettingV2#visualization_color_01}.
	VisualizationColor01 *string `field:"optional" json:"visualizationColor01" yaml:"visualizationColor01"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_02 WorkspaceSettingV2#visualization_color_02}.
	VisualizationColor02 *string `field:"optional" json:"visualizationColor02" yaml:"visualizationColor02"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_03 WorkspaceSettingV2#visualization_color_03}.
	VisualizationColor03 *string `field:"optional" json:"visualizationColor03" yaml:"visualizationColor03"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_04 WorkspaceSettingV2#visualization_color_04}.
	VisualizationColor04 *string `field:"optional" json:"visualizationColor04" yaml:"visualizationColor04"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_05 WorkspaceSettingV2#visualization_color_05}.
	VisualizationColor05 *string `field:"optional" json:"visualizationColor05" yaml:"visualizationColor05"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_06 WorkspaceSettingV2#visualization_color_06}.
	VisualizationColor06 *string `field:"optional" json:"visualizationColor06" yaml:"visualizationColor06"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_07 WorkspaceSettingV2#visualization_color_07}.
	VisualizationColor07 *string `field:"optional" json:"visualizationColor07" yaml:"visualizationColor07"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_08 WorkspaceSettingV2#visualization_color_08}.
	VisualizationColor08 *string `field:"optional" json:"visualizationColor08" yaml:"visualizationColor08"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_09 WorkspaceSettingV2#visualization_color_09}.
	VisualizationColor09 *string `field:"optional" json:"visualizationColor09" yaml:"visualizationColor09"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_10 WorkspaceSettingV2#visualization_color_10}.
	VisualizationColor10 *string `field:"optional" json:"visualizationColor10" yaml:"visualizationColor10"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_11 WorkspaceSettingV2#visualization_color_11}.
	VisualizationColor11 *string `field:"optional" json:"visualizationColor11" yaml:"visualizationColor11"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_12 WorkspaceSettingV2#visualization_color_12}.
	VisualizationColor12 *string `field:"optional" json:"visualizationColor12" yaml:"visualizationColor12"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_13 WorkspaceSettingV2#visualization_color_13}.
	VisualizationColor13 *string `field:"optional" json:"visualizationColor13" yaml:"visualizationColor13"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_14 WorkspaceSettingV2#visualization_color_14}.
	VisualizationColor14 *string `field:"optional" json:"visualizationColor14" yaml:"visualizationColor14"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_15 WorkspaceSettingV2#visualization_color_15}.
	VisualizationColor15 *string `field:"optional" json:"visualizationColor15" yaml:"visualizationColor15"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_16 WorkspaceSettingV2#visualization_color_16}.
	VisualizationColor16 *string `field:"optional" json:"visualizationColor16" yaml:"visualizationColor16"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_17 WorkspaceSettingV2#visualization_color_17}.
	VisualizationColor17 *string `field:"optional" json:"visualizationColor17" yaml:"visualizationColor17"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_18 WorkspaceSettingV2#visualization_color_18}.
	VisualizationColor18 *string `field:"optional" json:"visualizationColor18" yaml:"visualizationColor18"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_19 WorkspaceSettingV2#visualization_color_19}.
	VisualizationColor19 *string `field:"optional" json:"visualizationColor19" yaml:"visualizationColor19"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#visualization_color_20 WorkspaceSettingV2#visualization_color_20}.
	VisualizationColor20 *string `field:"optional" json:"visualizationColor20" yaml:"visualizationColor20"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#widget_background_color WorkspaceSettingV2#widget_background_color}.
	WidgetBackgroundColor *WorkspaceSettingV2WorkspaceDashboardThemeWidgetBackgroundColor `field:"optional" json:"widgetBackgroundColor" yaml:"widgetBackgroundColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#widget_border_color WorkspaceSettingV2#widget_border_color}.
	WidgetBorderColor *WorkspaceSettingV2WorkspaceDashboardThemeWidgetBorderColor `field:"optional" json:"widgetBorderColor" yaml:"widgetBorderColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#widget_corner_radius WorkspaceSettingV2#widget_corner_radius}.
	WidgetCornerRadius *float64 `field:"optional" json:"widgetCornerRadius" yaml:"widgetCornerRadius"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#widget_header_alignment WorkspaceSettingV2#widget_header_alignment}.
	WidgetHeaderAlignment *string `field:"optional" json:"widgetHeaderAlignment" yaml:"widgetHeaderAlignment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#widget_margin WorkspaceSettingV2#widget_margin}.
	WidgetMargin *float64 `field:"optional" json:"widgetMargin" yaml:"widgetMargin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#widget_padding WorkspaceSettingV2#widget_padding}.
	WidgetPadding *float64 `field:"optional" json:"widgetPadding" yaml:"widgetPadding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#widget_shadow WorkspaceSettingV2#widget_shadow}.
	WidgetShadow *float64 `field:"optional" json:"widgetShadow" yaml:"widgetShadow"`
}

