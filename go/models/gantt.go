package models

import "time"

type Gantt struct {
	Name string

	ComputedStart time.Time
	ComputedEnd   time.Time

	LaneHeight           float64
	RatioBarToLaneHeight float64
	YTopMargin           float64

	XLeftText  float64
	TextHeight float64

	XLeftLanes   float64
	XRightMargin float64

	ArrowLengthToTheRightOfStartBar float64
	ArrowTipLenght                  float64

	TimeLine_Color       string
	TimeLine_FillOpacity float64
	TimeLine_Stroke      string
	TimeLine_StrokeWidth float64

	Group_Stroke          string
	Group_StrokeWidth     float64
	Group_StrokeDashArray string

	DateYOffset float64

	AlignOnStartEndOnYearStart bool

	// List of Lanes
	Lanes []*Lane

	// list of Milestones
	Milestones []*Milestone

	// list of Groups
	Groups []*Group

	// list of Arrows
	Arrows []*Arrow
}

func (gantt *Gantt) ComputeStartAndEndDate() {

	firstBar := true
	for _, lane := range gantt.Lanes {
		for _, bar := range lane.Bars {

			if firstBar == true {
				gantt.ComputedStart = bar.Start
				gantt.ComputedEnd = bar.End

				firstBar = false
			} else {
				if gantt.ComputedStart.After(bar.Start) {
					gantt.ComputedStart = bar.Start
				}
				if gantt.ComputedEnd.Before(bar.End) {
					gantt.ComputedEnd = bar.End
				}
			}
		}
	}

	// align start on the beginning of the year
	if gantt.AlignOnStartEndOnYearStart {
		gantt.ComputedStart = time.Date(gantt.ComputedStart.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
		gantt.ComputedEnd = time.Date(gantt.ComputedEnd.Year(), time.December, 31, 0, 0, 0, 0, time.UTC)
	}
}
