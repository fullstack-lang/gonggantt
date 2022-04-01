package models

import "time"

type Milestone struct {
	Name string
	Date time.Time

	// a red diamond a text anchor will be displayed
	LanesToDisplayMilestone []*LaneUse
}
