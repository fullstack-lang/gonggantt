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
	case Animate:
		fieldCoder := Animate{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.AttributeName = "1"
		fieldCoder.Values = "2"
		fieldCoder.Dur = "3"
		fieldCoder.RepeatCount = "4"
		return (any)(fieldCoder).(Type)
	case Circle:
		fieldCoder := Circle{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.CX = 1.000000
		fieldCoder.CY = 2.000000
		fieldCoder.Radius = 3.000000
		fieldCoder.Color = "4"
		fieldCoder.FillOpacity = 5.000000
		fieldCoder.Stroke = "6"
		fieldCoder.StrokeWidth = 7.000000
		fieldCoder.StrokeDashArray = "8"
		fieldCoder.Transform = "9"
		return (any)(fieldCoder).(Type)
	case Ellipse:
		fieldCoder := Ellipse{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.CX = 1.000000
		fieldCoder.CY = 2.000000
		fieldCoder.RX = 3.000000
		fieldCoder.RY = 4.000000
		fieldCoder.Color = "5"
		fieldCoder.FillOpacity = 6.000000
		fieldCoder.Stroke = "7"
		fieldCoder.StrokeWidth = 8.000000
		fieldCoder.StrokeDashArray = "9"
		fieldCoder.Transform = "10"
		return (any)(fieldCoder).(Type)
	case Line:
		fieldCoder := Line{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.X1 = 1.000000
		fieldCoder.Y1 = 2.000000
		fieldCoder.X2 = 3.000000
		fieldCoder.Y2 = 4.000000
		fieldCoder.Color = "5"
		fieldCoder.FillOpacity = 6.000000
		fieldCoder.Stroke = "7"
		fieldCoder.StrokeWidth = 8.000000
		fieldCoder.StrokeDashArray = "9"
		fieldCoder.Transform = "10"
		return (any)(fieldCoder).(Type)
	case Path:
		fieldCoder := Path{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Definition = "1"
		fieldCoder.Color = "2"
		fieldCoder.FillOpacity = 3.000000
		fieldCoder.Stroke = "4"
		fieldCoder.StrokeWidth = 5.000000
		fieldCoder.StrokeDashArray = "6"
		fieldCoder.Transform = "7"
		return (any)(fieldCoder).(Type)
	case Polygone:
		fieldCoder := Polygone{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Points = "1"
		fieldCoder.Color = "2"
		fieldCoder.FillOpacity = 3.000000
		fieldCoder.Stroke = "4"
		fieldCoder.StrokeWidth = 5.000000
		fieldCoder.StrokeDashArray = "6"
		fieldCoder.Transform = "7"
		return (any)(fieldCoder).(Type)
	case Polyline:
		fieldCoder := Polyline{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Points = "1"
		fieldCoder.Color = "2"
		fieldCoder.FillOpacity = 3.000000
		fieldCoder.Stroke = "4"
		fieldCoder.StrokeWidth = 5.000000
		fieldCoder.StrokeDashArray = "6"
		fieldCoder.Transform = "7"
		return (any)(fieldCoder).(Type)
	case Rect:
		fieldCoder := Rect{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.X = 1.000000
		fieldCoder.Y = 2.000000
		fieldCoder.Width = 3.000000
		fieldCoder.Height = 4.000000
		fieldCoder.RX = 5.000000
		fieldCoder.Color = "6"
		fieldCoder.FillOpacity = 7.000000
		fieldCoder.Stroke = "8"
		fieldCoder.StrokeWidth = 9.000000
		fieldCoder.StrokeDashArray = "10"
		fieldCoder.Transform = "11"
		return (any)(fieldCoder).(Type)
	case SVG:
		fieldCoder := SVG{}
		// insertion point for field dependant code
		fieldCoder.Display = false
		fieldCoder.Name = "1"
		return (any)(fieldCoder).(Type)
	case Text:
		fieldCoder := Text{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.X = 1.000000
		fieldCoder.Y = 2.000000
		fieldCoder.Content = "3"
		fieldCoder.Color = "4"
		fieldCoder.FillOpacity = 5.000000
		fieldCoder.Stroke = "6"
		fieldCoder.StrokeWidth = 7.000000
		fieldCoder.StrokeDashArray = "8"
		fieldCoder.Transform = "9"
		return (any)(fieldCoder).(Type)
	default:
		return t
	}
}

type Gongfield interface {
	string | bool | int | float64 | time.Time | time.Duration | *Animate | []*Animate | *Circle | []*Circle | *Ellipse | []*Ellipse | *Line | []*Line | *Path | []*Path | *Polygone | []*Polygone | *Polyline | []*Polyline | *Rect | []*Rect | *SVG | []*SVG | *Text | []*Text
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
	case *Animate:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "AttributeName"
			}
			if field == "2" {
				return "Values"
			}
			if field == "3" {
				return "Dur"
			}
			if field == "4" {
				return "RepeatCount"
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
	case *Circle:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "4" {
				return "Color"
			}
			if field == "6" {
				return "Stroke"
			}
			if field == "8" {
				return "StrokeDashArray"
			}
			if field == "9" {
				return "Transform"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 1.000000 {
				return "CX"
			}
			if field == 2.000000 {
				return "CY"
			}
			if field == 3.000000 {
				return "Radius"
			}
			if field == 5.000000 {
				return "FillOpacity"
			}
			if field == 7.000000 {
				return "StrokeWidth"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Ellipse:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "5" {
				return "Color"
			}
			if field == "7" {
				return "Stroke"
			}
			if field == "9" {
				return "StrokeDashArray"
			}
			if field == "10" {
				return "Transform"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 1.000000 {
				return "CX"
			}
			if field == 2.000000 {
				return "CY"
			}
			if field == 3.000000 {
				return "RX"
			}
			if field == 4.000000 {
				return "RY"
			}
			if field == 6.000000 {
				return "FillOpacity"
			}
			if field == 8.000000 {
				return "StrokeWidth"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Line:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "5" {
				return "Color"
			}
			if field == "7" {
				return "Stroke"
			}
			if field == "9" {
				return "StrokeDashArray"
			}
			if field == "10" {
				return "Transform"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 1.000000 {
				return "X1"
			}
			if field == 2.000000 {
				return "Y1"
			}
			if field == 3.000000 {
				return "X2"
			}
			if field == 4.000000 {
				return "Y2"
			}
			if field == 6.000000 {
				return "FillOpacity"
			}
			if field == 8.000000 {
				return "StrokeWidth"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Path:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "Definition"
			}
			if field == "2" {
				return "Color"
			}
			if field == "4" {
				return "Stroke"
			}
			if field == "6" {
				return "StrokeDashArray"
			}
			if field == "7" {
				return "Transform"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 3.000000 {
				return "FillOpacity"
			}
			if field == 5.000000 {
				return "StrokeWidth"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Polygone:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "Points"
			}
			if field == "2" {
				return "Color"
			}
			if field == "4" {
				return "Stroke"
			}
			if field == "6" {
				return "StrokeDashArray"
			}
			if field == "7" {
				return "Transform"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 3.000000 {
				return "FillOpacity"
			}
			if field == 5.000000 {
				return "StrokeWidth"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Polyline:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "Points"
			}
			if field == "2" {
				return "Color"
			}
			if field == "4" {
				return "Stroke"
			}
			if field == "6" {
				return "StrokeDashArray"
			}
			if field == "7" {
				return "Transform"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 3.000000 {
				return "FillOpacity"
			}
			if field == 5.000000 {
				return "StrokeWidth"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Rect:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "6" {
				return "Color"
			}
			if field == "8" {
				return "Stroke"
			}
			if field == "10" {
				return "StrokeDashArray"
			}
			if field == "11" {
				return "Transform"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 1.000000 {
				return "X"
			}
			if field == 2.000000 {
				return "Y"
			}
			if field == 3.000000 {
				return "Width"
			}
			if field == 4.000000 {
				return "Height"
			}
			if field == 5.000000 {
				return "RX"
			}
			if field == 7.000000 {
				return "FillOpacity"
			}
			if field == 9.000000 {
				return "StrokeWidth"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *SVG:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "1" {
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
			if field == false {
				return "Display"
			}
		}
	case *Text:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "3" {
				return "Content"
			}
			if field == "4" {
				return "Color"
			}
			if field == "6" {
				return "Stroke"
			}
			if field == "8" {
				return "StrokeDashArray"
			}
			if field == "9" {
				return "Transform"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 1.000000 {
				return "X"
			}
			if field == 2.000000 {
				return "Y"
			}
			if field == 5.000000 {
				return "FillOpacity"
			}
			if field == 7.000000 {
				return "StrokeWidth"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	default:
		return ""
	}
	_ = field
	return ""
}
