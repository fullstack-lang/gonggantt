// insertion point for imports
import { LaneDB } from './lane-db'
import { MilestoneDB } from './milestone-db'
import { GroupDB } from './group-db'
import { ArrowDB } from './arrow-db'

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
	LaneHeight?: number
	RatioBarToLaneHeight?: number
	YTopMargin?: number
	XLeftText?: number
	TextHeight?: number
	XLeftLanes?: number
	XRightMargin?: number
	ArrowLengthToTheRightOfStartBar?: number
	ArrowTipLenght?: number
	TimeLine_Color?: string
	TimeLine_FillOpacity?: number
	TimeLine_Stroke?: string
	TimeLine_StrokeWidth?: number
	Group_Stroke?: string
	Group_StrokeWidth?: number
	Group_StrokeDashArray?: string
	DateYOffset?: number
	AlignOnStartEndOnYearStart?: string

	// insertion point for other declarations
	Lanes?: Array<LaneDB>
	Milestones?: Array<MilestoneDB>
	Groups?: Array<GroupDB>
	Arrows?: Array<ArrowDB>
}
