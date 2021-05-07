package models

import "time"

type Gantt struct {
	Name  string
	Start time.Time
	End   time.Time

	LaneHeight           float64
	RatioBarToLaneHeight float64
	YTopMargin           float64

	XLeftText  float64
	TextHeight float64

	XLeftLanes   float64
	XRightMargin float64

	TimeLine_Color       string
	TimeLine_FillOpacity float64
	TimeLine_Stroke      string
	TimeLine_StrokeWidth float64

	DateYOffset float64

	// List of Lanes
	Lanes []*Lane

	// list of Milestones
	Milestones []*Milestone
}

func (gantt *Gantt) ComputeStartAndEndDate() {

	firstBar := true
	for _, lane := range gantt.Lanes {
		for _, bar := range lane.Bars {

			if firstBar == true {
				gantt.Start = bar.Start
				gantt.End = bar.End

				firstBar = false
			} else {
				if gantt.Start.After(bar.Start) {
					gantt.Start = bar.Start
				}
				if gantt.End.Before(bar.End) {
					gantt.End = bar.End
				}
			}
		}
	}
}
