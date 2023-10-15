// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type Arrow_WOP struct {
	// insertion point
	Name string
	OptionnalColor string
	OptionnalStroke string
}

type Bar_WOP struct {
	// insertion point
	Name string
	Start time.Time
	End time.Time
	ComputedDuration time.Duration
	OptionnalColor string
	OptionnalStroke string
	FillOpacity float64
	StrokeWidth float64
	StrokeDashArray string
}

type Gantt_WOP struct {
	// insertion point
	Name string
	ComputedStart time.Time
	ComputedEnd time.Time
	ComputedDuration time.Duration
	UseManualStartAndEndDates bool
	ManualStart time.Time
	ManualEnd time.Time
	LaneHeight float64
	RatioBarToLaneHeight float64
	YTopMargin float64
	XLeftText float64
	TextHeight float64
	XLeftLanes float64
	XRightMargin float64
	ArrowLengthToTheRightOfStartBar float64
	ArrowTipLenght float64
	TimeLine_Color string
	TimeLine_FillOpacity float64
	TimeLine_Stroke string
	TimeLine_StrokeWidth float64
	Group_Stroke string
	Group_StrokeWidth float64
	Group_StrokeDashArray string
	DateYOffset float64
	AlignOnStartEndOnYearStart bool
}

type Group_WOP struct {
	// insertion point
	Name string
}

type Lane_WOP struct {
	// insertion point
	Name string
	Order int
}

type LaneUse_WOP struct {
	// insertion point
	Name string
}

type Milestone_WOP struct {
	// insertion point
	Name string
	Date time.Time
	DisplayVerticalBar bool
}

