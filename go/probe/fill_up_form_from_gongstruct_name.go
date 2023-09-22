// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gonggantt/go/models"
)

func FillUpFormFromGongstructName(
	playground *Playground,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := playground.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = "New"
	} else {
		prefix = "Update"
	}

	switch gongstructName {
	// insertion point
	case "Arrow":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Arrow Form",
			OnSave: __gong__New__ArrowFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		arrow := new(models.Arrow)
		FillUpForm(arrow, formGroup, playground)
	case "Bar":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Bar Form",
			OnSave: __gong__New__BarFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		bar := new(models.Bar)
		FillUpForm(bar, formGroup, playground)
	case "Gantt":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Gantt Form",
			OnSave: __gong__New__GanttFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		gantt := new(models.Gantt)
		FillUpForm(gantt, formGroup, playground)
	case "Group":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Group Form",
			OnSave: __gong__New__GroupFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		group := new(models.Group)
		FillUpForm(group, formGroup, playground)
	case "Lane":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Lane Form",
			OnSave: __gong__New__LaneFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		lane := new(models.Lane)
		FillUpForm(lane, formGroup, playground)
	case "LaneUse":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " LaneUse Form",
			OnSave: __gong__New__LaneUseFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		laneuse := new(models.LaneUse)
		FillUpForm(laneuse, formGroup, playground)
	case "Milestone":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Milestone Form",
			OnSave: __gong__New__MilestoneFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		milestone := new(models.Milestone)
		FillUpForm(milestone, formGroup, playground)
	}
	formStage.Commit()
}
