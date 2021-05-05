package models

import "time"

type Gantt struct {
	Name  string
	Start time.Time
	End   time.Time

	// List of Lanes
	Lanes []*Lane
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
