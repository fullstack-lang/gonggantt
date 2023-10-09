// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gonggantt/go/models"
	"github.com/fullstack-lang/gonggantt/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Arrow:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("From", instanceWithInferedType.From, formGroup, playground)
		AssociationFieldToForm("To", instanceWithInferedType.To, formGroup, playground)
		BasicFieldtoForm("OptionnalColor", instanceWithInferedType.OptionnalColor, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("OptionnalStroke", instanceWithInferedType.OptionnalStroke, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Gantt"
			rf.Fieldname = "Arrows"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Gantt),
					"Arrows",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Gantt, *models.Arrow](
					nil,
					"Arrows",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Bar:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Start", instanceWithInferedType.Start, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("End", instanceWithInferedType.End, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("ComputedDuration", instanceWithInferedType.ComputedDuration, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("OptionnalColor", instanceWithInferedType.OptionnalColor, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("OptionnalStroke", instanceWithInferedType.OptionnalStroke, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Lane"
			rf.Fieldname = "Bars"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Lane),
					"Bars",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Lane, *models.Bar](
					nil,
					"Bars",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Gantt:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("ComputedStart", instanceWithInferedType.ComputedStart, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("ComputedEnd", instanceWithInferedType.ComputedEnd, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("ComputedDuration", instanceWithInferedType.ComputedDuration, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("UseManualStartAndEndDates", instanceWithInferedType.UseManualStartAndEndDates, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("ManualStart", instanceWithInferedType.ManualStart, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("ManualEnd", instanceWithInferedType.ManualEnd, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("LaneHeight", instanceWithInferedType.LaneHeight, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("RatioBarToLaneHeight", instanceWithInferedType.RatioBarToLaneHeight, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("YTopMargin", instanceWithInferedType.YTopMargin, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("XLeftText", instanceWithInferedType.XLeftText, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("TextHeight", instanceWithInferedType.TextHeight, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("XLeftLanes", instanceWithInferedType.XLeftLanes, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("XRightMargin", instanceWithInferedType.XRightMargin, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("ArrowLengthToTheRightOfStartBar", instanceWithInferedType.ArrowLengthToTheRightOfStartBar, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("ArrowTipLenght", instanceWithInferedType.ArrowTipLenght, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("TimeLine_Color", instanceWithInferedType.TimeLine_Color, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("TimeLine_FillOpacity", instanceWithInferedType.TimeLine_FillOpacity, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("TimeLine_Stroke", instanceWithInferedType.TimeLine_Stroke, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("TimeLine_StrokeWidth", instanceWithInferedType.TimeLine_StrokeWidth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Group_Stroke", instanceWithInferedType.Group_Stroke, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Group_StrokeWidth", instanceWithInferedType.Group_StrokeWidth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Group_StrokeDashArray", instanceWithInferedType.Group_StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("DateYOffset", instanceWithInferedType.DateYOffset, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("AlignOnStartEndOnYearStart", instanceWithInferedType.AlignOnStartEndOnYearStart, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Lanes", instanceWithInferedType, &instanceWithInferedType.Lanes, formGroup, playground)
		AssociationSliceToForm("Milestones", instanceWithInferedType, &instanceWithInferedType.Milestones, formGroup, playground)
		AssociationSliceToForm("Groups", instanceWithInferedType, &instanceWithInferedType.Groups, formGroup, playground)
		AssociationSliceToForm("Arrows", instanceWithInferedType, &instanceWithInferedType.Arrows, formGroup, playground)

	case *models.Group:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("GroupLanes", instanceWithInferedType, &instanceWithInferedType.GroupLanes, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Gantt"
			rf.Fieldname = "Groups"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Gantt),
					"Groups",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Gantt, *models.Group](
					nil,
					"Groups",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Lane:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Order", instanceWithInferedType.Order, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Bars", instanceWithInferedType, &instanceWithInferedType.Bars, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Gantt"
			rf.Fieldname = "Lanes"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Gantt),
					"Lanes",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Gantt, *models.Lane](
					nil,
					"Lanes",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "GroupLanes"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Group),
					"GroupLanes",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Group, *models.Lane](
					nil,
					"GroupLanes",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.LaneUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Lane", instanceWithInferedType.Lane, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Milestone"
			rf.Fieldname = "LanesToDisplayMilestoneUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Milestone),
					"LanesToDisplayMilestoneUse",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Milestone, *models.LaneUse](
					nil,
					"LanesToDisplayMilestoneUse",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Milestone:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Date", instanceWithInferedType.Date, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("DisplayVerticalBar", instanceWithInferedType.DisplayVerticalBar, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("LanesToDisplayMilestoneUse", instanceWithInferedType, &instanceWithInferedType.LanesToDisplayMilestoneUse, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Gantt"
			rf.Fieldname = "Milestones"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Gantt),
					"Milestones",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Gantt, *models.Milestone](
					nil,
					"Milestones",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	default:
		_ = instanceWithInferedType
	}
}
