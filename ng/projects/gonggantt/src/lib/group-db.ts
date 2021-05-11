// insertion point for imports
import { LaneDB } from './lane-db'
import { GanttDB } from './gantt-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class GroupDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string

	// insertion point for other declarations
	GroupLanes?: Array<LaneDB>
	Gantt_GroupsDBID?: NullInt64
	Gantt_GroupsDBID_Index?: NullInt64 // store the index of the group instance in Gantt.Groups
	Gantt_Groups_reverse?: GanttDB

}
