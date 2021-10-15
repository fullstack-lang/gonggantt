// insertion point for imports
import { BarDB } from './bar-db'
import { GanttDB } from './gantt-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ArrowDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	OptionnalColor: string = ""
	OptionnalStroke: string = ""

	// insertion point for other declarations
	From?: BarDB
	FromID: NullInt64 = new NullInt64 // if pointer is null, From.ID = 0

	To?: BarDB
	ToID: NullInt64 = new NullInt64 // if pointer is null, To.ID = 0

	Gantt_ArrowsDBID: NullInt64 = new NullInt64
	Gantt_ArrowsDBID_Index: NullInt64  = new NullInt64 // store the index of the arrow instance in Gantt.Arrows
	Gantt_Arrows_reverse?: GanttDB 

}
