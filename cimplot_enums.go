// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package cimgui

// original name: ImAxis_
type ImAxis int

const (
	ImAxis_X1    = 0
	ImAxis_X2    = 1
	ImAxis_X3    = 2
	ImAxis_Y1    = 3
	ImAxis_Y2    = 4
	ImAxis_Y3    = 5
	ImAxis_COUNT = 6
)

// original name: ImPlotAxisFlags_
type PlotAxisFlags int

const (
	PlotAxisFlags_None          = 0
	PlotAxisFlags_NoLabel       = 1
	PlotAxisFlags_NoGridLines   = 2
	PlotAxisFlags_NoTickMarks   = 4
	PlotAxisFlags_NoTickLabels  = 8
	PlotAxisFlags_NoInitialFit  = 16
	PlotAxisFlags_NoMenus       = 32
	PlotAxisFlags_Opposite      = 64
	PlotAxisFlags_Foreground    = 128
	PlotAxisFlags_Invert        = 256
	PlotAxisFlags_AutoFit       = 512
	PlotAxisFlags_RangeFit      = 1024
	PlotAxisFlags_PanStretch    = 2048
	PlotAxisFlags_LockMin       = 4096
	PlotAxisFlags_LockMax       = 8192
	PlotAxisFlags_Lock          = 12288
	PlotAxisFlags_NoDecorations = 15
	PlotAxisFlags_AuxDefault    = 66
)

// original name: ImPlotBarGroupsFlags_
type PlotBarGroupsFlags int

const (
	PlotBarGroupsFlags_None       = 0
	PlotBarGroupsFlags_Horizontal = 1024
	PlotBarGroupsFlags_Stacked    = 2048
)

// original name: ImPlotBarsFlags_
type PlotBarsFlags int

const (
	PlotBarsFlags_None       = 0
	PlotBarsFlags_Horizontal = 1024
)

// original name: ImPlotBin_
type PlotBin int

const (
	PlotBin_Sqrt    = -1
	PlotBin_Sturges = -2
	PlotBin_Rice    = -3
	PlotBin_Scott   = -4
)

// original name: ImPlotCol_
type PlotCol int

const (
	PlotCol_Line          = 0
	PlotCol_Fill          = 1
	PlotCol_MarkerOutline = 2
	PlotCol_MarkerFill    = 3
	PlotCol_ErrorBar      = 4
	PlotCol_FrameBg       = 5
	PlotCol_PlotBg        = 6
	PlotCol_PlotBorder    = 7
	PlotCol_LegendBg      = 8
	PlotCol_LegendBorder  = 9
	PlotCol_LegendText    = 10
	PlotCol_TitleText     = 11
	PlotCol_InlayText     = 12
	PlotCol_AxisText      = 13
	PlotCol_AxisGrid      = 14
	PlotCol_AxisTick      = 15
	PlotCol_AxisBg        = 16
	PlotCol_AxisBgHovered = 17
	PlotCol_AxisBgActive  = 18
	PlotCol_Selection     = 19
	PlotCol_Crosshairs    = 20
	PlotCol_COUNT         = 21
)

// original name: ImPlotColormapScaleFlags_
type PlotColormapScaleFlags int

const (
	PlotColormapScaleFlags_None     = 0
	PlotColormapScaleFlags_NoLabel  = 1
	PlotColormapScaleFlags_Opposite = 2
	PlotColormapScaleFlags_Invert   = 4
)

// original name: ImPlotColormap_
type PlotColormap int

const (
	PlotColormap_Deep     = 0
	PlotColormap_Dark     = 1
	PlotColormap_Pastel   = 2
	PlotColormap_Paired   = 3
	PlotColormap_Viridis  = 4
	PlotColormap_Plasma   = 5
	PlotColormap_Hot      = 6
	PlotColormap_Cool     = 7
	PlotColormap_Pink     = 8
	PlotColormap_Jet      = 9
	PlotColormap_Twilight = 10
	PlotColormap_RdBu     = 11
	PlotColormap_BrBG     = 12
	PlotColormap_PiYG     = 13
	PlotColormap_Spectral = 14
	PlotColormap_Greys    = 15
)

// original name: ImPlotCond_
type PlotCond int

const (
	PlotCond_None   = 0
	PlotCond_Always = 1
	PlotCond_Once   = 2
)

// original name: ImPlotDateFmt_
type PlotDateFmt int

const (
	PlotDateFmt_None    = 0
	PlotDateFmt_DayMo   = 1
	PlotDateFmt_DayMoYr = 2
	PlotDateFmt_MoYr    = 3
	PlotDateFmt_Mo      = 4
	PlotDateFmt_Yr      = 5
)

