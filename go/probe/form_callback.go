// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gonggantt/go/models"
	"github.com/fullstack-lang/gonggantt/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__ArrowFormCallback(
	arrow *models.Arrow,
	playground *Playground,
) (arrowFormCallback *ArrowFormCallback) {
	arrowFormCallback = new(ArrowFormCallback)
	arrowFormCallback.playground = playground
	arrowFormCallback.arrow = arrow

	arrowFormCallback.CreationMode = (arrow == nil)

	return
}

type ArrowFormCallback struct {
	arrow *models.Arrow

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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
		case "Gantt:Arrows":
			// we need to retrieve the field owner before the change
			var pastGanttOwner *models.Gantt
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Gantt"
			rf.Fieldname = "Arrows"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				arrowFormCallback.playground.stageOfInterest,
				arrowFormCallback.playground.backRepoOfInterest,
				arrow_,
				&rf)

			if reverseFieldOwner != nil {
				pastGanttOwner = reverseFieldOwner.(*models.Gantt)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGanttOwner != nil {
					idx := slices.Index(pastGanttOwner.Arrows, arrow_)
					pastGanttOwner.Arrows = slices.Delete(pastGanttOwner.Arrows, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gantt := range *models.GetGongstructInstancesSet[models.Gantt](arrowFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _gantt.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGanttOwner := _gantt // we have a match
						if pastGanttOwner != nil {
							if newGanttOwner != pastGanttOwner {
								idx := slices.Index(pastGanttOwner.Arrows, arrow_)
								pastGanttOwner.Arrows = slices.Delete(pastGanttOwner.Arrows, idx, idx+1)
								newGanttOwner.Arrows = append(newGanttOwner.Arrows, arrow_)
							}
						} else {
							newGanttOwner.Arrows = append(newGanttOwner.Arrows, arrow_)
						}
					}
				}
			}
		}
	}

	arrowFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Arrow](
		arrowFormCallback.playground,
	)
	arrowFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if arrowFormCallback.CreationMode {
		arrowFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__ArrowFormCallback(
				nil,
				arrowFormCallback.playground,
			),
		}).Stage(arrowFormCallback.playground.formStage)
		arrow := new(models.Arrow)
		FillUpForm(arrow, newFormGroup, arrowFormCallback.playground)
		arrowFormCallback.playground.formStage.Commit()
	}

	fillUpTree(arrowFormCallback.playground)
}
func __gong__New__BarFormCallback(
	bar *models.Bar,
	playground *Playground,
) (barFormCallback *BarFormCallback) {
	barFormCallback = new(BarFormCallback)
	barFormCallback.playground = playground
	barFormCallback.bar = bar

	barFormCallback.CreationMode = (bar == nil)

	return
}

type BarFormCallback struct {
	bar *models.Bar

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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
		case "Lane:Bars":
			// we need to retrieve the field owner before the change
			var pastLaneOwner *models.Lane
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Lane"
			rf.Fieldname = "Bars"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				barFormCallback.playground.stageOfInterest,
				barFormCallback.playground.backRepoOfInterest,
				bar_,
				&rf)

			if reverseFieldOwner != nil {
				pastLaneOwner = reverseFieldOwner.(*models.Lane)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLaneOwner != nil {
					idx := slices.Index(pastLaneOwner.Bars, bar_)
					pastLaneOwner.Bars = slices.Delete(pastLaneOwner.Bars, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _lane := range *models.GetGongstructInstancesSet[models.Lane](barFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _lane.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLaneOwner := _lane // we have a match
						if pastLaneOwner != nil {
							if newLaneOwner != pastLaneOwner {
								idx := slices.Index(pastLaneOwner.Bars, bar_)
								pastLaneOwner.Bars = slices.Delete(pastLaneOwner.Bars, idx, idx+1)
								newLaneOwner.Bars = append(newLaneOwner.Bars, bar_)
							}
						} else {
							newLaneOwner.Bars = append(newLaneOwner.Bars, bar_)
						}
					}
				}
			}
		}
	}

	barFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Bar](
		barFormCallback.playground,
	)
	barFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if barFormCallback.CreationMode {
		barFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__BarFormCallback(
				nil,
				barFormCallback.playground,
			),
		}).Stage(barFormCallback.playground.formStage)
		bar := new(models.Bar)
		FillUpForm(bar, newFormGroup, barFormCallback.playground)
		barFormCallback.playground.formStage.Commit()
	}

	fillUpTree(barFormCallback.playground)
}
func __gong__New__GanttFormCallback(
	gantt *models.Gantt,
	playground *Playground,
) (ganttFormCallback *GanttFormCallback) {
	ganttFormCallback = new(GanttFormCallback)
	ganttFormCallback.playground = playground
	ganttFormCallback.gantt = gantt

	ganttFormCallback.CreationMode = (gantt == nil)

	return
}

