package models

import "time"

type Gantt struct {
	Name  string
	Start time.Time
	End   time.Time

	// List of Lanes
	Lanes []*Lane
}