// original name: ImPlotDigitalFlags_
type PlotDigitalFlags int

const (
	PlotDigitalFlags_None = 0
)

// original name: ImPlotDragToolFlags_
type PlotDragToolFlags int

const (
	PlotDragToolFlags_None      = 0
	PlotDragToolFlags_NoCursors = 1
	PlotDragToolFlags_NoFit     = 2
	PlotDragToolFlags_NoInputs  = 4
	PlotDragToolFlags_Delayed   = 8
)

// original name: ImPlotDummyFlags_
type PlotDummyFlags int

const (
	PlotDummyFlags_None = 0
)

// original name: ImPlotErrorBarsFlags_
type PlotErrorBarsFlags int

const (
	PlotErrorBarsFlags_None       = 0
	PlotErrorBarsFlags_Horizontal = 1024
)

// original name: ImPlotFlags_
type PlotFlags int

const (
	PlotFlags_None        = 0
	PlotFlags_NoTitle     = 1
	PlotFlags_NoLegend    = 2
	PlotFlags_NoMouseText = 4
	PlotFlags_NoInputs    = 8
	PlotFlags_NoMenus     = 16
	PlotFlags_NoBoxSelect = 32
	PlotFlags_NoChild     = 64
	PlotFlags_NoFrame     = 128
	PlotFlags_Equal       = 256
	PlotFlags_Crosshairs  = 512
	PlotFlags_CanvasOnly  = 55
)

// original name: ImPlotHeatmapFlags_
type PlotHeatmapFlags int

const (
	PlotHeatmapFlags_None     = 0
	PlotHeatmapFlags_ColMajor = 1024
)

// original name: ImPlotHistogramFlags_
type PlotHistogramFlags int

const (
	PlotHistogramFlags_None       = 0
	PlotHistogramFlags_Horizontal = 1024
	PlotHistogramFlags_Cumulative = 2048
	PlotHistogramFlags_Density    = 4096
	PlotHistogramFlags_NoOutliers = 8192
	PlotHistogramFlags_ColMajor   = 16384
)

// original name: ImPlotImageFlags_
type PlotImageFlags int

const (
	PlotImageFlags_None = 0
)

// original name: ImPlotInfLinesFlags_
type PlotInfLinesFlags int

const (
	PlotInfLinesFlags_None       = 0
	PlotInfLinesFlags_Horizontal = 1024
)

// original name: ImPlotItemFlags_
type PlotItemFlags int

const (
	PlotItemFlags_None     = 0
	PlotItemFlags_NoLegend = 1
	PlotItemFlags_NoFit    = 2
)

// original name: ImPlotLegendFlags_
type PlotLegendFlags int

const (
	PlotLegendFlags_None            = 0
	PlotLegendFlags_NoButtons       = 1
	PlotLegendFlags_NoHighlightItem = 2
	PlotLegendFlags_NoHighlightAxis = 4
	PlotLegendFlags_NoMenus         = 8
	PlotLegendFlags_Outside         = 16
	PlotLegendFlags_Horizontal      = 32
)

// original name: ImPlotLineFlags_
type PlotLineFlags int

const (
	PlotLineFlags_None     = 0
	PlotLineFlags_Segments = 1024
	PlotLineFlags_Loop     = 2048
	PlotLineFlags_SkipNaN  = 4096
	PlotLineFlags_NoClip   = 8192
)

// original name: ImPlotLocation_
type PlotLocation int

const (
	PlotLocation_Center    = 0
	PlotLocation_North     = 1
	PlotLocation_South     = 2
	PlotLocation_West      = 4
	PlotLocation_East      = 8
	PlotLocation_NorthWest = 5
	PlotLocation_NorthEast = 9
	PlotLocation_SouthWest = 6
	PlotLocation_SouthEast = 10
)

// original name: ImPlotMarker_
type PlotMarker int

const (
	PlotMarker_None     = -1
	PlotMarker_Circle   = 0
	PlotMarker_Square   = 1
	PlotMarker_Diamond  = 2
	PlotMarker_Up       = 3
	PlotMarker_Down     = 4
	PlotMarker_Left     = 5
	PlotMarker_Right    = 6
	PlotMarker_Cross    = 7
	PlotMarker_Plus     = 8
	PlotMarker_Asterisk = 9
	PlotMarker_COUNT    = 10
)

// original name: ImPlotMouseTextFlags_
type PlotMouseTextFlags int

const (
	PlotMouseTextFlags_None       = 0
	PlotMouseTextFlags_NoAuxAxes  = 1
	PlotMouseTextFlags_NoFormat   = 2
	PlotMouseTextFlags_ShowAlways = 4
)

