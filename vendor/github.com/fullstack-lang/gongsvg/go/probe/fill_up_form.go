// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongsvg/go/models"
	"github.com/fullstack-lang/gongsvg/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Animate:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("AttributeName", instanceWithInferedType.AttributeName, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Values", instanceWithInferedType.Values, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Dur", instanceWithInferedType.Dur, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("RepeatCount", instanceWithInferedType.RepeatCount, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Circle"
			rf.Fieldname = "Animations"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Circle),
					"Animations",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Circle, *models.Animate](
					nil,
					"Animations",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ellipse"
			rf.Fieldname = "Animates"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ellipse),
					"Animates",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Ellipse, *models.Animate](
					nil,
					"Animates",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Line"
			rf.Fieldname = "Animates"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Line),
					"Animates",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Line, *models.Animate](
					nil,
					"Animates",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "LinkAnchoredText"
			rf.Fieldname = "Animates"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.LinkAnchoredText),
					"Animates",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.LinkAnchoredText, *models.Animate](
					nil,
					"Animates",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Path"
			rf.Fieldname = "Animates"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Path),
					"Animates",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Path, *models.Animate](
					nil,
					"Animates",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Polygone"
			rf.Fieldname = "Animates"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Polygone),
					"Animates",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Polygone, *models.Animate](
					nil,
					"Animates",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Polyline"
			rf.Fieldname = "Animates"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Polyline),
					"Animates",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Polyline, *models.Animate](
					nil,
					"Animates",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Rect"
			rf.Fieldname = "Animations"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Rect),
					"Animations",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Rect, *models.Animate](
					nil,
					"Animations",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RectAnchoredText"
			rf.Fieldname = "Animates"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RectAnchoredText),
					"Animates",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.RectAnchoredText, *models.Animate](
					nil,
					"Animates",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Text"
			rf.Fieldname = "Animates"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Text),
					"Animates",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Text, *models.Animate](
					nil,
					"Animates",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Circle:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CX", instanceWithInferedType.CX, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CY", instanceWithInferedType.CY, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Radius", instanceWithInferedType.Radius, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Animations", instanceWithInferedType, &instanceWithInferedType.Animations, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Circles"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"Circles",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.Circle](
					nil,
					"Circles",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Ellipse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CX", instanceWithInferedType.CX, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CY", instanceWithInferedType.CY, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("RX", instanceWithInferedType.RX, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("RY", instanceWithInferedType.RY, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Ellipses"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"Ellipses",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.Ellipse](
					nil,
					"Ellipses",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Layer:
		// insertion point
		BasicFieldtoForm("Display", instanceWithInferedType.Display, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Rects", instanceWithInferedType, &instanceWithInferedType.Rects, formGroup, playground)
		AssociationSliceToForm("Texts", instanceWithInferedType, &instanceWithInferedType.Texts, formGroup, playground)
		AssociationSliceToForm("Circles", instanceWithInferedType, &instanceWithInferedType.Circles, formGroup, playground)
		AssociationSliceToForm("Lines", instanceWithInferedType, &instanceWithInferedType.Lines, formGroup, playground)
		AssociationSliceToForm("Ellipses", instanceWithInferedType, &instanceWithInferedType.Ellipses, formGroup, playground)
		AssociationSliceToForm("Polylines", instanceWithInferedType, &instanceWithInferedType.Polylines, formGroup, playground)
		AssociationSliceToForm("Polygones", instanceWithInferedType, &instanceWithInferedType.Polygones, formGroup, playground)
		AssociationSliceToForm("Paths", instanceWithInferedType, &instanceWithInferedType.Paths, formGroup, playground)
		AssociationSliceToForm("Links", instanceWithInferedType, &instanceWithInferedType.Links, formGroup, playground)
		AssociationSliceToForm("RectLinkLinks", instanceWithInferedType, &instanceWithInferedType.RectLinkLinks, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SVG"
			rf.Fieldname = "Layers"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SVG),
					"Layers",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.SVG, *models.Layer](
					nil,
					"Layers",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Line:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("X1", instanceWithInferedType.X1, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Y1", instanceWithInferedType.Y1, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("X2", instanceWithInferedType.X2, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Y2", instanceWithInferedType.Y2, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, playground)
		BasicFieldtoForm("MouseClickX", instanceWithInferedType.MouseClickX, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("MouseClickY", instanceWithInferedType.MouseClickY, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Lines"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"Lines",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.Line](
					nil,
					"Lines",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("Start", instanceWithInferedType.Start, formGroup, playground)
		EnumTypeStringToForm("StartAnchorType", instanceWithInferedType.StartAnchorType, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("End", instanceWithInferedType.End, formGroup, playground)
		EnumTypeStringToForm("EndAnchorType", instanceWithInferedType.EndAnchorType, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CornerRadius", instanceWithInferedType.CornerRadius, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("HasEndArrow", instanceWithInferedType.HasEndArrow, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("EndArrowSize", instanceWithInferedType.EndArrowSize, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("TextAtArrowEnd", instanceWithInferedType, &instanceWithInferedType.TextAtArrowEnd, formGroup, playground)
		AssociationSliceToForm("TextAtArrowStart", instanceWithInferedType, &instanceWithInferedType.TextAtArrowStart, formGroup, playground)
		AssociationSliceToForm("ControlPoints", instanceWithInferedType, &instanceWithInferedType.ControlPoints, formGroup, playground)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Links"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"Links",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.Link](
					nil,
					"Links",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.LinkAnchoredText:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("X_Offset", instanceWithInferedType.X_Offset, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Y_Offset", instanceWithInferedType.Y_Offset, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FontWeight", instanceWithInferedType.FontWeight, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Link"
			rf.Fieldname = "TextAtArrowEnd"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Link),
					"TextAtArrowEnd",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Link, *models.LinkAnchoredText](
					nil,
					"TextAtArrowEnd",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Link"
			rf.Fieldname = "TextAtArrowStart"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Link),
					"TextAtArrowStart",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Link, *models.LinkAnchoredText](
					nil,
					"TextAtArrowStart",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Path:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Definition", instanceWithInferedType.Definition, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Paths"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"Paths",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.Path](
					nil,
					"Paths",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Point:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Link"
			rf.Fieldname = "ControlPoints"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Link),
					"ControlPoints",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Link, *models.Point](
					nil,
					"ControlPoints",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Polygone:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Points", instanceWithInferedType.Points, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Polygones"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"Polygones",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.Polygone](
					nil,
					"Polygones",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Polyline:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Points", instanceWithInferedType.Points, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Polylines"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"Polylines",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.Polyline](
					nil,
					"Polylines",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Rect:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("RX", instanceWithInferedType.RX, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Animations", instanceWithInferedType, &instanceWithInferedType.Animations, formGroup, playground)
		BasicFieldtoForm("IsSelectable", instanceWithInferedType.IsSelectable, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("IsSelected", instanceWithInferedType.IsSelected, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CanHaveLeftHandle", instanceWithInferedType.CanHaveLeftHandle, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("HasLeftHandle", instanceWithInferedType.HasLeftHandle, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CanHaveRightHandle", instanceWithInferedType.CanHaveRightHandle, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("HasRightHandle", instanceWithInferedType.HasRightHandle, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CanHaveTopHandle", instanceWithInferedType.CanHaveTopHandle, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("HasTopHandle", instanceWithInferedType.HasTopHandle, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CanHaveBottomHandle", instanceWithInferedType.CanHaveBottomHandle, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("HasBottomHandle", instanceWithInferedType.HasBottomHandle, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CanMoveHorizontaly", instanceWithInferedType.CanMoveHorizontaly, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CanMoveVerticaly", instanceWithInferedType.CanMoveVerticaly, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("RectAnchoredTexts", instanceWithInferedType, &instanceWithInferedType.RectAnchoredTexts, formGroup, playground)
		AssociationSliceToForm("RectAnchoredRects", instanceWithInferedType, &instanceWithInferedType.RectAnchoredRects, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Rects"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"Rects",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.Rect](
					nil,
					"Rects",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.RectAnchoredRect:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("RX", instanceWithInferedType.RX, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("X_Offset", instanceWithInferedType.X_Offset, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Y_Offset", instanceWithInferedType.Y_Offset, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("RectAnchorType", instanceWithInferedType.RectAnchorType, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("WidthFollowRect", instanceWithInferedType.WidthFollowRect, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("HeightFollowRect", instanceWithInferedType.HeightFollowRect, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Rect"
			rf.Fieldname = "RectAnchoredRects"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Rect),
					"RectAnchoredRects",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Rect, *models.RectAnchoredRect](
					nil,
					"RectAnchoredRects",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.RectAnchoredText:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FontWeight", instanceWithInferedType.FontWeight, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FontSize", instanceWithInferedType.FontSize, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("X_Offset", instanceWithInferedType.X_Offset, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Y_Offset", instanceWithInferedType.Y_Offset, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("RectAnchorType", instanceWithInferedType.RectAnchorType, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("TextAnchorType", instanceWithInferedType.TextAnchorType, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Rect"
			rf.Fieldname = "RectAnchoredTexts"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Rect),
					"RectAnchoredTexts",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Rect, *models.RectAnchoredText](
					nil,
					"RectAnchoredTexts",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.RectLinkLink:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Start", instanceWithInferedType.Start, formGroup, playground)
		AssociationFieldToForm("End", instanceWithInferedType.End, formGroup, playground)
		BasicFieldtoForm("TargetAnchorPosition", instanceWithInferedType.TargetAnchorPosition, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "RectLinkLinks"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"RectLinkLinks",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.RectLinkLink](
					nil,
					"RectLinkLinks",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.SVG:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Layers", instanceWithInferedType, &instanceWithInferedType.Layers, formGroup, playground)
		EnumTypeStringToForm("DrawingState", instanceWithInferedType.DrawingState, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("StartRect", instanceWithInferedType.StartRect, formGroup, playground)
		AssociationFieldToForm("EndRect", instanceWithInferedType.EndRect, formGroup, playground)
		BasicFieldtoForm("IsEditable", instanceWithInferedType.IsEditable, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.Text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Animates", instanceWithInferedType, &instanceWithInferedType.Animates, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Texts"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Layer),
					"Texts",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Layer, *models.Text](
					nil,
					"Texts",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	default:
		_ = instanceWithInferedType
	}
}
