// generated code - do not edit
package probe

import (
	"log"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gonggantt/go/models"
)

const __dummmy__time = time.Nanosecond

// insertion point
func NewArrowFormCallback(
	arrow *models.Arrow,
	playground *Playground,
) (arrowFormCallback *ArrowFormCallback) {
	arrowFormCallback = new(ArrowFormCallback)
	arrowFormCallback.playground = playground
	arrowFormCallback.arrow = arrow
	return
}

type ArrowFormCallback struct {
	arrow *models.Arrow

	playground *Playground
}

func (arrowFormCallback *ArrowFormCallback) OnSave() {

	log.Println("ArrowFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	arrowFormCallback.playground.formStage.Checkout()

	if arrowFormCallback.arrow == nil {
		arrowFormCallback.arrow = new(models.Arrow).Stage(arrowFormCallback.playground.stageOfInterest)
	}
	arrow_ := arrowFormCallback.arrow
	_ = arrow_

	// get the formGroup
	formGroup := arrowFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(arrow_.Name), formDiv)
		case "From":
			FormDivSelectFieldToField(&(arrow_.From), arrowFormCallback.playground.stageOfInterest, formDiv)
		case "To":
			FormDivSelectFieldToField(&(arrow_.To), arrowFormCallback.playground.stageOfInterest, formDiv)
		case "OptionnalColor":
			FormDivBasicFieldToField(&(arrow_.OptionnalColor), formDiv)
		case "OptionnalStroke":
			FormDivBasicFieldToField(&(arrow_.OptionnalStroke), formDiv)
		}
	}

	arrowFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Arrow](
		arrowFormCallback.playground,
	)
	arrowFormCallback.playground.tableStage.Commit()
}
func NewBarFormCallback(
	bar *models.Bar,
	playground *Playground,
) (barFormCallback *BarFormCallback) {
	barFormCallback = new(BarFormCallback)
	barFormCallback.playground = playground
	barFormCallback.bar = bar
	return
}

type BarFormCallback struct {
	bar *models.Bar

	playground *Playground
}

func (barFormCallback *BarFormCallback) OnSave() {

	log.Println("BarFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	barFormCallback.playground.formStage.Checkout()

	if barFormCallback.bar == nil {
		barFormCallback.bar = new(models.Bar).Stage(barFormCallback.playground.stageOfInterest)
	}
	bar_ := barFormCallback.bar
	_ = bar_

	// get the formGroup
	formGroup := barFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bar_.Name), formDiv)
		case "OptionnalColor":
			FormDivBasicFieldToField(&(bar_.OptionnalColor), formDiv)
		case "OptionnalStroke":
			FormDivBasicFieldToField(&(bar_.OptionnalStroke), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(bar_.FillOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(bar_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(bar_.StrokeDashArray), formDiv)
		}
	}

	barFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Bar](
		barFormCallback.playground,
	)
	barFormCallback.playground.tableStage.Commit()
}
func NewGanttFormCallback(
	gantt *models.Gantt,
	playground *Playground,
) (ganttFormCallback *GanttFormCallback) {
	ganttFormCallback = new(GanttFormCallback)
	ganttFormCallback.playground = playground
	ganttFormCallback.gantt = gantt
	return
}

type GanttFormCallback struct {
	gantt *models.Gantt

	playground *Playground
}

func (ganttFormCallback *GanttFormCallback) OnSave() {

	log.Println("GanttFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	ganttFormCallback.playground.formStage.Checkout()

	if ganttFormCallback.gantt == nil {
		ganttFormCallback.gantt = new(models.Gantt).Stage(ganttFormCallback.playground.stageOfInterest)
	}
	gantt_ := ganttFormCallback.gantt
	_ = gantt_

	// get the formGroup
	formGroup := ganttFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gantt_.Name), formDiv)
		case "UseManualStartAndEndDates":
			FormDivBasicFieldToField(&(gantt_.UseManualStartAndEndDates), formDiv)
		case "LaneHeight":
			FormDivBasicFieldToField(&(gantt_.LaneHeight), formDiv)
		case "RatioBarToLaneHeight":
			FormDivBasicFieldToField(&(gantt_.RatioBarToLaneHeight), formDiv)
		case "YTopMargin":
			FormDivBasicFieldToField(&(gantt_.YTopMargin), formDiv)
		case "XLeftText":
			FormDivBasicFieldToField(&(gantt_.XLeftText), formDiv)
		case "TextHeight":
			FormDivBasicFieldToField(&(gantt_.TextHeight), formDiv)
		case "XLeftLanes":
			FormDivBasicFieldToField(&(gantt_.XLeftLanes), formDiv)
		case "XRightMargin":
			FormDivBasicFieldToField(&(gantt_.XRightMargin), formDiv)
		case "ArrowLengthToTheRightOfStartBar":
			FormDivBasicFieldToField(&(gantt_.ArrowLengthToTheRightOfStartBar), formDiv)
		case "ArrowTipLenght":
			FormDivBasicFieldToField(&(gantt_.ArrowTipLenght), formDiv)
		case "TimeLine_Color":
			FormDivBasicFieldToField(&(gantt_.TimeLine_Color), formDiv)
		case "TimeLine_FillOpacity":
			FormDivBasicFieldToField(&(gantt_.TimeLine_FillOpacity), formDiv)
		case "TimeLine_Stroke":
			FormDivBasicFieldToField(&(gantt_.TimeLine_Stroke), formDiv)
		case "TimeLine_StrokeWidth":
			FormDivBasicFieldToField(&(gantt_.TimeLine_StrokeWidth), formDiv)
		case "Group_Stroke":
			FormDivBasicFieldToField(&(gantt_.Group_Stroke), formDiv)
		case "Group_StrokeWidth":
			FormDivBasicFieldToField(&(gantt_.Group_StrokeWidth), formDiv)
		case "Group_StrokeDashArray":
			FormDivBasicFieldToField(&(gantt_.Group_StrokeDashArray), formDiv)
		case "DateYOffset":
			FormDivBasicFieldToField(&(gantt_.DateYOffset), formDiv)
		case "AlignOnStartEndOnYearStart":
			FormDivBasicFieldToField(&(gantt_.AlignOnStartEndOnYearStart), formDiv)
		}
	}

	ganttFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Gantt](
		ganttFormCallback.playground,
	)
	ganttFormCallback.playground.tableStage.Commit()
}
func NewGroupFormCallback(
	group *models.Group,
	playground *Playground,
) (groupFormCallback *GroupFormCallback) {
	groupFormCallback = new(GroupFormCallback)
	groupFormCallback.playground = playground
	groupFormCallback.group = group
	return
}

