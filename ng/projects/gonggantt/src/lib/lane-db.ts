// insertion point for imports
import { BarDB } from './bar-db'
import { GanttDB } from './gantt-db'

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
	Gantt_Lanes_reverse?: GanttDB

}