type GanttFormCallback struct {
	gantt *models.Gantt

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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

	// display a new form by reset the form stage
	if ganttFormCallback.CreationMode {
		ganttFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__GanttFormCallback(
				nil,
				ganttFormCallback.playground,
			),
		}).Stage(ganttFormCallback.playground.formStage)
		gantt := new(models.Gantt)
		FillUpForm(gantt, newFormGroup, ganttFormCallback.playground)
		ganttFormCallback.playground.formStage.Commit()
	}

	fillUpTree(ganttFormCallback.playground)
}
func __gong__New__GroupFormCallback(
	group *models.Group,
	playground *Playground,
) (groupFormCallback *GroupFormCallback) {
	groupFormCallback = new(GroupFormCallback)
	groupFormCallback.playground = playground
	groupFormCallback.group = group

	groupFormCallback.CreationMode = (group == nil)

	return
}

type GroupFormCallback struct {
	group *models.Group

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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
		case "Gantt:Groups":
			// we need to retrieve the field owner before the change
			var pastGanttOwner *models.Gantt
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Gantt"
			rf.Fieldname = "Groups"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				groupFormCallback.playground.stageOfInterest,
				groupFormCallback.playground.backRepoOfInterest,
				group_,
				&rf)

			if reverseFieldOwner != nil {
				pastGanttOwner = reverseFieldOwner.(*models.Gantt)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGanttOwner != nil {
					idx := slices.Index(pastGanttOwner.Groups, group_)
					pastGanttOwner.Groups = slices.Delete(pastGanttOwner.Groups, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gantt := range *models.GetGongstructInstancesSet[models.Gantt](groupFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _gantt.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGanttOwner := _gantt // we have a match
						if pastGanttOwner != nil {
							if newGanttOwner != pastGanttOwner {
								idx := slices.Index(pastGanttOwner.Groups, group_)
								pastGanttOwner.Groups = slices.Delete(pastGanttOwner.Groups, idx, idx+1)
								newGanttOwner.Groups = append(newGanttOwner.Groups, group_)
							}
						} else {
							newGanttOwner.Groups = append(newGanttOwner.Groups, group_)
						}
					}
				}
			}
		}
	}

	groupFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Group](
		groupFormCallback.playground,
	)
	groupFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if groupFormCallback.CreationMode {
		groupFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__GroupFormCallback(
				nil,
				groupFormCallback.playground,
			),
		}).Stage(groupFormCallback.playground.formStage)
		group := new(models.Group)
		FillUpForm(group, newFormGroup, groupFormCallback.playground)
		groupFormCallback.playground.formStage.Commit()
	}

	fillUpTree(groupFormCallback.playground)
}
func __gong__New__LaneFormCallback(
	lane *models.Lane,
	playground *Playground,
) (laneFormCallback *LaneFormCallback) {
	laneFormCallback = new(LaneFormCallback)
	laneFormCallback.playground = playground
	laneFormCallback.lane = lane

	laneFormCallback.CreationMode = (lane == nil)

	return
}

