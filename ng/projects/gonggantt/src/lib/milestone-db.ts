// insertion point for imports
import { LaneDB } from './lane-db'
import { GanttDB } from './gantt-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class MilestoneDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Date?: Date

	// insertion point for other declarations
	DiamonfAndTextAnchors?: Array<LaneDB>
	Gantt_MilestonesDBID?: NullInt64
	Gantt_Milestones_reverse?: GanttDB

}
