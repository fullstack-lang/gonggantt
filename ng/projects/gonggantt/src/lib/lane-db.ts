// insertion point for imports
import { BarDB } from './bar-db'
import { GanttDB } from './gantt-db'
import { GroupDB } from './group-db'
import { MilestoneDB } from './milestone-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class LaneDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Order?: number

	// insertion point for other declarations
	Bars?: Array<BarDB>
	Gantt_LanesDBID?: NullInt64
	Gantt_LanesDBID_Index?: NullInt64 // store the index of the lane instance in Gantt.Lanes
	Gantt_Lanes_reverse?: GanttDB

	Group_GroupLanesDBID?: NullInt64
	Group_GroupLanesDBID_Index?: NullInt64 // store the index of the lane instance in Group.GroupLanes
	Group_GroupLanes_reverse?: GroupDB

	Milestone_DiamonfAndTextAnchorsDBID?: NullInt64
	Milestone_DiamonfAndTextAnchorsDBID_Index?: NullInt64 // store the index of the lane instance in Milestone.DiamonfAndTextAnchors
	Milestone_DiamonfAndTextAnchors_reverse?: MilestoneDB

}
