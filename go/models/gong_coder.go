package models

import "time"

// GongfieldCoder return an instance of Type where each field 
// encodes the index of the field
//
// This allows for refactorable field names
// 
func GongfieldCoder[Type Gongstruct]() Type {
	var t Type

	switch any(t).(type) {
	// insertion point for cases
	case Arrow:
		fieldCoder := Arrow{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.OptionnalColor = "3"
		fieldCoder.OptionnalStroke = "4"
		return (any)(fieldCoder).(Type)
	case Bar:
		fieldCoder := Bar{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Start = time.Date(1, time.January, 0, 0, 0, 0, 0, time.UTC)
		fieldCoder.End = time.Date(2, time.January, 0, 0, 0, 0, 0, time.UTC)
		fieldCoder.OptionnalColor = "3"
		fieldCoder.OptionnalStroke = "4"
		fieldCoder.FillOpacity = 5.000000
		fieldCoder.StrokeWidth = 6.000000
		fieldCoder.StrokeDashArray = "7"
		return (any)(fieldCoder).(Type)
	case Gantt:
		fieldCoder := Gantt{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.ComputedStart = time.Date(1, time.January, 0, 0, 0, 0, 0, time.UTC)
		fieldCoder.ComputedEnd = time.Date(2, time.January, 0, 0, 0, 0, 0, time.UTC)
		fieldCoder.UseManualStartAndEndDates = false
		fieldCoder.ManualStart = time.Date(4, time.January, 0, 0, 0, 0, 0, time.UTC)
		fieldCoder.ManualEnd = time.Date(5, time.January, 0, 0, 0, 0, 0, time.UTC)
		fieldCoder.LaneHeight = 6.000000
		fieldCoder.RatioBarToLaneHeight = 7.000000
		fieldCoder.YTopMargin = 8.000000
		fieldCoder.XLeftText = 9.000000
		fieldCoder.TextHeight = 10.000000
		fieldCoder.XLeftLanes = 11.000000
		fieldCoder.XRightMargin = 12.000000
		fieldCoder.ArrowLengthToTheRightOfStartBar = 13.000000
		fieldCoder.ArrowTipLenght = 14.000000
		fieldCoder.TimeLine_Color = "15"
		fieldCoder.TimeLine_FillOpacity = 16.000000
		fieldCoder.TimeLine_Stroke = "17"
		fieldCoder.TimeLine_StrokeWidth = 18.000000
		fieldCoder.Group_Stroke = "19"
		fieldCoder.Group_StrokeWidth = 20.000000
		fieldCoder.Group_StrokeDashArray = "21"
		fieldCoder.DateYOffset = 22.000000
		fieldCoder.AlignOnStartEndOnYearStart = true
		return (any)(fieldCoder).(Type)
	case Group:
		fieldCoder := Group{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		return (any)(fieldCoder).(Type)
	case Lane:
		fieldCoder := Lane{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Order = 1
		return (any)(fieldCoder).(Type)
	case LaneUse:
		fieldCoder := LaneUse{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		return (any)(fieldCoder).(Type)
	case Milestone:
		fieldCoder := Milestone{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Date = time.Date(1, time.January, 0, 0, 0, 0, 0, time.UTC)
		fieldCoder.DisplayVerticalBar = false
		return (any)(fieldCoder).(Type)
	default:
		return t
	}
}

type Gongfield interface {
	string | bool | int | float64 | time.Time | time.Duration | *Arrow | []*Arrow | *Bar | []*Bar | *Gantt | []*Gantt | *Group | []*Group | *Lane | []*Lane | *LaneUse | []*LaneUse | *Milestone | []*Milestone
}

// GongfieldName provides the name of the field by passing the instance of the coder to
// the fonction.
//
// This allows for refactorable field name
//
// fieldCoder := models.GongfieldCoder[models.Astruct]()
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Name))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Booleanfield))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Intfield))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Floatfield))
// 
// limitations:
// 1. cannot encode boolean fields
// 2. for associations (pointer to gongstruct or slice of pointer to gongstruct, uses GetAssociationName)
func GongfieldName[Type PointerToGongstruct, FieldType Gongfield](field FieldType) string {
	var t Type

	switch any(t).(type) {
	// insertion point for cases
	case *Arrow:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "3" {
				return "OptionnalColor"
			}
			if field == "4" {
				return "OptionnalStroke"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Bar:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "3" {
				return "OptionnalColor"
			}
			if field == "4" {
				return "OptionnalStroke"
			}
			if field == "7" {
				return "StrokeDashArray"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 5.000000 {
				return "FillOpacity"
			}
			if field == 6.000000 {
				return "StrokeWidth"
			}
		case time.Time:
			// insertion point for field dependant name
			if field == time.Date(1, time.January, 0, 0, 0, 0, 0, time.UTC) {
				return "Start"
			}
			if field == time.Date(2, time.January, 0, 0, 0, 0, 0, time.UTC) {
				return "End"
			}
		case bool:
			// insertion point for field dependant name
		}
	case *Gantt:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "15" {
				return "TimeLine_Color"
			}
			if field == "17" {
				return "TimeLine_Stroke"
			}
			if field == "19" {
				return "Group_Stroke"
			}
			if field == "21" {
				return "Group_StrokeDashArray"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 6.000000 {
				return "LaneHeight"
			}
			if field == 7.000000 {
				return "RatioBarToLaneHeight"
			}
			if field == 8.000000 {
				return "YTopMargin"
			}
			if field == 9.000000 {
				return "XLeftText"
			}
			if field == 10.000000 {
				return "TextHeight"
			}
			if field == 11.000000 {
				return "XLeftLanes"
			}
			if field == 12.000000 {
				return "XRightMargin"
			}
			if field == 13.000000 {
				return "ArrowLengthToTheRightOfStartBar"
			}
			if field == 14.000000 {
				return "ArrowTipLenght"
			}
			if field == 16.000000 {
				return "TimeLine_FillOpacity"
			}
			if field == 18.000000 {
				return "TimeLine_StrokeWidth"
			}
			if field == 20.000000 {
				return "Group_StrokeWidth"
			}
			if field == 22.000000 {
				return "DateYOffset"
			}
		case time.Time:
			// insertion point for field dependant name
			if field == time.Date(1, time.January, 0, 0, 0, 0, 0, time.UTC) {
				return "ComputedStart"
			}
			if field == time.Date(2, time.January, 0, 0, 0, 0, 0, time.UTC) {
				return "ComputedEnd"
			}
			if field == time.Date(4, time.January, 0, 0, 0, 0, 0, time.UTC) {
				return "ManualStart"
			}
			if field == time.Date(5, time.January, 0, 0, 0, 0, 0, time.UTC) {
				return "ManualEnd"
			}
		case bool:
			// insertion point for field dependant name
			if field == false {
				return "UseManualStartAndEndDates"
			}
			if field == true {
				return "AlignOnStartEndOnYearStart"
			}
		}
	case *Group:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Lane:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 1 {
				return "Order"
			}
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *LaneUse:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Milestone:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
			if field == time.Date(1, time.January, 0, 0, 0, 0, 0, time.UTC) {
				return "Date"
			}
		case bool:
			// insertion point for field dependant name
			if field == false {
				return "DisplayVerticalBar"
			}
		}
	default:
		return ""
	}
	_ = field
	return ""
}