type GroupFormCallback struct {
	group *models.Group

	playground *Playground
}

func (groupFormCallback *GroupFormCallback) OnSave() {

	log.Println("GroupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	groupFormCallback.playground.formStage.Checkout()

	if groupFormCallback.group == nil {
		groupFormCallback.group = new(models.Group).Stage(groupFormCallback.playground.stageOfInterest)
	}
	group_ := groupFormCallback.group
	_ = group_

	// get the formGroup
	formGroup := groupFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(group_.Name), formDiv)
		}
	}

	groupFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Group](
		groupFormCallback.playground,
	)
	groupFormCallback.playground.tableStage.Commit()
}
func NewLaneFormCallback(
	lane *models.Lane,
	playground *Playground,
) (laneFormCallback *LaneFormCallback) {
	laneFormCallback = new(LaneFormCallback)
	laneFormCallback.playground = playground
	laneFormCallback.lane = lane
	return
}

type LaneFormCallback struct {
	lane *models.Lane

	playground *Playground
}

func (laneFormCallback *LaneFormCallback) OnSave() {

	log.Println("LaneFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	laneFormCallback.playground.formStage.Checkout()

	if laneFormCallback.lane == nil {
		laneFormCallback.lane = new(models.Lane).Stage(laneFormCallback.playground.stageOfInterest)
	}
	lane_ := laneFormCallback.lane
	_ = lane_

	// get the formGroup
	formGroup := laneFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(lane_.Name), formDiv)
		case "Order":
			FormDivBasicFieldToField(&(lane_.Order), formDiv)
		}
	}

	laneFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Lane](
		laneFormCallback.playground,
	)
	laneFormCallback.playground.tableStage.Commit()
}
func NewLaneUseFormCallback(
	laneuse *models.LaneUse,
	playground *Playground,
) (laneuseFormCallback *LaneUseFormCallback) {
	laneuseFormCallback = new(LaneUseFormCallback)
	laneuseFormCallback.playground = playground
	laneuseFormCallback.laneuse = laneuse
	return
}

type LaneUseFormCallback struct {
	laneuse *models.LaneUse

	playground *Playground
}

func (laneuseFormCallback *LaneUseFormCallback) OnSave() {

	log.Println("LaneUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	laneuseFormCallback.playground.formStage.Checkout()

	if laneuseFormCallback.laneuse == nil {
		laneuseFormCallback.laneuse = new(models.LaneUse).Stage(laneuseFormCallback.playground.stageOfInterest)
	}
	laneuse_ := laneuseFormCallback.laneuse
	_ = laneuse_

	// get the formGroup
	formGroup := laneuseFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(laneuse_.Name), formDiv)
		case "Lane":
			FormDivSelectFieldToField(&(laneuse_.Lane), laneuseFormCallback.playground.stageOfInterest, formDiv)
		}
	}

	laneuseFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.LaneUse](
		laneuseFormCallback.playground,
	)
	laneuseFormCallback.playground.tableStage.Commit()
}
func NewMilestoneFormCallback(
	milestone *models.Milestone,
	playground *Playground,
) (milestoneFormCallback *MilestoneFormCallback) {
	milestoneFormCallback = new(MilestoneFormCallback)
	milestoneFormCallback.playground = playground
	milestoneFormCallback.milestone = milestone
	return
}

type MilestoneFormCallback struct {
	milestone *models.Milestone

	playground *Playground
}

func (milestoneFormCallback *MilestoneFormCallback) OnSave() {

	log.Println("MilestoneFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	milestoneFormCallback.playground.formStage.Checkout()

	if milestoneFormCallback.milestone == nil {
		milestoneFormCallback.milestone = new(models.Milestone).Stage(milestoneFormCallback.playground.stageOfInterest)
	}
	milestone_ := milestoneFormCallback.milestone
	_ = milestone_

	// get the formGroup
	formGroup := milestoneFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(milestone_.Name), formDiv)
		case "DisplayVerticalBar":
			FormDivBasicFieldToField(&(milestone_.DisplayVerticalBar), formDiv)
		}
	}

	milestoneFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Milestone](
		milestoneFormCallback.playground,
	)
	milestoneFormCallback.playground.tableStage.Commit()
}
