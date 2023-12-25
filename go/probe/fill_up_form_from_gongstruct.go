// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gonggantt/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Arrow:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Arrow Form",
			OnSave: __gong__New__ArrowFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Bar:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Bar Form",
			OnSave: __gong__New__BarFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Gantt:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Gantt Form",
			OnSave: __gong__New__GanttFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Group:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Group Form",
			OnSave: __gong__New__GroupFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Lane:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Lane Form",
			OnSave: __gong__New__LaneFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.LaneUse:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "LaneUse Form",
			OnSave: __gong__New__LaneUseFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Milestone:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Milestone Form",
			OnSave: __gong__New__MilestoneFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
