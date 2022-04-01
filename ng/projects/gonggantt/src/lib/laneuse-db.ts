// insertion point for imports
import { LaneDB } from './lane-db'
import { MilestoneDB } from './milestone-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LaneUseDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Lane?: LaneDB
	LaneID: NullInt64 = new NullInt64 // if pointer is null, Lane.ID = 0

	Milestone_LanesToDisplayMilestoneDBID: NullInt64 = new NullInt64
	Milestone_LanesToDisplayMilestoneDBID_Index: NullInt64  = new NullInt64 // store the index of the laneuse instance in Milestone.LanesToDisplayMilestone
	Milestone_LanesToDisplayMilestone_reverse?: MilestoneDB 

}
