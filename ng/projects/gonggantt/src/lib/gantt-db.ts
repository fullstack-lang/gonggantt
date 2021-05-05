// insertion point for imports
import { LaneDB } from './lane-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class GanttDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Start?: Date
	End?: Date

	// insertion point for other declarations
	Lanes?: Array<LaneDB>
}
