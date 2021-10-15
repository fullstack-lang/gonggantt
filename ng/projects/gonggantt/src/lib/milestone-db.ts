// insertion point for imports
import { LaneDB } from './lane-db'
import { GanttDB } from './gantt-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MilestoneDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Date: Date = new Date

	// insertion point for other declarations
	DiamonfAndTextAnchors?: Array<LaneDB>
	Gantt_MilestonesDBID: NullInt64 = new NullInt64
	Gantt_MilestonesDBID_Index: NullInt64  = new NullInt64 // store the index of the milestone instance in Gantt.Milestones
	Gantt_Milestones_reverse?: GanttDB 

}