type LaneFormCallback struct {
	lane *models.Lane

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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
		case "Gantt:Lanes":
			// we need to retrieve the field owner before the change
			var pastGanttOwner *models.Gantt
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Gantt"
			rf.Fieldname = "Lanes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				laneFormCallback.playground.stageOfInterest,
				laneFormCallback.playground.backRepoOfInterest,
				lane_,
				&rf)

			if reverseFieldOwner != nil {
				pastGanttOwner = reverseFieldOwner.(*models.Gantt)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGanttOwner != nil {
					idx := slices.Index(pastGanttOwner.Lanes, lane_)
					pastGanttOwner.Lanes = slices.Delete(pastGanttOwner.Lanes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gantt := range *models.GetGongstructInstancesSet[models.Gantt](laneFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _gantt.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGanttOwner := _gantt // we have a match
						if pastGanttOwner != nil {
							if newGanttOwner != pastGanttOwner {
								idx := slices.Index(pastGanttOwner.Lanes, lane_)
								pastGanttOwner.Lanes = slices.Delete(pastGanttOwner.Lanes, idx, idx+1)
								newGanttOwner.Lanes = append(newGanttOwner.Lanes, lane_)
							}
						} else {
							newGanttOwner.Lanes = append(newGanttOwner.Lanes, lane_)
						}
					}
				}
			}
		case "Group:GroupLanes":
			// we need to retrieve the field owner before the change
			var pastGroupOwner *models.Group
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "GroupLanes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				laneFormCallback.playground.stageOfInterest,
				laneFormCallback.playground.backRepoOfInterest,
				lane_,
				&rf)

			if reverseFieldOwner != nil {
				pastGroupOwner = reverseFieldOwner.(*models.Group)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGroupOwner != nil {
					idx := slices.Index(pastGroupOwner.GroupLanes, lane_)
					pastGroupOwner.GroupLanes = slices.Delete(pastGroupOwner.GroupLanes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _group := range *models.GetGongstructInstancesSet[models.Group](laneFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _group.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGroupOwner := _group // we have a match
						if pastGroupOwner != nil {
							if newGroupOwner != pastGroupOwner {
								idx := slices.Index(pastGroupOwner.GroupLanes, lane_)
								pastGroupOwner.GroupLanes = slices.Delete(pastGroupOwner.GroupLanes, idx, idx+1)
								newGroupOwner.GroupLanes = append(newGroupOwner.GroupLanes, lane_)
							}
						} else {
							newGroupOwner.GroupLanes = append(newGroupOwner.GroupLanes, lane_)
						}
					}
				}
			}
		}
	}

	laneFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Lane](
		laneFormCallback.playground,
	)
	laneFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if laneFormCallback.CreationMode {
		laneFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__LaneFormCallback(
				nil,
				laneFormCallback.playground,
			),
		}).Stage(laneFormCallback.playground.formStage)
		lane := new(models.Lane)
		FillUpForm(lane, newFormGroup, laneFormCallback.playground)
		laneFormCallback.playground.formStage.Commit()
	}

	fillUpTree(laneFormCallback.playground)
}
func __gong__New__LaneUseFormCallback(
	laneuse *models.LaneUse,
	playground *Playground,
) (laneuseFormCallback *LaneUseFormCallback) {
	laneuseFormCallback = new(LaneUseFormCallback)
	laneuseFormCallback.playground = playground
	laneuseFormCallback.laneuse = laneuse

	laneuseFormCallback.CreationMode = (laneuse == nil)

	return
}

