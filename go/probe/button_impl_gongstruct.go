// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	form "github.com/fullstack-lang/gongtable/go/models"
	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gonggantt/go/models"
)

type ButtonImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	Icon       gongtree_buttons.ButtonType
	playground *Playground
}

func NewButtonImplGongstruct(
	gongStruct *gong_models.GongStruct,
	icon gongtree_buttons.ButtonType,
	playground *Playground,
) (buttonImplGongstruct *ButtonImplGongstruct) {

	buttonImplGongstruct = new(ButtonImplGongstruct)
	buttonImplGongstruct.Icon = icon
	buttonImplGongstruct.gongStruct = gongStruct
	buttonImplGongstruct.playground = playground

	return
}

func (buttonImpl *ButtonImplGongstruct) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	log.Println("ButtonImplGongstruct: ButtonUpdated")

	formStage := buttonImpl.playground.formStage
	formStage.Reset()
	formStage.Commit()

	switch buttonImpl.gongStruct.Name {
	// insertion point
	case "Arrow":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewArrowFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		arrow := new(models.Arrow)
		FillUpForm(arrow, formGroup, buttonImpl.playground)
	case "Bar":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewBarFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		bar := new(models.Bar)
		FillUpForm(bar, formGroup, buttonImpl.playground)
	case "Gantt":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGanttFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		gantt := new(models.Gantt)
		FillUpForm(gantt, formGroup, buttonImpl.playground)
	case "Group":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGroupFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		group := new(models.Group)
		FillUpForm(group, formGroup, buttonImpl.playground)
	case "Lane":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewLaneFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		lane := new(models.Lane)
		FillUpForm(lane, formGroup, buttonImpl.playground)
	case "LaneUse":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewLaneUseFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		laneuse := new(models.LaneUse)
		FillUpForm(laneuse, formGroup, buttonImpl.playground)
	case "Milestone":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewMilestoneFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		milestone := new(models.Milestone)
		FillUpForm(milestone, formGroup, buttonImpl.playground)
	}
	formStage.Commit()
}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Arrow:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("From", instanceWithInferedType.From, formGroup, playground)
		AssociationFieldToForm("To", instanceWithInferedType.To, formGroup, playground)
		BasicFieldtoForm("OptionnalColor", instanceWithInferedType.OptionnalColor, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("OptionnalStroke", instanceWithInferedType.OptionnalStroke, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Bar:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("OptionnalColor", instanceWithInferedType.OptionnalColor, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("OptionnalStroke", instanceWithInferedType.OptionnalStroke, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Gantt:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("UseManualStartAndEndDates", instanceWithInferedType.UseManualStartAndEndDates, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("LaneHeight", instanceWithInferedType.LaneHeight, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("RatioBarToLaneHeight", instanceWithInferedType.RatioBarToLaneHeight, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("YTopMargin", instanceWithInferedType.YTopMargin, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("XLeftText", instanceWithInferedType.XLeftText, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("TextHeight", instanceWithInferedType.TextHeight, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("XLeftLanes", instanceWithInferedType.XLeftLanes, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("XRightMargin", instanceWithInferedType.XRightMargin, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("ArrowLengthToTheRightOfStartBar", instanceWithInferedType.ArrowLengthToTheRightOfStartBar, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("ArrowTipLenght", instanceWithInferedType.ArrowTipLenght, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("TimeLine_Color", instanceWithInferedType.TimeLine_Color, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("TimeLine_FillOpacity", instanceWithInferedType.TimeLine_FillOpacity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("TimeLine_Stroke", instanceWithInferedType.TimeLine_Stroke, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("TimeLine_StrokeWidth", instanceWithInferedType.TimeLine_StrokeWidth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Group_Stroke", instanceWithInferedType.Group_Stroke, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Group_StrokeWidth", instanceWithInferedType.Group_StrokeWidth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Group_StrokeDashArray", instanceWithInferedType.Group_StrokeDashArray, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("DateYOffset", instanceWithInferedType.DateYOffset, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("AlignOnStartEndOnYearStart", instanceWithInferedType.AlignOnStartEndOnYearStart, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Lanes", instanceWithInferedType, &instanceWithInferedType.Lanes, formGroup, playground)
		AssociationSliceToForm("Milestones", instanceWithInferedType, &instanceWithInferedType.Milestones, formGroup, playground)
		AssociationSliceToForm("Groups", instanceWithInferedType, &instanceWithInferedType.Groups, formGroup, playground)
		AssociationSliceToForm("Arrows", instanceWithInferedType, &instanceWithInferedType.Arrows, formGroup, playground)

	case *models.Group:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("GroupLanes", instanceWithInferedType, &instanceWithInferedType.GroupLanes, formGroup, playground)

	case *models.Lane:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Order", instanceWithInferedType.Order, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Bars", instanceWithInferedType, &instanceWithInferedType.Bars, formGroup, playground)

	case *models.LaneUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("Lane", instanceWithInferedType.Lane, formGroup, playground)

	case *models.Milestone:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("DisplayVerticalBar", instanceWithInferedType.DisplayVerticalBar, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("LanesToDisplayMilestoneUse", instanceWithInferedType, &instanceWithInferedType.LanesToDisplayMilestoneUse, formGroup, playground)

	default:
		_ = instanceWithInferedType
	}
}

