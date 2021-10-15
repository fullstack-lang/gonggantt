// insertion point for imports
import { LaneDB } from './lane-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BarDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Start: Date = new Date
	End: Date = new Date
	OptionnalColor: string = ""
	OptionnalStroke: string = ""

	// insertion point for other declarations
	Lane_BarsDBID: NullInt64 = new NullInt64
	Lane_BarsDBID_Index: NullInt64  = new NullInt64 // store the index of the bar instance in Lane.Bars
	Lane_Bars_reverse?: LaneDB 

}
