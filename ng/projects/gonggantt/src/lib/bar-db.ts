// insertion point for imports
import { LaneDB } from './lane-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class BarDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Start?: Date
	End?: Date

	// insertion point for other declarations
	Lane_BarsDBID?: NullInt64
	Lane_Bars_reverse?: LaneDB

}
