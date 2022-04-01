// insertion point for imports
import { BarDB } from './bar-db'
import { GanttDB } from './gantt-db'
import { GroupDB } from './group-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LaneDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Order: number = 0

	// insertion point for other declarations
	Bars?: Array<BarDB>
	Gantt_LanesDBID: NullInt64 = new NullInt64
	Gantt_LanesDBID_Index: NullInt64  = new NullInt64 // store the index of the lane instance in Gantt.Lanes
	Gantt_Lanes_reverse?: GanttDB 

	Group_GroupLanesDBID: NullInt64 = new NullInt64
	Group_GroupLanesDBID_Index: NullInt64  = new NullInt64 // store the index of the lane instance in Group.GroupLanes
	Group_GroupLanes_reverse?: GroupDB 

}
