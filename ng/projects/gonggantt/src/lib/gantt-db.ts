// insertion point for imports
import { LaneDB } from './lane-db'
import { MilestoneDB } from './milestone-db'
import { GroupDB } from './group-db'
import { ArrowDB } from './arrow-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GanttDB {

	static GONGSTRUCT_NAME = "Gantt"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	ComputedStart: Date = new Date
	ComputedEnd: Date = new Date
	UseManualStartAndEndDates: boolean = false
	ManualStart: Date = new Date
	ManualEnd: Date = new Date
	LaneHeight: number = 0
	RatioBarToLaneHeight: number = 0
	YTopMargin: number = 0
	XLeftText: number = 0
	TextHeight: number = 0
	XLeftLanes: number = 0
	XRightMargin: number = 0
	ArrowLengthToTheRightOfStartBar: number = 0
	ArrowTipLenght: number = 0
	TimeLine_Color: string = ""
	TimeLine_FillOpacity: number = 0
	TimeLine_Stroke: string = ""
	TimeLine_StrokeWidth: number = 0
	Group_Stroke: string = ""
	Group_StrokeWidth: number = 0
	Group_StrokeDashArray: string = ""
	DateYOffset: number = 0
	AlignOnStartEndOnYearStart: boolean = false

	// insertion point for other declarations
	Lanes?: Array<LaneDB>
	Milestones?: Array<MilestoneDB>
	Groups?: Array<GroupDB>
	Arrows?: Array<ArrowDB>
}
