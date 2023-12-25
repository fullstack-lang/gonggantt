// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gonggantt/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Arrow":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Arrow Form",
			OnSave: __gong__New__ArrowFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		arrow := new(models.Arrow)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(arrow, formGroup, probe)
	case "Bar":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Bar Form",
			OnSave: __gong__New__BarFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		bar := new(models.Bar)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bar, formGroup, probe)
	case "Gantt":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Gantt Form",
			OnSave: __gong__New__GanttFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		gantt := new(models.Gantt)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gantt, formGroup, probe)
	case "Group":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Group Form",
			OnSave: __gong__New__GroupFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		group := new(models.Group)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(group, formGroup, probe)
	case "Lane":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Lane Form",
			OnSave: __gong__New__LaneFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		lane := new(models.Lane)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(lane, formGroup, probe)
	case "LaneUse":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "LaneUse Form",
			OnSave: __gong__New__LaneUseFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		laneuse := new(models.LaneUse)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(laneuse, formGroup, probe)
	case "Milestone":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Milestone Form",
			OnSave: __gong__New__MilestoneFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		milestone := new(models.Milestone)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(milestone, formGroup, probe)
	}
	formStage.Commit()
}
