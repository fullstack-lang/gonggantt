// insertion point for imports
import { AnimateDB } from './animate-db'
import { LayerDB } from './layer-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0
	Width: number = 0
	Height: number = 0
	RX: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	Transform: string = ""
	Selected: boolean = false

	// insertion point for other declarations
	Animations?: Array<AnimateDB>
	Layer_RectsDBID: NullInt64 = new NullInt64
	Layer_RectsDBID_Index: NullInt64  = new NullInt64 // store the index of the rect instance in Layer.Rects
	Layer_Rects_reverse?: LayerDB 

}