// original name: ImPlotPieChartFlags_
type PlotPieChartFlags int

const (
	PlotPieChartFlags_None      = 0
	PlotPieChartFlags_Normalize = 1024
)

// original name: ImPlotScale_
type PlotScale int

const (
	PlotScale_Linear = 0
	PlotScale_Time   = 1
	PlotScale_Log10  = 2
	PlotScale_SymLog = 3
)

// original name: ImPlotScatterFlags_
type PlotScatterFlags int

const (
	PlotScatterFlags_None   = 0
	PlotScatterFlags_NoClip = 1024
)

// original name: ImPlotShadedFlags_
type PlotShadedFlags int

const (
	PlotShadedFlags_None = 0
)

// original name: ImPlotStairsFlags_
type PlotStairsFlags int

const (
	PlotStairsFlags_None    = 0
	PlotStairsFlags_PreStep = 1024
)

// original name: ImPlotStemsFlags_
type PlotStemsFlags int

const (
	PlotStemsFlags_None       = 0
	PlotStemsFlags_Horizontal = 1024
)

// original name: ImPlotStyleVar_
type PlotStyleVar int

const (
	PlotStyleVar_LineWeight         = 0
	PlotStyleVar_Marker             = 1
	PlotStyleVar_MarkerSize         = 2
	PlotStyleVar_MarkerWeight       = 3
	PlotStyleVar_FillAlpha          = 4
	PlotStyleVar_ErrorBarSize       = 5
	PlotStyleVar_ErrorBarWeight     = 6
	PlotStyleVar_DigitalBitHeight   = 7
	PlotStyleVar_DigitalBitGap      = 8
	PlotStyleVar_PlotBorderSize     = 9
	PlotStyleVar_MinorAlpha         = 10
	PlotStyleVar_MajorTickLen       = 11
	PlotStyleVar_MinorTickLen       = 12
	PlotStyleVar_MajorTickSize      = 13
	PlotStyleVar_MinorTickSize      = 14
	PlotStyleVar_MajorGridSize      = 15
	PlotStyleVar_MinorGridSize      = 16
	PlotStyleVar_PlotPadding        = 17
	PlotStyleVar_LabelPadding       = 18
	PlotStyleVar_LegendPadding      = 19
	PlotStyleVar_LegendInnerPadding = 20
	PlotStyleVar_LegendSpacing      = 21
	PlotStyleVar_MousePosPadding    = 22
	PlotStyleVar_AnnotationPadding  = 23
	PlotStyleVar_FitPadding         = 24
	PlotStyleVar_PlotDefaultSize    = 25
	PlotStyleVar_PlotMinSize        = 26
	PlotStyleVar_COUNT              = 27
)

// original name: ImPlotSubplotFlags_
type PlotSubplotFlags int

const (
	PlotSubplotFlags_None       = 0
	PlotSubplotFlags_NoTitle    = 1
	PlotSubplotFlags_NoLegend   = 2
	PlotSubplotFlags_NoMenus    = 4
	PlotSubplotFlags_NoResize   = 8
	PlotSubplotFlags_NoAlign    = 16
	PlotSubplotFlags_ShareItems = 32
	PlotSubplotFlags_LinkRows   = 64
	PlotSubplotFlags_LinkCols   = 128
	PlotSubplotFlags_LinkAllX   = 256
	PlotSubplotFlags_LinkAllY   = 512
	PlotSubplotFlags_ColMajor   = 1024
)

// original name: ImPlotTextFlags_
type PlotTextFlags int

const (
	PlotTextFlags_None     = 0
	PlotTextFlags_Vertical = 1024
)

// original name: ImPlotTimeFmt_
type PlotTimeFmt int

const (
	PlotTimeFmt_None     = 0
	PlotTimeFmt_Us       = 1
	PlotTimeFmt_SUs      = 2
	PlotTimeFmt_SMs      = 3
	PlotTimeFmt_S        = 4
	PlotTimeFmt_HrMinSMs = 5
	PlotTimeFmt_HrMinS   = 6
	PlotTimeFmt_HrMin    = 7
	PlotTimeFmt_Hr       = 8
)

// original name: ImPlotTimeUnit_
type PlotTimeUnit int

const (
	PlotTimeUnit_Us    = 0
	PlotTimeUnit_Ms    = 1
	PlotTimeUnit_S     = 2
	PlotTimeUnit_Min   = 3
	PlotTimeUnit_Hr    = 4
	PlotTimeUnit_Day   = 5
	PlotTimeUnit_Mo    = 6
	PlotTimeUnit_Yr    = 7
	PlotTimeUnit_COUNT = 8
)
