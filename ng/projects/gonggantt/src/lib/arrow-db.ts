// insertion point for imports
import { BarDB } from './bar-db'
import { GanttDB } from './gantt-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class ArrowDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	OptionnalColor?: string
	OptionnalStroke?: string

	// insertion point for other declarations
	From?: BarDB
	FromID?: NullInt64

	To?: BarDB
	ToID?: NullInt64

	Gantt_ArrowsDBID?: NullInt64
	Gantt_ArrowsDBID_Index?: NullInt64 // store the index of the arrow instance in Gantt.Arrows
	Gantt_Arrows_reverse?: GanttDB

}