type LaneUseFormCallback struct {
	laneuse *models.LaneUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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
		case "Milestone:LanesToDisplayMilestoneUse":
			// we need to retrieve the field owner before the change
			var pastMilestoneOwner *models.Milestone
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Milestone"
			rf.Fieldname = "LanesToDisplayMilestoneUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				laneuseFormCallback.playground.stageOfInterest,
				laneuseFormCallback.playground.backRepoOfInterest,
				laneuse_,
				&rf)

			if reverseFieldOwner != nil {
				pastMilestoneOwner = reverseFieldOwner.(*models.Milestone)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastMilestoneOwner != nil {
					idx := slices.Index(pastMilestoneOwner.LanesToDisplayMilestoneUse, laneuse_)
					pastMilestoneOwner.LanesToDisplayMilestoneUse = slices.Delete(pastMilestoneOwner.LanesToDisplayMilestoneUse, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _milestone := range *models.GetGongstructInstancesSet[models.Milestone](laneuseFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _milestone.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newMilestoneOwner := _milestone // we have a match
						if pastMilestoneOwner != nil {
							if newMilestoneOwner != pastMilestoneOwner {
								idx := slices.Index(pastMilestoneOwner.LanesToDisplayMilestoneUse, laneuse_)
								pastMilestoneOwner.LanesToDisplayMilestoneUse = slices.Delete(pastMilestoneOwner.LanesToDisplayMilestoneUse, idx, idx+1)
								newMilestoneOwner.LanesToDisplayMilestoneUse = append(newMilestoneOwner.LanesToDisplayMilestoneUse, laneuse_)
							}
						} else {
							newMilestoneOwner.LanesToDisplayMilestoneUse = append(newMilestoneOwner.LanesToDisplayMilestoneUse, laneuse_)
						}
					}
				}
			}
		}
	}

	laneuseFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.LaneUse](
		laneuseFormCallback.playground,
	)
	laneuseFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if laneuseFormCallback.CreationMode {
		laneuseFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__LaneUseFormCallback(
				nil,
				laneuseFormCallback.playground,
			),
		}).Stage(laneuseFormCallback.playground.formStage)
		laneuse := new(models.LaneUse)
		FillUpForm(laneuse, newFormGroup, laneuseFormCallback.playground)
		laneuseFormCallback.playground.formStage.Commit()
	}

	fillUpTree(laneuseFormCallback.playground)
}
func __gong__New__MilestoneFormCallback(
	milestone *models.Milestone,
	playground *Playground,
) (milestoneFormCallback *MilestoneFormCallback) {
	milestoneFormCallback = new(MilestoneFormCallback)
	milestoneFormCallback.playground = playground
	milestoneFormCallback.milestone = milestone

	milestoneFormCallback.CreationMode = (milestone == nil)

	return
}

type MilestoneFormCallback struct {
	milestone *models.Milestone

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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
		case "Gantt:Milestones":
			// we need to retrieve the field owner before the change
			var pastGanttOwner *models.Gantt
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Gantt"
			rf.Fieldname = "Milestones"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				milestoneFormCallback.playground.stageOfInterest,
				milestoneFormCallback.playground.backRepoOfInterest,
				milestone_,
				&rf)

			if reverseFieldOwner != nil {
				pastGanttOwner = reverseFieldOwner.(*models.Gantt)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGanttOwner != nil {
					idx := slices.Index(pastGanttOwner.Milestones, milestone_)
					pastGanttOwner.Milestones = slices.Delete(pastGanttOwner.Milestones, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gantt := range *models.GetGongstructInstancesSet[models.Gantt](milestoneFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _gantt.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGanttOwner := _gantt // we have a match
						if pastGanttOwner != nil {
							if newGanttOwner != pastGanttOwner {
								idx := slices.Index(pastGanttOwner.Milestones, milestone_)
								pastGanttOwner.Milestones = slices.Delete(pastGanttOwner.Milestones, idx, idx+1)
								newGanttOwner.Milestones = append(newGanttOwner.Milestones, milestone_)
							}
						} else {
							newGanttOwner.Milestones = append(newGanttOwner.Milestones, milestone_)
						}
					}
				}
			}
		}
	}

	milestoneFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Milestone](
		milestoneFormCallback.playground,
	)
	milestoneFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if milestoneFormCallback.CreationMode {
		milestoneFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__MilestoneFormCallback(
				nil,
				milestoneFormCallback.playground,
			),
		}).Stage(milestoneFormCallback.playground.formStage)
		milestone := new(models.Milestone)
		FillUpForm(milestone, newFormGroup, milestoneFormCallback.playground)
		milestoneFormCallback.playground.formStage.Commit()
	}

	fillUpTree(milestoneFormCallback.playground)
}
