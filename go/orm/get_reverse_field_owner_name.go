// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gonggantt/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Arrow:
		tmp := GetInstanceDBFromInstance[models.Arrow, ArrowDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			case "Arrows":
				if tmp != nil && tmp.Gantt_ArrowsDBID.Int64 != 0 {
					id := uint(tmp.Gantt_ArrowsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGantt.Map_GanttDBID_GanttPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}

	case *models.Bar:
		tmp := GetInstanceDBFromInstance[models.Bar, BarDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			case "Bars":
				if tmp != nil && tmp.Lane_BarsDBID.Int64 != 0 {
					id := uint(tmp.Lane_BarsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLane.Map_LaneDBID_LanePtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}

	case *models.Gantt:
		tmp := GetInstanceDBFromInstance[models.Gantt, GanttDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}

	case *models.Group:
		tmp := GetInstanceDBFromInstance[models.Group, GroupDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			case "Groups":
				if tmp != nil && tmp.Gantt_GroupsDBID.Int64 != 0 {
					id := uint(tmp.Gantt_GroupsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGantt.Map_GanttDBID_GanttPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}

	case *models.Lane:
		tmp := GetInstanceDBFromInstance[models.Lane, LaneDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			case "Lanes":
				if tmp != nil && tmp.Gantt_LanesDBID.Int64 != 0 {
					id := uint(tmp.Gantt_LanesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGantt.Map_GanttDBID_GanttPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Group":
			switch reverseField.Fieldname {
			case "GroupLanes":
				if tmp != nil && tmp.Group_GroupLanesDBID.Int64 != 0 {
					id := uint(tmp.Group_GroupLanesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGroup.Map_GroupDBID_GroupPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}

	case *models.LaneUse:
		tmp := GetInstanceDBFromInstance[models.LaneUse, LaneUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			case "LanesToDisplayMilestoneUse":
				if tmp != nil && tmp.Milestone_LanesToDisplayMilestoneUseDBID.Int64 != 0 {
					id := uint(tmp.Milestone_LanesToDisplayMilestoneUseDBID.Int64)
					reservePointerTarget := backRepo.BackRepoMilestone.Map_MilestoneDBID_MilestonePtr[id]
					res = reservePointerTarget.Name
				}
			}
		}

	case *models.Milestone:
		tmp := GetInstanceDBFromInstance[models.Milestone, MilestoneDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			case "Milestones":
				if tmp != nil && tmp.Gantt_MilestonesDBID.Int64 != 0 {
					id := uint(tmp.Gantt_MilestonesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGantt.Map_GanttDBID_GanttPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Arrow:
		tmp := GetInstanceDBFromInstance[models.Arrow, ArrowDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			case "Arrows":
				if tmp != nil && tmp.Gantt_ArrowsDBID.Int64 != 0 {
					id := uint(tmp.Gantt_ArrowsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGantt.Map_GanttDBID_GanttPtr[id]
					res = reservePointerTarget
				}
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.Bar:
		tmp := GetInstanceDBFromInstance[models.Bar, BarDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			case "Bars":
				if tmp != nil && tmp.Lane_BarsDBID.Int64 != 0 {
					id := uint(tmp.Lane_BarsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLane.Map_LaneDBID_LanePtr[id]
					res = reservePointerTarget
				}
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.Gantt:
		tmp := GetInstanceDBFromInstance[models.Gantt, GanttDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.Group:
		tmp := GetInstanceDBFromInstance[models.Group, GroupDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			case "Groups":
				if tmp != nil && tmp.Gantt_GroupsDBID.Int64 != 0 {
					id := uint(tmp.Gantt_GroupsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGantt.Map_GanttDBID_GanttPtr[id]
					res = reservePointerTarget
				}
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.Lane:
		tmp := GetInstanceDBFromInstance[models.Lane, LaneDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			case "Lanes":
				if tmp != nil && tmp.Gantt_LanesDBID.Int64 != 0 {
					id := uint(tmp.Gantt_LanesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGantt.Map_GanttDBID_GanttPtr[id]
					res = reservePointerTarget
				}
			}
		case "Group":
			switch reverseField.Fieldname {
			case "GroupLanes":
				if tmp != nil && tmp.Group_GroupLanesDBID.Int64 != 0 {
					id := uint(tmp.Group_GroupLanesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGroup.Map_GroupDBID_GroupPtr[id]
					res = reservePointerTarget
				}
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.LaneUse:
		tmp := GetInstanceDBFromInstance[models.LaneUse, LaneUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			case "LanesToDisplayMilestoneUse":
				if tmp != nil && tmp.Milestone_LanesToDisplayMilestoneUseDBID.Int64 != 0 {
					id := uint(tmp.Milestone_LanesToDisplayMilestoneUseDBID.Int64)
					reservePointerTarget := backRepo.BackRepoMilestone.Map_MilestoneDBID_MilestonePtr[id]
					res = reservePointerTarget
				}
			}
		}
	
	case *models.Milestone:
		tmp := GetInstanceDBFromInstance[models.Milestone, MilestoneDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			case "Milestones":
				if tmp != nil && tmp.Gantt_MilestonesDBID.Int64 != 0 {
					id := uint(tmp.Gantt_MilestonesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoGantt.Map_GanttDBID_GanttPtr[id]
					res = reservePointerTarget
				}
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}
	
	default:
		_ = inst
	}
	return res
}
